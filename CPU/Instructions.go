package CPU

const Z_BIT uint8 = 0x80
const N_BIT uint8 = 0x40
const H_BIT uint8 = 0x20
const C_BIT uint8 = 0x10

func (cpu *CPU_struct) Innit_Instruction_maps() {
	cpu.Instructions_maps.Instructions_map_uint8 = make(map[uint8]func(uint8))
	cpu.Instructions_maps.Instructions_map_uint16 = make(map[uint8]func(uint16))
	cpu.Instructions_maps.Instructions_map_ = make(map[uint8]func())
	cpu.Instructions_maps.Instructions_map_CB = make(map[uint8]func())

	// 8
	// LD
	cpu.Instructions_maps.Instructions_map_uint8[0x06] = cpu.op_0x06
	cpu.Instructions_maps.Instructions_map_uint8[0x0E] = cpu.op_0x0e
	cpu.Instructions_maps.Instructions_map_uint8[0x16] = cpu.op_0x16
	cpu.Instructions_maps.Instructions_map_uint8[0x1E] = cpu.op_0x1e
	cpu.Instructions_maps.Instructions_map_uint8[0x26] = cpu.op_0x26
	cpu.Instructions_maps.Instructions_map_uint8[0x2E] = cpu.op_0x2e
	cpu.Instructions_maps.Instructions_map_uint8[0x36] = cpu.op_0x36
	cpu.Instructions_maps.Instructions_map_uint8[0x3E] = cpu.op_0x3e
	cpu.Instructions_maps.Instructions_map_uint8[0xe0] = cpu.op_0xe0
	cpu.Instructions_maps.Instructions_map_uint8[0xf0] = cpu.op_0xf0
	// JP
	cpu.Instructions_maps.Instructions_map_uint8[0x18] = cpu.op_0x18
	cpu.Instructions_maps.Instructions_map_uint8[0x20] = cpu.op_0x20
	cpu.Instructions_maps.Instructions_map_uint8[0x28] = cpu.op_0x28
	cpu.Instructions_maps.Instructions_map_uint8[0x30] = cpu.op_0x30
	cpu.Instructions_maps.Instructions_map_uint8[0x38] = cpu.op_0x38
	// CP
	cpu.Instructions_maps.Instructions_map_uint8[0xfe] = cpu.op_0xfe
	// Bitwise
	cpu.Instructions_maps.Instructions_map_uint8[0xc6] = cpu.op_0xc6
	cpu.Instructions_maps.Instructions_map_uint8[0xd6] = cpu.op_0xd6
	cpu.Instructions_maps.Instructions_map_uint8[0xe6] = cpu.op_0xe6
	cpu.Instructions_maps.Instructions_map_uint8[0xf6] = cpu.op_0xf6
	cpu.Instructions_maps.Instructions_map_uint8[0xee] = cpu.op_0xee
	// ADD/SUB
	cpu.Instructions_maps.Instructions_map_uint8[0xce] = cpu.op_0xce
	cpu.Instructions_maps.Instructions_map_uint8[0xde] = cpu.op_0xde
	cpu.Instructions_maps.Instructions_map_uint8[0xe8] = cpu.op_0xe8
	cpu.Instructions_maps.Instructions_map_uint8[0xf8] = cpu.op_0xf8

	// 16
	//LD
	cpu.Instructions_maps.Instructions_map_uint16[0x01] = cpu.op_0x01
	cpu.Instructions_maps.Instructions_map_uint16[0x11] = cpu.op_0x11
	cpu.Instructions_maps.Instructions_map_uint16[0x21] = cpu.op_0x21
	cpu.Instructions_maps.Instructions_map_uint16[0x31] = cpu.op_0x31
	cpu.Instructions_maps.Instructions_map_uint16[0xea] = cpu.op_0xea
	cpu.Instructions_maps.Instructions_map_uint16[0xfa] = cpu.op_0xfa
	//JP
	cpu.Instructions_maps.Instructions_map_uint16[0x08] = cpu.op_0x08
	cpu.Instructions_maps.Instructions_map_uint16[0xc2] = cpu.op_0xc2
	cpu.Instructions_maps.Instructions_map_uint16[0xc3] = cpu.op_0xc3
	cpu.Instructions_maps.Instructions_map_uint16[0xca] = cpu.op_0xca
	cpu.Instructions_maps.Instructions_map_uint16[0xd2] = cpu.op_0xd2
	cpu.Instructions_maps.Instructions_map_uint16[0xda] = cpu.op_0xda
	//CALL
	cpu.Instructions_maps.Instructions_map_uint16[0xc4] = cpu.op_0xc4
	cpu.Instructions_maps.Instructions_map_uint16[0xcc] = cpu.op_0xcc
	cpu.Instructions_maps.Instructions_map_uint16[0xcd] = cpu.op_0xcd
	cpu.Instructions_maps.Instructions_map_uint16[0xd4] = cpu.op_0xd4
	cpu.Instructions_maps.Instructions_map_uint16[0xdc] = cpu.op_0xdc

	// 0
	//LD
	cpu.Instructions_maps.Instructions_map_[0x02] = cpu.op_0x02
	cpu.Instructions_maps.Instructions_map_[0x0a] = cpu.op_0x0a
	cpu.Instructions_maps.Instructions_map_[0x12] = cpu.op_0x12
	cpu.Instructions_maps.Instructions_map_[0x1a] = cpu.op_0x1a
	cpu.Instructions_maps.Instructions_map_[0x22] = cpu.op_0x22
	cpu.Instructions_maps.Instructions_map_[0x2a] = cpu.op_0x2a
	cpu.Instructions_maps.Instructions_map_[0x32] = cpu.op_0x32
	cpu.Instructions_maps.Instructions_map_[0x3a] = cpu.op_0x3a
	cpu.Instructions_maps.Instructions_map_[0x40] = cpu.op_0x40
	cpu.Instructions_maps.Instructions_map_[0x41] = cpu.op_0x41
	cpu.Instructions_maps.Instructions_map_[0x42] = cpu.op_0x42
	cpu.Instructions_maps.Instructions_map_[0x43] = cpu.op_0x43
	cpu.Instructions_maps.Instructions_map_[0x44] = cpu.op_0x44
	cpu.Instructions_maps.Instructions_map_[0x45] = cpu.op_0x45
	cpu.Instructions_maps.Instructions_map_[0x46] = cpu.op_0x46
	cpu.Instructions_maps.Instructions_map_[0x47] = cpu.op_0x47
	cpu.Instructions_maps.Instructions_map_[0x48] = cpu.op_0x48
	cpu.Instructions_maps.Instructions_map_[0x49] = cpu.op_0x49
	cpu.Instructions_maps.Instructions_map_[0x4a] = cpu.op_0x4a
	cpu.Instructions_maps.Instructions_map_[0x4b] = cpu.op_0x4b
	cpu.Instructions_maps.Instructions_map_[0x4c] = cpu.op_0x4c
	cpu.Instructions_maps.Instructions_map_[0x4d] = cpu.op_0x4d
	cpu.Instructions_maps.Instructions_map_[0x4e] = cpu.op_0x4e
	cpu.Instructions_maps.Instructions_map_[0x4f] = cpu.op_0x4f
	cpu.Instructions_maps.Instructions_map_[0x50] = cpu.op_0x50
	cpu.Instructions_maps.Instructions_map_[0x51] = cpu.op_0x51
	cpu.Instructions_maps.Instructions_map_[0x52] = cpu.op_0x52
	cpu.Instructions_maps.Instructions_map_[0x53] = cpu.op_0x53
	cpu.Instructions_maps.Instructions_map_[0x54] = cpu.op_0x54
	cpu.Instructions_maps.Instructions_map_[0x55] = cpu.op_0x55
	cpu.Instructions_maps.Instructions_map_[0x56] = cpu.op_0x56
	cpu.Instructions_maps.Instructions_map_[0x57] = cpu.op_0x57
	cpu.Instructions_maps.Instructions_map_[0x58] = cpu.op_0x58
	cpu.Instructions_maps.Instructions_map_[0x59] = cpu.op_0x59
	cpu.Instructions_maps.Instructions_map_[0x5a] = cpu.op_0x5a
	cpu.Instructions_maps.Instructions_map_[0x5b] = cpu.op_0x5b
	cpu.Instructions_maps.Instructions_map_[0x5c] = cpu.op_0x5c
	cpu.Instructions_maps.Instructions_map_[0x5d] = cpu.op_0x5d
	cpu.Instructions_maps.Instructions_map_[0x5e] = cpu.op_0x5e
	cpu.Instructions_maps.Instructions_map_[0x5f] = cpu.op_0x5f
	cpu.Instructions_maps.Instructions_map_[0x60] = cpu.op_0x60
	cpu.Instructions_maps.Instructions_map_[0x61] = cpu.op_0x61
	cpu.Instructions_maps.Instructions_map_[0x62] = cpu.op_0x62
	cpu.Instructions_maps.Instructions_map_[0x63] = cpu.op_0x63
	cpu.Instructions_maps.Instructions_map_[0x64] = cpu.op_0x64
	cpu.Instructions_maps.Instructions_map_[0x65] = cpu.op_0x65
	cpu.Instructions_maps.Instructions_map_[0x66] = cpu.op_0x66
	cpu.Instructions_maps.Instructions_map_[0x67] = cpu.op_0x67
	cpu.Instructions_maps.Instructions_map_[0x68] = cpu.op_0x68
	cpu.Instructions_maps.Instructions_map_[0x69] = cpu.op_0x69
	cpu.Instructions_maps.Instructions_map_[0x6a] = cpu.op_0x6a
	cpu.Instructions_maps.Instructions_map_[0x6b] = cpu.op_0x6b
	cpu.Instructions_maps.Instructions_map_[0x6c] = cpu.op_0x6c
	cpu.Instructions_maps.Instructions_map_[0x6d] = cpu.op_0x6d
	cpu.Instructions_maps.Instructions_map_[0x6e] = cpu.op_0x6e
	cpu.Instructions_maps.Instructions_map_[0x6f] = cpu.op_0x6f
	cpu.Instructions_maps.Instructions_map_[0x70] = cpu.op_0x70
	cpu.Instructions_maps.Instructions_map_[0x71] = cpu.op_0x71
	cpu.Instructions_maps.Instructions_map_[0x72] = cpu.op_0x72
	cpu.Instructions_maps.Instructions_map_[0x73] = cpu.op_0x73
	cpu.Instructions_maps.Instructions_map_[0x74] = cpu.op_0x74
	cpu.Instructions_maps.Instructions_map_[0x75] = cpu.op_0x75
	cpu.Instructions_maps.Instructions_map_[0x76] = cpu.op_0x76
	cpu.Instructions_maps.Instructions_map_[0x77] = cpu.op_0x77
	cpu.Instructions_maps.Instructions_map_[0x78] = cpu.op_0x78
	cpu.Instructions_maps.Instructions_map_[0x79] = cpu.op_0x79
	cpu.Instructions_maps.Instructions_map_[0x7a] = cpu.op_0x7a
	cpu.Instructions_maps.Instructions_map_[0x7b] = cpu.op_0x7b
	cpu.Instructions_maps.Instructions_map_[0x7c] = cpu.op_0x7c
	cpu.Instructions_maps.Instructions_map_[0x7d] = cpu.op_0x7d
	cpu.Instructions_maps.Instructions_map_[0x7e] = cpu.op_0x7e
	cpu.Instructions_maps.Instructions_map_[0x7f] = cpu.op_0x7f
	cpu.Instructions_maps.Instructions_map_[0xe2] = cpu.op_0xe2
	cpu.Instructions_maps.Instructions_map_[0xf2] = cpu.op_0xf2
	cpu.Instructions_maps.Instructions_map_[0xf9] = cpu.op_0xf9
	// Bitwise
	cpu.Instructions_maps.Instructions_map_[0x0f] = cpu.op_0x0f
	cpu.Instructions_maps.Instructions_map_[0x1f] = cpu.op_0x1f
	cpu.Instructions_maps.Instructions_map_[0x37] = cpu.op_0x37
	cpu.Instructions_maps.Instructions_map_[0x3f] = cpu.op_0x3f
	cpu.Instructions_maps.Instructions_map_[0xa0] = cpu.op_0xa0
	cpu.Instructions_maps.Instructions_map_[0xa1] = cpu.op_0xa1
	cpu.Instructions_maps.Instructions_map_[0xa2] = cpu.op_0xa2
	cpu.Instructions_maps.Instructions_map_[0xa3] = cpu.op_0xa3
	cpu.Instructions_maps.Instructions_map_[0xa4] = cpu.op_0xa4
	cpu.Instructions_maps.Instructions_map_[0xa5] = cpu.op_0xa5
	cpu.Instructions_maps.Instructions_map_[0xa6] = cpu.op_0xa6
	cpu.Instructions_maps.Instructions_map_[0xa7] = cpu.op_0xa7
	cpu.Instructions_maps.Instructions_map_[0xa8] = cpu.op_0xa8
	cpu.Instructions_maps.Instructions_map_[0xa9] = cpu.op_0xa9
	cpu.Instructions_maps.Instructions_map_[0xaa] = cpu.op_0xaa
	cpu.Instructions_maps.Instructions_map_[0xab] = cpu.op_0xab
	cpu.Instructions_maps.Instructions_map_[0xac] = cpu.op_0xac
	cpu.Instructions_maps.Instructions_map_[0xad] = cpu.op_0xad
	cpu.Instructions_maps.Instructions_map_[0xae] = cpu.op_0xae
	cpu.Instructions_maps.Instructions_map_[0xaf] = cpu.op_0xaf
	cpu.Instructions_maps.Instructions_map_[0xb0] = cpu.op_0xb0
	cpu.Instructions_maps.Instructions_map_[0xb1] = cpu.op_0xb1
	cpu.Instructions_maps.Instructions_map_[0xb2] = cpu.op_0xb2
	cpu.Instructions_maps.Instructions_map_[0xb3] = cpu.op_0xb3
	cpu.Instructions_maps.Instructions_map_[0xb4] = cpu.op_0xb4
	cpu.Instructions_maps.Instructions_map_[0xb5] = cpu.op_0xb5
	cpu.Instructions_maps.Instructions_map_[0xb6] = cpu.op_0xb6
	cpu.Instructions_maps.Instructions_map_[0xb7] = cpu.op_0xb7
	// CP
	cpu.Instructions_maps.Instructions_map_[0x2f] = cpu.op_0x2f
	cpu.Instructions_maps.Instructions_map_[0xb7] = cpu.op_0xb7
	cpu.Instructions_maps.Instructions_map_[0xb8] = cpu.op_0xb8
	cpu.Instructions_maps.Instructions_map_[0xb9] = cpu.op_0xb9
	cpu.Instructions_maps.Instructions_map_[0xba] = cpu.op_0xba
	cpu.Instructions_maps.Instructions_map_[0xbb] = cpu.op_0xbb
	cpu.Instructions_maps.Instructions_map_[0xbc] = cpu.op_0xbc
	cpu.Instructions_maps.Instructions_map_[0xbd] = cpu.op_0xbd
	cpu.Instructions_maps.Instructions_map_[0xbe] = cpu.op_0xbe
	cpu.Instructions_maps.Instructions_map_[0xbf] = cpu.op_0xbf
	// DEC
	cpu.Instructions_maps.Instructions_map_[0x05] = cpu.op_0x05
	cpu.Instructions_maps.Instructions_map_[0x0b] = cpu.op_0x0b
	cpu.Instructions_maps.Instructions_map_[0x0d] = cpu.op_0x0d
	cpu.Instructions_maps.Instructions_map_[0x15] = cpu.op_0x15
	cpu.Instructions_maps.Instructions_map_[0x1b] = cpu.op_0x1b
	cpu.Instructions_maps.Instructions_map_[0x1d] = cpu.op_0x1d
	cpu.Instructions_maps.Instructions_map_[0x25] = cpu.op_0x25
	cpu.Instructions_maps.Instructions_map_[0x2b] = cpu.op_0x2b
	cpu.Instructions_maps.Instructions_map_[0x2d] = cpu.op_0x2d
	cpu.Instructions_maps.Instructions_map_[0x35] = cpu.op_0x35
	cpu.Instructions_maps.Instructions_map_[0x3b] = cpu.op_0x3b
	cpu.Instructions_maps.Instructions_map_[0x3d] = cpu.op_0x3d
	// INC
	cpu.Instructions_maps.Instructions_map_[0x03] = cpu.op_0x03
	cpu.Instructions_maps.Instructions_map_[0x04] = cpu.op_0x04
	cpu.Instructions_maps.Instructions_map_[0x0c] = cpu.op_0x0c
	cpu.Instructions_maps.Instructions_map_[0x13] = cpu.op_0x13
	cpu.Instructions_maps.Instructions_map_[0x14] = cpu.op_0x14
	cpu.Instructions_maps.Instructions_map_[0x1c] = cpu.op_0x1c
	cpu.Instructions_maps.Instructions_map_[0x23] = cpu.op_0x23
	cpu.Instructions_maps.Instructions_map_[0x24] = cpu.op_0x24
	cpu.Instructions_maps.Instructions_map_[0x2c] = cpu.op_0x2c
	cpu.Instructions_maps.Instructions_map_[0x33] = cpu.op_0x33
	cpu.Instructions_maps.Instructions_map_[0x34] = cpu.op_0x34
	cpu.Instructions_maps.Instructions_map_[0x3c] = cpu.op_0x3c
	// SUB
	cpu.Instructions_maps.Instructions_map_[0x90] = cpu.op_0x90
	cpu.Instructions_maps.Instructions_map_[0x95] = cpu.op_0x95
	cpu.Instructions_maps.Instructions_map_[0x97] = cpu.op_0x97
	// NOP
	cpu.Instructions_maps.Instructions_map_[0x00] = cpu.op_0x00
	// HALT
	cpu.Instructions_maps.Instructions_map_[0x76] = cpu.op_0x76
	// PUSH/POP
	cpu.Instructions_maps.Instructions_map_[0xc1] = cpu.op_0xc1
	cpu.Instructions_maps.Instructions_map_[0xc5] = cpu.op_0xc5
	cpu.Instructions_maps.Instructions_map_[0xd1] = cpu.op_0xd1
	cpu.Instructions_maps.Instructions_map_[0xd5] = cpu.op_0xd5
	cpu.Instructions_maps.Instructions_map_[0xe1] = cpu.op_0xe1
	cpu.Instructions_maps.Instructions_map_[0xe5] = cpu.op_0xe5
	cpu.Instructions_maps.Instructions_map_[0xf1] = cpu.op_0xf1
	cpu.Instructions_maps.Instructions_map_[0xf5] = cpu.op_0xf5
	// RL
	cpu.Instructions_maps.Instructions_map_[0x07] = cpu.op_0x07
	cpu.Instructions_maps.Instructions_map_[0x17] = cpu.op_0x17
	// RET
	cpu.Instructions_maps.Instructions_map_[0xc0] = cpu.op_0xc0
	cpu.Instructions_maps.Instructions_map_[0xc8] = cpu.op_0xc8
	cpu.Instructions_maps.Instructions_map_[0xc9] = cpu.op_0xc9
	cpu.Instructions_maps.Instructions_map_[0xd0] = cpu.op_0xd0
	cpu.Instructions_maps.Instructions_map_[0xd8] = cpu.op_0xd8
	cpu.Instructions_maps.Instructions_map_[0xd9] = cpu.op_0xd9
	// ADD/SUB
	cpu.Instructions_maps.Instructions_map_[0x09] = cpu.op_0x09
	cpu.Instructions_maps.Instructions_map_[0x19] = cpu.op_0x19
	cpu.Instructions_maps.Instructions_map_[0x29] = cpu.op_0x29
	cpu.Instructions_maps.Instructions_map_[0x39] = cpu.op_0x39
	cpu.Instructions_maps.Instructions_map_[0x80] = cpu.op_0x80
	cpu.Instructions_maps.Instructions_map_[0x81] = cpu.op_0x81
	cpu.Instructions_maps.Instructions_map_[0x82] = cpu.op_0x82
	cpu.Instructions_maps.Instructions_map_[0x83] = cpu.op_0x83
	cpu.Instructions_maps.Instructions_map_[0x84] = cpu.op_0x84
	cpu.Instructions_maps.Instructions_map_[0x85] = cpu.op_0x85
	cpu.Instructions_maps.Instructions_map_[0x86] = cpu.op_0x86
	cpu.Instructions_maps.Instructions_map_[0x87] = cpu.op_0x87
	cpu.Instructions_maps.Instructions_map_[0x88] = cpu.op_0x88
	cpu.Instructions_maps.Instructions_map_[0x89] = cpu.op_0x89
	cpu.Instructions_maps.Instructions_map_[0x8a] = cpu.op_0x8a
	cpu.Instructions_maps.Instructions_map_[0x8b] = cpu.op_0x8b
	cpu.Instructions_maps.Instructions_map_[0x8c] = cpu.op_0x8c
	cpu.Instructions_maps.Instructions_map_[0x8d] = cpu.op_0x8d
	cpu.Instructions_maps.Instructions_map_[0x8e] = cpu.op_0x8e
	cpu.Instructions_maps.Instructions_map_[0x8f] = cpu.op_0x8f
	cpu.Instructions_maps.Instructions_map_[0x91] = cpu.op_0x91
	cpu.Instructions_maps.Instructions_map_[0x92] = cpu.op_0x92
	cpu.Instructions_maps.Instructions_map_[0x93] = cpu.op_0x93
	cpu.Instructions_maps.Instructions_map_[0x94] = cpu.op_0x94
	cpu.Instructions_maps.Instructions_map_[0x95] = cpu.op_0x95
	cpu.Instructions_maps.Instructions_map_[0x96] = cpu.op_0x96
	cpu.Instructions_maps.Instructions_map_[0x97] = cpu.op_0x97
	cpu.Instructions_maps.Instructions_map_[0x98] = cpu.op_0x98
	cpu.Instructions_maps.Instructions_map_[0x99] = cpu.op_0x99
	cpu.Instructions_maps.Instructions_map_[0x9a] = cpu.op_0x9a
	cpu.Instructions_maps.Instructions_map_[0x9b] = cpu.op_0x9b
	cpu.Instructions_maps.Instructions_map_[0x9c] = cpu.op_0x9c
	cpu.Instructions_maps.Instructions_map_[0x9d] = cpu.op_0x9d
	cpu.Instructions_maps.Instructions_map_[0x9e] = cpu.op_0x9e
	cpu.Instructions_maps.Instructions_map_[0x9f] = cpu.op_0x9f
	// DI/EI
	cpu.Instructions_maps.Instructions_map_[0xf3] = cpu.op_0xf3
	cpu.Instructions_maps.Instructions_map_[0xfb] = cpu.op_0xfb
	// JP
	cpu.Instructions_maps.Instructions_map_[0xe9] = cpu.op_0xe9
	// RST
	cpu.Instructions_maps.Instructions_map_[0xc7] = cpu.op_0xc7
	cpu.Instructions_maps.Instructions_map_[0xcf] = cpu.op_0xcf
	cpu.Instructions_maps.Instructions_map_[0xd7] = cpu.op_0xd7
	cpu.Instructions_maps.Instructions_map_[0xdf] = cpu.op_0xdf
	cpu.Instructions_maps.Instructions_map_[0xe7] = cpu.op_0xe7
	cpu.Instructions_maps.Instructions_map_[0xef] = cpu.op_0xef
	cpu.Instructions_maps.Instructions_map_[0xf7] = cpu.op_0xf7
	cpu.Instructions_maps.Instructions_map_[0xff] = cpu.op_0xff
	//DAA
	cpu.Instructions_maps.Instructions_map_[0x27] = cpu.op_0x27

	// CB
	// cpu.Instructions_maps.Instructions_map_CB[0x01] = cpu.op_cb_0x01
	// cpu.Instructions_maps.Instructions_map_CB[0x02] = cpu.op_cb_0x02
	// cpu.Instructions_maps.Instructions_map_CB[0x03] = cpu.op_cb_0x03
	cpu.Instructions_maps.Instructions_map_CB[0x07] = cpu.op_cb_0x07
	cpu.Instructions_maps.Instructions_map_CB[0x11] = cpu.op_cb_0x11
	cpu.Instructions_maps.Instructions_map_CB[0x37] = cpu.op_cb_0x37
	cpu.Instructions_maps.Instructions_map_CB[0x38] = cpu.op_cb_0x38
	cpu.Instructions_maps.Instructions_map_CB[0x47] = cpu.op_cb_0x47
	cpu.Instructions_maps.Instructions_map_CB[0x6f] = cpu.op_cb_0x6f
	cpu.Instructions_maps.Instructions_map_CB[0x7c] = cpu.op_cb_0x7c
	cpu.Instructions_maps.Instructions_map_CB[0x77] = cpu.op_cb_0x77
	cpu.Instructions_maps.Instructions_map_CB[0x87] = cpu.op_cb_0x87
	cpu.Instructions_maps.Instructions_map_CB[0xcf] = cpu.op_cb_0xcf

}

