package GPU

import "sort"

const V_BLANK_BIT uint8 = 0x01
const LCD_BIT uint8 = 0x02

type ICPU interface {
	SetIE(uint8)
	SetIF(uint8)
	SetInterrupt(uint8)
}

type IMMU interface {
	ReadByte(uint16) uint8
	ReadWord(uint16) uint16
	WriteByte(uint16, uint8)
	WriteWord(uint16, uint16)
	Get_VRAM_modified() bool
	Set_VRAM_modified(bool)
	Get_OAM_modified() bool
	Set_OAM_modified(bool)
}

type OAM_Struct struct {
	sprites []Sprite_Struct
}

type Sprite_Struct struct {
	x, y, tile_no, flags              uint8
	priority, y_flip, x_flip, palette uint8
}

type GPU_struct struct {
	// GPU internal values
	mode       uint8
	modeclock  int
	gpu_clock  int
	Screen     [][]uint8
	Background [][]uint8
	Tile_data  [][]uint8
	tile_map   map[uint16][]uint8

	// GPU registers
	Lcd_control uint8
	Scroll_y    uint8
	Scroll_x    uint8
	Line        uint8
	BGP         uint8
	OBP0        uint8
	OBP1        uint8
	BGP_map     map[uint8]uint8
	OBP0_map    map[uint8]uint8
	OBP1_map    map[uint8]uint8

	// LCDC
	Background_palette     uint8
	WindowTileMap          uint8
	WindowDisplay          uint8
	BGWindowTileDataSelect uint8
	BGWindowTileMapSelect  uint8
	SpriteSize             uint8
	SpriteDisplay          uint8
	BGWindowDisplay        uint8
	// STAT
	LYC        uint8
	LYC_flag   uint8
	LYC_select uint8
	Mode00     uint8
	Mode01     uint8
	Mode10     uint8

	// Channels
	hblank_channel chan bool

	// OAM
	OAM OAM_Struct

	CPU ICPU
	MMU IMMU
}

func GPUFactory(hblank_channel chan bool) *GPU_struct {
	var gpu = GPU_struct{}
	gpu.Innit(hblank_channel)
	return &gpu
}

func (gpu *GPU_struct) Innit(hblank_channel chan bool) {
	gpu.Screen = make([][]uint8, 144)
	for i := range gpu.Screen {
		gpu.Screen[i] = make([]uint8, 160)
	}

	gpu.Background = make([][]uint8, 256)
	for i := range gpu.Background {
		gpu.Background[i] = make([]uint8, 256)
	}

	gpu.Tile_data = make([][]uint8, 256)
	for i := range gpu.Tile_data {
		gpu.Tile_data[i] = make([]uint8, 256)
	}

	gpu.OAM = OAM_Struct{}
	gpu.OAM.sprites = make([]Sprite_Struct, 40)

	gpu.hblank_channel = hblank_channel
	gpu.BGP_map = make(map[uint8]uint8)
	gpu.OBP0_map = make(map[uint8]uint8)
	gpu.OBP1_map = make(map[uint8]uint8)
}

func (gpu *GPU_struct) Update_clock(cpu_clock_delta int, cpu_clock int) {
	if (gpu.Lcd_control & 1) != 0 {
		gpu.modeclock += cpu_clock_delta
	}
	gpu.gpu_clock = cpu_clock
}

func (gpu *GPU_struct) decode_tile(tile [16]uint8) []uint8 {
	tile_dec := make([]uint8, 0)

	for i := 0; i < 16; i += 2 {
		for j := 0; j < 8; j++ {
			var l_pix uint8 = 0
			if (tile[i] & (1 << (7 - j))) > 0 {
				l_pix = 1
			}
			var r_pix uint8 = 0
			if (tile[i+1] & (1 << (7 - j))) > 0 {
				r_pix = 2
			}
			var pixel uint8 = l_pix + r_pix
			tile_dec = append(tile_dec, pixel)
		}
	}

	return tile_dec
}

