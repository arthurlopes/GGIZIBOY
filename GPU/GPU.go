package GPU

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
}

type GPU_struct struct {
	// GPU internal values
	mode       uint8
	modeclock  int
	gpu_clock  int
	Screen     [][]uint8
	Background [][]uint8

	// GPU registers
	Lcd_control uint8
	Scroll_y    uint8
	Scroll_x    uint8
	Line        uint8
	BGP         uint8
	BGP_map     map[uint8]uint8

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

	CPU ICPU
	MMU IMMU
}

func GPUFactory(hblank_channel chan bool) *GPU_struct {
	var gpu = GPU_struct{}
	gpu.Innit(hblank_channel)
	return &gpu
}

func (gpu *GPU_struct) Innit(hblank_channel chan bool) {
	gpu.Screen = make([][]uint8, 160)
	for i := range gpu.Screen {
		gpu.Screen[i] = make([]uint8, 144)
	}

	gpu.Background = make([][]uint8, 256)
	for i := range gpu.Background {
		gpu.Background[i] = make([]uint8, 256)
	}

	gpu.hblank_channel = hblank_channel
	gpu.BGP_map = make(map[uint8]uint8)
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
	var tile_map = gpu.Create_tile_map()
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

		var tile = tile_map[tile_start_offset+uint16(tile_idx)]

		for i = 0; i < 8; i++ {
			for j = 0; j < 8; j++ {
				gpu.Background[tile_i+i][tile_j+j] = gpu.BGP_map[tile[8*i+j]]
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
	var tile_map = gpu.Create_tile_map()
	var tile_i, tile_j, i, j uint8
	for tile_idx := 0; tile_idx < 384; tile_idx++ {
		if tile, ok := tile_map[uint16(tile_idx)]; ok {
			for i = 0; i < 8; i++ {
				for j = 0; j < 8; j++ {
					gpu.Background[tile_i+i][tile_j+j] = gpu.BGP_map[tile[8*i+j]]
				}
			}
		}
		tile_j = (tile_j + 8) % 128
		if tile_j == 0 {
			tile_i = (tile_i + 8)
		}
	}

	return gpu.Background
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
					gpu.Render_Background()
					// gpu.Render_TileMap()
					gpu.MMU.Set_VRAM_modified(false)
					gpu.hblank_channel <- true
					// select {
					// case gpu.hblank_channel <- true:
					// 	// fmt.Println("sent message", msg)
					// default:
					// 	// fmt.Println("no message sent")
					// }
				}
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

func (gpu *GPU_struct) GetBGP() uint8 {
	return gpu.BGP
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