func (cpu *CPU_struct) getAF() uint16 {
	return (uint16(cpu.Registers.A) << 8) | uint16(cpu.Registers.F)
}

func (cpu *CPU_struct) getBC() uint16 {
	return (uint16(cpu.Registers.B) << 8) | uint16(cpu.Registers.C)
}

func (cpu *CPU_struct) getDE() uint16 {
	return (uint16(cpu.Registers.D) << 8) | uint16(cpu.Registers.E)
}

func (cpu *CPU_struct) getHL() uint16 {
	return (uint16(cpu.Registers.H) << 8) | uint16(cpu.Registers.L)
}

func (cpu *CPU_struct) setAF(nn uint16) {
	cpu.Registers.A = uint8(nn >> 8)
	cpu.Registers.F = uint8(nn & 0xff)
}

func (cpu *CPU_struct) setBC(nn uint16) {
	cpu.Registers.B = uint8(nn >> 8)
	cpu.Registers.C = uint8(nn & 0xff)
}

func (cpu *CPU_struct) setDE(nn uint16) {
	cpu.Registers.D = uint8(nn >> 8)
	cpu.Registers.E = uint8(nn & 0xff)
}

func (cpu *CPU_struct) setHL(nn uint16) {
	cpu.Registers.H = uint8(nn >> 8)
	cpu.Registers.L = uint8(nn & 0xff)
}

