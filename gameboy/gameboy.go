package gameboy

import (
	"GGIZIBOY/CPU"
	"GGIZIBOY/GPU"
	"GGIZIBOY/MMU"
)

type Gameboy_struct struct {
	mmu *MMU.MMU_struct
	cpu *CPU.CPU_struct
	gpu *GPU.GPU_struct
}

func GameboyFactory() *Gameboy_struct {
	var gb = Gameboy_struct{}
	gb.Innit()
	return &gb
}

func (gb *Gameboy_struct) Innit() {
	gb.mmu = MMU.MMUFactory()
	gb.cpu = CPU.CPUFactory()
	gb.gpu = GPU.GPUFactory()

	gb.gpu.CPU = gb.cpu
	gb.gpu.MMU = gb.mmu

	gb.mmu.CPU = gb.cpu
	gb.mmu.GPU = gb.gpu

	gb.cpu.MMU = gb.mmu
	gb.cpu.GPU = gb.gpu
}

func (gb *Gameboy_struct) Run(limit int) {
	gb.cpu.Run(limit)
}
