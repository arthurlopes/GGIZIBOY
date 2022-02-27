package GPU

type ICPU interface {
	SetIE(uint8)
}

type IMMU interface {
	ReadByte(uint16) uint8
	ReadWord(uint16) uint16
	WriteByte(uint16, uint8)
	WriteWord(uint16, uint16)
}

type GPU_struct struct {
	// GPU internal values
	mode      int
	modeclock int
	gpu_clock int
	screen    [][]uint8

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

	CPU ICPU
	MMU IMMU
}

func GPUFactory() *GPU_struct {
	var gpu = GPU_struct{}
	gpu.Innit()
	return &gpu
}

func (gpu *GPU_struct) Innit() {
	gpu.screen = make([][]uint8, 160)
	for i := range gpu.screen {
		gpu.screen[i] = make([]uint8, 144)
	}
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
				r_pix = 1
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

// func (GPU *GPU_struct) render_screen() *[]uint8 {

// 	// print(tile_map)
// 	// for tile_idx in mmu.memory[0x9800:0x9bff+1] {
// 	// 	screen.append(tile_map[tile_idx])
// 	// }

// 	// plt.figure()
// 	// plt.imshow(np.array(screen).reshape(32, 32, 8, 8).reshape(32, 256, 8).reshape(256, 256))
// 	// plt.show()

// 	return &GPU.screen
// }

func (gpu *GPU_struct) Step() {
	// render_screen()
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
			// renderscan()
		}
		// Hblank
	} else if gpu.mode == 0 {
		if gpu.modeclock >= 204 {
			gpu.modeclock = 0
			gpu.Line += 1
			if gpu.Line == 143 {
				gpu.mode = 1
				// CPU.IE |= 0x01
				// render_screen()
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
