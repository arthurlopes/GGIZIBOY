package gameboy

import (
	"GGIZIBOY/CPU"
	"GGIZIBOY/GPU"
	"GGIZIBOY/MMU"
)

type Gameboy_struct struct {
	MMU *MMU.MMU_struct
	CPU *CPU.CPU_struct
	GPU *GPU.GPU_struct
}

func GameboyFactory(hblank_channel chan bool, input_channel chan uint8) *Gameboy_struct {
	var gb = Gameboy_struct{}
	gb.Innit(hblank_channel, input_channel)
	return &gb
}

func (gb *Gameboy_struct) Innit(hblank_channel chan bool, input_channel chan uint8) {
	gb.MMU = MMU.MMUFactory(input_channel)
	gb.CPU = CPU.CPUFactory()
	gb.GPU = GPU.GPUFactory(hblank_channel)

	gb.GPU.CPU = gb.CPU
	gb.GPU.MMU = gb.MMU

	gb.MMU.CPU = gb.CPU
	gb.MMU.GPU = gb.GPU

	gb.CPU.MMU = gb.MMU
	gb.CPU.GPU = gb.GPU
}

func (gb *Gameboy_struct) Run(limit int) {
	gb.CPU.Run(limit)
}