func (cpu *CPU_struct) setZ(n uint8) {
	cpu.Registers.F &= (Z_BIT ^ 0xff)
	if n == 0 {
		cpu.Registers.F |= Z_BIT
	}
}

// 0x00 NOP
func (cpu *CPU_struct) op_0x00() {
	cpu.Cycle += 4
}

// 0x10 - STOP
func (cpu *CPU_struct) op_0x10() {
	panic("Not implemented")
}

// 0x20 - JP NZ, *
func (cpu *CPU_struct) op_0x20(n uint8) {
	if (cpu.Registers.F & Z_BIT) == 0 {
		cpu.Registers.PC = uint16(int16(cpu.Registers.PC) + int16(int8(n)))
		cpu.Cycle += 12
		return
	}

	// TODO: double check cpu.cycles
	cpu.Cycle += 8
}

// 0x30 - JP NC, *
func (cpu *CPU_struct) op_0x30(n uint8) {
	if (cpu.Registers.F & C_BIT) == 0 {
		cpu.Registers.PC = uint16(int16(cpu.Registers.PC) + int16(int8(n)))
		cpu.Cycle += 12
		return
	}

	// TODO: double check cpu.cycles
	cpu.Cycle += 8
}

// 0x01 LD BC, nn
func (cpu *CPU_struct) op_0x01(nn uint16) {
	cpu.setBC(nn)

	cpu.Cycle += 12
}

// 0x11 LD DE,nn
func (cpu *CPU_struct) op_0x11(nn uint16) {
	cpu.setDE(nn)

	cpu.Cycle += 12
}

// 0x21 - LD HL,nn
func (cpu *CPU_struct) op_0x21(nn uint16) {
	cpu.setHL(nn)

	cpu.Cycle += 12
}

// 0x31 - LD SP,nn
func (cpu *CPU_struct) op_0x31(nn uint16) {
	cpu.Registers.SP = nn

	cpu.Cycle += 12
}

// 0x02 LD (BC), A
func (cpu *CPU_struct) op_0x02() {
	var address uint16 = cpu.getBC()
	cpu.MMU.WriteByte(address, cpu.Registers.A)

	cpu.Cycle += 8
}

// 0x12 LD (DE),A
func (cpu *CPU_struct) op_0x12() {
	var address uint16 = cpu.getDE()
	cpu.MMU.WriteByte(address, cpu.Registers.A)

	cpu.Cycle += 8
}

// 0x22 LD (HL+),A
func (cpu *CPU_struct) op_0x22() {
	var address uint16 = cpu.getHL()
	cpu.MMU.WriteByte(address, cpu.Registers.A)
	address += 1
	cpu.setHL(address)

	cpu.Cycle += 8
}

// 0x32 LD (HL-),A
func (cpu *CPU_struct) op_0x32() {
	var address uint16 = cpu.getHL()
	cpu.MMU.WriteByte(address, cpu.Registers.A)
	address -= 1
	cpu.setHL(address)

	cpu.Cycle += 8
}

// 0x03 INC BC
func (cpu *CPU_struct) op_0x03() {
	var nn uint16 = cpu.getBC()
	nn += 1
	cpu.setBC(nn)

	cpu.Cycle += 8
}

// 0x13 INC DE
func (cpu *CPU_struct) op_0x13() {
	var nn uint16 = cpu.getDE()
	nn += 1
	cpu.setDE(nn)

	cpu.Cycle += 8
}

// 0x23 INC HL
func (cpu *CPU_struct) op_0x23() {
	var nn uint16 = cpu.getHL()
	nn += 1
	cpu.setHL(nn)

	cpu.Cycle += 8
}

// 0x33 INC SP
func (cpu *CPU_struct) op_0x33() {
	cpu.Registers.SP++

	cpu.Cycle += 8
}

