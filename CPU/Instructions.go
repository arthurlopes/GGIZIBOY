package CPU

const Z_BIT = 0x80
const N_BIT = 0x40
const H_BIT = 0x20
const C_BIT = 0x10

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
	// ADD
	cpu.Instructions_maps.Instructions_map_uint8[0xe8] = cpu.op_0xe8

	// 16
	//LD
	cpu.Instructions_maps.Instructions_map_uint16[0x31] = cpu.op_0x31
	cpu.Instructions_maps.Instructions_map_uint16[0x21] = cpu.op_0x21
	cpu.Instructions_maps.Instructions_map_uint16[0xea] = cpu.op_0xea
	cpu.Instructions_maps.Instructions_map_uint16[0x11] = cpu.op_0x11
	cpu.Instructions_maps.Instructions_map_uint16[0x01] = cpu.op_0x01
	//JP
	cpu.Instructions_maps.Instructions_map_uint16[0xc3] = cpu.op_0xc3
	//CALL
	cpu.Instructions_maps.Instructions_map_uint16[0xcd] = cpu.op_0xcd

	// 0
	//LD
	cpu.Instructions_maps.Instructions_map_[0x22] = cpu.op_0x22
	cpu.Instructions_maps.Instructions_map_[0x32] = cpu.op_0x32
	cpu.Instructions_maps.Instructions_map_[0xe2] = cpu.op_0xe2
	cpu.Instructions_maps.Instructions_map_[0x57] = cpu.op_0x57
	cpu.Instructions_maps.Instructions_map_[0x67] = cpu.op_0x67
	cpu.Instructions_maps.Instructions_map_[0x77] = cpu.op_0x77
	cpu.Instructions_maps.Instructions_map_[0x02] = cpu.op_0x02
	cpu.Instructions_maps.Instructions_map_[0x4f] = cpu.op_0x4f
	cpu.Instructions_maps.Instructions_map_[0x1a] = cpu.op_0x1a
	cpu.Instructions_maps.Instructions_map_[0x7b] = cpu.op_0x7b
	cpu.Instructions_maps.Instructions_map_[0x7c] = cpu.op_0x7c
	cpu.Instructions_maps.Instructions_map_[0x7d] = cpu.op_0x7d
	cpu.Instructions_maps.Instructions_map_[0x78] = cpu.op_0x78
	cpu.Instructions_maps.Instructions_map_[0x47] = cpu.op_0x47
	cpu.Instructions_maps.Instructions_map_[0x2A] = cpu.op_0x2a
	// Bitwise
	cpu.Instructions_maps.Instructions_map_[0xaf] = cpu.op_0xaf
	cpu.Instructions_maps.Instructions_map_[0xb1] = cpu.op_0xb1
	cpu.Instructions_maps.Instructions_map_[0xa8] = cpu.op_0xa8
	// CP
	cpu.Instructions_maps.Instructions_map_[0x2f] = cpu.op_0x2f
	cpu.Instructions_maps.Instructions_map_[0xbe] = cpu.op_0xbe
	// DEC
	cpu.Instructions_maps.Instructions_map_[0x05] = cpu.op_0x05
	cpu.Instructions_maps.Instructions_map_[0x15] = cpu.op_0x15
	cpu.Instructions_maps.Instructions_map_[0x1d] = cpu.op_0x1d
	cpu.Instructions_maps.Instructions_map_[0x3d] = cpu.op_0x3d
	cpu.Instructions_maps.Instructions_map_[0x0d] = cpu.op_0x0d
	cpu.Instructions_maps.Instructions_map_[0x0b] = cpu.op_0x0b
	// INC
	cpu.Instructions_maps.Instructions_map_[0x04] = cpu.op_0x04
	cpu.Instructions_maps.Instructions_map_[0x0c] = cpu.op_0x0c
	cpu.Instructions_maps.Instructions_map_[0x24] = cpu.op_0x24
	cpu.Instructions_maps.Instructions_map_[0x03] = cpu.op_0x03
	cpu.Instructions_maps.Instructions_map_[0x13] = cpu.op_0x13
	cpu.Instructions_maps.Instructions_map_[0x23] = cpu.op_0x23
	// SUB
	cpu.Instructions_maps.Instructions_map_[0x95] = cpu.op_0x95
	cpu.Instructions_maps.Instructions_map_[0x90] = cpu.op_0x90
	cpu.Instructions_maps.Instructions_map_[0x97] = cpu.op_0x97
	// NOP
	cpu.Instructions_maps.Instructions_map_[0x00] = cpu.op_0x00
	// PUSH/POP
	cpu.Instructions_maps.Instructions_map_[0xc5] = cpu.op_0xc5
	cpu.Instructions_maps.Instructions_map_[0xc1] = cpu.op_0xc1
	// RL
	cpu.Instructions_maps.Instructions_map_[0x17] = cpu.op_0x17
	// RET
	cpu.Instructions_maps.Instructions_map_[0xc9] = cpu.op_0xc9
	// ADD
	cpu.Instructions_maps.Instructions_map_[0x86] = cpu.op_0x86
	// DI
	cpu.Instructions_maps.Instructions_map_[0xf3] = cpu.op_0xf3

	// CB
	cpu.Instructions_maps.Instructions_map_CB[0x11] = cpu.op_cb_0x11
	cpu.Instructions_maps.Instructions_map_CB[0x07] = cpu.op_cb_0x07
	cpu.Instructions_maps.Instructions_map_CB[0x7c] = cpu.op_cb_0x7c
	cpu.Instructions_maps.Instructions_map_CB[0x87] = cpu.op_cb_0x87

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

