package gameboy

import (
	"testing"
)

func TestLD(t *testing.T) {
	var gb = GameboyFactory()

	gb.mmu.Bootstrap_path = "../" + gb.mmu.Bootstrap_path
	gb.mmu.Innit()

	gb.cpu.Registers.PC = 0x0050
	gb.mmu.WriteByte(0x0050, 0x06)
	gb.mmu.WriteByte(0x0051, 0x45)

	gb.Run(1)

	if gb.cpu.Registers.B != 0x45 {
		t.Errorf("Issue on instruction 0x06, value in B 0x%X", gb.cpu.Registers.B)
	}
}
