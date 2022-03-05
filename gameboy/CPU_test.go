package gameboy

import (
	"testing"
)

func TestLD(t *testing.T) {
	ch := make(chan bool)
	var gb = GameboyFactory(ch)

	gb.CPU.Registers.PC = 0x0050
	gb.MMU.WriteByte(0x0050, 0x06)
	gb.MMU.WriteByte(0x0051, 0x45)

	gb.Run(1)

	if gb.CPU.Registers.B != 0x45 {
		t.Errorf("Issue on instruction 0x06, value in B 0x%X", gb.CPU.Registers.B)
	}
}

func TestADD(t *testing.T) {
	ch := make(chan bool)
	var gb = GameboyFactory(ch)

	gb.CPU.Registers.PC = 0x0050
	gb.CPU.Registers.A = 0x01
	gb.CPU.Registers.H = 0xf2
	gb.CPU.Registers.L = 0x34
	gb.MMU.WriteByte(0x0050, 0x86)
	gb.MMU.WriteByte(0xf234, 0x45)

	gb.Run(1)

	if gb.CPU.Registers.A != 0x46 {
		t.Errorf("Issue on instruction 0x86,  0x%X, 0x%X", gb.CPU.Registers.A, gb.MMU.Memory[0x1234])
	}

	if gb.CPU.Registers.F != 0x00 {
		t.Errorf("Issue on instruction 0x86, value in F 0x%X", gb.CPU.Registers.B)
	}
}

func TestADDOverflow(t *testing.T) {
	ch := make(chan bool)
	var gb = GameboyFactory(ch)

	gb.CPU.Registers.PC = 0x0050
	gb.CPU.Registers.A = 0x02
	gb.CPU.Registers.H = 0xf2
	gb.CPU.Registers.L = 0x34
	gb.MMU.WriteByte(0x0050, 0x86)
	gb.MMU.WriteByte(0xf234, 0xff)

	gb.Run(1)

	// if gb.CPU.Registers.A != 0x01 {
	// 	t.Errorf("Issue on instruction 0x86,  0x%X, 0x%X", gb.CPU.Registers.A, gb.MMU.Memory[0x1234])
	// }

	// if gb.CPU.Registers.F != 0x10 {
	// 	t.Errorf("Issue on instruction 0x86, value in F 0x%X", gb.CPU.Registers.B)
	// }
}

func TestJR(t *testing.T) {
	ch := make(chan bool)
	var gb = GameboyFactory(ch)

	gb.CPU.Registers.PC = 0x0050
	gb.MMU.WriteByte(0x0050, 0x18)
	gb.MMU.WriteByte(0x0051, 0x05)

	gb.Run(1)

	if gb.CPU.Registers.PC != 0x57 {
		t.Errorf("Issue on instruction 0x18,  0x%X, %d", gb.CPU.Registers.PC, gb.MMU.Memory[0x0051])
	}

	gb.CPU.Registers.PC = 0x0050
	gb.MMU.WriteByte(0x0050, 0x18)
	gb.MMU.WriteByte(0x0051, 0xfe)

	gb.Run(1)

	if gb.CPU.Registers.PC != 0x50 {
		t.Errorf("Issue on instruction 0x18,  0x%X, %d", gb.CPU.Registers.PC, int8(gb.MMU.Memory[0x0051]))
	}

}

func TestLD_nn_A_0xEA(t *testing.T) {
	ch := make(chan bool)
	var gb = GameboyFactory(ch)

	gb.CPU.Registers.PC = 0x0050
	gb.MMU.WriteByte(0x0050, 0xEA)
	gb.MMU.WriteByte(0x0051, 0x34)
	gb.MMU.WriteByte(0x0052, 0x12)
	gb.CPU.Registers.A = 0x66

	gb.Run(1)

	if gb.MMU.Memory[0x1234] != 0x66 {
		t.Errorf("Issue on instruction 0xEA")
	}
}

func TestLD_SP_nn_0x31(t *testing.T) {
	ch := make(chan bool)
	var gb = GameboyFactory(ch)

	gb.CPU.Registers.PC = 0x0050
	gb.MMU.WriteByte(0x0050, 0x31)
	gb.MMU.WriteByte(0x0051, 0x34)
	gb.MMU.WriteByte(0x0052, 0x12)

	gb.Run(1)

	if gb.CPU.Registers.SP != 0x1234 {
		t.Errorf("Issue on instruction 0x31")
	}
}

func TestLD_HL_0x21(t *testing.T) {
	ch := make(chan bool)
	var gb = GameboyFactory(ch)

	gb.CPU.Registers.PC = 0x0050
	gb.MMU.WriteByte(0x0050, 0x21)
	gb.MMU.WriteByte(0x0051, 0x34)
	gb.MMU.WriteByte(0x0052, 0x12)

	gb.Run(1)

	if gb.CPU.Registers.L != 0x34 || gb.CPU.Registers.H != 0x12 {
		t.Errorf("Issue on instruction 0x21")
	}
}

func TestLD_BC_0x01(t *testing.T) {
	ch := make(chan bool)
	var gb = GameboyFactory(ch)

	gb.CPU.Registers.PC = 0x0050
	gb.MMU.WriteByte(0x0050, 0x01)
	gb.MMU.WriteByte(0x0051, 0x34)
	gb.MMU.WriteByte(0x0052, 0x12)

	gb.Run(1)

	if gb.CPU.Registers.C != 0x34 || gb.CPU.Registers.B != 0x12 {
		t.Errorf("Issue on instruction 0x21")
	}
}