// 0xc3 - JP nn
func (cpu *CPU_struct) op_0xc3(nn uint16) {
	cpu.Registers.PC = nn
	cpu.Cycle += 12
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

// 0x67 LD H, A
func (cpu *CPU_struct) op_0x67() {
	cpu.Registers.H = cpu.Registers.A

	cpu.Cycle += 4
}

// 0x77 LD (HL),A
func (cpu *CPU_struct) op_0x77() {
	var address uint16 = (uint16(cpu.Registers.H) << 8) | uint16(cpu.Registers.L)
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

// 0x78 LD A, B
func (cpu *CPU_struct) op_0x78() {
	cpu.Registers.A = cpu.Registers.B

	cpu.Cycle += 4
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

// 0xB1 OR C
func (cpu *CPU_struct) op_0xb1() {
	cpu.Registers.A |= cpu.Registers.C
	cpu.Registers.F = 0x00

	if cpu.Registers.A == 0 {
		cpu.Registers.F = Z_BIT
	}

	cpu.Cycle += 4
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
	cpu.Registers.H |= 0b10000000
	cpu.Registers.F &= C_BIT
	cpu.Registers.F |= H_BIT

	if cpu.Registers.H == 0 {
		cpu.Registers.F |= Z_BIT
	}

	cpu.Cycle += 8
}

// 0xcb 0x87 RES b, A
func (cpu *CPU_struct) op_cb_0x87() {
	cpu.Registers.A &= 0b11111110

	// TODO review 4 or 8, is 0xCB included here?
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

// 0xc1 POP BC
func (cpu *CPU_struct) op_0xc1() {
	var data uint16 = cpu.MMU.ReadWord(cpu.Registers.SP)
	cpu.Registers.C = uint8(data >> 8)
	cpu.Registers.B = uint8(data & 0xff)
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

	if ((old_a&0x0f)+(n&0x0f))&0xf0 != 0 {
		cpu.Registers.F |= H_BIT
	}

	if ((uint16(old_a))+(uint16(n)))&0xff00 != 0 {
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

	if ((old_sp&0x0f)+(uint16(n)&0x0f))&0xf0 != 0 {
		cpu.Registers.F |= H_BIT
	}

	if ((old_sp&0x00ff)+(uint16(n)&0x00ff))&0xff00 != 0 {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 16
}

// 0xf3 DI
func (cpu *CPU_struct) op_0xf3() {
	cpu.Registers.IME = 0

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

// 0x2A LDI A, (HL)
func (cpu *CPU_struct) op_0x2a() {
	var address uint16 = (uint16(cpu.Registers.H) << 8) + uint16(cpu.Registers.L)
	cpu.Registers.A = cpu.MMU.ReadByte(address)

	// INC HL
	cpu.op_0x23()

	// cpu.Cycle += 8
}