func (gpu *GPU_struct) Create_tile_map() map[uint16][]uint8 {
	var tile_map = make(map[uint16][]uint8)
	var tile_start_address uint16 = 0x8000

	for address := tile_start_address; address < tile_start_address+0x1800; address += 16 {
		var tile_raw [16]uint8
		for i := uint16(0); i < 16; i++ {
			tile_raw[i] = gpu.MMU.ReadByte(address + i)
		}
		var tile []uint8 = gpu.decode_tile(tile_raw)
		tile_map[uint16((address-tile_start_address)/16)] = tile
	}

	return tile_map
}

func (gpu *GPU_struct) Render_Background() [][]uint8 {
	var tile_i, tile_j, i, j uint8

	var tile_map_adress uint16
	if gpu.BGWindowTileMapSelect == 0 {
		tile_map_adress = 0x9800
	} else {
		tile_map_adress = 0x9C00
	}

	var tile_start_offset uint16
	for address := tile_map_adress; address < tile_map_adress+0x400; address++ {
		var tile_idx = gpu.MMU.ReadByte(uint16(address))

		tile_start_offset = 0
		if (gpu.BGWindowTileDataSelect == 0) && (tile_idx < 128) {
			tile_start_offset = 256
		}

		var tile = gpu.tile_map[tile_start_offset+uint16(tile_idx)]

		for i = 0; i < 8; i++ {
			for j = 0; j < 8; j++ {
				gpu.Background[tile_i+i][tile_j+j] = tile[8*i+j]
			}
		}

		tile_j = (tile_j + 8)
		if tile_j == 0 {
			tile_i = (tile_i + 8)
		}
	}

	return gpu.Background
}

func (gpu *GPU_struct) Render_TileMap() [][]uint8 {
	var tile_i, tile_j, i, j uint8
	for tile_idx := 0; tile_idx < 384; tile_idx++ {
		if tile, ok := gpu.tile_map[uint16(tile_idx)]; ok {
			for i = 0; i < 8; i++ {
				for j = 0; j < 8; j++ {
					gpu.Tile_data[tile_i+i][tile_j+j] = gpu.BGP_map[tile[8*i+j]]
				}
			}
		}
		tile_j = (tile_j + 8) % 128
		if tile_j == 0 {
			tile_i = (tile_i + 8)
		}
	}

	return gpu.Tile_data
}

func (gpu *GPU_struct) Render_OAM() {
	for i := uint16(0); i < 40; i++ {
		spriteAddress := 0xFE00 + 4*i
		gpu.OAM.sprites[i].y = gpu.MMU.ReadByte(spriteAddress)
		gpu.OAM.sprites[i].x = gpu.MMU.ReadByte(spriteAddress + 1)
		gpu.OAM.sprites[i].tile_no = gpu.MMU.ReadByte(spriteAddress + 2)
		gpu.OAM.sprites[i].flags = gpu.MMU.ReadByte(spriteAddress + 3)

		gpu.OAM.sprites[i].priority = (gpu.OAM.sprites[i].flags & 0b10000000) >> 7
		gpu.OAM.sprites[i].y_flip = (gpu.OAM.sprites[i].flags & 0b01000000) >> 6
		gpu.OAM.sprites[i].x_flip = (gpu.OAM.sprites[i].flags & 0b00100000) >> 5
		gpu.OAM.sprites[i].palette = (gpu.OAM.sprites[i].flags & 0b00010000) >> 4
	}

	// Sort based on x
	sort.Slice(gpu.OAM.sprites, func(i, j int) bool {
		return gpu.OAM.sprites[i].x > gpu.OAM.sprites[j].x
	})

}

