package GPU

const V_BLANK_BIT uint8 = 0x01

type ICPU interface {
	SetIE(uint8)
	SetIF(uint8)
}

type IMMU interface {
	ReadByte(uint16) uint8
	ReadWord(uint16) uint16
	WriteByte(uint16, uint8)
	WriteWord(uint16, uint16)
}

type GPU_struct struct {
	// GPU internal values
	mode       int
	modeclock  int
	gpu_clock  int
	Screen     [][]uint8
	Background [][]uint8

	// GPU registers
	Lcd_control        uint8
	Gpu_control        uint8
	Scroll_y           uint8
	Scroll_x           uint8
	Cur_scanline       uint8
	Line               uint8
	Background_palette uint8
	Switchbg           uint8
	Bgmap              uint8
	Bgtile             uint8
	Switchlcd          uint8

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
}

func (gpu *GPU_struct) Update_clock(cpu_clock int) {
	gpu.modeclock += cpu_clock - gpu.gpu_clock
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

	for address := 0x8000; address < 0x9000; address += 16 {
		var tile_raw [16]uint8
		for i := 0; i < 16; i++ {
			tile_raw[i] = gpu.MMU.ReadByte(uint16(address + i))
		}
		var tile []uint8 = gpu.decode_tile(tile_raw)
		tile_map[uint16((address-0x8000)/16)] = tile
	}

	return tile_map
}

func (gpu *GPU_struct) Render_Background() [][]uint8 {
	var tile_map = gpu.Create_tile_map()
	var tile_i, tile_j, i, j uint8
	for address := 0x9800; address <= 0x9bff; address++ {
		var tile_idx = gpu.MMU.ReadByte(uint16(address))
		var tile = tile_map[uint16(tile_idx)]

		for i = 0; i < 8; i++ {
			for j = 0; j < 8; j++ {
				gpu.Background[tile_i+i][tile_j+j] = tile[8*i+j]
			}
		}

		tile_j = (tile_j + 8) % 255
		if tile_j == 0 {
			tile_i = (tile_i + 8) % 255
		}
	}

	return gpu.Background
}

func (gpu *GPU_struct) Render_TileMap() [][]uint8 {
	var tile_map = gpu.Create_tile_map()
	var tile_i, tile_j, i, j uint8
	for tile_idx := 0; tile_idx < 1024; tile_idx++ {
		if tile, ok := tile_map[uint16(tile_idx)]; ok {
			for i = 0; i < 8; i++ {
				for j = 0; j < 8; j++ {
					gpu.Background[tile_i+i][tile_j+j] = tile[8*i+j]
				}
			}
		}

		tile_j = (tile_j + 8) % 255
		if tile_j == 0 {
			tile_i = (tile_i + 8) % 255
		}
	}

	return gpu.Background
}

func (gpu *GPU_struct) Step() {
	// render_Screen()
	// Scanline - OAM
	if gpu.mode == 2 {
		if gpu.modeclock >= 80 {
			gpu.modeclock = 0
			gpu.mode = 3
		}
		// Scanline - VRAM
	} else if gpu.mode == 3 {
		if gpu.modeclock >= 172 {
			gpu.modeclock = 0
			gpu.mode = 0
		}
		// Hblank
	} else if gpu.mode == 0 {
		if gpu.modeclock >= 204 {
			gpu.modeclock = 0
			gpu.Line += 1
			if gpu.Line == 143 {
				gpu.mode = 1
				gpu.CPU.SetIE(V_BLANK_BIT)
				gpu.CPU.SetIF(V_BLANK_BIT)

				gpu.Render_Background()
				// gpu.Render_TileMap()

				gpu.hblank_channel <- true
				// select {
				// case gpu.hblank_channel <- true:
				// 	// fmt.Println("sent message", msg)
				// default:
				// 	// fmt.Println("no message sent")
				// }
			} else {
				gpu.mode = 2
			}
		}
		// Vblank
	} else if gpu.mode == 1 {
		if gpu.modeclock >= 456 {
			gpu.modeclock = 0
			gpu.Line += 1
			if gpu.Line > 153 {
				gpu.mode = 2
				gpu.Line = 0
			}
		}
	}
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

func (gpu *GPU_struct) SetLine(n uint8) {
	gpu.Line = n
}