func TestCall_nn_0xcd(t *testing.T) {
	ch := make(chan bool)
	var gb = GameboyFactory(ch)

	gb.CPU.Registers.PC = 0x8150
	gb.CPU.Registers.SP = 0xaaaa
	gb.MMU.WriteByte(0x8150, 0xcd)
	gb.MMU.WriteByte(0x8151, 0x34)
	gb.MMU.WriteByte(0x8152, 0x12)

	gb.Run(1)

	if gb.CPU.Registers.PC != 0x1234 || gb.CPU.Registers.SP != 0xaaa8 || gb.MMU.ReadWord(gb.CPU.Registers.SP) != 0x8153 {
		t.Errorf("Issue on instruction 0xcd, %X, %X, %X", gb.CPU.Registers.PC, gb.CPU.Registers.SP, gb.MMU.ReadWord(gb.CPU.Registers.SP))
	}
}

func TestDEC_B_0x05(t *testing.T) {
	ch := make(chan bool)
	var gb = GameboyFactory(ch)

	gb.CPU.Registers.PC = 0x0050
	gb.CPU.Registers.B = 0
	gb.CPU.Registers.F = 0
	gb.MMU.WriteByte(0x0050, 0x05)

	gb.Run(1)

	if gb.CPU.Registers.B != 0xff || gb.CPU.Registers.F&0x20 != 0 {
		t.Errorf("Issue1 on instruction 0x05, %X, %X", gb.CPU.Registers.B, gb.CPU.Registers.F)
	}

	gb.CPU.Registers.PC = 0x0050
	gb.CPU.Registers.B = 1
	gb.CPU.Registers.F = 0
	gb.MMU.WriteByte(0x0050, 0x05)

	gb.Run(1)

	if gb.CPU.Registers.B != 0 || gb.CPU.Registers.F&0x80 == 0 {
		t.Errorf("Issue2 on instruction 0x05, %X, %X", gb.CPU.Registers.B, gb.CPU.Registers.F)
	}

	gb.CPU.Registers.PC = 0x0050
	gb.CPU.Registers.B = 0x05
	gb.CPU.Registers.F = 0
	gb.MMU.WriteByte(0x0050, 0x05)

	gb.Run(1)

	if gb.CPU.Registers.B != 0x04 || gb.CPU.Registers.F&0x80 != 0 || gb.CPU.Registers.F&0x20 == 0 {
		t.Errorf("Issue3 on instruction 0x05, %X, %X", gb.CPU.Registers.B, gb.CPU.Registers.F)
	}
}

func TestPush_Pop(t *testing.T) {
	ch := make(chan bool)
	var gb = GameboyFactory(ch)

	gb.CPU.Registers.PC = 0x0050
	gb.CPU.Registers.SP = 0xef02
	gb.CPU.Registers.B = 0x66
	gb.CPU.Registers.C = 0x44
	gb.MMU.WriteByte(0x0050, 0xc5)
	gb.MMU.WriteByte(0x0051, 0xc1)

	gb.Run(1)

	if gb.CPU.Registers.SP != 0xef00 || gb.MMU.ReadByte(0xef00) != 0x0066 || gb.MMU.ReadByte(0xef01) != 0x0044 {
		t.Errorf("Issue on instruction 0xc5, %X, %X, %X", gb.CPU.Registers.SP, gb.MMU.ReadByte(0xef00), gb.MMU.ReadByte(0xef01))
	}

	gb.CPU.Registers.B = 0
	gb.CPU.Registers.C = 0

	gb.Run(1)

	if gb.CPU.Registers.SP != 0xef02 || gb.CPU.Registers.B != 0x0066 || gb.CPU.Registers.C != 0x0044 {
		t.Errorf("Issue on instruction 0xc1, %X, %X, %X", gb.CPU.Registers.PC, gb.CPU.Registers.SP, gb.MMU.ReadByte(gb.CPU.Registers.SP))
	}
}

func TestSWAP(t *testing.T) {
	ch := make(chan bool)
	var gb = GameboyFactory(ch)

	gb.CPU.Registers.PC = 0x0050
	gb.CPU.Registers.SP = 0xef02
	gb.CPU.Registers.A = 0x3A
	gb.MMU.WriteByte(0x0050, 0xcb)
	gb.MMU.WriteByte(0x0051, 0x37)

	gb.Run(1)

	if gb.CPU.Registers.A != 0xA3 {
		t.Errorf("Issue on instruction 0xcb 0x37")
	}
}

func TestINC_BC_overflow(t *testing.T) {
	ch := make(chan bool)
	var gb = GameboyFactory(ch)

	gb.CPU.Registers.PC = 0x0050
	gb.CPU.Registers.B = 0xff
	gb.CPU.Registers.C = 0xff
	gb.MMU.WriteByte(0x0050, 0x03)

	gb.Run(1)

	if gb.CPU.Registers.B != 0x00 || gb.CPU.Registers.C != 0x00 {
		t.Errorf("Issue on instruction 0x03")
	}
}

func TestINC_B_half_carry(t *testing.T) {
	ch := make(chan bool)
	var gb = GameboyFactory(ch)

	gb.CPU.Registers.PC = 0x0050
	gb.CPU.Registers.B = 0x1f
	gb.MMU.WriteByte(0x0050, 0x04)

	gb.Run(1)

	if gb.CPU.Registers.B != 0x20 || (gb.CPU.Registers.F&0x20) == 0 {
		t.Errorf("Issue on instruction 0x04")
	}
}