func (gpu *GPU_struct) Render_Screen_Line() {
	if gpu.MMU.Get_OAM_modified() {
		gpu.Render_OAM()
		gpu.MMU.Set_OAM_modified(false)
	}

	// Load Background
	y := (gpu.Scroll_y + gpu.Line)
	for i := uint8(0); i < 160; i++ {
		x := (gpu.Scroll_x + i)
		gpu.Screen[gpu.Line][i] = gpu.BGP_map[gpu.Background[y][x]]
	}

	// Load sprites
	sprites_on_current_line := make([]uint8, 0)
	for i := uint8(0); i < 40; i++ {
		sprite := gpu.OAM.sprites[i]
		if (sprite.x < 8) || (sprite.y < 16) || (sprite.x >= 160+8) || (sprite.y >= 144+16) {
			continue
		}
		if (sprite.y-16 <= gpu.Line) && (gpu.Line < sprite.y-8) {
			sprites_on_current_line = append(sprites_on_current_line, i)
		}
	}

	for _, i := range sprites_on_current_line {
		sprite := gpu.OAM.sprites[i]
		for j := uint8(0); j < 8; j++ {
			tile := gpu.tile_map[uint16(sprite.tile_no)]
			tile_pixel_idx := ((gpu.Line - (sprite.y - 16)) * 8)
			var x_offset uint8
			if sprite.x_flip != 0 {
				x_offset = (7 - j)
			} else {
				x_offset = j
			}
			tile_pixel := tile[tile_pixel_idx+x_offset]

			var tile_paletted uint8
			if sprite.palette == 0 {
				tile_paletted = gpu.OBP0_map[tile_pixel]
			} else {
				tile_paletted = gpu.OBP1_map[tile_pixel]
			}

			x_pos := sprite.x - 8 + j
			if ((sprite.priority == 0) && tile_pixel != 0) || (gpu.Background[y][gpu.Scroll_x+x_pos] == 0) {
				gpu.Screen[gpu.Line][x_pos] = tile_paletted
			}
		}
	}
}

func (gpu *GPU_struct) check_LYC_interrupt() {
	if gpu.LYC_select != 0 {
		if gpu.LYC_flag != 0 {
			gpu.CPU.SetInterrupt(LCD_BIT)
		}
	}
}

func (gpu *GPU_struct) Step() {
	if gpu.Lcd_control == 0 {
		return
	}

	OAM_TIME := 80
	VRAM_TIME := 170
	HBLANK_TIME := 206
	VBLANK_TIME := 456

	// Scanline - OAM
	if gpu.mode == 2 {
		if gpu.modeclock >= OAM_TIME {
			gpu.modeclock = 0
			gpu.mode = 3
		}
		// Scanline - VRAM
	} else if gpu.mode == 3 {
		if gpu.modeclock >= VRAM_TIME {
			gpu.modeclock = 0
			gpu.mode = 0
			if gpu.Mode00 != 0 {
				gpu.CPU.SetInterrupt(LCD_BIT)
			}
			gpu.Render_Screen_Line()
		}
		// Hblank
	} else if gpu.mode == 0 {
		if gpu.modeclock >= HBLANK_TIME {
			gpu.modeclock = 0
			gpu.Line += 1
			if gpu.Line == 143 {
				gpu.mode = 1
				if gpu.Mode01 != 0 {
					gpu.CPU.SetInterrupt(LCD_BIT)
				}
				gpu.CPU.SetInterrupt(V_BLANK_BIT)

				if gpu.MMU.Get_VRAM_modified() {
					gpu.tile_map = gpu.Create_tile_map()
					gpu.Render_TileMap()
					gpu.Render_Background()

					gpu.MMU.Set_VRAM_modified(false)

				}
				gpu.hblank_channel <- true
			} else {
				gpu.mode = 2
				if gpu.Mode10 != 0 {
					gpu.CPU.SetInterrupt(LCD_BIT)
				}
			}
			if gpu.LYC == gpu.Line {
				gpu.LYC_flag = 1
			} else {
				gpu.LYC_flag = 0
			}
			gpu.check_LYC_interrupt()
		}
		// Vblank
	} else if gpu.mode == 1 {
		if gpu.modeclock >= VBLANK_TIME {
			gpu.modeclock = 0
			gpu.Line += 1
			if gpu.Line == 153 {
				gpu.mode = 2
				gpu.Line = 0
				if (gpu.Mode10 != 0) || (gpu.Mode01 != 0) {
					gpu.CPU.SetInterrupt(LCD_BIT)
				}
			}
			if gpu.LYC == gpu.Line {
				gpu.LYC_flag = 1
			} else {
				gpu.LYC_flag = 0
			}
			gpu.check_LYC_interrupt()
		}
	}
}