//0x04 INC B
func (cpu *CPU_struct) op_0x04() {
	var n uint8 = cpu.Registers.B
	cpu.Registers.B++
	cpu.Registers.F &= C_BIT

	cpu.setZ(n)

	if (n & 0x0f) == 0x0f {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 4
}

//0x14 INC D
func (cpu *CPU_struct) op_0x14() {
	var n uint8 = cpu.Registers.D
	cpu.Registers.D++
	cpu.Registers.F &= C_BIT

	cpu.setZ(n)

	if (n & 0x0f) == 0x0f {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 4
}

//0x24 INC H
func (cpu *CPU_struct) op_0x24() {
	var n uint8 = cpu.Registers.H
	cpu.Registers.H++
	cpu.Registers.F &= C_BIT

	cpu.setZ(n)

	if (n & 0x0f) == 0x0f {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 4
}

//0x34 INC (HL)
func (cpu *CPU_struct) op_0x34() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	n++
	cpu.MMU.WriteByte(address, n)

	cpu.Registers.F &= C_BIT

	cpu.setZ(n)

	if ((n - 1) & 0x0f) == 0x0f {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 12
}

// 0x05 DEC B
func (cpu *CPU_struct) op_0x05() {
	cpu.Registers.B -= 1
	var r uint8 = cpu.Registers.B
	cpu.Registers.F &= C_BIT
	cpu.Registers.F |= N_BIT

	cpu.setZ(r)

	if ((r + 1) & 0x0f) != 0x00 {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 4
}

// 0x15 DEC D
func (cpu *CPU_struct) op_0x15() {
	cpu.Registers.D -= 1
	var r uint8 = cpu.Registers.D
	cpu.Registers.F &= C_BIT
	cpu.Registers.F |= N_BIT

	cpu.setZ(r)

	if ((r + 1) & 0x0f) != 0x00 {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 4
}

// 0x25 DEC H
func (cpu *CPU_struct) op_0x25() {
	cpu.Registers.H -= 1
	var r uint8 = cpu.Registers.H
	cpu.Registers.F &= C_BIT
	cpu.Registers.F |= N_BIT

	cpu.setZ(r)

	if ((r + 1) & 0x0f) != 0x00 {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 4
}

// 0x35 DEC (HL)
func (cpu *CPU_struct) op_0x35() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	n--
	cpu.MMU.WriteByte(address, n)

	cpu.Registers.F &= C_BIT
	cpu.Registers.F |= N_BIT

	cpu.setZ(n)

	if ((n + 1) & 0x0f) != 0x00 {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 12
}

// 0x06 - LD B,n
func (cpu *CPU_struct) op_0x06(n uint8) {
	cpu.Registers.B = n

	cpu.Cycle += 8
}

// 0x16 - LD D,n
func (cpu *CPU_struct) op_0x16(n uint8) {
	cpu.Registers.D = n

	cpu.Cycle += 8
}

// 0x26 - LD H,n
func (cpu *CPU_struct) op_0x26(n uint8) {
	cpu.Registers.H = n

	cpu.Cycle += 8
}

// 0x36 - LD (HL),n
func (cpu *CPU_struct) op_0x36(n uint8) {
	var address uint16 = cpu.getHL()
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 12
}

func (cpu *CPU_struct) rla() {
	var carry uint8 = cpu.Registers.F & C_BIT
	cpu.Registers.F = 0
	// bit 7 from A to carry
	cpu.Registers.F |= ((cpu.Registers.A & 0x80) >> 3)

	cpu.Registers.A = (cpu.Registers.A << 1)
	if carry != 0 {
		cpu.Registers.A |= 0x01
	}

	cpu.setZ(cpu.Registers.A)

	cpu.Cycle += 8
}

// 0x07 RLA
func (cpu *CPU_struct) op_0x07() {
	cpu.rla()
}

// 0x17 RLA
func (cpu *CPU_struct) op_0x17() {
	cpu.rla()
}

// 0x27 DAA
// TODO: double check
func (cpu *CPU_struct) op_0x27() {
	var n_flag uint8 = cpu.Registers.F & N_BIT
	var c_flag uint8 = cpu.Registers.F & C_BIT
	var h_flag uint8 = cpu.Registers.F & H_BIT

	// note: assumes a is a uint8_t and wraps from 0xff to 0
	if n_flag == 0 { // after an addition, adjust if (half-)carry occurred or if result is out of bounds
		if (c_flag != 0) || (cpu.Registers.A > 0x99) {
			cpu.Registers.A += 0x60
			c_flag = 1
		}
		if (h_flag != 0) || ((cpu.Registers.A & 0x0f) > 0x09) {
			cpu.Registers.A += 0x6
		}
	} else { // after a subtraction, only adjust if (half-)carry occurred
		if c_flag != 0 {
			cpu.Registers.A -= 0x60
		}
		if h_flag != 0 {
			cpu.Registers.A -= 0x6
		}
	}
	// these flags are always updated
	cpu.Registers.F &= N_BIT
	cpu.setZ(cpu.Registers.A)

	cpu.Cycle += 4
}

// 0x37 SCF
func (cpu *CPU_struct) op_0x37() {
	cpu.Registers.F &= Z_BIT
	cpu.Registers.F |= C_BIT

	cpu.Cycle += 4
}

// 0x08 - LD (nn), SP
func (cpu *CPU_struct) op_0x08(nn uint16) {
	cpu.MMU.WriteWord(nn, cpu.Registers.SP)

	cpu.Cycle += 20
}

func (cpu *CPU_struct) jr(n uint8) {
	cpu.Registers.PC = uint16(int16(cpu.Registers.PC) + int16(int8(n)))
}

// 0x18 - JR n
func (cpu *CPU_struct) op_0x18(n uint8) {
	cpu.jr(n)

	cpu.Cycle += 8
}

// 0x28 - JP Z, *
func (cpu *CPU_struct) op_0x28(n uint8) {
	if (cpu.Registers.F & Z_BIT) != 0 {
		cpu.jr(n)
	}

	cpu.Cycle += 8
}

// 0x38 - JP C, nn
func (cpu *CPU_struct) op_0x38(n uint8) {
	if (cpu.Registers.F & C_BIT) != 0 {
		cpu.jr(n)
	}
}

// 0x09 ADD HL, BC
func (cpu *CPU_struct) op_0x09() {
	var hl uint16 = cpu.getHL()
	var bc uint16 = cpu.getBC()
	cpu.setHL(hl + bc)

	cpu.Registers.F &= Z_BIT

	if (((hl & 0x0fff) + (bc & 0x0fff)) & 0xf000) != 0 {
		cpu.Registers.F |= H_BIT
	}

	if ((uint32(hl) + uint32(bc)) & 0xf0000) != 0 {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 8
}

// 0x19 ADD HL, DE
func (cpu *CPU_struct) op_0x19() {
	var hl uint16 = cpu.getHL()
	var de uint16 = cpu.getDE()
	cpu.setHL(hl + de)

	cpu.Registers.F &= Z_BIT

	if (((hl & 0x0fff) + (de & 0x0fff)) & 0xf000) != 0 {
		cpu.Registers.F |= H_BIT
	}

	if ((uint32(hl) + uint32(de)) & 0xf0000) != 0 {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 8
}

// 0x29 ADD HL, HL
func (cpu *CPU_struct) op_0x29() {
	var hl uint16 = cpu.getHL()
	cpu.setHL(hl + hl)

	cpu.Registers.F &= Z_BIT

	if (((hl & 0x0fff) + (hl & 0x0fff)) & 0xf000) != 0 {
		cpu.Registers.F |= H_BIT
	}

	if ((uint32(hl) + uint32(hl)) & 0xf0000) != 0 {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 8
}

// 0x39 ADD HL, SP
func (cpu *CPU_struct) op_0x39() {
	var hl uint16 = cpu.getHL()
	cpu.setHL(hl + cpu.Registers.SP)

	cpu.Registers.F &= Z_BIT

	if (((hl & 0x0fff) + (cpu.Registers.SP & 0x0fff)) & 0xf000) != 0 {
		cpu.Registers.F |= H_BIT
	}

	if ((uint32(hl) + uint32(cpu.Registers.SP)) & 0xf0000) != 0 {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 8
}

//  0x0a LD A,(BC)
func (cpu *CPU_struct) op_0x0a() {
	var address uint16 = cpu.getBC()
	cpu.Registers.A = cpu.MMU.ReadByte(address)

	cpu.Cycle += 8
}

//  0x1a LD A,(DE)
func (cpu *CPU_struct) op_0x1a() {
	var address uint16 = cpu.getDE()
	cpu.Registers.A = cpu.MMU.ReadByte(address)

	cpu.Cycle += 8
}

// 0x2A LDI A, (HL)
func (cpu *CPU_struct) op_0x2a() {
	var address uint16 = cpu.getHL()
	cpu.Registers.A = cpu.MMU.ReadByte(address)

	// INC HL
	cpu.op_0x23()

	// cpu.Cycle += 8 // alreay in 0x23
}

// 0x3A LDD A, (HL)
func (cpu *CPU_struct) op_0x3a() {
	var address uint16 = cpu.getHL()
	cpu.Registers.A = cpu.MMU.ReadByte(address)

	// DEC HL
	cpu.op_0x2b()

	// cpu.Cycle += 8 // alreay in 0x2B
}

// 0x0B DEC BC
func (cpu *CPU_struct) op_0x0b() {
	var nn uint16 = cpu.getBC()
	nn -= 1
	cpu.setBC(nn)

	cpu.Cycle += 8
}

// 0x1B DEC DE
func (cpu *CPU_struct) op_0x1b() {
	var nn uint16 = cpu.getDE()
	nn -= 1
	cpu.setDE(nn)

	cpu.Cycle += 8
}

// 0x2B DEC HL
func (cpu *CPU_struct) op_0x2b() {
	var nn uint16 = cpu.getHL()
	nn -= 1
	cpu.setHL(nn)

	cpu.Cycle += 8
}

// 0x3B DEC SP
func (cpu *CPU_struct) op_0x3b() {
	cpu.Registers.SP -= 1

	cpu.Cycle += 8
}

//0x0c INC C
func (cpu *CPU_struct) op_0x0c() {
	cpu.Registers.C += 1
	cpu.Registers.F &= C_BIT

	cpu.setZ(cpu.Registers.C)

	if ((cpu.Registers.C - 1) & 0x0f) == 0x0f {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 4
}

//0x1c INC E
func (cpu *CPU_struct) op_0x1c() {
	cpu.Registers.E += 1
	cpu.Registers.F &= C_BIT

	cpu.setZ(cpu.Registers.E)

	if ((cpu.Registers.E - 1) & 0x0f) == 0x0f {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 4
}

//0x2c INC L
func (cpu *CPU_struct) op_0x2c() {
	cpu.Registers.L += 1
	cpu.Registers.F &= C_BIT

	cpu.setZ(cpu.Registers.L)

	if ((cpu.Registers.L - 1) & 0x0f) == 0x0f {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 4
}

//0x3c INC A
func (cpu *CPU_struct) op_0x3c() {
	cpu.Registers.A += 1
	cpu.Registers.F &= C_BIT

	cpu.setZ(cpu.Registers.A)

	if ((cpu.Registers.A - 1) & 0x0f) == 0x0f {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 4
}

// 0x0d DEC C
func (cpu *CPU_struct) op_0x0d() {
	cpu.Registers.C -= 1
	cpu.Registers.F &= C_BIT
	cpu.Registers.F |= N_BIT

	cpu.setZ(cpu.Registers.C)

	if ((cpu.Registers.C + 1) & 0xf) != 0 {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 4
}

// 0x1D DEC E
func (cpu *CPU_struct) op_0x1d() {
	cpu.Registers.E -= 1
	cpu.Registers.F &= C_BIT
	cpu.Registers.F |= N_BIT

	cpu.setZ(cpu.Registers.E)

	if ((cpu.Registers.E + 1) & 0xf) != 0 {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 4
}

// 0x2D DEC L
func (cpu *CPU_struct) op_0x2d() {
	cpu.Registers.L -= 1
	cpu.Registers.F &= C_BIT
	cpu.Registers.F |= N_BIT

	cpu.setZ(cpu.Registers.L)

	if ((cpu.Registers.L + 1) & 0xf) != 0 {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 4
}

// 0x3d DEC A
func (cpu *CPU_struct) op_0x3d() {
	cpu.Registers.A -= 1
	cpu.Registers.F &= C_BIT
	cpu.Registers.F |= N_BIT

	cpu.setZ(cpu.Registers.A)

	if ((cpu.Registers.A + 1) & 0xf) != 0 {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 4
}

// 0x0e - LD C,n
func (cpu *CPU_struct) op_0x0e(n uint8) {
	cpu.Registers.C = n

	cpu.Cycle += 8
}

// 0x1e - LD E,n
func (cpu *CPU_struct) op_0x1e(n uint8) {
	cpu.Registers.E = n

	cpu.Cycle += 8
}

// 0x2e - LD L,n
func (cpu *CPU_struct) op_0x2e(n uint8) {
	cpu.Registers.L = n

	cpu.Cycle += 8
}

// 0x3e - LD A,n
func (cpu *CPU_struct) op_0x3e(n uint8) {
	cpu.Registers.A = n

	cpu.Cycle += 8
}

func (cpu *CPU_struct) rra() {
	cpu.Registers.F = 0
	// First bit from A on carry
	cpu.Registers.F |= ((cpu.Registers.A & 0x01) << 4)

	cpu.Registers.A = cpu.Registers.A >> 1

	cpu.setZ(cpu.Registers.A)

	cpu.Cycle += 4
}

// 0x0f RRCA
func (cpu *CPU_struct) op_0x0f() {
	cpu.rra()
}

// 0x1f RRA
func (cpu *CPU_struct) op_0x1f() {
	cpu.rra()
}

// 0x2f CPL
func (cpu *CPU_struct) op_0x2f() {
	cpu.Registers.A = cpu.Registers.A ^ 0xff
	cpu.Registers.F |= (N_BIT | H_BIT)

	cpu.Cycle += 4
}

// 0x3f CCF
func (cpu *CPU_struct) op_0x3f() {
	cpu.Registers.F &= Z_BIT
	cpu.Registers.F ^= C_BIT

	cpu.Cycle += 4
}

// 0x40 LD B, B
func (cpu *CPU_struct) op_0x40() {
	cpu.Cycle += 4
}

// 0x41 LD B, C
func (cpu *CPU_struct) op_0x41() {
	cpu.Registers.B = cpu.Registers.C

	cpu.Cycle += 4
}

// 0x42 LD B, D
func (cpu *CPU_struct) op_0x42() {
	cpu.Registers.B = cpu.Registers.D

	cpu.Cycle += 4
}

// 0x43 LD B, E
func (cpu *CPU_struct) op_0x43() {
	cpu.Registers.B = cpu.Registers.E

	cpu.Cycle += 4
}

// 0x44 LD B, H
func (cpu *CPU_struct) op_0x44() {
	cpu.Registers.B = cpu.Registers.H

	cpu.Cycle += 4
}

// 0x45 LD B, L
func (cpu *CPU_struct) op_0x45() {
	cpu.Registers.B = cpu.Registers.L

	cpu.Cycle += 4
}

// 0x46 LD B, (HL)
func (cpu *CPU_struct) op_0x46() {
	var address uint16 = cpu.getHL()
	cpu.Registers.B = cpu.MMU.ReadByte(address)

	cpu.Cycle += 8
}

// 0x47 LD B,A
func (cpu *CPU_struct) op_0x47() {
	cpu.Registers.B = cpu.Registers.A

	cpu.Cycle += 4
}

// 0x48 LD C, B
func (cpu *CPU_struct) op_0x48() {
	cpu.Registers.C = cpu.Registers.B

	cpu.Cycle += 4
}

// 0x49 LD C, C
func (cpu *CPU_struct) op_0x49() {
	cpu.Cycle += 4
}

// 0x4a LD C, D
func (cpu *CPU_struct) op_0x4a() {
	cpu.Registers.C = cpu.Registers.D

	cpu.Cycle += 4
}

// 0x4b LD C, E
func (cpu *CPU_struct) op_0x4b() {
	cpu.Registers.C = cpu.Registers.E

	cpu.Cycle += 4
}

// 0x4c LD C, H
func (cpu *CPU_struct) op_0x4c() {
	cpu.Registers.C = cpu.Registers.H

	cpu.Cycle += 4
}

// 0x4d LD C, L
func (cpu *CPU_struct) op_0x4d() {
	cpu.Registers.C = cpu.Registers.L

	cpu.Cycle += 4
}

// 0x4e LD C, (HL)
func (cpu *CPU_struct) op_0x4e() {
	var address uint16 = cpu.getHL()
	cpu.Registers.C = cpu.MMU.ReadByte(address)

	cpu.Cycle += 8
}

// 0x4f LD C,A
func (cpu *CPU_struct) op_0x4f() {
	cpu.Registers.C = cpu.Registers.A

	cpu.Cycle += 4
}

// 0x50 LD D, B
func (cpu *CPU_struct) op_0x50() {
	cpu.Registers.D = cpu.Registers.B

	cpu.Cycle += 4
}

// 0x51 LD D, C
func (cpu *CPU_struct) op_0x51() {
	cpu.Registers.D = cpu.Registers.C

	cpu.Cycle += 4
}

// 0x52 LD D, D
func (cpu *CPU_struct) op_0x52() {
	cpu.Cycle += 4
}

// 0x53 LD D, E
func (cpu *CPU_struct) op_0x53() {
	cpu.Registers.D = cpu.Registers.E

	cpu.Cycle += 4
}

// 0x54 LD D, H
func (cpu *CPU_struct) op_0x54() {
	cpu.Registers.D = cpu.Registers.H

	cpu.Cycle += 4
}

// 0x55 LD D, L
func (cpu *CPU_struct) op_0x55() {
	cpu.Registers.D = cpu.Registers.L

	cpu.Cycle += 4
}

// 0x56 LD D, (HL)
func (cpu *CPU_struct) op_0x56() {
	var address uint16 = cpu.getHL()
	cpu.Registers.D = cpu.MMU.ReadByte(address)

	cpu.Cycle += 8
}

// 0x57 LD D, A
func (cpu *CPU_struct) op_0x57() {
	cpu.Registers.D = cpu.Registers.A

	cpu.Cycle += 4
}

// 0x58 LD E, B
func (cpu *CPU_struct) op_0x58() {
	cpu.Registers.E = cpu.Registers.B

	cpu.Cycle += 4
}

// 0x59 LD E, C
func (cpu *CPU_struct) op_0x59() {
	cpu.Registers.E = cpu.Registers.C

	cpu.Cycle += 4
}

// 0x5a LD E, D
func (cpu *CPU_struct) op_0x5a() {
	cpu.Registers.E = cpu.Registers.D

	cpu.Cycle += 4
}

// 0x5b LD E, E
func (cpu *CPU_struct) op_0x5b() {
	cpu.Cycle += 4
}

// 0x5c LD E, H
func (cpu *CPU_struct) op_0x5c() {
	cpu.Registers.E = cpu.Registers.H

	cpu.Cycle += 4
}

// 0x5d LD E, L
func (cpu *CPU_struct) op_0x5d() {
	cpu.Registers.E = cpu.Registers.L

	cpu.Cycle += 4
}

// 0x5e LD E, (HL)
func (cpu *CPU_struct) op_0x5e() {
	var address uint16 = cpu.getHL()
	cpu.Registers.E = cpu.MMU.ReadByte(address)

	cpu.Cycle += 8
}

// 0x5f LD E, A
func (cpu *CPU_struct) op_0x5f() {
	cpu.Registers.E = cpu.Registers.A

	cpu.Cycle += 4
}

// 0x60 LD H, B
func (cpu *CPU_struct) op_0x60() {
	cpu.Registers.H = cpu.Registers.B

	cpu.Cycle += 4
}

// 0x61 LD H, C
func (cpu *CPU_struct) op_0x61() {
	cpu.Registers.H = cpu.Registers.C

	cpu.Cycle += 4
}

// 0x62 LD H, D
func (cpu *CPU_struct) op_0x62() {
	cpu.Registers.H = cpu.Registers.D

	cpu.Cycle += 4
}

// 0x63 LD H, E
func (cpu *CPU_struct) op_0x63() {
	cpu.Registers.H = cpu.Registers.E

	cpu.Cycle += 4
}

// 0x64 LD H, H
func (cpu *CPU_struct) op_0x64() {
	cpu.Cycle += 4
}

// 0x65 LD H, L
func (cpu *CPU_struct) op_0x65() {
	cpu.Registers.H = cpu.Registers.L

	cpu.Cycle += 4
}

// 0x66 LD H, (HL)
func (cpu *CPU_struct) op_0x66() {
	var address uint16 = cpu.getHL()
	cpu.Registers.H = cpu.MMU.ReadByte(address)

	cpu.Cycle += 8
}

// 0x67 LD H, A
func (cpu *CPU_struct) op_0x67() {
	cpu.Registers.H = cpu.Registers.A

	cpu.Cycle += 4
}

// 0x68 LD L, B
func (cpu *CPU_struct) op_0x68() {
	cpu.Registers.L = cpu.Registers.B

	cpu.Cycle += 4
}

// 0x69 LD L, C
func (cpu *CPU_struct) op_0x69() {
	cpu.Registers.L = cpu.Registers.C

	cpu.Cycle += 4
}

// 0x6a LD L, D
func (cpu *CPU_struct) op_0x6a() {
	cpu.Registers.L = cpu.Registers.D

	cpu.Cycle += 4
}

// 0x6b LD L, E
func (cpu *CPU_struct) op_0x6b() {
	cpu.Registers.L = cpu.Registers.E

	cpu.Cycle += 4
}

// 0x6c LD L, H
func (cpu *CPU_struct) op_0x6c() {
	cpu.Registers.L = cpu.Registers.H

	cpu.Cycle += 4
}

// 0x6d LD L, L
func (cpu *CPU_struct) op_0x6d() {
	cpu.Cycle += 8
}

// 0x6e LD L, (HL)
func (cpu *CPU_struct) op_0x6e() {
	var address uint16 = cpu.getHL()
	cpu.Registers.L = cpu.MMU.ReadByte(address)

	cpu.Cycle += 8
}

// 0x6f LD L, A
func (cpu *CPU_struct) op_0x6f() {
	cpu.Registers.L = cpu.Registers.A

	cpu.Cycle += 4
}

// 0x70 LD (HL), B
func (cpu *CPU_struct) op_0x70() {
	var address uint16 = cpu.getHL()
	cpu.MMU.WriteByte(address, cpu.Registers.B)

	cpu.Cycle += 8
}

// 0x71 LD (HL), C
func (cpu *CPU_struct) op_0x71() {
	var address uint16 = cpu.getHL()
	cpu.MMU.WriteByte(address, cpu.Registers.C)

	cpu.Cycle += 8
}

// 0x72 LD (HL), D
func (cpu *CPU_struct) op_0x72() {
	var address uint16 = cpu.getHL()
	cpu.MMU.WriteByte(address, cpu.Registers.D)

	cpu.Cycle += 8
}

// 0x73 LD (HL), E
func (cpu *CPU_struct) op_0x73() {
	var address uint16 = cpu.getHL()
	cpu.MMU.WriteByte(address, cpu.Registers.E)

	cpu.Cycle += 8
}

// 0x74 LD (HL), H
func (cpu *CPU_struct) op_0x74() {
	var address uint16 = cpu.getHL()
	cpu.MMU.WriteByte(address, cpu.Registers.H)

	cpu.Cycle += 8
}

// 0x75 LD (HL), L
func (cpu *CPU_struct) op_0x75() {
	var address uint16 = cpu.getHL()
	cpu.MMU.WriteByte(address, cpu.Registers.L)

	cpu.Cycle += 8
}

// 0x76 HALT
func (cpu *CPU_struct) op_0x76() {
	cpu.Registers.Halt = 1
	cpu.Registers.IME = 1
	cpu.Cycle += 4
}

// 0x77 LD (HL),A
func (cpu *CPU_struct) op_0x77() {
	var address uint16 = cpu.getHL()
	cpu.MMU.WriteByte(address, cpu.Registers.A)

	cpu.Cycle += 8
}

// 0x78 LD A, B
func (cpu *CPU_struct) op_0x78() {
	cpu.Registers.A = cpu.Registers.B

	cpu.Cycle += 4
}

// 0x79 LD A, C
func (cpu *CPU_struct) op_0x79() {
	cpu.Registers.A = cpu.Registers.C

	cpu.Cycle += 4
}

// 0x7A LD A, D
func (cpu *CPU_struct) op_0x7a() {
	cpu.Registers.A = cpu.Registers.D

	cpu.Cycle += 4
}

// 0x7b LD A, E
func (cpu *CPU_struct) op_0x7b() {
	cpu.Registers.A = cpu.Registers.E

	cpu.Cycle += 4
}

// 0x7c LD A, H
func (cpu *CPU_struct) op_0x7c() {
	cpu.Registers.A = cpu.Registers.H

	cpu.Cycle += 4
}

// 0x7d LD A, L
func (cpu *CPU_struct) op_0x7d() {
	cpu.Registers.A = cpu.Registers.L

	cpu.Cycle += 4
}

// 0x7e LD A, (HL)
func (cpu *CPU_struct) op_0x7e() {
	var address uint16 = cpu.getHL()
	cpu.Registers.A = cpu.MMU.ReadByte(address)

	cpu.Cycle += 8
}

// 0x7f LD A, A
func (cpu *CPU_struct) op_0x7f() {
	cpu.Cycle += 4
}

func (cpu *CPU_struct) add(n uint8) {
	old_a := cpu.Registers.A
	cpu.Registers.A += n
	cpu.Registers.F = 0

	cpu.setZ(cpu.Registers.A)

	if (((old_a & 0x0f) + (n & 0x0f)) & 0xf0) != 0 {
		cpu.Registers.F |= H_BIT
	}

	if (((uint16(old_a)) + (uint16(n))) & 0xff00) != 0 {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 4
}

// 0x80 - ADD A, B
func (cpu *CPU_struct) op_0x80() {
	cpu.add(cpu.Registers.B)
}

// 0x81 - ADD A, C
func (cpu *CPU_struct) op_0x81() {
	cpu.add(cpu.Registers.C)
}

// 0x82 - ADD A, D
func (cpu *CPU_struct) op_0x82() {
	cpu.add(cpu.Registers.D)
}

// 0x83 - ADD A, E
func (cpu *CPU_struct) op_0x83() {
	cpu.add(cpu.Registers.E)
}

// 0x84 - ADD A, H
func (cpu *CPU_struct) op_0x84() {
	cpu.add(cpu.Registers.H)
}

// 0x85 - ADD A, L
func (cpu *CPU_struct) op_0x85() {
	cpu.add(cpu.Registers.L)
}

// 0x86 - ADD A, (HL)
func (cpu *CPU_struct) op_0x86() {
	var address uint16 = cpu.getHL()

	cpu.add(cpu.MMU.ReadByte(address))

	cpu.Cycle += 4
}

// 0x87 - ADD A, A
func (cpu *CPU_struct) op_0x87() {
	cpu.add(cpu.Registers.A)
}

func (cpu *CPU_struct) adc(n uint8) {
	var carry uint8 = 0

	if (cpu.Registers.F & C_BIT) > 0 {
		carry += 1
	}

	cpu.add(n + carry)
}

// 0x88 - ADC A, B
func (cpu *CPU_struct) op_0x88() {
	cpu.adc(cpu.Registers.B)
}

// 0x89 - ADC A, C
func (cpu *CPU_struct) op_0x89() {
	cpu.adc(cpu.Registers.C)
}

// 0x8a - ADC A, D
func (cpu *CPU_struct) op_0x8a() {
	cpu.adc(cpu.Registers.D)
}

// 0x8b - ADC A, E
func (cpu *CPU_struct) op_0x8b() {
	cpu.adc(cpu.Registers.E)
}

// 0x8c - ADC A, H
func (cpu *CPU_struct) op_0x8c() {
	cpu.adc(cpu.Registers.H)
}

// 0x8d - ADC A, L
func (cpu *CPU_struct) op_0x8d() {
	cpu.adc(cpu.Registers.L)
}

// 0x8e - ADC A, (HL)
func (cpu *CPU_struct) op_0x8e() {
	var address uint16 = cpu.getHL()

	cpu.adc(cpu.MMU.ReadByte(address))

	cpu.Cycle += 4
}

// 0x8f - ADC A, A
func (cpu *CPU_struct) op_0x8f() {
	cpu.adc(cpu.Registers.A)
}

func (cpu *CPU_struct) sub(n uint8) {
	old := cpu.Registers.A
	cpu.Registers.A = cpu.Registers.A - n
	cpu.Registers.F = N_BIT

	cpu.setZ(cpu.Registers.A)

	if (old & 0x0f) < (n & 0x0f) {
		cpu.Registers.F |= H_BIT
	}

	if old < n {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 4
}

// 0x90 - SUB A, B
func (cpu *CPU_struct) op_0x90() {
	cpu.sub(cpu.Registers.B)
}

// 0x91 - SUB A, C
func (cpu *CPU_struct) op_0x91() {
	cpu.sub(cpu.Registers.C)
}

// 0x92 - SUB A, D
func (cpu *CPU_struct) op_0x92() {
	cpu.sub(cpu.Registers.D)
}

// 0x93 - SUB A, E
func (cpu *CPU_struct) op_0x93() {
	cpu.sub(cpu.Registers.E)
}

// 0x94 - SUB A, H
func (cpu *CPU_struct) op_0x94() {
	cpu.sub(cpu.Registers.H)
}

// 0x95 - SUB A, L
func (cpu *CPU_struct) op_0x95() {
	cpu.sub(cpu.Registers.L)
}

// 0x96 - SUB A, (HL)
func (cpu *CPU_struct) op_0x96() {
	var address uint16 = cpu.getHL()

	cpu.sub(cpu.MMU.ReadByte(address))

	cpu.Cycle += 4
}

// 0x97 - SUB A, A
func (cpu *CPU_struct) op_0x97() {
	cpu.sub(cpu.Registers.A)
}

func (cpu *CPU_struct) sbc(n uint8) {
	var carry uint8 = 0

	if (cpu.Registers.F & C_BIT) > 0 {
		carry += 1
	}

	cpu.sub(n + carry)
}

// 0x98 - SBC A, B
func (cpu *CPU_struct) op_0x98() {
	cpu.sbc(cpu.Registers.B)
}

// 0x99 - SBC A, C
func (cpu *CPU_struct) op_0x99() {
	cpu.sbc(cpu.Registers.C)
}

// 0x9a - SBC A, D
func (cpu *CPU_struct) op_0x9a() {
	cpu.sbc(cpu.Registers.D)
}

// 0x9b - SBC A, E
func (cpu *CPU_struct) op_0x9b() {
	cpu.sbc(cpu.Registers.E)
}

// 0x9c - SBC A, H
func (cpu *CPU_struct) op_0x9c() {
	cpu.sbc(cpu.Registers.H)
}

// 0x9d - SBC A, L
func (cpu *CPU_struct) op_0x9d() {
	cpu.sbc(cpu.Registers.L)
}

// 0x9e - SBC A, B
func (cpu *CPU_struct) op_0x9e() {
	var address uint16 = cpu.getHL()

	cpu.sbc(cpu.MMU.ReadByte(address))

	cpu.Cycle += 4
}

// 0x9f - SBC A, A
func (cpu *CPU_struct) op_0x9f() {
	cpu.sbc(cpu.Registers.A)
}

func (cpu *CPU_struct) and(n uint8) {
	cpu.Registers.A &= n

	cpu.Registers.F = H_BIT

	cpu.setZ(cpu.Registers.A)

	cpu.Cycle += 4
}

// 0xa0 - AND A, B
func (cpu *CPU_struct) op_0xa0() {
	cpu.and(cpu.Registers.B)
}

// 0xa1 - AND A, C
func (cpu *CPU_struct) op_0xa1() {
	cpu.and(cpu.Registers.C)
}

// 0xa2 - AND A, D
func (cpu *CPU_struct) op_0xa2() {
	cpu.and(cpu.Registers.D)
}

// 0xa3 - AND A, E
func (cpu *CPU_struct) op_0xa3() {
	cpu.and(cpu.Registers.E)
}

// 0xa4 - AND A, H
func (cpu *CPU_struct) op_0xa4() {
	cpu.and(cpu.Registers.H)
}

// 0xa5 - AND A, L
func (cpu *CPU_struct) op_0xa5() {
	cpu.and(cpu.Registers.L)
}

// 0xa6 - AND A, (HL)
func (cpu *CPU_struct) op_0xa6() {
	var address uint16 = cpu.getHL()
	cpu.and(cpu.MMU.ReadByte(address))

	cpu.Cycle += 4
}

// 0xa7 - AND A, A
func (cpu *CPU_struct) op_0xa7() {
	cpu.and(cpu.Registers.A)
}

func (cpu *CPU_struct) xor(n uint8) {
	cpu.Registers.A ^= n

	cpu.Registers.F = 0

	cpu.setZ(cpu.Registers.A)

	cpu.Cycle += 4
}

// 0xa8 - XOR A, B
func (cpu *CPU_struct) op_0xa8() {
	cpu.xor(cpu.Registers.B)
}

// 0xa9 - XOR A, C
func (cpu *CPU_struct) op_0xa9() {
	cpu.xor(cpu.Registers.C)
}

// 0xaa - XOR A, D
func (cpu *CPU_struct) op_0xaa() {
	cpu.xor(cpu.Registers.D)
}

// 0xab - XOR A, E
func (cpu *CPU_struct) op_0xab() {
	cpu.xor(cpu.Registers.E)
}

// 0xac - XOR A, H
func (cpu *CPU_struct) op_0xac() {
	cpu.xor(cpu.Registers.H)
}

// 0xad - XOR A, L
func (cpu *CPU_struct) op_0xad() {
	cpu.xor(cpu.Registers.L)
}

// 0xae - XOR A, (HL)
func (cpu *CPU_struct) op_0xae() {
	var address uint16 = cpu.getHL()
	cpu.xor(cpu.MMU.ReadByte(address))

	cpu.Cycle += 4
}

// 0xaf - XOR A, A
func (cpu *CPU_struct) op_0xaf() {
	cpu.xor(cpu.Registers.A)
}

func (cpu *CPU_struct) or(n uint8) {
	cpu.Registers.A |= n
	cpu.Registers.F = 0

	cpu.setZ(cpu.Registers.A)

	cpu.Cycle += 4
}

// 0xb0 - OR A, B
func (cpu *CPU_struct) op_0xb0() {
	cpu.or(cpu.Registers.B)
}

// 0xb1 - OR A, C
func (cpu *CPU_struct) op_0xb1() {
	cpu.or(cpu.Registers.C)
}

// 0xb2 - OR A, D
func (cpu *CPU_struct) op_0xb2() {
	cpu.or(cpu.Registers.D)
}

// 0xb3 - OR A, E
func (cpu *CPU_struct) op_0xb3() {
	cpu.or(cpu.Registers.E)
}

// 0xb4 - OR A, H
func (cpu *CPU_struct) op_0xb4() {
	cpu.or(cpu.Registers.H)
}

// 0xb5 - OR A, L
func (cpu *CPU_struct) op_0xb5() {
	cpu.or(cpu.Registers.L)
}

// 0xb6 - OR A, (HL)
func (cpu *CPU_struct) op_0xb6() {
	var address uint16 = cpu.getHL()
	cpu.or(cpu.MMU.ReadByte(address))

	cpu.Cycle += 4
}

// 0xb7 - OR A, A
func (cpu *CPU_struct) op_0xb7() {
	cpu.or(cpu.Registers.A)
}

func (cpu *CPU_struct) cp(n uint8) {
	cpu.Registers.F = N_BIT

	cpu.setZ(cpu.Registers.A - n)

	if (cpu.Registers.A & 0x0f) < (n & 0x0f) {
		cpu.Registers.F |= H_BIT
	}

	if cpu.Registers.A < n {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 4
}

// 0xb8 - CP A, B
func (cpu *CPU_struct) op_0xb8() {
	cpu.cp(cpu.Registers.B)
}

// 0xb9 - CP A, C
func (cpu *CPU_struct) op_0xb9() {
	cpu.cp(cpu.Registers.C)
}

// 0xba - CP A, D
func (cpu *CPU_struct) op_0xba() {
	cpu.cp(cpu.Registers.D)
}

// 0xbb - CP A, E
func (cpu *CPU_struct) op_0xbb() {
	cpu.cp(cpu.Registers.E)
}

// 0xbc - CP A, H
func (cpu *CPU_struct) op_0xbc() {
	cpu.cp(cpu.Registers.H)
}

// 0xbd - CP A, L
func (cpu *CPU_struct) op_0xbd() {
	cpu.cp(cpu.Registers.L)
}

// 0xbe - CP A, (HL)
func (cpu *CPU_struct) op_0xbe() {
	var address uint16 = cpu.getHL()
	cpu.cp(cpu.MMU.ReadByte(address))

	cpu.Cycle += 4
}

// 0xbf - CP A, A
func (cpu *CPU_struct) op_0xbf() {
	cpu.cp(cpu.Registers.A)
}

func (cpu *CPU_struct) ret() {
	var addr uint16 = cpu.MMU.ReadWord(cpu.Registers.SP)
	cpu.Registers.SP += 2
	cpu.Registers.PC = addr

	cpu.Cycle += 8
}

// 0xc0 RET NZ
func (cpu *CPU_struct) op_0xc0() {
	if (cpu.Registers.F & Z_BIT) == 0 {
		cpu.ret()
	}
}

// 0xd0 RET NC
func (cpu *CPU_struct) op_0xd0() {
	if (cpu.Registers.F & C_BIT) == 0 {
		cpu.ret()
	}
}

// 0xe0 LDH (n),A
func (cpu *CPU_struct) op_0xe0(n uint8) {
	cpu.MMU.WriteByte(0xff00|uint16(n), cpu.Registers.A)

	cpu.Cycle += 12
}

// 0xf0 LDH A,(n)
func (cpu *CPU_struct) op_0xf0(n uint8) {
	cpu.Registers.A = cpu.MMU.ReadByte(0xff00 | uint16(n))

	cpu.Cycle += 12
}

// 0xc1 POP BC
func (cpu *CPU_struct) op_0xc1() {
	var data uint16 = cpu.MMU.ReadWord(cpu.Registers.SP)

	cpu.setBC(data)
	cpu.Registers.SP += 2

	cpu.Cycle += 12
}

// 0xd1 POP DE
func (cpu *CPU_struct) op_0xd1() {
	var data uint16 = cpu.MMU.ReadWord(cpu.Registers.SP)

	cpu.setDE(data)
	cpu.Registers.SP += 2

	cpu.Cycle += 12
}

// 0xe1 POP HL
func (cpu *CPU_struct) op_0xe1() {
	var data uint16 = cpu.MMU.ReadWord(cpu.Registers.SP)

	cpu.setHL(data)
	cpu.Registers.SP += 2

	cpu.Cycle += 12
}

// 0xf1 POP AF
func (cpu *CPU_struct) op_0xf1() {
	var data uint16 = cpu.MMU.ReadWord(cpu.Registers.SP)

	cpu.setAF(data)
	cpu.Registers.SP += 2

	cpu.Cycle += 12
}

func (cpu *CPU_struct) jp(nn uint16) {
	cpu.Registers.PC = nn
	cpu.Cycle += 12
}

// 0xc2 - JP NZ, nn
func (cpu *CPU_struct) op_0xc2(nn uint16) {
	if (cpu.Registers.F & Z_BIT) == 0 {
		cpu.jp(nn)
	}
}

// 0xd2 - JP NC, nn
func (cpu *CPU_struct) op_0xd2(nn uint16) {
	if (cpu.Registers.F & C_BIT) == 0 {
		cpu.jp(nn)
	}
}

// 0xe2 LD (C), A
func (cpu *CPU_struct) op_0xe2() {
	cpu.MMU.WriteByte(0xff00|uint16(cpu.Registers.C), cpu.Registers.A)

	cpu.Cycle += 8
}

// 0xf2 LD A, (C)
func (cpu *CPU_struct) op_0xf2() {
	cpu.Registers.A = cpu.MMU.ReadByte(0xff00 | uint16(cpu.Registers.C))

	cpu.Cycle += 8
}

// 0xc3 - JP nn
func (cpu *CPU_struct) op_0xc3(nn uint16) {
	cpu.jp(nn)
}

// 0xf3 DI
func (cpu *CPU_struct) op_0xf3() {
	cpu.Registers.IME = 0

	cpu.Cycle += 4
}

// 0xc4 CALL NZ, nn
func (cpu *CPU_struct) op_0xc4(nn uint16) {
	if (cpu.Registers.F & Z_BIT) == 0 {
		cpu.call(nn)
	}
}

// 0xd4 CALL NC, nn
func (cpu *CPU_struct) op_0xd4(nn uint16) {
	if (cpu.Registers.F & C_BIT) == 0 {
		cpu.call(nn)
	}
}

// 0xc5 PUSH BC
func (cpu *CPU_struct) op_0xc5() {
	var nn uint16 = cpu.getBC()
	cpu.Registers.SP -= 2
	cpu.MMU.WriteWord(cpu.Registers.SP, nn)

	cpu.Cycle += 16
}

// 0xd5 PUSH DE
func (cpu *CPU_struct) op_0xd5() {
	var de uint16 = cpu.getDE()
	cpu.Registers.SP -= 2
	cpu.MMU.WriteWord(cpu.Registers.SP, de)

	cpu.Cycle += 16
}

// 0xe5 PUSH HL
func (cpu *CPU_struct) op_0xe5() {
	var hl uint16 = cpu.getHL()
	cpu.Registers.SP -= 2
	cpu.MMU.WriteWord(cpu.Registers.SP, hl)

	cpu.Cycle += 16
}

// 0xf5 PUSH AF
func (cpu *CPU_struct) op_0xf5() {
	var af uint16 = cpu.getAF()
	cpu.Registers.SP -= 2
	cpu.MMU.WriteWord(cpu.Registers.SP, af)

	cpu.Cycle += 16
}

// 0xc6 ADD #
func (cpu *CPU_struct) op_0xc6(n uint8) {
	cpu.add(n)

	cpu.Cycle += 4
}

// 0xd6 SUB #
func (cpu *CPU_struct) op_0xd6(n uint8) {
	cpu.sub(n)

	cpu.Cycle += 4
}

// 0xe6 AND #
func (cpu *CPU_struct) op_0xe6(n uint8) {
	cpu.and(n)

	cpu.Cycle += 4
}

// 0xf6 OR #
func (cpu *CPU_struct) op_0xf6(n uint8) {
	cpu.or(n)

	cpu.Cycle += 4
}

func (cpu *CPU_struct) rst(nn uint16) {
	cpu.Registers.SP -= 2
	cpu.MMU.WriteWord(cpu.Registers.SP, cpu.Registers.PC)
	cpu.Registers.PC = nn

	cpu.Cycle += 16
}

// 0xc7 RST 00h
func (cpu *CPU_struct) op_0xc7() {
	cpu.rst(0x0000)
}

// 0xd7 RST 10h
func (cpu *CPU_struct) op_0xd7() {
	cpu.rst(0x0010)
}

// 0xe7 RST 20h
func (cpu *CPU_struct) op_0xe7() {
	cpu.rst(0x0020)
}

// 0xf7 RST 30h
func (cpu *CPU_struct) op_0xf7() {
	cpu.rst(0x0030)
}

// 0xc8 RET Z
func (cpu *CPU_struct) op_0xc8() {
	if (cpu.Registers.F & Z_BIT) > 0 {
		cpu.ret()
	}
}

// 0xd8 RET Z
func (cpu *CPU_struct) op_0xd8() {
	if (cpu.Registers.F & C_BIT) > 0 {
		cpu.ret()
	}
}

// 0xE8 ADD SP, n
// TODO double check
func (cpu *CPU_struct) op_0xe8(n uint8) {
	old_sp := cpu.Registers.SP
	cpu.Registers.SP = uint16(int16(cpu.Registers.SP) + int16(n))
	cpu.Registers.F = 0

	if (int16(n) > 0) && ((((old_sp & 0x0f) + (uint16(n) & 0x0f)) & 0xf0) != 0) {
		cpu.Registers.F |= H_BIT
	}

	if (int16(n) > 0) && ((((old_sp & 0x00ff) + (uint16(n) & 0x00ff)) & 0xff00) != 0) {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 16
}

// LDHL SP, n
// TODO double check
func (cpu *CPU_struct) op_0xf8(n uint8) {
	cpu.setHL(uint16(int16(cpu.Registers.SP) + int16(n)))

	cpu.Registers.F = 0

	if (int16(n) > 0) && ((((cpu.Registers.SP & 0x0f) + (uint16(n) & 0x0f)) & 0xf0) != 0) {
		cpu.Registers.F |= H_BIT
	}

	if (int16(n) > 0) && ((((cpu.Registers.SP & 0x00ff) + (uint16(n) & 0x00ff)) & 0xff00) != 0) {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 12
}

// 0xc9 RET
func (cpu *CPU_struct) op_0xc9() {
	cpu.ret()
}

// 0xd9 RETI
func (cpu *CPU_struct) op_0xd9() {
	cpu.ret()
	cpu.Registers.IME = 1
}

// 0xe9 - JP (HL)
func (cpu *CPU_struct) op_0xe9() {
	var address uint16 = cpu.getHL()
	cpu.Registers.PC = address

	cpu.Cycle += 4
}

// 0xf9 - LD SP, HL
func (cpu *CPU_struct) op_0xf9() {
	cpu.Registers.SP = cpu.getHL()

	cpu.Cycle += 8
}

// 0xca - JP Z, nn
func (cpu *CPU_struct) op_0xca(nn uint16) {
	if (cpu.Registers.F & Z_BIT) > 0 {
		cpu.jp(nn)
	}
}

// 0xda - JP C, nn
func (cpu *CPU_struct) op_0xda(nn uint16) {
	if (cpu.Registers.F & C_BIT) > 0 {
		cpu.jp(nn)
	}
}

// 0xea LD (nn),A
func (cpu *CPU_struct) op_0xea(nn uint16) {
	cpu.MMU.WriteByte(nn, cpu.Registers.A)

	cpu.Cycle += 16
}

//  0xfa LD A,(nn)
func (cpu *CPU_struct) op_0xfa(nn uint16) {
	cpu.Registers.A = cpu.MMU.ReadByte(nn)

	cpu.Cycle += 16
}

// 0xfb EI
func (cpu *CPU_struct) op_0xfb() {
	cpu.Registers.IME = 1

	cpu.Cycle += 4
}

// 0xcc CALL Z, nn
func (cpu *CPU_struct) op_0xcc(nn uint16) {
	if (cpu.Registers.F & Z_BIT) != 0 {
		cpu.call(nn)
	}
}

// 0xdc CALL C, nn
func (cpu *CPU_struct) op_0xdc(nn uint16) {
	if (cpu.Registers.F & C_BIT) != 0 {
		cpu.call(nn)
	}
}

// 0xcd CALL nn
func (cpu *CPU_struct) op_0xcd(nn uint16) {
	cpu.call(nn)
}

// 0xce ADC A, n
func (cpu *CPU_struct) op_0xce(n uint8) {
	cpu.adc(n)

	cpu.Cycle += 4
}

// 0xde SBC A, n
func (cpu *CPU_struct) op_0xde(n uint8) {
	cpu.sbc(n)

	cpu.Cycle += 4
}

// 0xee XOR A, n
func (cpu *CPU_struct) op_0xee(n uint8) {
	cpu.xor(n)

	cpu.Cycle += 4
}

// 0xfe CP A, n
func (cpu *CPU_struct) op_0xfe(n uint8) {
	cpu.cp(n)

	cpu.Cycle += 4
}

// 0xcf RST 0x08
func (cpu *CPU_struct) op_0xcf() {
	cpu.rst(0x0008)
}

// 0xdf RST 0x18
func (cpu *CPU_struct) op_0xdf() {
	cpu.rst(0x0018)
}

// 0xef RST 0x28
func (cpu *CPU_struct) op_0xef() {
	cpu.rst(0x0028)
}

// 0xff RST 0x38
func (cpu *CPU_struct) op_0xff() {
	cpu.rst(0x0038)
}

func (cpu *CPU_struct) call(nn uint16) {
	cpu.MMU.WriteWord(cpu.Registers.SP-2, cpu.Registers.PC)
	cpu.Registers.SP -= 2
	cpu.Registers.PC = nn

	cpu.Cycle += 12
}

//===============================================================
//===============================================================
//===============================================================
//===============================================================
//===============================================================
//===============================================================
//===============================================================

// 0xcb 0x11 RL  C
func (cpu *CPU_struct) op_cb_0x11() {
	var carry uint8 = cpu.Registers.F & C_BIT
	cpu.Registers.F = 0
	cpu.Registers.F |= ((cpu.Registers.C & 0x80) >> 3)

	cpu.Registers.C = (cpu.Registers.C << 1) & 0xff
	if carry != 0 {
		cpu.Registers.C |= 0x01
	}

	if cpu.Registers.C == 0 {
		cpu.Registers.F |= Z_BIT
	}

	cpu.Cycle += 8
}

// 0xcb 0x07 RLC A
func (cpu *CPU_struct) op_cb_0x07() {
	cpu.rla()
}

// 0xcb 0x38 SRL B
func (cpu *CPU_struct) op_cb_0x38() {
	cpu.Registers.F = 0
	cpu.Registers.F |= ((cpu.Registers.B & 0x01) << 4)

	cpu.Registers.B = cpu.Registers.B >> 1

	if cpu.Registers.B == 0 {
		cpu.Registers.F |= Z_BIT
	}

	cpu.Cycle += 8
}

// 0xcb 0x7c BIT 7, H
func (cpu *CPU_struct) op_cb_0x7c() {
	cpu.Registers.F &= C_BIT
	cpu.Registers.F |= H_BIT
	if (cpu.Registers.H & 0b10000000) == 0 {
		cpu.Registers.F |= Z_BIT
	}

	cpu.Cycle += 8
}

// 0xcb 0x6f BIT 5, A
func (cpu *CPU_struct) op_cb_0x6f() {
	cpu.Registers.F &= C_BIT
	cpu.Registers.F |= H_BIT
	if (cpu.Registers.A & 0b00100000) == 0 {
		cpu.Registers.F |= Z_BIT
	}

	cpu.Cycle += 8
}

// 0xcb 0x77 BIT 6, A
func (cpu *CPU_struct) op_cb_0x77() {
	cpu.Registers.F &= C_BIT
	cpu.Registers.F |= H_BIT
	if (cpu.Registers.A & 0b01000000) == 0 {
		cpu.Registers.F |= Z_BIT
	}

	cpu.Cycle += 8
}

// 0xcb 0x47 BIT 0, A
func (cpu *CPU_struct) op_cb_0x47() {
	cpu.Registers.F &= C_BIT
	cpu.Registers.F |= H_BIT
	if (cpu.Registers.A & 0b00000001) == 0 {
		cpu.Registers.F |= Z_BIT
	}

	cpu.Cycle += 8
}

// 0xcb 0xcf SET 1, A
func (cpu *CPU_struct) op_cb_0xcf() {
	cpu.Registers.A |= 0b00000010

	cpu.Cycle += 8
}

// 0xcb 0x87 RES 0, A
func (cpu *CPU_struct) op_cb_0x87() {
	cpu.Registers.A &= 0b11111110

	// TODO review 4 or 8, is 0xCB included here?
	cpu.Cycle += 8
}

// 0xcb 0x37 SWAP A
func (cpu *CPU_struct) op_cb_0x37() {
	cpu.Registers.A = (cpu.Registers.A << 4) | (cpu.Registers.A >> 4)

	cpu.Registers.F = 0
	if cpu.Registers.A == 0 {
		cpu.Registers.F |= Z_BIT
	}

	cpu.Cycle += 8
}
