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
	cpu.Instructions_maps.Instructions_map_uint8[0x3E] = cpu.op_0x3e
	cpu.Instructions_maps.Instructions_map_uint8[0x36] = cpu.op_0x36
	cpu.Instructions_maps.Instructions_map_uint8[0xe0] = cpu.op_0xe0
	cpu.Instructions_maps.Instructions_map_uint8[0xf0] = cpu.op_0xf0
	// JP
	cpu.Instructions_maps.Instructions_map_uint8[0x18] = cpu.op_0x18
	cpu.Instructions_maps.Instructions_map_uint8[0x20] = cpu.op_0x20
	cpu.Instructions_maps.Instructions_map_uint8[0x28] = cpu.op_0x28
	// CP
	cpu.Instructions_maps.Instructions_map_uint8[0xfe] = cpu.op_0xfe
	// Bitwise
	cpu.Instructions_maps.Instructions_map_uint8[0xe6] = cpu.op_0xe6
	cpu.Instructions_maps.Instructions_map_uint8[0xee] = cpu.op_0xee
	// ADD
	cpu.Instructions_maps.Instructions_map_uint8[0xe8] = cpu.op_0xe8
	cpu.Instructions_maps.Instructions_map_uint8[0xde] = cpu.op_0xde

	// 16
	//LD
	cpu.Instructions_maps.Instructions_map_uint16[0x31] = cpu.op_0x31
	cpu.Instructions_maps.Instructions_map_uint16[0x21] = cpu.op_0x21
	cpu.Instructions_maps.Instructions_map_uint16[0xea] = cpu.op_0xea
	cpu.Instructions_maps.Instructions_map_uint16[0x11] = cpu.op_0x11
	cpu.Instructions_maps.Instructions_map_uint16[0x01] = cpu.op_0x01
	cpu.Instructions_maps.Instructions_map_uint16[0xfa] = cpu.op_0xfa
	//JP
	cpu.Instructions_maps.Instructions_map_uint16[0xc3] = cpu.op_0xc3
	cpu.Instructions_maps.Instructions_map_uint16[0xca] = cpu.op_0xca
	cpu.Instructions_maps.Instructions_map_uint16[0xd2] = cpu.op_0xd2
	//CALL
	cpu.Instructions_maps.Instructions_map_uint16[0xcd] = cpu.op_0xcd

	// 0
	//LD
	cpu.Instructions_maps.Instructions_map_[0x22] = cpu.op_0x22
	cpu.Instructions_maps.Instructions_map_[0x32] = cpu.op_0x32
	cpu.Instructions_maps.Instructions_map_[0xe2] = cpu.op_0xe2
	cpu.Instructions_maps.Instructions_map_[0x57] = cpu.op_0x57
	cpu.Instructions_maps.Instructions_map_[0x5f] = cpu.op_0x5f
	cpu.Instructions_maps.Instructions_map_[0x67] = cpu.op_0x67
	cpu.Instructions_maps.Instructions_map_[0x6f] = cpu.op_0x6f
	cpu.Instructions_maps.Instructions_map_[0x77] = cpu.op_0x77
	cpu.Instructions_maps.Instructions_map_[0x12] = cpu.op_0x12
	cpu.Instructions_maps.Instructions_map_[0x02] = cpu.op_0x02
	cpu.Instructions_maps.Instructions_map_[0x4f] = cpu.op_0x4f
	cpu.Instructions_maps.Instructions_map_[0x1a] = cpu.op_0x1a
	cpu.Instructions_maps.Instructions_map_[0x7b] = cpu.op_0x7b
	cpu.Instructions_maps.Instructions_map_[0x7c] = cpu.op_0x7c
	cpu.Instructions_maps.Instructions_map_[0x7d] = cpu.op_0x7d
	cpu.Instructions_maps.Instructions_map_[0x7e] = cpu.op_0x7e
	cpu.Instructions_maps.Instructions_map_[0x78] = cpu.op_0x78
	cpu.Instructions_maps.Instructions_map_[0x79] = cpu.op_0x79
	cpu.Instructions_maps.Instructions_map_[0x6b] = cpu.op_0x6b
	cpu.Instructions_maps.Instructions_map_[0x44] = cpu.op_0x44
	cpu.Instructions_maps.Instructions_map_[0x4d] = cpu.op_0x4d
	cpu.Instructions_maps.Instructions_map_[0x60] = cpu.op_0x60
	cpu.Instructions_maps.Instructions_map_[0x62] = cpu.op_0x62
	cpu.Instructions_maps.Instructions_map_[0x66] = cpu.op_0x66
	cpu.Instructions_maps.Instructions_map_[0x56] = cpu.op_0x56
	cpu.Instructions_maps.Instructions_map_[0x5e] = cpu.op_0x5e
	cpu.Instructions_maps.Instructions_map_[0x69] = cpu.op_0x69
	cpu.Instructions_maps.Instructions_map_[0x7a] = cpu.op_0x7A
	cpu.Instructions_maps.Instructions_map_[0x47] = cpu.op_0x47
	cpu.Instructions_maps.Instructions_map_[0x2a] = cpu.op_0x2a
	// Bitwise
	cpu.Instructions_maps.Instructions_map_[0xaf] = cpu.op_0xaf
	cpu.Instructions_maps.Instructions_map_[0xab] = cpu.op_0xab
	cpu.Instructions_maps.Instructions_map_[0xb0] = cpu.op_0xb0
	cpu.Instructions_maps.Instructions_map_[0xb1] = cpu.op_0xb1
	cpu.Instructions_maps.Instructions_map_[0xb2] = cpu.op_0xb2
	cpu.Instructions_maps.Instructions_map_[0xb6] = cpu.op_0xb6
	cpu.Instructions_maps.Instructions_map_[0xa8] = cpu.op_0xa8
	cpu.Instructions_maps.Instructions_map_[0xa9] = cpu.op_0xa9
	cpu.Instructions_maps.Instructions_map_[0xa7] = cpu.op_0xa7
	cpu.Instructions_maps.Instructions_map_[0xa1] = cpu.op_0xa1
	// CP
	cpu.Instructions_maps.Instructions_map_[0x2f] = cpu.op_0x2f
	cpu.Instructions_maps.Instructions_map_[0xbe] = cpu.op_0xbe
	cpu.Instructions_maps.Instructions_map_[0xb8] = cpu.op_0xb8
	cpu.Instructions_maps.Instructions_map_[0xb9] = cpu.op_0xb9
	// DEC
	cpu.Instructions_maps.Instructions_map_[0x05] = cpu.op_0x05
	cpu.Instructions_maps.Instructions_map_[0x15] = cpu.op_0x15
	cpu.Instructions_maps.Instructions_map_[0x1d] = cpu.op_0x1d
	cpu.Instructions_maps.Instructions_map_[0x25] = cpu.op_0x25
	cpu.Instructions_maps.Instructions_map_[0x3d] = cpu.op_0x3d
	cpu.Instructions_maps.Instructions_map_[0x0d] = cpu.op_0x0d
	cpu.Instructions_maps.Instructions_map_[0x0b] = cpu.op_0x0b
	cpu.Instructions_maps.Instructions_map_[0x1b] = cpu.op_0x1b
	cpu.Instructions_maps.Instructions_map_[0x2b] = cpu.op_0x2b
	// INC
	cpu.Instructions_maps.Instructions_map_[0x04] = cpu.op_0x04
	cpu.Instructions_maps.Instructions_map_[0x0c] = cpu.op_0x0c
	cpu.Instructions_maps.Instructions_map_[0x24] = cpu.op_0x24
	cpu.Instructions_maps.Instructions_map_[0x2c] = cpu.op_0x2c
	cpu.Instructions_maps.Instructions_map_[0x34] = cpu.op_0x34
	cpu.Instructions_maps.Instructions_map_[0x3c] = cpu.op_0x3c
	cpu.Instructions_maps.Instructions_map_[0x03] = cpu.op_0x03
	cpu.Instructions_maps.Instructions_map_[0x13] = cpu.op_0x13
	cpu.Instructions_maps.Instructions_map_[0x23] = cpu.op_0x23
	// SUB
	cpu.Instructions_maps.Instructions_map_[0x95] = cpu.op_0x95
	cpu.Instructions_maps.Instructions_map_[0x90] = cpu.op_0x90
	cpu.Instructions_maps.Instructions_map_[0x97] = cpu.op_0x97
	// NOP
	cpu.Instructions_maps.Instructions_map_[0x00] = cpu.op_0x00
	// HALT
	// cpu.Instructions_maps.Instructions_map_[0x76] = cpu.op_0x76
	// PUSH/POP
	cpu.Instructions_maps.Instructions_map_[0xc5] = cpu.op_0xc5
	cpu.Instructions_maps.Instructions_map_[0xd5] = cpu.op_0xd5
	cpu.Instructions_maps.Instructions_map_[0xe5] = cpu.op_0xe5
	cpu.Instructions_maps.Instructions_map_[0xf5] = cpu.op_0xf5
	cpu.Instructions_maps.Instructions_map_[0xc1] = cpu.op_0xc1
	cpu.Instructions_maps.Instructions_map_[0xd1] = cpu.op_0xd1
	cpu.Instructions_maps.Instructions_map_[0xe1] = cpu.op_0xe1
	cpu.Instructions_maps.Instructions_map_[0xf1] = cpu.op_0xf1
	// RL
	cpu.Instructions_maps.Instructions_map_[0x07] = cpu.op_0x17
	cpu.Instructions_maps.Instructions_map_[0x17] = cpu.op_0x17
	// RET
	cpu.Instructions_maps.Instructions_map_[0xc9] = cpu.op_0xc9
	cpu.Instructions_maps.Instructions_map_[0xd9] = cpu.op_0xd9
	cpu.Instructions_maps.Instructions_map_[0xc8] = cpu.op_0xc8
	cpu.Instructions_maps.Instructions_map_[0xc0] = cpu.op_0xc0
	// ADD/SUB
	cpu.Instructions_maps.Instructions_map_[0x86] = cpu.op_0x86
	cpu.Instructions_maps.Instructions_map_[0x87] = cpu.op_0x87
	cpu.Instructions_maps.Instructions_map_[0x80] = cpu.op_0x80
	cpu.Instructions_maps.Instructions_map_[0x09] = cpu.op_0x09
	cpu.Instructions_maps.Instructions_map_[0x19] = cpu.op_0x19
	cpu.Instructions_maps.Instructions_map_[0x89] = cpu.op_0x89
	cpu.Instructions_maps.Instructions_map_[0x99] = cpu.op_0x99
	// DI/EI
	cpu.Instructions_maps.Instructions_map_[0xf3] = cpu.op_0xf3
	cpu.Instructions_maps.Instructions_map_[0xfb] = cpu.op_0xfb
	// JP
	cpu.Instructions_maps.Instructions_map_[0xe9] = cpu.op_0xe9
	// RST
	cpu.Instructions_maps.Instructions_map_[0xef] = cpu.op_0xef

	// CB
	cpu.Instructions_maps.Instructions_map_CB[0x11] = cpu.op_cb_0x11
	cpu.Instructions_maps.Instructions_map_CB[0x07] = cpu.op_cb_0x07
	cpu.Instructions_maps.Instructions_map_CB[0x7c] = cpu.op_cb_0x7c
	cpu.Instructions_maps.Instructions_map_CB[0x87] = cpu.op_cb_0x87
	cpu.Instructions_maps.Instructions_map_CB[0xcf] = cpu.op_cb_0xcf
	cpu.Instructions_maps.Instructions_map_CB[0x37] = cpu.op_cb_0x37
	cpu.Instructions_maps.Instructions_map_CB[0x47] = cpu.op_cb_0x47
	cpu.Instructions_maps.Instructions_map_CB[0x77] = cpu.op_cb_0x77
	cpu.Instructions_maps.Instructions_map_CB[0x6f] = cpu.op_cb_0x6f

}