func (gpu *GPU_struct) SetBGP(n uint8) {
	gpu.BGP = n
	gpu.BGP_map[0] = n & 0b00000011
	gpu.BGP_map[1] = n & 0b00001100 >> 2
	gpu.BGP_map[2] = n & 0b00110000 >> 4
	gpu.BGP_map[3] = n & 0b11000000 >> 6
}

func (gpu *GPU_struct) SetOBP0(n uint8) {
	gpu.OBP0 = n
	gpu.OBP0_map[0] = n & 0b00000011
	gpu.OBP0_map[1] = n & 0b00001100 >> 2
	gpu.OBP0_map[2] = n & 0b00110000 >> 4
	gpu.OBP0_map[3] = n & 0b11000000 >> 6
}

func (gpu *GPU_struct) SetOBP1(n uint8) {
	gpu.OBP1 = n
	gpu.OBP1_map[0] = n & 0b00000011
	gpu.OBP1_map[1] = n & 0b00001100 >> 2
	gpu.OBP1_map[2] = n & 0b00110000 >> 4
	gpu.OBP1_map[3] = n & 0b11000000 >> 6
}

func (gpu *GPU_struct) GetBGP() uint8 {
	return gpu.BGP
}

func (gpu *GPU_struct) GetOBP0() uint8 {
	return gpu.OBP0
}

func (gpu *GPU_struct) GetOBP1() uint8 {
	return gpu.OBP1
}

func (gpu *GPU_struct) GetLine() uint8 {
	return gpu.Line
}

func (gpu *GPU_struct) GetScroll_x() uint8 {
	return gpu.Scroll_x
}

func (gpu *GPU_struct) GetScroll_y() uint8 {
	return gpu.Scroll_y
}

func (gpu *GPU_struct) SetScroll_x(n uint8) {
	gpu.Scroll_x = n
}

func (gpu *GPU_struct) SetScroll_y(n uint8) {
	gpu.Scroll_y = n
}

func (gpu *GPU_struct) GetLYC() uint8 {
	return gpu.LYC
}

func (gpu *GPU_struct) SetLYC(n uint8) {
	gpu.LYC = n
}

func (gpu *GPU_struct) GetLCDC() uint8 {
	var lcdc uint8 = (gpu.Lcd_control << 7) |
		(gpu.WindowTileMap << 6) |
		(gpu.WindowDisplay << 5) |
		(gpu.BGWindowTileDataSelect << 4) |
		(gpu.BGWindowTileMapSelect << 3) |
		(gpu.SpriteSize << 2) |
		(gpu.SpriteDisplay << 1) |
		gpu.BGWindowDisplay

	return lcdc
}

func (gpu *GPU_struct) GetSTAT() uint8 {
	var stat uint8 = (1 << 7) |
		(gpu.LYC_select << 6) |
		(gpu.Mode10 << 5) |
		(gpu.Mode01 << 4) |
		(gpu.Mode00 << 3) |
		(gpu.LYC_flag << 2) |
		gpu.mode

	return stat
}

func (gpu *GPU_struct) SetLCDC(n uint8) {
	if (gpu.Lcd_control != 0) && ((n & 0b10000000) == 0) {
		gpu.mode = 0
		gpu.Line = 0
	}

	gpu.Lcd_control = (n & 0b10000000) >> 7
	gpu.WindowTileMap = (n & 0b01000000) >> 6
	gpu.WindowDisplay = (n & 0b00100000) >> 5
	gpu.BGWindowTileDataSelect = (n & 0b00010000) >> 4
	gpu.BGWindowTileMapSelect = (n & 0b00001000) >> 3
	gpu.SpriteSize = (n & 0b00000100) >> 2
	gpu.SpriteDisplay = (n & 0b00000010) >> 1
	gpu.BGWindowDisplay = (n & 0b00000001)
}

func (gpu *GPU_struct) SetSTAT(n uint8) {
	gpu.LYC_select = (n & 0b01000000) >> 6
	gpu.Mode10 = (n & 0b00100000) >> 5
	gpu.Mode01 = (n & 0b00010000) >> 4
	gpu.Mode00 = (n & 0b00001000) >> 3
	// gpu.LYC_flag = (n & 0b00000100) >> 2
	// gpu.mode = (n & 0b00000011)
}
