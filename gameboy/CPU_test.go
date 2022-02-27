package gameboy

import (
	"testing"
)

func TestLD(t *testing.T) {
	ch := make(chan bool)
	var gb = GameboyFactory(ch)

	gb.MMU.Bootstrap_path = "../" + gb.MMU.Bootstrap_path
	gb.MMU.Innit()

	gb.CPU.Registers.PC = 0x0050
	gb.MMU.WriteByte(0x0050, 0x06)
	gb.MMU.WriteByte(0x0051, 0x45)

	gb.Run(1)

	if gb.CPU.Registers.B != 0x45 {
		t.Errorf("Issue on instruction 0x06, value in B 0x%X", gb.CPU.Registers.B)
	}
}