// 0x06 - LD B,n
func (cpu *CPU_struct) op_0x06(n uint8) {
	cpu.Registers.B = n
	cpu.Cycle += 8
}

// 0x0e - LD C,n
func (cpu *CPU_struct) op_0x0e(n uint8) {
	cpu.Registers.C = n
	cpu.Cycle += 8
}

// 0x16 - LD D,n
func (cpu *CPU_struct) op_0x16(n uint8) {
	cpu.Registers.D = n
	cpu.Cycle += 8
}

// 0x1e - LD E,n
func (cpu *CPU_struct) op_0x1e(n uint8) {
	cpu.Registers.E = n
	cpu.Cycle += 8
}

// 0x26 - LD H,n
func (cpu *CPU_struct) op_0x26(n uint8) {
	cpu.Registers.H = n
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

// 0x36 - LD (HL),n
func (cpu *CPU_struct) op_0x36(n uint8) {
	var address uint16 = (uint16(cpu.Registers.H) << 8) | uint16(cpu.Registers.L)
	cpu.MMU.WriteByte(address, n)
	cpu.Cycle += 12
}

// 0x31 - LD SP,nn
func (cpu *CPU_struct) op_0x31(nn uint16) {
	cpu.Registers.SP = nn
	cpu.Cycle += 12
}

// 0x21 - LD HL,nn
func (cpu *CPU_struct) op_0x21(nn uint16) {
	cpu.Registers.H = uint8(nn >> 8)
	cpu.Registers.L = uint8(nn & 0xff)
	cpu.Cycle += 12
}

// 0xe9 - JP (HL)
func (cpu *CPU_struct) op_0xe9() {
	var address uint16 = (uint16(cpu.Registers.H) << 8) | uint16(cpu.Registers.L)

	cpu.Registers.PC = address
	cpu.Cycle += 4
}

// 0xc3 - JP nn
func (cpu *CPU_struct) op_0xc3(nn uint16) {
	cpu.Registers.PC = nn
	cpu.Cycle += 12
}

// 0xca - JP Z, nn
func (cpu *CPU_struct) op_0xca(nn uint16) {
	if (cpu.Registers.F & Z_BIT) > 0 {
		cpu.op_0xc3(nn)
	}
}

// 0xd2 - JP NC, nn
func (cpu *CPU_struct) op_0xd2(nn uint16) {
	if (cpu.Registers.F & C_BIT) == 0 {
		cpu.op_0xc3(nn)
	}
}

// 0x18 - JR n
func (cpu *CPU_struct) op_0x18(n uint8) {
	cpu.Registers.PC = uint16(int16(cpu.Registers.PC) + int16(int8(n)))
	cpu.Cycle += 8
}

// 0x20 - JP NZ, *
func (cpu *CPU_struct) op_0x20(n uint8) {
	if (cpu.Registers.F & Z_BIT) == 0 {
		cpu.Registers.PC = uint16(int16(cpu.Registers.PC) + int16(int8(n)))
	}

	// TODO: double check cpu.cycles
	cpu.Cycle += 8
}

// 0x28 - JP Z, *
func (cpu *CPU_struct) op_0x28(n uint8) {
	if (cpu.Registers.F & Z_BIT) != 0 {
		cpu.Registers.PC = uint16(int16(cpu.Registers.PC) + int16(int8(n)))
	}

	// TODO: double check cpu.cycles
	cpu.Cycle += 8
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

// 0xe2 LD (C),A
func (cpu *CPU_struct) op_0xe2() {
	cpu.MMU.WriteByte(0xff00|uint16(cpu.Registers.C), cpu.Registers.A)

	cpu.Cycle += 8
}

// 0x22 LD (HL+),A
func (cpu *CPU_struct) op_0x22() {
	var address uint16 = (uint16(cpu.Registers.H) << 8) | uint16(cpu.Registers.L)
	cpu.MMU.WriteByte(address, cpu.Registers.A)
	address += 1
	cpu.Registers.H = uint8(address >> 8)
	cpu.Registers.L = uint8(address & 0xff)

	cpu.Cycle += 8
}

// 0x32 LD (HL-),A
func (cpu *CPU_struct) op_0x32() {
	var address uint16 = (uint16(cpu.Registers.H) << 8) | uint16(cpu.Registers.L)
	cpu.MMU.WriteByte(address, cpu.Registers.A)
	address -= 1
	cpu.Registers.H = uint8(address >> 8)
	cpu.Registers.L = uint8(address & 0xff)
	cpu.Cycle += 8
}

// 0x57 LD D, A
func (cpu *CPU_struct) op_0x57() {
	cpu.Registers.D = cpu.Registers.A

	cpu.Cycle += 4
}

// 0x5f LD E, A
func (cpu *CPU_struct) op_0x5f() {
	cpu.Registers.E = cpu.Registers.A

	cpu.Cycle += 4
}

// 0x67 LD H, A
func (cpu *CPU_struct) op_0x67() {
	cpu.Registers.H = cpu.Registers.A

	cpu.Cycle += 4
}

// 0x6f LD L, A
func (cpu *CPU_struct) op_0x6f() {
	cpu.Registers.L = cpu.Registers.A

	cpu.Cycle += 4
}

// 0x77 LD (HL),A
func (cpu *CPU_struct) op_0x77() {
	var address uint16 = (uint16(cpu.Registers.H) << 8) | uint16(cpu.Registers.L)
	cpu.MMU.WriteByte(address, cpu.Registers.A)

	cpu.Cycle += 8
}

// 0x12 LD (DE),A
func (cpu *CPU_struct) op_0x12() {
	var address uint16 = (uint16(cpu.Registers.D) << 8) | uint16(cpu.Registers.E)
	cpu.MMU.WriteByte(address, cpu.Registers.A)

	cpu.Cycle += 8
}

// 0x02 LD (HL),A
func (cpu *CPU_struct) op_0x02() {
	var address uint16 = (uint16(cpu.Registers.B) << 8) | uint16(cpu.Registers.C)
	cpu.MMU.WriteByte(address, cpu.Registers.A)

	cpu.Cycle += 8
}

// 0x4f LD C,A
func (cpu *CPU_struct) op_0x4f() {
	cpu.Registers.C = cpu.Registers.A

	cpu.Cycle += 4
}

// 0x47 LD B,A
func (cpu *CPU_struct) op_0x47() {
	cpu.Registers.B = cpu.Registers.A

	cpu.Cycle += 4
}

//  0x1a LD A,(DE)
func (cpu *CPU_struct) op_0x1a() {
	var address uint16 = (uint16(cpu.Registers.D) << 8) | uint16(cpu.Registers.E)
	cpu.Registers.A = cpu.MMU.ReadByte(address)

	cpu.Cycle += 8
}

//  0xfa LD A,(nn)
func (cpu *CPU_struct) op_0xfa(nn uint16) {
	var address uint16 = nn
	cpu.Registers.A = cpu.MMU.ReadByte(address)

	cpu.Cycle += 8
}

// 0x11 LD DE,nn
func (cpu *CPU_struct) op_0x11(nn uint16) {
	cpu.Registers.D = uint8(nn >> 8)
	cpu.Registers.E = uint8(nn & 0xff)

	cpu.Cycle += 12
}

// 0x01 LD BC,nn
func (cpu *CPU_struct) op_0x01(nn uint16) {
	cpu.Registers.B = uint8(nn >> 8)
	cpu.Registers.C = uint8(nn & 0xff)

	cpu.Cycle += 12
}

// 0xea LD (nn),A
func (cpu *CPU_struct) op_0xea(nn uint16) {
	cpu.MMU.WriteByte(nn, cpu.Registers.A)

	cpu.Cycle += 16
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
	var address uint16 = (uint16(cpu.Registers.H) << 8) | uint16(cpu.Registers.L)
	cpu.Registers.A = cpu.MMU.ReadByte(address)

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
func (cpu *CPU_struct) op_0x7A() {
	cpu.Registers.A = cpu.Registers.D

	cpu.Cycle += 4
}

// 0x6b LD L, E
func (cpu *CPU_struct) op_0x6b() {
	cpu.Registers.L = cpu.Registers.E

	cpu.Cycle += 4
}

// 0x60 LD H, B
func (cpu *CPU_struct) op_0x60() {
	cpu.Registers.H = cpu.Registers.B

	cpu.Cycle += 4
}

// 0x62 LD H, D
func (cpu *CPU_struct) op_0x62() {
	cpu.Registers.H = cpu.Registers.D

	cpu.Cycle += 4
}

// 0x69 LD L, C
func (cpu *CPU_struct) op_0x69() {
	cpu.Registers.L = cpu.Registers.C

	cpu.Cycle += 4
}

// 0x44 LD B, H
func (cpu *CPU_struct) op_0x44() {
	cpu.Registers.B = cpu.Registers.H

	cpu.Cycle += 4
}

// 0x4D LD C, L
func (cpu *CPU_struct) op_0x4d() {
	cpu.Registers.C = cpu.Registers.L

	cpu.Cycle += 4
}

// 0x66 LD H, (HL)
func (cpu *CPU_struct) op_0x66() {
	var address uint16 = (uint16(cpu.Registers.H) << 8) | uint16(cpu.Registers.L)
	cpu.Registers.H = cpu.MMU.ReadByte(address)

	cpu.Cycle += 8
}

// 0x56 LD E, (HL)
func (cpu *CPU_struct) op_0x56() {
	var address uint16 = (uint16(cpu.Registers.H) << 8) | uint16(cpu.Registers.L)
	cpu.Registers.E = cpu.MMU.ReadByte(address)

	cpu.Cycle += 8
}

// 0x5e LD E, (HL)
func (cpu *CPU_struct) op_0x5e() {
	var address uint16 = (uint16(cpu.Registers.H) << 8) | uint16(cpu.Registers.L)
	cpu.Registers.E = cpu.MMU.ReadByte(address)

	cpu.Cycle += 8
}

// 0xcd CALL nn
func (cpu *CPU_struct) op_0xcd(nn uint16) {
	cpu.MMU.WriteWord(cpu.Registers.SP-2, cpu.Registers.PC)
	cpu.Registers.SP -= 2
	cpu.Registers.PC = nn

	cpu.Cycle += 12
}

// 0xaf XOR A
func (cpu *CPU_struct) op_0xaf() {
	cpu.Registers.A = 0x00
	cpu.Registers.F = Z_BIT

	cpu.Cycle += 4
}

// 0xa8 XOR B
func (cpu *CPU_struct) op_0xa8() {
	cpu.Registers.A = cpu.Registers.A ^ cpu.Registers.B

	if cpu.Registers.A == 0 {
		cpu.Registers.F = Z_BIT
	}

	cpu.Cycle += 4
}

// 0xa9 XOR C
func (cpu *CPU_struct) op_0xa9() {
	cpu.Registers.A = cpu.Registers.A ^ cpu.Registers.C

	if cpu.Registers.A == 0 {
		cpu.Registers.F = Z_BIT
	}

	cpu.Cycle += 4
}

// 0xab XOR E
func (cpu *CPU_struct) op_0xab() {
	cpu.Registers.A = cpu.Registers.A ^ cpu.Registers.E

	if cpu.Registers.A == 0 {
		cpu.Registers.F = Z_BIT
	}

	cpu.Cycle += 4
}

// 0xee XOR n
func (cpu *CPU_struct) op_0xee(n uint8) {
	cpu.Registers.A = cpu.Registers.A ^ n

	if cpu.Registers.A == 0 {
		cpu.Registers.F = Z_BIT
	}

	cpu.Cycle += 4
}

// 0xB0 OR B
func (cpu *CPU_struct) op_0xb0() {
	cpu.Registers.A |= cpu.Registers.B
	cpu.Registers.F = 0

	if cpu.Registers.A == 0 {
		cpu.Registers.F = Z_BIT
	}

	cpu.Cycle += 4
}

// 0xB1 OR C
func (cpu *CPU_struct) op_0xb1() {
	cpu.Registers.A |= cpu.Registers.C
	cpu.Registers.F = 0x00

	if cpu.Registers.A == 0 {
		cpu.Registers.F = Z_BIT
	}

	cpu.Cycle += 4
}

// 0xB2 OR D
func (cpu *CPU_struct) op_0xb2() {
	cpu.Registers.A |= cpu.Registers.D
	cpu.Registers.F = 0

	if cpu.Registers.A == 0 {
		cpu.Registers.F = Z_BIT
	}

	cpu.Cycle += 4
}

// 0xB6 OR (HL)
func (cpu *CPU_struct) op_0xb6() {
	var address uint16 = (uint16(cpu.Registers.H) << 8) | uint16(cpu.Registers.L)

	cpu.Registers.A |= cpu.MMU.ReadByte(address)
	cpu.Registers.F = 0

	if cpu.Registers.A == 0 {
		cpu.Registers.F = Z_BIT
	}

	cpu.Cycle += 8
}

// 0x2f CPL
func (cpu *CPU_struct) op_0x2f() {
	cpu.Registers.A = cpu.Registers.A ^ 0xff
	cpu.Registers.F = cpu.Registers.F | N_BIT | H_BIT

	cpu.Cycle += 4
}

// 0xfe CP #
// TODO : double check
func (cpu *CPU_struct) op_0xfe(n uint8) {
	cpu.Registers.F = 0x00
	cpu.Registers.F = cpu.Registers.F | N_BIT

	if cpu.Registers.A == n {
		cpu.Registers.F = cpu.Registers.F | Z_BIT
	}

	if (cpu.Registers.A & 0x0f) < (n & 0x0f) {
		cpu.Registers.F = cpu.Registers.F | H_BIT
	}

	if cpu.Registers.A < n {
		cpu.Registers.F = cpu.Registers.F | C_BIT
	}

	cpu.Cycle += 8
}

// 0xbe CP A,(HL)
// TODO : double check
func (cpu *CPU_struct) op_0xbe() {
	var address uint16 = (uint16(cpu.Registers.H) << 8) + uint16(cpu.Registers.L)
	var n uint8 = cpu.MMU.ReadByte(address)

	cpu.Registers.F = 0
	cpu.Registers.F = cpu.Registers.F | N_BIT

	if cpu.Registers.A == n {
		cpu.Registers.F = cpu.Registers.F | Z_BIT
	}

	if (cpu.Registers.A & 0x0f) < (n & 0x0f) {
		cpu.Registers.F = cpu.Registers.F | H_BIT
	}

	if cpu.Registers.A < n {
		cpu.Registers.F = cpu.Registers.F | C_BIT
	}

	cpu.Cycle += 8
}

// 0xb8 CP A,B
// TODO : double check
func (cpu *CPU_struct) op_0xb8() {
	var n uint8 = cpu.Registers.B

	cpu.Registers.F = 0
	cpu.Registers.F = cpu.Registers.F | N_BIT

	if cpu.Registers.A == n {
		cpu.Registers.F = cpu.Registers.F | Z_BIT
	}

	if (cpu.Registers.A & 0x0f) < (n & 0x0f) {
		cpu.Registers.F = cpu.Registers.F | H_BIT
	}

	if cpu.Registers.A < n {
		cpu.Registers.F = cpu.Registers.F | C_BIT
	}

	cpu.Cycle += 4
}

// 0xb9 CP A,C
// TODO : double check
func (cpu *CPU_struct) op_0xb9() {
	var n uint8 = cpu.Registers.C

	cpu.Registers.F = 0
	cpu.Registers.F = cpu.Registers.F | N_BIT

	if cpu.Registers.A == n {
		cpu.Registers.F = cpu.Registers.F | Z_BIT
	}

	if (cpu.Registers.A & 0x0f) < (n & 0x0f) {
		cpu.Registers.F = cpu.Registers.F | H_BIT
	}

	if cpu.Registers.A < n {
		cpu.Registers.F = cpu.Registers.F | C_BIT
	}

	cpu.Cycle += 4
}

// 0x05 DEC B
func (cpu *CPU_struct) op_0x05() {
	cpu.Registers.B -= 1
	cpu.Registers.F = cpu.Registers.F & C_BIT
	cpu.Registers.F = cpu.Registers.F | N_BIT

	if cpu.Registers.B == 0 {
		cpu.Registers.F = cpu.Registers.F | Z_BIT
	}

	if ((cpu.Registers.B + 1) & 0xf) != 0 {
		cpu.Registers.F = cpu.Registers.F | H_BIT
	}

	cpu.Cycle += 4
}

// 0x15 DEC D
func (cpu *CPU_struct) op_0x15() {
	cpu.Registers.D -= 1
	cpu.Registers.F = cpu.Registers.F & C_BIT
	cpu.Registers.F = cpu.Registers.F | N_BIT

	if cpu.Registers.D == 0 {
		cpu.Registers.F = cpu.Registers.F | Z_BIT
	}

	if ((cpu.Registers.D + 1) & 0xf) != 0 {
		cpu.Registers.F = cpu.Registers.F | H_BIT
	}

	cpu.Cycle += 4
}

// 0x1D DEC E
func (cpu *CPU_struct) op_0x1d() {
	cpu.Registers.E -= 1
	cpu.Registers.F = cpu.Registers.F & C_BIT
	cpu.Registers.F = cpu.Registers.F | N_BIT

	if cpu.Registers.E == 0 {
		cpu.Registers.F = cpu.Registers.F | Z_BIT
	}

	if ((cpu.Registers.E + 1) & 0xf) != 0 {
		cpu.Registers.F = cpu.Registers.F | H_BIT
	}

	cpu.Cycle += 4
}

// 0x25 DEC H
func (cpu *CPU_struct) op_0x25() {
	cpu.Registers.H -= 1
	cpu.Registers.F = cpu.Registers.F & C_BIT
	cpu.Registers.F = cpu.Registers.F | N_BIT

	if cpu.Registers.H == 0 {
		cpu.Registers.F = cpu.Registers.F | Z_BIT
	}

	if ((cpu.Registers.H + 1) & 0xf) != 0 {
		cpu.Registers.F = cpu.Registers.F | H_BIT
	}

	cpu.Cycle += 4
}

// 0x3d DEC A
func (cpu *CPU_struct) op_0x3d() {
	cpu.Registers.A -= 1
	cpu.Registers.F = cpu.Registers.F & C_BIT
	cpu.Registers.F = cpu.Registers.F | N_BIT

	if cpu.Registers.A == 0 {
		cpu.Registers.F = cpu.Registers.F | Z_BIT
	}

	if ((cpu.Registers.A + 1) & 0xf) != 0 {
		cpu.Registers.F = cpu.Registers.F | H_BIT
	}

	cpu.Cycle += 4
}

// 0x0d DEC C
func (cpu *CPU_struct) op_0x0d() {
	cpu.Registers.C -= 1
	cpu.Registers.F = cpu.Registers.F & C_BIT
	cpu.Registers.F = cpu.Registers.F | N_BIT

	if cpu.Registers.C == 0 {
		cpu.Registers.F = cpu.Registers.F | Z_BIT
	}

	if ((cpu.Registers.C + 1) & 0xf) != 0 {
		cpu.Registers.F = cpu.Registers.F | H_BIT
	}

	cpu.Cycle += 4
}

// 0x0B DEC BC
func (cpu *CPU_struct) op_0x0b() {
	var bc uint16 = (uint16(cpu.Registers.B) << 8) | uint16(cpu.Registers.C)
	bc -= 1
	cpu.Registers.B = uint8(bc >> 8)
	cpu.Registers.C = uint8(bc & 0xff)

	cpu.Cycle += 8
}

// 0x1B DEC DE
func (cpu *CPU_struct) op_0x1b() {
	var de uint16 = (uint16(cpu.Registers.D) << 8) | uint16(cpu.Registers.E)
	de -= 1
	cpu.Registers.D = uint8(de >> 8)
	cpu.Registers.E = uint8(de & 0xff)

	cpu.Cycle += 8
}

// 0x2B DEC HL
func (cpu *CPU_struct) op_0x2b() {
	var hl uint16 = (uint16(cpu.Registers.H) << 8) | uint16(cpu.Registers.L)
	hl -= 1
	cpu.Registers.H = uint8(hl >> 8)
	cpu.Registers.L = uint8(hl & 0xff)

	cpu.Cycle += 8
}

//0x04 INC B
func (cpu *CPU_struct) op_0x04() {
	cpu.Registers.B += 1
	cpu.Registers.F &= C_BIT

	if cpu.Registers.B == 0 {
		cpu.Registers.F |= Z_BIT
	}

	if (cpu.Registers.B & 0x0f) == 0x0f {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 4
}

//0x0c INC C
func (cpu *CPU_struct) op_0x0c() {
	cpu.Registers.C += 1
	cpu.Registers.F &= C_BIT

	if cpu.Registers.C == 0 {
		cpu.Registers.F |= Z_BIT
	}

	if (cpu.Registers.C & 0x0f) == 0x0f {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 4
}

//0x24 INC H
func (cpu *CPU_struct) op_0x24() {
	cpu.Registers.H += 1
	cpu.Registers.F &= C_BIT

	if cpu.Registers.H == 0 {
		cpu.Registers.F |= Z_BIT
	}

	if (cpu.Registers.H & 0x0f) == 0x0f {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 4
}

//0x2c INC L
func (cpu *CPU_struct) op_0x2c() {
	cpu.Registers.L += 1
	cpu.Registers.F &= C_BIT

	if cpu.Registers.L == 0 {
		cpu.Registers.F |= Z_BIT
	}

	if (cpu.Registers.L & 0x0f) == 0x0f {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 4
}

//0x34 INC (HL)
func (cpu *CPU_struct) op_0x34() {
	var address uint16 = (uint16(cpu.Registers.H) << 8) | uint16(cpu.Registers.L)
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.MMU.WriteByte(address, n+1)

	cpu.Registers.F &= C_BIT

	if n == 0 {
		cpu.Registers.F |= Z_BIT
	}

	if (n & 0x0f) == 0x0f {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 12
}

//0x3c INC A
func (cpu *CPU_struct) op_0x3c() {
	cpu.Registers.A += 1
	cpu.Registers.F &= C_BIT

	if cpu.Registers.A == 0 {
		cpu.Registers.F |= Z_BIT
	}

	if (cpu.Registers.A & 0x0f) == 0x0f {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 4
}

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
	var carry uint8 = cpu.Registers.F & C_BIT
	cpu.Registers.F = 0
	cpu.Registers.F |= ((cpu.Registers.A & 0x80) >> 3)

	cpu.Registers.A = (cpu.Registers.A << 1) & 0xff
	if carry != 0 {
		cpu.Registers.A |= 0x01
	}

	if cpu.Registers.A == 0 {
		cpu.Registers.F |= Z_BIT
	}

	cpu.Cycle += 8
}

// 0x17 RLA
func (cpu *CPU_struct) op_0x17() {
	cpu.op_cb_0x07()
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

// 0xcb 0x87 RES b, A
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

// 0x90 SUB B
func (cpu *CPU_struct) op_0x90() {
	old := cpu.Registers.A
	cpu.Registers.A = cpu.Registers.A - cpu.Registers.B
	cpu.Registers.F = N_BIT

	if cpu.Registers.A == 0 {
		cpu.Registers.F |= Z_BIT
	}

	if (old & 0x0f) > (cpu.Registers.B & 0x0f) {
		cpu.Registers.F |= H_BIT
	}

	if old > cpu.Registers.B {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 4
}

// 0x97 SUB A
func (cpu *CPU_struct) op_0x97() {
	old := cpu.Registers.A
	cpu.Registers.A = 0
	cpu.Registers.F = N_BIT

	if cpu.Registers.A == 0 {
		cpu.Registers.F |= Z_BIT
	}

	if (old & 0x0f) > (cpu.Registers.A & 0x0f) {
		cpu.Registers.F |= H_BIT
	}

	if old > cpu.Registers.A {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 4
}

// 0x95 SUB L
func (cpu *CPU_struct) op_0x95() {
	old := cpu.Registers.A
	cpu.Registers.A = cpu.Registers.A - cpu.Registers.L
	cpu.Registers.F = N_BIT

	if cpu.Registers.A == 0 {
		cpu.Registers.F |= Z_BIT
	}

	if (old & 0x0f) > (cpu.Registers.L & 0x0f) {
		cpu.Registers.F |= H_BIT
	}

	if old > cpu.Registers.L {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 4
}

// 0x00 NOP
func (cpu *CPU_struct) op_0x00() {
	cpu.Cycle += 4
}

// 0xc5 PUSH BC
func (cpu *CPU_struct) op_0xc5() {
	var bc uint16 = (uint16(cpu.Registers.C) << 8) | uint16(cpu.Registers.B)
	cpu.Registers.SP -= 2
	cpu.MMU.WriteWord(cpu.Registers.SP, bc)

	cpu.Cycle += 16
}

// 0xd5 PUSH DE
func (cpu *CPU_struct) op_0xd5() {
	var de uint16 = (uint16(cpu.Registers.E) << 8) | uint16(cpu.Registers.D)
	cpu.Registers.SP -= 2
	cpu.MMU.WriteWord(cpu.Registers.SP, de)

	cpu.Cycle += 16
}

// 0xe5 PUSH HL
func (cpu *CPU_struct) op_0xe5() {
	var hl uint16 = (uint16(cpu.Registers.L) << 8) | uint16(cpu.Registers.H)
	cpu.Registers.SP -= 2
	cpu.MMU.WriteWord(cpu.Registers.SP, hl)

	cpu.Cycle += 16
}

// 0xf5 PUSH AF
func (cpu *CPU_struct) op_0xf5() {
	var af uint16 = (uint16(cpu.Registers.F) << 8) | uint16(cpu.Registers.A)
	cpu.Registers.SP -= 2
	cpu.MMU.WriteWord(cpu.Registers.SP, af)

	cpu.Cycle += 16
}

// 0xc1 POP BC
func (cpu *CPU_struct) op_0xc1() {
	var data uint16 = cpu.MMU.ReadWord(cpu.Registers.SP)
	cpu.Registers.C = uint8(data >> 8)
	cpu.Registers.B = uint8(data & 0xff)
	cpu.Registers.SP += 2

	cpu.Cycle += 12
}

// 0xd1 POP DE
func (cpu *CPU_struct) op_0xd1() {
	var data uint16 = cpu.MMU.ReadWord(cpu.Registers.SP)
	cpu.Registers.E = uint8(data >> 8)
	cpu.Registers.D = uint8(data & 0xff)
	cpu.Registers.SP += 2

	cpu.Cycle += 12
}

// 0xe1 POP HL
func (cpu *CPU_struct) op_0xe1() {
	var data uint16 = cpu.MMU.ReadWord(cpu.Registers.SP)
	cpu.Registers.L = uint8(data >> 8)
	cpu.Registers.H = uint8(data & 0xff)
	cpu.Registers.SP += 2

	cpu.Cycle += 12
}

// 0xf1 POP AF
func (cpu *CPU_struct) op_0xf1() {
	var data uint16 = cpu.MMU.ReadWord(cpu.Registers.SP)
	cpu.Registers.F = uint8(data >> 8)
	cpu.Registers.A = uint8(data & 0xff)
	cpu.Registers.SP += 2

	cpu.Cycle += 12
}

// 0x13 INC DE
func (cpu *CPU_struct) op_0x13() {
	var de uint16 = (uint16(cpu.Registers.D) << 8) + uint16(cpu.Registers.E)
	de += 1
	cpu.Registers.D = uint8(de >> 8)
	cpu.Registers.E = uint8(de & 0xff)

	cpu.Cycle += 8
}

// 0x03 INC BC
func (cpu *CPU_struct) op_0x03() {
	var bc uint16 = (uint16(cpu.Registers.B) << 8) + uint16(cpu.Registers.C)
	bc += 1
	cpu.Registers.B = uint8(bc >> 8)
	cpu.Registers.C = uint8(bc & 0xff)

	cpu.Cycle += 8
}

// 0x23 INC HL
func (cpu *CPU_struct) op_0x23() {
	var hl uint16 = (uint16(cpu.Registers.H) << 8) + uint16(cpu.Registers.L)
	hl += 1
	cpu.Registers.H = uint8(hl >> 8)
	cpu.Registers.L = uint8(hl & 0xff)

	cpu.Cycle += 8
}

// 0xc9 RET
func (cpu *CPU_struct) op_0xc9() {
	var addr uint16 = cpu.MMU.ReadWord(cpu.Registers.SP)
	cpu.Registers.SP += 2
	cpu.Registers.PC = addr

	cpu.Cycle += 8
}

// 0xc8 RET Z
func (cpu *CPU_struct) op_0xc8() {
	if (cpu.Registers.F & Z_BIT) > 0 {
		cpu.op_0xc9()
	}
}

// 0xc0 RET NZ
func (cpu *CPU_struct) op_0xc0() {
	if (cpu.Registers.F & Z_BIT) == 0 {
		cpu.op_0xc9()
	}
}

// 0xd9 RETI
func (cpu *CPU_struct) op_0xd9() {
	var addr uint16 = cpu.MMU.ReadWord(cpu.Registers.SP)
	cpu.Registers.SP += 2
	cpu.Registers.PC = addr

	cpu.Registers.IME = 1

	cpu.Cycle += 8
}

// 0x87 - ADD A, A
func (cpu *CPU_struct) op_0x87() {
	var n uint8 = cpu.Registers.A

	old_a := cpu.Registers.A
	cpu.Registers.A += n
	cpu.Registers.F = 0

	if cpu.Registers.A == 0 {
		cpu.Registers.F |= Z_BIT
	}

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
	var n uint8 = cpu.Registers.B

	old_a := cpu.Registers.A
	cpu.Registers.A += n
	cpu.Registers.F = 0

	if cpu.Registers.A == 0 {
		cpu.Registers.F |= Z_BIT
	}

	if (((old_a & 0x0f) + (n & 0x0f)) & 0xf0) != 0 {
		cpu.Registers.F |= H_BIT
	}

	if (((uint16(old_a)) + (uint16(n))) & 0xff00) != 0 {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 4
}

// 0x86 - ADD A, (HL)
func (cpu *CPU_struct) op_0x86() {
	var address uint16 = (uint16(cpu.Registers.H) << 8) + uint16(cpu.Registers.L)
	var n uint8 = cpu.MMU.ReadByte(address)

	old_a := cpu.Registers.A
	cpu.Registers.A += n
	cpu.Registers.F = 0

	if cpu.Registers.A == 0 {
		cpu.Registers.F |= Z_BIT
	}

	if (((old_a & 0x0f) + (n & 0x0f)) & 0xf0) != 0 {
		cpu.Registers.F |= H_BIT
	}

	if (((uint16(old_a)) + (uint16(n))) & 0xff00) != 0 {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 8
}

// 0xE8 ADD SP, n
// TODO double check
func (cpu *CPU_struct) op_0xe8(n uint8) {
	old_sp := cpu.Registers.SP
	cpu.Registers.SP = uint16(int16(cpu.Registers.SP) + int16(n))
	cpu.Registers.F = 0

	// if cpu.Registers.SP == 0 {
	// 	cpu.Registers.F = cpu.Registers.F | 0x80
	// }

	if (((old_sp & 0x0f) + (uint16(n) & 0x0f)) & 0xf0) != 0 {
		cpu.Registers.F |= H_BIT
	}

	if (((old_sp & 0x00ff) + (uint16(n) & 0x00ff)) & 0xff00) != 0 {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 16
}

// 0x09 ADD HL, BC
func (cpu *CPU_struct) op_0x09() {
	var hl uint16 = (uint16(cpu.Registers.H) << 8) + uint16(cpu.Registers.L)
	var bc uint16 = (uint16(cpu.Registers.B) << 8) + uint16(cpu.Registers.C)
	var sum uint16 = hl + bc
	cpu.Registers.H = uint8(sum >> 8)
	cpu.Registers.L = uint8(sum & 0xff)

	cpu.Registers.F &= Z_BIT

	if (((hl & 0x0fff) + (bc & 0x0fff)) & 0xf000) != 0 {
		cpu.Registers.F |= H_BIT
	}

	if ((uint32(hl) + uint32(bc)) & 0xffff) != 0 {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 8
}

// 0x19 ADD HL, DE
func (cpu *CPU_struct) op_0x19() {
	var hl uint16 = (uint16(cpu.Registers.H) << 8) + uint16(cpu.Registers.L)
	var de uint16 = (uint16(cpu.Registers.D) << 8) + uint16(cpu.Registers.E)
	var sum uint16 = hl + de
	cpu.Registers.H = uint8(sum >> 8)
	cpu.Registers.L = uint8(sum & 0xff)

	cpu.Registers.F &= Z_BIT

	if (((hl & 0x0fff) + (de & 0x0fff)) & 0xf000) != 0 {
		cpu.Registers.F |= H_BIT
	}

	if ((uint32(hl) + uint32(de)) & 0xffff) != 0 {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 8
}

// 0xf3 DI
func (cpu *CPU_struct) op_0xf3() {
	cpu.Registers.IME = 0

	cpu.Cycle += 4
}

// 0xfb EI
func (cpu *CPU_struct) op_0xfb() {
	cpu.Registers.IME = 1

	cpu.Cycle += 4
}

// 0xe6 AND #
func (cpu *CPU_struct) op_0xe6(n uint8) {
	cpu.Registers.A &= n

	cpu.Registers.F = H_BIT

	if cpu.Registers.A == 0 {
		cpu.Registers.F |= Z_BIT
	}

	cpu.Cycle += 8
}

// 0xa7 AND A
func (cpu *CPU_struct) op_0xa7() {
	// cpu.Registers.A &= cpu.Registers.A

	cpu.Registers.F = H_BIT

	if cpu.Registers.A == 0 {
		cpu.Registers.F |= Z_BIT
	}

	cpu.Cycle += 8
}

// 0xa1 AND C
func (cpu *CPU_struct) op_0xa1() {
	cpu.Registers.A &= cpu.Registers.C

	cpu.Registers.F = H_BIT

	if cpu.Registers.A == 0 {
		cpu.Registers.F |= Z_BIT
	}

	cpu.Cycle += 8
}

// 0x2A LDI A, (HL)
func (cpu *CPU_struct) op_0x2a() {
	var address uint16 = (uint16(cpu.Registers.H) << 8) + uint16(cpu.Registers.L)
	cpu.Registers.A = cpu.MMU.ReadByte(address)

	// INC HL
	cpu.op_0x23()

	// cpu.Cycle += 8
}

// 0x76 HALT
func (cpu *CPU_struct) op_0x76() {
	cpu.Cycle += 4
}

// 0xef RST 0x28
func (cpu *CPU_struct) op_0xef() {
	cpu.Registers.SP -= 2
	cpu.MMU.WriteWord(cpu.Registers.SP, cpu.Registers.PC)
	cpu.Registers.PC = 0x0028

	cpu.Cycle += 32
}

// 0x89 - ADC A, C
func (cpu *CPU_struct) op_0x89() {
	var n uint8 = cpu.Registers.C
	var carry uint8 = 0

	if (cpu.Registers.F & C_BIT) > 0 {
		carry = 1
	}

	old_a := cpu.Registers.A
	cpu.Registers.A += n + carry
	cpu.Registers.F = 0

	if cpu.Registers.A == 0 {
		cpu.Registers.F |= Z_BIT
	}

	if (((old_a & 0x0f) + ((n + carry) & 0x0f)) & 0xf0) != 0 {
		cpu.Registers.F |= H_BIT
	}

	if (((uint16(old_a)) + (uint16(n + carry))) & 0xff00) != 0 {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 4
}

// 0x99 - SBC A, C
// TODO double check
func (cpu *CPU_struct) op_0x99() {
	var n uint8 = cpu.Registers.C
	var carry uint8 = 0

	if (cpu.Registers.F & C_BIT) > 0 {
		carry += 1
	}

	old := cpu.Registers.A
	cpu.Registers.A = cpu.Registers.A - n - carry
	cpu.Registers.F = N_BIT

	if cpu.Registers.A == 0 {
		cpu.Registers.F |= Z_BIT
	}

	if (old & 0x0f) > ((n + carry) & 0x0f) {
		cpu.Registers.F |= H_BIT
	}

	if old > (n + carry) {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 4
}

// 0xde - SBC A, n
// TODO double check
func (cpu *CPU_struct) op_0xde(n uint8) {
	if (cpu.Registers.F & C_BIT) != 0 {
		n += 1
	}

	old := cpu.Registers.A
	cpu.Registers.A = cpu.Registers.A - n
	cpu.Registers.F = N_BIT

	if cpu.Registers.A == 0 {
		cpu.Registers.F |= Z_BIT
	}

	if (old & 0x0f) > (n & 0x0f) {
		cpu.Registers.F |= H_BIT
	}

	if old > n {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 4
}
