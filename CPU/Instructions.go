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

	cpu.Instructions_maps.Instructions_map_str = make(map[uint8]string)
	cpu.Instructions_maps.Instructions_map_CB_str = make(map[uint8]string)

	cpu.Instructions_maps.Instructions_map_str[0x00] = "NOP"
	cpu.Instructions_maps.Instructions_map_str[0x01] = "LD BC, d16"
	cpu.Instructions_maps.Instructions_map_str[0x02] = "LD BC, A"
	cpu.Instructions_maps.Instructions_map_str[0x03] = "INC BC"
	cpu.Instructions_maps.Instructions_map_str[0x04] = "INC B"
	cpu.Instructions_maps.Instructions_map_str[0x05] = "DEC B"
	cpu.Instructions_maps.Instructions_map_str[0x06] = "LD B, d8"
	cpu.Instructions_maps.Instructions_map_str[0x07] = "RLCA"
	cpu.Instructions_maps.Instructions_map_str[0x08] = "LD a16, SP"
	cpu.Instructions_maps.Instructions_map_str[0x09] = "ADD HL, BC"
	cpu.Instructions_maps.Instructions_map_str[0x0A] = "LD A, BC"
	cpu.Instructions_maps.Instructions_map_str[0x0B] = "DEC BC"
	cpu.Instructions_maps.Instructions_map_str[0x0C] = "INC C"
	cpu.Instructions_maps.Instructions_map_str[0x0D] = "DEC C"
	cpu.Instructions_maps.Instructions_map_str[0x0E] = "LD C, d8"
	cpu.Instructions_maps.Instructions_map_str[0x0F] = "RRCA"
	cpu.Instructions_maps.Instructions_map_str[0x10] = "STOP d8"
	cpu.Instructions_maps.Instructions_map_str[0x11] = "LD DE, d16"
	cpu.Instructions_maps.Instructions_map_str[0x12] = "LD DE, A"
	cpu.Instructions_maps.Instructions_map_str[0x13] = "INC DE"
	cpu.Instructions_maps.Instructions_map_str[0x14] = "INC D"
	cpu.Instructions_maps.Instructions_map_str[0x15] = "DEC D"
	cpu.Instructions_maps.Instructions_map_str[0x16] = "LD D, d8"
	cpu.Instructions_maps.Instructions_map_str[0x17] = "RLA"
	cpu.Instructions_maps.Instructions_map_str[0x18] = "JR r8"
	cpu.Instructions_maps.Instructions_map_str[0x19] = "ADD HL, DE"
	cpu.Instructions_maps.Instructions_map_str[0x1A] = "LD A, DE"
	cpu.Instructions_maps.Instructions_map_str[0x1B] = "DEC DE"
	cpu.Instructions_maps.Instructions_map_str[0x1C] = "INC E"
	cpu.Instructions_maps.Instructions_map_str[0x1D] = "DEC E"
	cpu.Instructions_maps.Instructions_map_str[0x1E] = "LD E, d8"
	cpu.Instructions_maps.Instructions_map_str[0x1F] = "RRA"
	cpu.Instructions_maps.Instructions_map_str[0x20] = "JR NZ, r8"
	cpu.Instructions_maps.Instructions_map_str[0x21] = "LD HL, d16"
	cpu.Instructions_maps.Instructions_map_str[0x22] = "LD HL, A"
	cpu.Instructions_maps.Instructions_map_str[0x23] = "INC HL"
	cpu.Instructions_maps.Instructions_map_str[0x24] = "INC H"
	cpu.Instructions_maps.Instructions_map_str[0x25] = "DEC H"
	cpu.Instructions_maps.Instructions_map_str[0x26] = "LD H, d8"
	cpu.Instructions_maps.Instructions_map_str[0x27] = "DAA"
	cpu.Instructions_maps.Instructions_map_str[0x28] = "JR Z, r8"
	cpu.Instructions_maps.Instructions_map_str[0x29] = "ADD HL, HL"
	cpu.Instructions_maps.Instructions_map_str[0x2A] = "LD A, HL"
	cpu.Instructions_maps.Instructions_map_str[0x2B] = "DEC HL"
	cpu.Instructions_maps.Instructions_map_str[0x2C] = "INC L"
	cpu.Instructions_maps.Instructions_map_str[0x2D] = "DEC L"
	cpu.Instructions_maps.Instructions_map_str[0x2E] = "LD L, d8"
	cpu.Instructions_maps.Instructions_map_str[0x2F] = "CPL"
	cpu.Instructions_maps.Instructions_map_str[0x30] = "JR NC, r8"
	cpu.Instructions_maps.Instructions_map_str[0x31] = "LD SP, d16"
	cpu.Instructions_maps.Instructions_map_str[0x32] = "LD HL, A"
	cpu.Instructions_maps.Instructions_map_str[0x33] = "INC SP"
	cpu.Instructions_maps.Instructions_map_str[0x34] = "INC HL"
	cpu.Instructions_maps.Instructions_map_str[0x35] = "DEC HL"
	cpu.Instructions_maps.Instructions_map_str[0x36] = "LD HL, d8"
	cpu.Instructions_maps.Instructions_map_str[0x37] = "SCF"
	cpu.Instructions_maps.Instructions_map_str[0x38] = "JR C, r8"
	cpu.Instructions_maps.Instructions_map_str[0x39] = "ADD HL, SP"
	cpu.Instructions_maps.Instructions_map_str[0x3A] = "LD A, HL"
	cpu.Instructions_maps.Instructions_map_str[0x3B] = "DEC SP"
	cpu.Instructions_maps.Instructions_map_str[0x3C] = "INC A"
	cpu.Instructions_maps.Instructions_map_str[0x3D] = "DEC A"
	cpu.Instructions_maps.Instructions_map_str[0x3E] = "LD A, d8"
	cpu.Instructions_maps.Instructions_map_str[0x3F] = "CCF"
	cpu.Instructions_maps.Instructions_map_str[0x40] = "LD B, B"
	cpu.Instructions_maps.Instructions_map_str[0x41] = "LD B, C"
	cpu.Instructions_maps.Instructions_map_str[0x42] = "LD B, D"
	cpu.Instructions_maps.Instructions_map_str[0x43] = "LD B, E"
	cpu.Instructions_maps.Instructions_map_str[0x44] = "LD B, H"
	cpu.Instructions_maps.Instructions_map_str[0x45] = "LD B, L"
	cpu.Instructions_maps.Instructions_map_str[0x46] = "LD B, HL"
	cpu.Instructions_maps.Instructions_map_str[0x47] = "LD B, A"
	cpu.Instructions_maps.Instructions_map_str[0x48] = "LD C, B"
	cpu.Instructions_maps.Instructions_map_str[0x49] = "LD C, C"
	cpu.Instructions_maps.Instructions_map_str[0x4A] = "LD C, D"
	cpu.Instructions_maps.Instructions_map_str[0x4B] = "LD C, E"
	cpu.Instructions_maps.Instructions_map_str[0x4C] = "LD C, H"
	cpu.Instructions_maps.Instructions_map_str[0x4D] = "LD C, L"
	cpu.Instructions_maps.Instructions_map_str[0x4E] = "LD C, HL"
	cpu.Instructions_maps.Instructions_map_str[0x4F] = "LD C, A"
	cpu.Instructions_maps.Instructions_map_str[0x50] = "LD D, B"
	cpu.Instructions_maps.Instructions_map_str[0x51] = "LD D, C"
	cpu.Instructions_maps.Instructions_map_str[0x52] = "LD D, D"
	cpu.Instructions_maps.Instructions_map_str[0x53] = "LD D, E"
	cpu.Instructions_maps.Instructions_map_str[0x54] = "LD D, H"
	cpu.Instructions_maps.Instructions_map_str[0x55] = "LD D, L"
	cpu.Instructions_maps.Instructions_map_str[0x56] = "LD D, HL"
	cpu.Instructions_maps.Instructions_map_str[0x57] = "LD D, A"
	cpu.Instructions_maps.Instructions_map_str[0x58] = "LD E, B"
	cpu.Instructions_maps.Instructions_map_str[0x59] = "LD E, C"
	cpu.Instructions_maps.Instructions_map_str[0x5A] = "LD E, D"
	cpu.Instructions_maps.Instructions_map_str[0x5B] = "LD E, E"
	cpu.Instructions_maps.Instructions_map_str[0x5C] = "LD E, H"
	cpu.Instructions_maps.Instructions_map_str[0x5D] = "LD E, L"
	cpu.Instructions_maps.Instructions_map_str[0x5E] = "LD E, HL"
	cpu.Instructions_maps.Instructions_map_str[0x5F] = "LD E, A"
	cpu.Instructions_maps.Instructions_map_str[0x60] = "LD H, B"
	cpu.Instructions_maps.Instructions_map_str[0x61] = "LD H, C"
	cpu.Instructions_maps.Instructions_map_str[0x62] = "LD H, D"
	cpu.Instructions_maps.Instructions_map_str[0x63] = "LD H, E"
	cpu.Instructions_maps.Instructions_map_str[0x64] = "LD H, H"
	cpu.Instructions_maps.Instructions_map_str[0x65] = "LD H, L"
	cpu.Instructions_maps.Instructions_map_str[0x66] = "LD H, HL"
	cpu.Instructions_maps.Instructions_map_str[0x67] = "LD H, A"
	cpu.Instructions_maps.Instructions_map_str[0x68] = "LD L, B"
	cpu.Instructions_maps.Instructions_map_str[0x69] = "LD L, C"
	cpu.Instructions_maps.Instructions_map_str[0x6A] = "LD L, D"
	cpu.Instructions_maps.Instructions_map_str[0x6B] = "LD L, E"
	cpu.Instructions_maps.Instructions_map_str[0x6C] = "LD L, H"
	cpu.Instructions_maps.Instructions_map_str[0x6D] = "LD L, L"
	cpu.Instructions_maps.Instructions_map_str[0x6E] = "LD L, HL"
	cpu.Instructions_maps.Instructions_map_str[0x6F] = "LD L, A"
	cpu.Instructions_maps.Instructions_map_str[0x70] = "LD HL, B"
	cpu.Instructions_maps.Instructions_map_str[0x71] = "LD HL, C"
	cpu.Instructions_maps.Instructions_map_str[0x72] = "LD HL, D"
	cpu.Instructions_maps.Instructions_map_str[0x73] = "LD HL, E"
	cpu.Instructions_maps.Instructions_map_str[0x74] = "LD HL, H"
	cpu.Instructions_maps.Instructions_map_str[0x75] = "LD HL, L"
	cpu.Instructions_maps.Instructions_map_str[0x76] = "HALT"
	cpu.Instructions_maps.Instructions_map_str[0x77] = "LD HL, A"
	cpu.Instructions_maps.Instructions_map_str[0x78] = "LD A, B"
	cpu.Instructions_maps.Instructions_map_str[0x79] = "LD A, C"
	cpu.Instructions_maps.Instructions_map_str[0x7A] = "LD A, D"
	cpu.Instructions_maps.Instructions_map_str[0x7B] = "LD A, E"
	cpu.Instructions_maps.Instructions_map_str[0x7C] = "LD A, H"
	cpu.Instructions_maps.Instructions_map_str[0x7D] = "LD A, L"
	cpu.Instructions_maps.Instructions_map_str[0x7E] = "LD A, HL"
	cpu.Instructions_maps.Instructions_map_str[0x7F] = "LD A, A"
	cpu.Instructions_maps.Instructions_map_str[0x80] = "ADD A, B"
	cpu.Instructions_maps.Instructions_map_str[0x81] = "ADD A, C"
	cpu.Instructions_maps.Instructions_map_str[0x82] = "ADD A, D"
	cpu.Instructions_maps.Instructions_map_str[0x83] = "ADD A, E"
	cpu.Instructions_maps.Instructions_map_str[0x84] = "ADD A, H"
	cpu.Instructions_maps.Instructions_map_str[0x85] = "ADD A, L"
	cpu.Instructions_maps.Instructions_map_str[0x86] = "ADD A, HL"
	cpu.Instructions_maps.Instructions_map_str[0x87] = "ADD A, A"
	cpu.Instructions_maps.Instructions_map_str[0x88] = "ADC A, B"
	cpu.Instructions_maps.Instructions_map_str[0x89] = "ADC A, C"
	cpu.Instructions_maps.Instructions_map_str[0x8A] = "ADC A, D"
	cpu.Instructions_maps.Instructions_map_str[0x8B] = "ADC A, E"
	cpu.Instructions_maps.Instructions_map_str[0x8C] = "ADC A, H"
	cpu.Instructions_maps.Instructions_map_str[0x8D] = "ADC A, L"
	cpu.Instructions_maps.Instructions_map_str[0x8E] = "ADC A, HL"
	cpu.Instructions_maps.Instructions_map_str[0x8F] = "ADC A, A"
	cpu.Instructions_maps.Instructions_map_str[0x90] = "SUB B"
	cpu.Instructions_maps.Instructions_map_str[0x91] = "SUB C"
	cpu.Instructions_maps.Instructions_map_str[0x92] = "SUB D"
	cpu.Instructions_maps.Instructions_map_str[0x93] = "SUB E"
	cpu.Instructions_maps.Instructions_map_str[0x94] = "SUB H"
	cpu.Instructions_maps.Instructions_map_str[0x95] = "SUB L"
	cpu.Instructions_maps.Instructions_map_str[0x96] = "SUB HL"
	cpu.Instructions_maps.Instructions_map_str[0x97] = "SUB A"
	cpu.Instructions_maps.Instructions_map_str[0x98] = "SBC A, B"
	cpu.Instructions_maps.Instructions_map_str[0x99] = "SBC A, C"
	cpu.Instructions_maps.Instructions_map_str[0x9A] = "SBC A, D"
	cpu.Instructions_maps.Instructions_map_str[0x9B] = "SBC A, E"
	cpu.Instructions_maps.Instructions_map_str[0x9C] = "SBC A, H"
	cpu.Instructions_maps.Instructions_map_str[0x9D] = "SBC A, L"
	cpu.Instructions_maps.Instructions_map_str[0x9E] = "SBC A, HL"
	cpu.Instructions_maps.Instructions_map_str[0x9F] = "SBC A, A"
	cpu.Instructions_maps.Instructions_map_str[0xA0] = "AND B"
	cpu.Instructions_maps.Instructions_map_str[0xA1] = "AND C"
	cpu.Instructions_maps.Instructions_map_str[0xA2] = "AND D"
	cpu.Instructions_maps.Instructions_map_str[0xA3] = "AND E"
	cpu.Instructions_maps.Instructions_map_str[0xA4] = "AND H"
	cpu.Instructions_maps.Instructions_map_str[0xA5] = "AND L"
	cpu.Instructions_maps.Instructions_map_str[0xA6] = "AND HL"
	cpu.Instructions_maps.Instructions_map_str[0xA7] = "AND A"
	cpu.Instructions_maps.Instructions_map_str[0xA8] = "XOR B"
	cpu.Instructions_maps.Instructions_map_str[0xA9] = "XOR C"
	cpu.Instructions_maps.Instructions_map_str[0xAA] = "XOR D"
	cpu.Instructions_maps.Instructions_map_str[0xAB] = "XOR E"
	cpu.Instructions_maps.Instructions_map_str[0xAC] = "XOR H"
	cpu.Instructions_maps.Instructions_map_str[0xAD] = "XOR L"
	cpu.Instructions_maps.Instructions_map_str[0xAE] = "XOR HL"
	cpu.Instructions_maps.Instructions_map_str[0xAF] = "XOR A"
	cpu.Instructions_maps.Instructions_map_str[0xB0] = "OR B"
	cpu.Instructions_maps.Instructions_map_str[0xB1] = "OR C"
	cpu.Instructions_maps.Instructions_map_str[0xB2] = "OR D"
	cpu.Instructions_maps.Instructions_map_str[0xB3] = "OR E"
	cpu.Instructions_maps.Instructions_map_str[0xB4] = "OR H"
	cpu.Instructions_maps.Instructions_map_str[0xB5] = "OR L"
	cpu.Instructions_maps.Instructions_map_str[0xB6] = "OR HL"
	cpu.Instructions_maps.Instructions_map_str[0xB7] = "OR A"
	cpu.Instructions_maps.Instructions_map_str[0xB8] = "CP B"
	cpu.Instructions_maps.Instructions_map_str[0xB9] = "CP C"
	cpu.Instructions_maps.Instructions_map_str[0xBA] = "CP D"
	cpu.Instructions_maps.Instructions_map_str[0xBB] = "CP E"
	cpu.Instructions_maps.Instructions_map_str[0xBC] = "CP H"
	cpu.Instructions_maps.Instructions_map_str[0xBD] = "CP L"
	cpu.Instructions_maps.Instructions_map_str[0xBE] = "CP HL"
	cpu.Instructions_maps.Instructions_map_str[0xBF] = "CP A"
	cpu.Instructions_maps.Instructions_map_str[0xC0] = "RET NZ"
	cpu.Instructions_maps.Instructions_map_str[0xC1] = "POP BC"
	cpu.Instructions_maps.Instructions_map_str[0xC2] = "JP NZ, a16"
	cpu.Instructions_maps.Instructions_map_str[0xC3] = "JP a16"
	cpu.Instructions_maps.Instructions_map_str[0xC4] = "CALL NZ, a16"
	cpu.Instructions_maps.Instructions_map_str[0xC5] = "PUSH BC"
	cpu.Instructions_maps.Instructions_map_str[0xC6] = "ADD A, d8"
	cpu.Instructions_maps.Instructions_map_str[0xC7] = "RST 00H"
	cpu.Instructions_maps.Instructions_map_str[0xC8] = "RET Z"
	cpu.Instructions_maps.Instructions_map_str[0xC9] = "RET"
	cpu.Instructions_maps.Instructions_map_str[0xCA] = "JP Z, a16"
	cpu.Instructions_maps.Instructions_map_str[0xCB] = "PREFIX"
	cpu.Instructions_maps.Instructions_map_str[0xCC] = "CALL Z, a16"
	cpu.Instructions_maps.Instructions_map_str[0xCD] = "CALL a16"
	cpu.Instructions_maps.Instructions_map_str[0xCE] = "ADC A, d8"
	cpu.Instructions_maps.Instructions_map_str[0xCF] = "RST 08H"
	cpu.Instructions_maps.Instructions_map_str[0xD0] = "RET NC"
	cpu.Instructions_maps.Instructions_map_str[0xD1] = "POP DE"
	cpu.Instructions_maps.Instructions_map_str[0xD2] = "JP NC, a16"
	cpu.Instructions_maps.Instructions_map_str[0xD3] = "ILLEGAL_D3"
	cpu.Instructions_maps.Instructions_map_str[0xD4] = "CALL NC, a16"
	cpu.Instructions_maps.Instructions_map_str[0xD5] = "PUSH DE"
	cpu.Instructions_maps.Instructions_map_str[0xD6] = "SUB d8"
	cpu.Instructions_maps.Instructions_map_str[0xD7] = "RST 10H"
	cpu.Instructions_maps.Instructions_map_str[0xD8] = "RET C"
	cpu.Instructions_maps.Instructions_map_str[0xD9] = "RETI"
	cpu.Instructions_maps.Instructions_map_str[0xDA] = "JP C, a16"
	cpu.Instructions_maps.Instructions_map_str[0xDB] = "ILLEGAL_DB"
	cpu.Instructions_maps.Instructions_map_str[0xDC] = "CALL C, a16"
	cpu.Instructions_maps.Instructions_map_str[0xDD] = "ILLEGAL_DD"
	cpu.Instructions_maps.Instructions_map_str[0xDE] = "SBC A, d8"
	cpu.Instructions_maps.Instructions_map_str[0xDF] = "RST 18H"
	cpu.Instructions_maps.Instructions_map_str[0xE0] = "LDH a8, A"
	cpu.Instructions_maps.Instructions_map_str[0xE1] = "POP HL"
	cpu.Instructions_maps.Instructions_map_str[0xE2] = "LD C, A"
	cpu.Instructions_maps.Instructions_map_str[0xE3] = "ILLEGAL_E3"
	cpu.Instructions_maps.Instructions_map_str[0xE4] = "ILLEGAL_E4"
	cpu.Instructions_maps.Instructions_map_str[0xE5] = "PUSH HL"
	cpu.Instructions_maps.Instructions_map_str[0xE6] = "AND d8"
	cpu.Instructions_maps.Instructions_map_str[0xE7] = "RST 20H"
	cpu.Instructions_maps.Instructions_map_str[0xE8] = "ADD SP, r8"
	cpu.Instructions_maps.Instructions_map_str[0xE9] = "JP HL"
	cpu.Instructions_maps.Instructions_map_str[0xEA] = "LD a16, A"
	cpu.Instructions_maps.Instructions_map_str[0xEB] = "ILLEGAL_EB"
	cpu.Instructions_maps.Instructions_map_str[0xEC] = "ILLEGAL_EC"
	cpu.Instructions_maps.Instructions_map_str[0xED] = "ILLEGAL_ED"
	cpu.Instructions_maps.Instructions_map_str[0xEE] = "XOR d8"
	cpu.Instructions_maps.Instructions_map_str[0xEF] = "RST 28H"
	cpu.Instructions_maps.Instructions_map_str[0xF0] = "LDH A, a8"
	cpu.Instructions_maps.Instructions_map_str[0xF1] = "POP AF"
	cpu.Instructions_maps.Instructions_map_str[0xF2] = "LD A, C"
	cpu.Instructions_maps.Instructions_map_str[0xF3] = "DI"
	cpu.Instructions_maps.Instructions_map_str[0xF4] = "ILLEGAL_F4"
	cpu.Instructions_maps.Instructions_map_str[0xF5] = "PUSH AF"
	cpu.Instructions_maps.Instructions_map_str[0xF6] = "OR d8"
	cpu.Instructions_maps.Instructions_map_str[0xF7] = "RST 30H"
	cpu.Instructions_maps.Instructions_map_str[0xF8] = "LD"
	cpu.Instructions_maps.Instructions_map_str[0xF9] = "LD SP, HL"
	cpu.Instructions_maps.Instructions_map_str[0xFA] = "LD A, a16"
	cpu.Instructions_maps.Instructions_map_str[0xFB] = "EI"
	cpu.Instructions_maps.Instructions_map_str[0xFC] = "ILLEGAL_FC"
	cpu.Instructions_maps.Instructions_map_str[0xFD] = "ILLEGAL_FD"
	cpu.Instructions_maps.Instructions_map_str[0xFE] = "CP d8"
	cpu.Instructions_maps.Instructions_map_str[0xFF] = "RST 38H"

	cpu.Instructions_maps.Instructions_map_CB_str[0x00] = "RLC B"
	cpu.Instructions_maps.Instructions_map_CB_str[0x01] = "RLC C"
	cpu.Instructions_maps.Instructions_map_CB_str[0x02] = "RLC D"
	cpu.Instructions_maps.Instructions_map_CB_str[0x03] = "RLC E"
	cpu.Instructions_maps.Instructions_map_CB_str[0x04] = "RLC H"
	cpu.Instructions_maps.Instructions_map_CB_str[0x05] = "RLC L"
	cpu.Instructions_maps.Instructions_map_CB_str[0x06] = "RLC HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0x07] = "RLC A"
	cpu.Instructions_maps.Instructions_map_CB_str[0x08] = "RRC B"
	cpu.Instructions_maps.Instructions_map_CB_str[0x09] = "RRC C"
	cpu.Instructions_maps.Instructions_map_CB_str[0x0A] = "RRC D"
	cpu.Instructions_maps.Instructions_map_CB_str[0x0B] = "RRC E"
	cpu.Instructions_maps.Instructions_map_CB_str[0x0C] = "RRC H"
	cpu.Instructions_maps.Instructions_map_CB_str[0x0D] = "RRC L"
	cpu.Instructions_maps.Instructions_map_CB_str[0x0E] = "RRC HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0x0F] = "RRC A"
	cpu.Instructions_maps.Instructions_map_CB_str[0x10] = "RL B"
	cpu.Instructions_maps.Instructions_map_CB_str[0x11] = "RL C"
	cpu.Instructions_maps.Instructions_map_CB_str[0x12] = "RL D"
	cpu.Instructions_maps.Instructions_map_CB_str[0x13] = "RL E"
	cpu.Instructions_maps.Instructions_map_CB_str[0x14] = "RL H"
	cpu.Instructions_maps.Instructions_map_CB_str[0x15] = "RL L"
	cpu.Instructions_maps.Instructions_map_CB_str[0x16] = "RL HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0x17] = "RL A"
	cpu.Instructions_maps.Instructions_map_CB_str[0x18] = "RR B"
	cpu.Instructions_maps.Instructions_map_CB_str[0x19] = "RR C"
	cpu.Instructions_maps.Instructions_map_CB_str[0x1A] = "RR D"
	cpu.Instructions_maps.Instructions_map_CB_str[0x1B] = "RR E"
	cpu.Instructions_maps.Instructions_map_CB_str[0x1C] = "RR H"
	cpu.Instructions_maps.Instructions_map_CB_str[0x1D] = "RR L"
	cpu.Instructions_maps.Instructions_map_CB_str[0x1E] = "RR HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0x1F] = "RR A"
	cpu.Instructions_maps.Instructions_map_CB_str[0x20] = "SLA B"
	cpu.Instructions_maps.Instructions_map_CB_str[0x21] = "SLA C"
	cpu.Instructions_maps.Instructions_map_CB_str[0x22] = "SLA D"
	cpu.Instructions_maps.Instructions_map_CB_str[0x23] = "SLA E"
	cpu.Instructions_maps.Instructions_map_CB_str[0x24] = "SLA H"
	cpu.Instructions_maps.Instructions_map_CB_str[0x25] = "SLA L"
	cpu.Instructions_maps.Instructions_map_CB_str[0x26] = "SLA HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0x27] = "SLA A"
	cpu.Instructions_maps.Instructions_map_CB_str[0x28] = "SRA B"
	cpu.Instructions_maps.Instructions_map_CB_str[0x29] = "SRA C"
	cpu.Instructions_maps.Instructions_map_CB_str[0x2A] = "SRA D"
	cpu.Instructions_maps.Instructions_map_CB_str[0x2B] = "SRA E"
	cpu.Instructions_maps.Instructions_map_CB_str[0x2C] = "SRA H"
	cpu.Instructions_maps.Instructions_map_CB_str[0x2D] = "SRA L"
	cpu.Instructions_maps.Instructions_map_CB_str[0x2E] = "SRA HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0x2F] = "SRA A"
	cpu.Instructions_maps.Instructions_map_CB_str[0x30] = "SWAP B"
	cpu.Instructions_maps.Instructions_map_CB_str[0x31] = "SWAP C"
	cpu.Instructions_maps.Instructions_map_CB_str[0x32] = "SWAP D"
	cpu.Instructions_maps.Instructions_map_CB_str[0x33] = "SWAP E"
	cpu.Instructions_maps.Instructions_map_CB_str[0x34] = "SWAP H"
	cpu.Instructions_maps.Instructions_map_CB_str[0x35] = "SWAP L"
	cpu.Instructions_maps.Instructions_map_CB_str[0x36] = "SWAP HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0x37] = "SWAP A"
	cpu.Instructions_maps.Instructions_map_CB_str[0x38] = "SRL B"
	cpu.Instructions_maps.Instructions_map_CB_str[0x39] = "SRL C"
	cpu.Instructions_maps.Instructions_map_CB_str[0x3A] = "SRL D"
	cpu.Instructions_maps.Instructions_map_CB_str[0x3B] = "SRL E"
	cpu.Instructions_maps.Instructions_map_CB_str[0x3C] = "SRL H"
	cpu.Instructions_maps.Instructions_map_CB_str[0x3D] = "SRL L"
	cpu.Instructions_maps.Instructions_map_CB_str[0x3E] = "SRL HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0x3F] = "SRL A"
	cpu.Instructions_maps.Instructions_map_CB_str[0x40] = "BIT 0, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0x41] = "BIT 0, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0x42] = "BIT 0, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0x43] = "BIT 0, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0x44] = "BIT 0, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0x45] = "BIT 0, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0x46] = "BIT 0, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0x47] = "BIT 0, A"
	cpu.Instructions_maps.Instructions_map_CB_str[0x48] = "BIT 1, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0x49] = "BIT 1, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0x4A] = "BIT 1, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0x4B] = "BIT 1, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0x4C] = "BIT 1, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0x4D] = "BIT 1, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0x4E] = "BIT 1, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0x4F] = "BIT 1, A"
	cpu.Instructions_maps.Instructions_map_CB_str[0x50] = "BIT 2, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0x51] = "BIT 2, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0x52] = "BIT 2, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0x53] = "BIT 2, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0x54] = "BIT 2, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0x55] = "BIT 2, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0x56] = "BIT 2, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0x57] = "BIT 2, A"
	cpu.Instructions_maps.Instructions_map_CB_str[0x58] = "BIT 3, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0x59] = "BIT 3, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0x5A] = "BIT 3, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0x5B] = "BIT 3, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0x5C] = "BIT 3, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0x5D] = "BIT 3, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0x5E] = "BIT 3, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0x5F] = "BIT 3, A"
	cpu.Instructions_maps.Instructions_map_CB_str[0x60] = "BIT 4, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0x61] = "BIT 4, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0x62] = "BIT 4, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0x63] = "BIT 4, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0x64] = "BIT 4, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0x65] = "BIT 4, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0x66] = "BIT 4, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0x67] = "BIT 4, A"
	cpu.Instructions_maps.Instructions_map_CB_str[0x68] = "BIT 5, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0x69] = "BIT 5, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0x6A] = "BIT 5, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0x6B] = "BIT 5, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0x6C] = "BIT 5, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0x6D] = "BIT 5, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0x6E] = "BIT 5, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0x6F] = "BIT 5, A"
	cpu.Instructions_maps.Instructions_map_CB_str[0x70] = "BIT 6, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0x71] = "BIT 6, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0x72] = "BIT 6, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0x73] = "BIT 6, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0x74] = "BIT 6, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0x75] = "BIT 6, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0x76] = "BIT 6, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0x77] = "BIT 6, A"
	cpu.Instructions_maps.Instructions_map_CB_str[0x78] = "BIT 7, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0x79] = "BIT 7, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0x7A] = "BIT 7, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0x7B] = "BIT 7, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0x7C] = "BIT 7, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0x7D] = "BIT 7, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0x7E] = "BIT 7, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0x7F] = "BIT 7, A"
	cpu.Instructions_maps.Instructions_map_CB_str[0x80] = "RES 0, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0x81] = "RES 0, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0x82] = "RES 0, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0x83] = "RES 0, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0x84] = "RES 0, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0x85] = "RES 0, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0x86] = "RES 0, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0x87] = "RES 0, A"
	cpu.Instructions_maps.Instructions_map_CB_str[0x88] = "RES 1, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0x89] = "RES 1, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0x8A] = "RES 1, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0x8B] = "RES 1, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0x8C] = "RES 1, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0x8D] = "RES 1, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0x8E] = "RES 1, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0x8F] = "RES 1, A"
	cpu.Instructions_maps.Instructions_map_CB_str[0x90] = "RES 2, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0x91] = "RES 2, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0x92] = "RES 2, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0x93] = "RES 2, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0x94] = "RES 2, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0x95] = "RES 2, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0x96] = "RES 2, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0x97] = "RES 2, A"
	cpu.Instructions_maps.Instructions_map_CB_str[0x98] = "RES 3, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0x99] = "RES 3, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0x9A] = "RES 3, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0x9B] = "RES 3, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0x9C] = "RES 3, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0x9D] = "RES 3, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0x9E] = "RES 3, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0x9F] = "RES 3, A"
	cpu.Instructions_maps.Instructions_map_CB_str[0xA0] = "RES 4, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0xA1] = "RES 4, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0xA2] = "RES 4, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0xA3] = "RES 4, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0xA4] = "RES 4, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0xA5] = "RES 4, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0xA6] = "RES 4, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0xA7] = "RES 4, A"
	cpu.Instructions_maps.Instructions_map_CB_str[0xA8] = "RES 5, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0xA9] = "RES 5, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0xAA] = "RES 5, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0xAB] = "RES 5, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0xAC] = "RES 5, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0xAD] = "RES 5, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0xAE] = "RES 5, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0xAF] = "RES 5, A"
	cpu.Instructions_maps.Instructions_map_CB_str[0xB0] = "RES 6, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0xB1] = "RES 6, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0xB2] = "RES 6, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0xB3] = "RES 6, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0xB4] = "RES 6, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0xB5] = "RES 6, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0xB6] = "RES 6, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0xB7] = "RES 6, A"
	cpu.Instructions_maps.Instructions_map_CB_str[0xB8] = "RES 7, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0xB9] = "RES 7, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0xBA] = "RES 7, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0xBB] = "RES 7, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0xBC] = "RES 7, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0xBD] = "RES 7, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0xBE] = "RES 7, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0xBF] = "RES 7, A"
	cpu.Instructions_maps.Instructions_map_CB_str[0xC0] = "SET 0, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0xC1] = "SET 0, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0xC2] = "SET 0, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0xC3] = "SET 0, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0xC4] = "SET 0, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0xC5] = "SET 0, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0xC6] = "SET 0, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0xC7] = "SET 0, A"
	cpu.Instructions_maps.Instructions_map_CB_str[0xC8] = "SET 1, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0xC9] = "SET 1, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0xCA] = "SET 1, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0xCB] = "SET 1, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0xCC] = "SET 1, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0xCD] = "SET 1, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0xCE] = "SET 1, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0xCF] = "SET 1, A"
	cpu.Instructions_maps.Instructions_map_CB_str[0xD0] = "SET 2, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0xD1] = "SET 2, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0xD2] = "SET 2, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0xD3] = "SET 2, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0xD4] = "SET 2, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0xD5] = "SET 2, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0xD6] = "SET 2, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0xD7] = "SET 2, A"
	cpu.Instructions_maps.Instructions_map_CB_str[0xD8] = "SET 3, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0xD9] = "SET 3, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0xDA] = "SET 3, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0xDB] = "SET 3, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0xDC] = "SET 3, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0xDD] = "SET 3, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0xDE] = "SET 3, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0xDF] = "SET 3, A"
	cpu.Instructions_maps.Instructions_map_CB_str[0xE0] = "SET 4, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0xE1] = "SET 4, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0xE2] = "SET 4, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0xE3] = "SET 4, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0xE4] = "SET 4, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0xE5] = "SET 4, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0xE6] = "SET 4, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0xE7] = "SET 4, A"
	cpu.Instructions_maps.Instructions_map_CB_str[0xE8] = "SET 5, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0xE9] = "SET 5, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0xEA] = "SET 5, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0xEB] = "SET 5, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0xEC] = "SET 5, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0xED] = "SET 5, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0xEE] = "SET 5, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0xEF] = "SET 5, A"
	cpu.Instructions_maps.Instructions_map_CB_str[0xF0] = "SET 6, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0xF1] = "SET 6, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0xF2] = "SET 6, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0xF3] = "SET 6, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0xF4] = "SET 6, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0xF5] = "SET 6, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0xF6] = "SET 6, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0xF7] = "SET 6, A"
	cpu.Instructions_maps.Instructions_map_CB_str[0xF8] = "SET 7, B"
	cpu.Instructions_maps.Instructions_map_CB_str[0xF9] = "SET 7, C"
	cpu.Instructions_maps.Instructions_map_CB_str[0xFA] = "SET 7, D"
	cpu.Instructions_maps.Instructions_map_CB_str[0xFB] = "SET 7, E"
	cpu.Instructions_maps.Instructions_map_CB_str[0xFC] = "SET 7, H"
	cpu.Instructions_maps.Instructions_map_CB_str[0xFD] = "SET 7, L"
	cpu.Instructions_maps.Instructions_map_CB_str[0xFE] = "SET 7, HL"
	cpu.Instructions_maps.Instructions_map_CB_str[0xFF] = "SET 7, A"

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
	cpu.Instructions_maps.Instructions_map_[0x10] = cpu.op_0x10
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
	cpu.Instructions_maps.Instructions_map_CB[0x00] = cpu.op_cb_0x00
	cpu.Instructions_maps.Instructions_map_CB[0x01] = cpu.op_cb_0x01
	cpu.Instructions_maps.Instructions_map_CB[0x02] = cpu.op_cb_0x02
	cpu.Instructions_maps.Instructions_map_CB[0x03] = cpu.op_cb_0x03
	cpu.Instructions_maps.Instructions_map_CB[0x04] = cpu.op_cb_0x04
	cpu.Instructions_maps.Instructions_map_CB[0x05] = cpu.op_cb_0x05
	cpu.Instructions_maps.Instructions_map_CB[0x06] = cpu.op_cb_0x06
	cpu.Instructions_maps.Instructions_map_CB[0x07] = cpu.op_cb_0x07
	cpu.Instructions_maps.Instructions_map_CB[0x08] = cpu.op_cb_0x08
	cpu.Instructions_maps.Instructions_map_CB[0x09] = cpu.op_cb_0x09
	cpu.Instructions_maps.Instructions_map_CB[0x0a] = cpu.op_cb_0x0a
	cpu.Instructions_maps.Instructions_map_CB[0x0b] = cpu.op_cb_0x0b
	cpu.Instructions_maps.Instructions_map_CB[0x0c] = cpu.op_cb_0x0c
	cpu.Instructions_maps.Instructions_map_CB[0x0d] = cpu.op_cb_0x0d
	cpu.Instructions_maps.Instructions_map_CB[0x0e] = cpu.op_cb_0x0e
	cpu.Instructions_maps.Instructions_map_CB[0x0f] = cpu.op_cb_0x0f
	cpu.Instructions_maps.Instructions_map_CB[0x10] = cpu.op_cb_0x10
	cpu.Instructions_maps.Instructions_map_CB[0x11] = cpu.op_cb_0x11
	cpu.Instructions_maps.Instructions_map_CB[0x12] = cpu.op_cb_0x12
	cpu.Instructions_maps.Instructions_map_CB[0x13] = cpu.op_cb_0x13
	cpu.Instructions_maps.Instructions_map_CB[0x14] = cpu.op_cb_0x14
	cpu.Instructions_maps.Instructions_map_CB[0x15] = cpu.op_cb_0x15
	cpu.Instructions_maps.Instructions_map_CB[0x16] = cpu.op_cb_0x16
	cpu.Instructions_maps.Instructions_map_CB[0x17] = cpu.op_cb_0x17
	cpu.Instructions_maps.Instructions_map_CB[0x18] = cpu.op_cb_0x18
	cpu.Instructions_maps.Instructions_map_CB[0x19] = cpu.op_cb_0x19
	cpu.Instructions_maps.Instructions_map_CB[0x1a] = cpu.op_cb_0x1a
	cpu.Instructions_maps.Instructions_map_CB[0x1b] = cpu.op_cb_0x1b
	cpu.Instructions_maps.Instructions_map_CB[0x1c] = cpu.op_cb_0x1c
	cpu.Instructions_maps.Instructions_map_CB[0x1d] = cpu.op_cb_0x1d
	cpu.Instructions_maps.Instructions_map_CB[0x1e] = cpu.op_cb_0x1e
	cpu.Instructions_maps.Instructions_map_CB[0x1f] = cpu.op_cb_0x1f
	cpu.Instructions_maps.Instructions_map_CB[0x20] = cpu.op_cb_0x20
	cpu.Instructions_maps.Instructions_map_CB[0x21] = cpu.op_cb_0x21
	cpu.Instructions_maps.Instructions_map_CB[0x22] = cpu.op_cb_0x22
	cpu.Instructions_maps.Instructions_map_CB[0x23] = cpu.op_cb_0x23
	cpu.Instructions_maps.Instructions_map_CB[0x24] = cpu.op_cb_0x24
	cpu.Instructions_maps.Instructions_map_CB[0x25] = cpu.op_cb_0x25
	cpu.Instructions_maps.Instructions_map_CB[0x26] = cpu.op_cb_0x26
	cpu.Instructions_maps.Instructions_map_CB[0x27] = cpu.op_cb_0x27
	cpu.Instructions_maps.Instructions_map_CB[0x28] = cpu.op_cb_0x28
	cpu.Instructions_maps.Instructions_map_CB[0x29] = cpu.op_cb_0x29
	cpu.Instructions_maps.Instructions_map_CB[0x2a] = cpu.op_cb_0x2a
	cpu.Instructions_maps.Instructions_map_CB[0x2b] = cpu.op_cb_0x2b
	cpu.Instructions_maps.Instructions_map_CB[0x2c] = cpu.op_cb_0x2c
	cpu.Instructions_maps.Instructions_map_CB[0x2d] = cpu.op_cb_0x2d
	cpu.Instructions_maps.Instructions_map_CB[0x2e] = cpu.op_cb_0x2e
	cpu.Instructions_maps.Instructions_map_CB[0x2f] = cpu.op_cb_0x2f
	cpu.Instructions_maps.Instructions_map_CB[0x30] = cpu.op_cb_0x30
	cpu.Instructions_maps.Instructions_map_CB[0x31] = cpu.op_cb_0x31
	cpu.Instructions_maps.Instructions_map_CB[0x32] = cpu.op_cb_0x32
	cpu.Instructions_maps.Instructions_map_CB[0x33] = cpu.op_cb_0x33
	cpu.Instructions_maps.Instructions_map_CB[0x34] = cpu.op_cb_0x34
	cpu.Instructions_maps.Instructions_map_CB[0x35] = cpu.op_cb_0x35
	cpu.Instructions_maps.Instructions_map_CB[0x36] = cpu.op_cb_0x36
	cpu.Instructions_maps.Instructions_map_CB[0x37] = cpu.op_cb_0x37
	cpu.Instructions_maps.Instructions_map_CB[0x38] = cpu.op_cb_0x38
	cpu.Instructions_maps.Instructions_map_CB[0x39] = cpu.op_cb_0x39
	cpu.Instructions_maps.Instructions_map_CB[0x3a] = cpu.op_cb_0x3a
	cpu.Instructions_maps.Instructions_map_CB[0x3b] = cpu.op_cb_0x3b
	cpu.Instructions_maps.Instructions_map_CB[0x3c] = cpu.op_cb_0x3c
	cpu.Instructions_maps.Instructions_map_CB[0x3d] = cpu.op_cb_0x3d
	cpu.Instructions_maps.Instructions_map_CB[0x3e] = cpu.op_cb_0x3e
	cpu.Instructions_maps.Instructions_map_CB[0x3f] = cpu.op_cb_0x3f
	cpu.Instructions_maps.Instructions_map_CB[0x40] = cpu.op_cb_0x40
	cpu.Instructions_maps.Instructions_map_CB[0x41] = cpu.op_cb_0x41
	cpu.Instructions_maps.Instructions_map_CB[0x42] = cpu.op_cb_0x42
	cpu.Instructions_maps.Instructions_map_CB[0x43] = cpu.op_cb_0x43
	cpu.Instructions_maps.Instructions_map_CB[0x44] = cpu.op_cb_0x44
	cpu.Instructions_maps.Instructions_map_CB[0x45] = cpu.op_cb_0x45
	cpu.Instructions_maps.Instructions_map_CB[0x46] = cpu.op_cb_0x46
	cpu.Instructions_maps.Instructions_map_CB[0x47] = cpu.op_cb_0x47
	cpu.Instructions_maps.Instructions_map_CB[0x48] = cpu.op_cb_0x48
	cpu.Instructions_maps.Instructions_map_CB[0x49] = cpu.op_cb_0x49
	cpu.Instructions_maps.Instructions_map_CB[0x4a] = cpu.op_cb_0x4a
	cpu.Instructions_maps.Instructions_map_CB[0x4b] = cpu.op_cb_0x4b
	cpu.Instructions_maps.Instructions_map_CB[0x4c] = cpu.op_cb_0x4c
	cpu.Instructions_maps.Instructions_map_CB[0x4d] = cpu.op_cb_0x4d
	cpu.Instructions_maps.Instructions_map_CB[0x4e] = cpu.op_cb_0x4e
	cpu.Instructions_maps.Instructions_map_CB[0x4f] = cpu.op_cb_0x4f
	cpu.Instructions_maps.Instructions_map_CB[0x50] = cpu.op_cb_0x50
	cpu.Instructions_maps.Instructions_map_CB[0x51] = cpu.op_cb_0x51
	cpu.Instructions_maps.Instructions_map_CB[0x52] = cpu.op_cb_0x52
	cpu.Instructions_maps.Instructions_map_CB[0x53] = cpu.op_cb_0x53
	cpu.Instructions_maps.Instructions_map_CB[0x54] = cpu.op_cb_0x54
	cpu.Instructions_maps.Instructions_map_CB[0x55] = cpu.op_cb_0x55
	cpu.Instructions_maps.Instructions_map_CB[0x56] = cpu.op_cb_0x56
	cpu.Instructions_maps.Instructions_map_CB[0x57] = cpu.op_cb_0x57
	cpu.Instructions_maps.Instructions_map_CB[0x58] = cpu.op_cb_0x58
	cpu.Instructions_maps.Instructions_map_CB[0x59] = cpu.op_cb_0x59
	cpu.Instructions_maps.Instructions_map_CB[0x5a] = cpu.op_cb_0x5a
	cpu.Instructions_maps.Instructions_map_CB[0x5b] = cpu.op_cb_0x5b
	cpu.Instructions_maps.Instructions_map_CB[0x5c] = cpu.op_cb_0x5c
	cpu.Instructions_maps.Instructions_map_CB[0x5d] = cpu.op_cb_0x5d
	cpu.Instructions_maps.Instructions_map_CB[0x5e] = cpu.op_cb_0x5e
	cpu.Instructions_maps.Instructions_map_CB[0x5f] = cpu.op_cb_0x5f
	cpu.Instructions_maps.Instructions_map_CB[0x60] = cpu.op_cb_0x60
	cpu.Instructions_maps.Instructions_map_CB[0x61] = cpu.op_cb_0x61
	cpu.Instructions_maps.Instructions_map_CB[0x62] = cpu.op_cb_0x62
	cpu.Instructions_maps.Instructions_map_CB[0x63] = cpu.op_cb_0x63
	cpu.Instructions_maps.Instructions_map_CB[0x64] = cpu.op_cb_0x64
	cpu.Instructions_maps.Instructions_map_CB[0x65] = cpu.op_cb_0x65
	cpu.Instructions_maps.Instructions_map_CB[0x66] = cpu.op_cb_0x66
	cpu.Instructions_maps.Instructions_map_CB[0x67] = cpu.op_cb_0x67
	cpu.Instructions_maps.Instructions_map_CB[0x68] = cpu.op_cb_0x68
	cpu.Instructions_maps.Instructions_map_CB[0x69] = cpu.op_cb_0x69
	cpu.Instructions_maps.Instructions_map_CB[0x6a] = cpu.op_cb_0x6a
	cpu.Instructions_maps.Instructions_map_CB[0x6b] = cpu.op_cb_0x6b
	cpu.Instructions_maps.Instructions_map_CB[0x6c] = cpu.op_cb_0x6c
	cpu.Instructions_maps.Instructions_map_CB[0x6d] = cpu.op_cb_0x6d
	cpu.Instructions_maps.Instructions_map_CB[0x6e] = cpu.op_cb_0x6e
	cpu.Instructions_maps.Instructions_map_CB[0x6f] = cpu.op_cb_0x6f
	cpu.Instructions_maps.Instructions_map_CB[0x70] = cpu.op_cb_0x70
	cpu.Instructions_maps.Instructions_map_CB[0x71] = cpu.op_cb_0x71
	cpu.Instructions_maps.Instructions_map_CB[0x72] = cpu.op_cb_0x72
	cpu.Instructions_maps.Instructions_map_CB[0x73] = cpu.op_cb_0x73
	cpu.Instructions_maps.Instructions_map_CB[0x74] = cpu.op_cb_0x74
	cpu.Instructions_maps.Instructions_map_CB[0x75] = cpu.op_cb_0x75
	cpu.Instructions_maps.Instructions_map_CB[0x76] = cpu.op_cb_0x76
	cpu.Instructions_maps.Instructions_map_CB[0x77] = cpu.op_cb_0x77
	cpu.Instructions_maps.Instructions_map_CB[0x78] = cpu.op_cb_0x78
	cpu.Instructions_maps.Instructions_map_CB[0x79] = cpu.op_cb_0x79
	cpu.Instructions_maps.Instructions_map_CB[0x7a] = cpu.op_cb_0x7a
	cpu.Instructions_maps.Instructions_map_CB[0x7b] = cpu.op_cb_0x7b
	cpu.Instructions_maps.Instructions_map_CB[0x7c] = cpu.op_cb_0x7c
	cpu.Instructions_maps.Instructions_map_CB[0x7d] = cpu.op_cb_0x7d
	cpu.Instructions_maps.Instructions_map_CB[0x7e] = cpu.op_cb_0x7e
	cpu.Instructions_maps.Instructions_map_CB[0x7f] = cpu.op_cb_0x7f
	cpu.Instructions_maps.Instructions_map_CB[0x80] = cpu.op_cb_0x80
	cpu.Instructions_maps.Instructions_map_CB[0x81] = cpu.op_cb_0x81
	cpu.Instructions_maps.Instructions_map_CB[0x82] = cpu.op_cb_0x82
	cpu.Instructions_maps.Instructions_map_CB[0x83] = cpu.op_cb_0x83
	cpu.Instructions_maps.Instructions_map_CB[0x84] = cpu.op_cb_0x84
	cpu.Instructions_maps.Instructions_map_CB[0x85] = cpu.op_cb_0x85
	cpu.Instructions_maps.Instructions_map_CB[0x86] = cpu.op_cb_0x86
	cpu.Instructions_maps.Instructions_map_CB[0x87] = cpu.op_cb_0x87
	cpu.Instructions_maps.Instructions_map_CB[0x88] = cpu.op_cb_0x88
	cpu.Instructions_maps.Instructions_map_CB[0x89] = cpu.op_cb_0x89
	cpu.Instructions_maps.Instructions_map_CB[0x8a] = cpu.op_cb_0x8a
	cpu.Instructions_maps.Instructions_map_CB[0x8b] = cpu.op_cb_0x8b
	cpu.Instructions_maps.Instructions_map_CB[0x8c] = cpu.op_cb_0x8c
	cpu.Instructions_maps.Instructions_map_CB[0x8d] = cpu.op_cb_0x8d
	cpu.Instructions_maps.Instructions_map_CB[0x8e] = cpu.op_cb_0x8e
	cpu.Instructions_maps.Instructions_map_CB[0x8f] = cpu.op_cb_0x8f
	cpu.Instructions_maps.Instructions_map_CB[0x90] = cpu.op_cb_0x90
	cpu.Instructions_maps.Instructions_map_CB[0x91] = cpu.op_cb_0x91
	cpu.Instructions_maps.Instructions_map_CB[0x92] = cpu.op_cb_0x92
	cpu.Instructions_maps.Instructions_map_CB[0x93] = cpu.op_cb_0x93
	cpu.Instructions_maps.Instructions_map_CB[0x94] = cpu.op_cb_0x94
	cpu.Instructions_maps.Instructions_map_CB[0x95] = cpu.op_cb_0x95
	cpu.Instructions_maps.Instructions_map_CB[0x96] = cpu.op_cb_0x96
	cpu.Instructions_maps.Instructions_map_CB[0x97] = cpu.op_cb_0x97
	cpu.Instructions_maps.Instructions_map_CB[0x98] = cpu.op_cb_0x98
	cpu.Instructions_maps.Instructions_map_CB[0x99] = cpu.op_cb_0x99
	cpu.Instructions_maps.Instructions_map_CB[0x9a] = cpu.op_cb_0x9a
	cpu.Instructions_maps.Instructions_map_CB[0x9b] = cpu.op_cb_0x9b
	cpu.Instructions_maps.Instructions_map_CB[0x9c] = cpu.op_cb_0x9c
	cpu.Instructions_maps.Instructions_map_CB[0x9d] = cpu.op_cb_0x9d
	cpu.Instructions_maps.Instructions_map_CB[0x9e] = cpu.op_cb_0x9e
	cpu.Instructions_maps.Instructions_map_CB[0x9f] = cpu.op_cb_0x9f
	cpu.Instructions_maps.Instructions_map_CB[0xa0] = cpu.op_cb_0xa0
	cpu.Instructions_maps.Instructions_map_CB[0xa1] = cpu.op_cb_0xa1
	cpu.Instructions_maps.Instructions_map_CB[0xa2] = cpu.op_cb_0xa2
	cpu.Instructions_maps.Instructions_map_CB[0xa3] = cpu.op_cb_0xa3
	cpu.Instructions_maps.Instructions_map_CB[0xa4] = cpu.op_cb_0xa4
	cpu.Instructions_maps.Instructions_map_CB[0xa5] = cpu.op_cb_0xa5
	cpu.Instructions_maps.Instructions_map_CB[0xa6] = cpu.op_cb_0xa6
	cpu.Instructions_maps.Instructions_map_CB[0xa7] = cpu.op_cb_0xa7
	cpu.Instructions_maps.Instructions_map_CB[0xa8] = cpu.op_cb_0xa8
	cpu.Instructions_maps.Instructions_map_CB[0xa9] = cpu.op_cb_0xa9
	cpu.Instructions_maps.Instructions_map_CB[0xaa] = cpu.op_cb_0xaa
	cpu.Instructions_maps.Instructions_map_CB[0xab] = cpu.op_cb_0xab
	cpu.Instructions_maps.Instructions_map_CB[0xac] = cpu.op_cb_0xac
	cpu.Instructions_maps.Instructions_map_CB[0xad] = cpu.op_cb_0xad
	cpu.Instructions_maps.Instructions_map_CB[0xae] = cpu.op_cb_0xae
	cpu.Instructions_maps.Instructions_map_CB[0xaf] = cpu.op_cb_0xaf
	cpu.Instructions_maps.Instructions_map_CB[0xb0] = cpu.op_cb_0xb0
	cpu.Instructions_maps.Instructions_map_CB[0xb1] = cpu.op_cb_0xb1
	cpu.Instructions_maps.Instructions_map_CB[0xb2] = cpu.op_cb_0xb2
	cpu.Instructions_maps.Instructions_map_CB[0xb3] = cpu.op_cb_0xb3
	cpu.Instructions_maps.Instructions_map_CB[0xb4] = cpu.op_cb_0xb4
	cpu.Instructions_maps.Instructions_map_CB[0xb5] = cpu.op_cb_0xb5
	cpu.Instructions_maps.Instructions_map_CB[0xb6] = cpu.op_cb_0xb6
	cpu.Instructions_maps.Instructions_map_CB[0xb7] = cpu.op_cb_0xb7
	cpu.Instructions_maps.Instructions_map_CB[0xb8] = cpu.op_cb_0xb8
	cpu.Instructions_maps.Instructions_map_CB[0xb9] = cpu.op_cb_0xb9
	cpu.Instructions_maps.Instructions_map_CB[0xba] = cpu.op_cb_0xba
	cpu.Instructions_maps.Instructions_map_CB[0xbb] = cpu.op_cb_0xbb
	cpu.Instructions_maps.Instructions_map_CB[0xbc] = cpu.op_cb_0xbc
	cpu.Instructions_maps.Instructions_map_CB[0xbd] = cpu.op_cb_0xbd
	cpu.Instructions_maps.Instructions_map_CB[0xbe] = cpu.op_cb_0xbe
	cpu.Instructions_maps.Instructions_map_CB[0xbf] = cpu.op_cb_0xbf
	cpu.Instructions_maps.Instructions_map_CB[0xc0] = cpu.op_cb_0xc0
	cpu.Instructions_maps.Instructions_map_CB[0xc1] = cpu.op_cb_0xc1
	cpu.Instructions_maps.Instructions_map_CB[0xc2] = cpu.op_cb_0xc2
	cpu.Instructions_maps.Instructions_map_CB[0xc3] = cpu.op_cb_0xc3
	cpu.Instructions_maps.Instructions_map_CB[0xc4] = cpu.op_cb_0xc4
	cpu.Instructions_maps.Instructions_map_CB[0xc5] = cpu.op_cb_0xc5
	cpu.Instructions_maps.Instructions_map_CB[0xc6] = cpu.op_cb_0xc6
	cpu.Instructions_maps.Instructions_map_CB[0xc7] = cpu.op_cb_0xc7
	cpu.Instructions_maps.Instructions_map_CB[0xc8] = cpu.op_cb_0xc8
	cpu.Instructions_maps.Instructions_map_CB[0xc9] = cpu.op_cb_0xc9
	cpu.Instructions_maps.Instructions_map_CB[0xca] = cpu.op_cb_0xca
	cpu.Instructions_maps.Instructions_map_CB[0xcb] = cpu.op_cb_0xcb
	cpu.Instructions_maps.Instructions_map_CB[0xcc] = cpu.op_cb_0xcc
	cpu.Instructions_maps.Instructions_map_CB[0xcd] = cpu.op_cb_0xcd
	cpu.Instructions_maps.Instructions_map_CB[0xce] = cpu.op_cb_0xce
	cpu.Instructions_maps.Instructions_map_CB[0xcf] = cpu.op_cb_0xcf
	cpu.Instructions_maps.Instructions_map_CB[0xd0] = cpu.op_cb_0xd0
	cpu.Instructions_maps.Instructions_map_CB[0xd1] = cpu.op_cb_0xd1
	cpu.Instructions_maps.Instructions_map_CB[0xd2] = cpu.op_cb_0xd2
	cpu.Instructions_maps.Instructions_map_CB[0xd3] = cpu.op_cb_0xd3
	cpu.Instructions_maps.Instructions_map_CB[0xd4] = cpu.op_cb_0xd4
	cpu.Instructions_maps.Instructions_map_CB[0xd5] = cpu.op_cb_0xd5
	cpu.Instructions_maps.Instructions_map_CB[0xd6] = cpu.op_cb_0xd6
	cpu.Instructions_maps.Instructions_map_CB[0xd7] = cpu.op_cb_0xd7
	cpu.Instructions_maps.Instructions_map_CB[0xd8] = cpu.op_cb_0xd8
	cpu.Instructions_maps.Instructions_map_CB[0xd9] = cpu.op_cb_0xd9
	cpu.Instructions_maps.Instructions_map_CB[0xda] = cpu.op_cb_0xda
	cpu.Instructions_maps.Instructions_map_CB[0xdb] = cpu.op_cb_0xdb
	cpu.Instructions_maps.Instructions_map_CB[0xdc] = cpu.op_cb_0xdc
	cpu.Instructions_maps.Instructions_map_CB[0xdd] = cpu.op_cb_0xdd
	cpu.Instructions_maps.Instructions_map_CB[0xde] = cpu.op_cb_0xde
	cpu.Instructions_maps.Instructions_map_CB[0xdf] = cpu.op_cb_0xdf
	cpu.Instructions_maps.Instructions_map_CB[0xe0] = cpu.op_cb_0xe0
	cpu.Instructions_maps.Instructions_map_CB[0xe1] = cpu.op_cb_0xe1
	cpu.Instructions_maps.Instructions_map_CB[0xe2] = cpu.op_cb_0xe2
	cpu.Instructions_maps.Instructions_map_CB[0xe3] = cpu.op_cb_0xe3
	cpu.Instructions_maps.Instructions_map_CB[0xe4] = cpu.op_cb_0xe4
	cpu.Instructions_maps.Instructions_map_CB[0xe5] = cpu.op_cb_0xe5
	cpu.Instructions_maps.Instructions_map_CB[0xe6] = cpu.op_cb_0xe6
	cpu.Instructions_maps.Instructions_map_CB[0xe7] = cpu.op_cb_0xe7
	cpu.Instructions_maps.Instructions_map_CB[0xe8] = cpu.op_cb_0xe8
	cpu.Instructions_maps.Instructions_map_CB[0xe9] = cpu.op_cb_0xe9
	cpu.Instructions_maps.Instructions_map_CB[0xea] = cpu.op_cb_0xea
	cpu.Instructions_maps.Instructions_map_CB[0xeb] = cpu.op_cb_0xeb
	cpu.Instructions_maps.Instructions_map_CB[0xec] = cpu.op_cb_0xec
	cpu.Instructions_maps.Instructions_map_CB[0xed] = cpu.op_cb_0xed
	cpu.Instructions_maps.Instructions_map_CB[0xee] = cpu.op_cb_0xee
	cpu.Instructions_maps.Instructions_map_CB[0xef] = cpu.op_cb_0xef
	cpu.Instructions_maps.Instructions_map_CB[0xf0] = cpu.op_cb_0xf0
	cpu.Instructions_maps.Instructions_map_CB[0xf1] = cpu.op_cb_0xf1
	cpu.Instructions_maps.Instructions_map_CB[0xf2] = cpu.op_cb_0xf2
	cpu.Instructions_maps.Instructions_map_CB[0xf3] = cpu.op_cb_0xf3
	cpu.Instructions_maps.Instructions_map_CB[0xf4] = cpu.op_cb_0xf4
	cpu.Instructions_maps.Instructions_map_CB[0xf5] = cpu.op_cb_0xf5
	cpu.Instructions_maps.Instructions_map_CB[0xf6] = cpu.op_cb_0xf6
	cpu.Instructions_maps.Instructions_map_CB[0xf7] = cpu.op_cb_0xf7
	cpu.Instructions_maps.Instructions_map_CB[0xf8] = cpu.op_cb_0xf8
	cpu.Instructions_maps.Instructions_map_CB[0xf9] = cpu.op_cb_0xf9
	cpu.Instructions_maps.Instructions_map_CB[0xfa] = cpu.op_cb_0xfa
	cpu.Instructions_maps.Instructions_map_CB[0xfb] = cpu.op_cb_0xfb
	cpu.Instructions_maps.Instructions_map_CB[0xfc] = cpu.op_cb_0xfc
	cpu.Instructions_maps.Instructions_map_CB[0xfd] = cpu.op_cb_0xfd
	cpu.Instructions_maps.Instructions_map_CB[0xfe] = cpu.op_cb_0xfe
	cpu.Instructions_maps.Instructions_map_CB[0xff] = cpu.op_cb_0xff
}

func (cpu *CPU_struct) getAF() uint16 {
	return (uint16(cpu.Registers.A) << 8) | uint16(cpu.Registers.F&0xf0)
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
	cpu.Registers.F = uint8(nn & 0xf0)
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
	cpu.Registers.Halt = 1
	cpu.Cycle += 4
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

func (cpu *CPU_struct) inc(n *uint8) {
	*n += 1
	cpu.Registers.F &= C_BIT

	cpu.setZ(*n)

	if (((*n - 1) & 0x0f) + (1 & 0x0f)) > 0x0f {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 4
}

//0x04 INC B
func (cpu *CPU_struct) op_0x04() {
	cpu.inc(&cpu.Registers.B)
}

//0x14 INC D
func (cpu *CPU_struct) op_0x14() {
	cpu.inc(&cpu.Registers.D)
}

//0x24 INC H
func (cpu *CPU_struct) op_0x24() {
	cpu.inc(&cpu.Registers.H)
}

//0x34 INC (HL)
func (cpu *CPU_struct) op_0x34() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.inc(&n)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 4
}

func (cpu *CPU_struct) dec(n *uint8) {
	*n -= 1
	cpu.Registers.F &= C_BIT
	cpu.Registers.F |= N_BIT

	cpu.setZ(*n)

	if ((int8(*n+1) & 0xF) - (1 & 0xF)) < 0 {
		// if ((*n + 1) & 0x0f) == 0x00 {
		cpu.Registers.F |= H_BIT
	}

	cpu.Cycle += 4
}

// 0x05 DEC B
func (cpu *CPU_struct) op_0x05() {
	cpu.dec(&cpu.Registers.B)
}

// 0x15 DEC D
func (cpu *CPU_struct) op_0x15() {
	cpu.dec(&cpu.Registers.D)
}

// 0x25 DEC H
func (cpu *CPU_struct) op_0x25() {
	cpu.dec(&cpu.Registers.H)
}

// 0x35 DEC (HL)
func (cpu *CPU_struct) op_0x35() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.dec(&n)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 4
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

func (cpu *CPU_struct) rlca() {
	cpu.Registers.F = 0
	// bit 7 from A to carry
	cpu.Registers.F |= ((cpu.Registers.A & 0x80) >> 3)

	cpu.Registers.A = (cpu.Registers.A << 1) | (cpu.Registers.A >> 7)

	// cpu.setZ(cpu.Registers.A)

	cpu.Cycle += 8
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

	// cpu.setZ(cpu.Registers.A)

	cpu.Cycle += 8
}

// 0x07 RLCA
func (cpu *CPU_struct) op_0x07() {
	cpu.rlca()
}

// 0x17 RLA
func (cpu *CPU_struct) op_0x17() {
	cpu.rla()
}

// 0x27 DAA
// TODO: double check
func (cpu *CPU_struct) op_0x27() {
	var temp uint8 = cpu.Registers.A
	var corr uint8 = 0
	if (cpu.Registers.F & H_BIT) != 0 {
		corr |= 0x06
	}
	if (cpu.Registers.F & C_BIT) != 0 {
		corr |= 0x60
	}
	if (cpu.Registers.F & N_BIT) != 0 {
		temp -= corr
	} else {
		if (temp & 0x0F) > 0x09 {
			corr |= 0x06
		}

		if temp > 0x99 {
			corr |= 0x60
		}
		temp += corr
	}
	var flag uint8 = 0
	if (temp & 0xFF) == 0 {
		flag |= Z_BIT
	}
	if (corr & 0x60) != 0 {
		flag |= C_BIT
	}
	cpu.Registers.F &= 0b01000000
	cpu.Registers.F |= flag
	temp &= 0xFF
	cpu.Registers.A = temp

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
		cpu.Cycle += 12
		return
	}

	cpu.Cycle += 8
}

// 0x38 - JP C, nn
func (cpu *CPU_struct) op_0x38(n uint8) {
	if (cpu.Registers.F & C_BIT) != 0 {
		cpu.jr(n)
		cpu.Cycle += 12
		return
	}

	cpu.Cycle += 8
}

func (cpu *CPU_struct) add_nn(nn uint16) {
	var hl uint16 = cpu.getHL()
	cpu.setHL(hl + nn)

	cpu.Registers.F &= Z_BIT

	if ((hl & 0x0fff) + (nn & 0x0fff)) > 0xfff {
		cpu.Registers.F |= H_BIT
	}

	if (uint32(hl) + uint32(nn)) > 0xffff {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 8
}

// 0x09 ADD HL, BC
func (cpu *CPU_struct) op_0x09() {
	var nn uint16 = cpu.getBC()
	cpu.add_nn(nn)
}

// 0x19 ADD HL, DE
func (cpu *CPU_struct) op_0x19() {
	var nn uint16 = cpu.getDE()
	cpu.add_nn(nn)
}

// 0x29 ADD HL, HL
func (cpu *CPU_struct) op_0x29() {
	var nn uint16 = cpu.getHL()
	cpu.add_nn(nn)
}

// 0x39 ADD HL, SP
func (cpu *CPU_struct) op_0x39() {
	var nn uint16 = cpu.Registers.SP
	cpu.add_nn(nn)
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
	cpu.inc(&cpu.Registers.C)
}

//0x1c INC E
func (cpu *CPU_struct) op_0x1c() {
	cpu.inc(&cpu.Registers.E)
}

//0x2c INC L
func (cpu *CPU_struct) op_0x2c() {
	cpu.inc(&cpu.Registers.L)
}

//0x3c INC A
func (cpu *CPU_struct) op_0x3c() {
	cpu.inc(&cpu.Registers.A)
}

// 0x0d DEC C
func (cpu *CPU_struct) op_0x0d() {
	cpu.dec(&cpu.Registers.C)
}

// 0x1D DEC E
func (cpu *CPU_struct) op_0x1d() {
	cpu.dec(&cpu.Registers.E)
}

// 0x2D DEC L
func (cpu *CPU_struct) op_0x2d() {
	cpu.dec(&cpu.Registers.L)
}

// 0x3d DEC A
func (cpu *CPU_struct) op_0x3d() {
	cpu.dec(&cpu.Registers.A)
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

func (cpu *CPU_struct) rrca() {
	cpu.Registers.F = 0
	// bit 0 from A to carry
	cpu.Registers.F |= ((cpu.Registers.A & 0x01) << 4)

	cpu.Registers.A = (cpu.Registers.A >> 1) | ((cpu.Registers.A & 0x01) << 7)

	// cpu.setZ(cpu.Registers.A)

	cpu.Cycle += 8
}

func (cpu *CPU_struct) rra() {
	var carry uint8 = cpu.Registers.F & C_BIT
	cpu.Registers.F = 0
	// bit 0 from A to carry
	cpu.Registers.F |= ((cpu.Registers.A & 0x01) << 4)

	cpu.Registers.A = cpu.Registers.A >> 1

	if carry != 0 {
		cpu.Registers.A |= 0x80
	}

	cpu.setZ(cpu.Registers.A)

	cpu.Cycle += 4
}

// 0x0f RRCA
func (cpu *CPU_struct) op_0x0f() {
	cpu.rrca()
}

// 0x1f RRA
func (cpu *CPU_struct) op_0x1f() {
	var carry uint8 = cpu.Registers.F & C_BIT
	cpu.Registers.F = 0
	// bit 0 from n to carry
	cpu.Registers.F |= ((cpu.Registers.A & 0x01) << 4)

	cpu.Registers.A = (cpu.Registers.A >> 1)
	if carry != 0 {
		cpu.Registers.A |= 0x80
	}

	cpu.Cycle += 8
}

// 0x2f CPL
func (cpu *CPU_struct) op_0x2f() {
	cpu.Registers.A = cpu.Registers.A ^ 0xff
	cpu.Registers.F |= (N_BIT | H_BIT)

	cpu.Cycle += 4
}

// 0x3f CCF
func (cpu *CPU_struct) op_0x3f() {
	var carry uint8 = cpu.Registers.F & C_BIT
	cpu.Registers.F &= Z_BIT
	cpu.Registers.F |= carry ^ C_BIT

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

func (cpu *CPU_struct) add(n uint8, carry uint8) {
	old_a := cpu.Registers.A
	cpu.Registers.A += n + carry
	cpu.Registers.F = 0

	cpu.setZ(cpu.Registers.A)

	if ((old_a & 0x0f) + (n & 0x0f) + carry) > 0x0f {
		cpu.Registers.F |= H_BIT
	}

	if (uint16(old_a) + uint16(n) + uint16(carry)) > 0xff {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 4
}

// 0x80 - ADD A, B
func (cpu *CPU_struct) op_0x80() {
	cpu.add(cpu.Registers.B, 0)
}

// 0x81 - ADD A, C
func (cpu *CPU_struct) op_0x81() {
	cpu.add(cpu.Registers.C, 0)
}

// 0x82 - ADD A, D
func (cpu *CPU_struct) op_0x82() {
	cpu.add(cpu.Registers.D, 0)
}

// 0x83 - ADD A, E
func (cpu *CPU_struct) op_0x83() {
	cpu.add(cpu.Registers.E, 0)
}

// 0x84 - ADD A, H
func (cpu *CPU_struct) op_0x84() {
	cpu.add(cpu.Registers.H, 0)
}

// 0x85 - ADD A, L
func (cpu *CPU_struct) op_0x85() {
	cpu.add(cpu.Registers.L, 0)
}

// 0x86 - ADD A, (HL)
func (cpu *CPU_struct) op_0x86() {
	var address uint16 = cpu.getHL()

	cpu.add(cpu.MMU.ReadByte(address), 0)

	cpu.Cycle += 4
}

// 0x87 - ADD A, A
func (cpu *CPU_struct) op_0x87() {
	cpu.add(cpu.Registers.A, 0)
}

func (cpu *CPU_struct) adc(n uint8) {
	var carry uint8 = 0

	if (cpu.Registers.F & C_BIT) > 0 {
		carry += 1
	}

	cpu.add(n, carry)
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

func (cpu *CPU_struct) sub(n uint8, carry uint8) {
	old := cpu.Registers.A
	cpu.Registers.A = cpu.Registers.A - n - carry
	cpu.Registers.F = N_BIT

	cpu.setZ(cpu.Registers.A)

	if (old & 0x0f) < ((n & 0x0f) + carry) {
		cpu.Registers.F |= H_BIT
	}

	if uint32(old) < uint32(n)+uint32(carry) {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 4
}

// 0x90 - SUB A, B
func (cpu *CPU_struct) op_0x90() {
	cpu.sub(cpu.Registers.B, 0)
}

// 0x91 - SUB A, C
func (cpu *CPU_struct) op_0x91() {
	cpu.sub(cpu.Registers.C, 0)
}

// 0x92 - SUB A, D
func (cpu *CPU_struct) op_0x92() {
	cpu.sub(cpu.Registers.D, 0)
}

// 0x93 - SUB A, E
func (cpu *CPU_struct) op_0x93() {
	cpu.sub(cpu.Registers.E, 0)
}

// 0x94 - SUB A, H
func (cpu *CPU_struct) op_0x94() {
	cpu.sub(cpu.Registers.H, 0)
}

// 0x95 - SUB A, L
func (cpu *CPU_struct) op_0x95() {
	cpu.sub(cpu.Registers.L, 0)
}

// 0x96 - SUB A, (HL)
func (cpu *CPU_struct) op_0x96() {
	var address uint16 = cpu.getHL()

	cpu.sub(cpu.MMU.ReadByte(address), 0)

	cpu.Cycle += 4
}

// 0x97 - SUB A, A
func (cpu *CPU_struct) op_0x97() {
	cpu.sub(cpu.Registers.A, 0)
}

func (cpu *CPU_struct) sbc(n uint8) {
	var carry uint8 = 0

	if (cpu.Registers.F & C_BIT) > 0 {
		carry += 1
	}

	cpu.sub(n, carry)
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
	cpu.add(n, 0)

	cpu.Cycle += 4
}

// 0xd6 SUB #
func (cpu *CPU_struct) op_0xd6(n uint8) {
	cpu.sub(n, 0)

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

// 0xd8 RET C
func (cpu *CPU_struct) op_0xd8() {
	if (cpu.Registers.F & C_BIT) > 0 {
		cpu.ret()
	}
}

// 0xE8 ADD SP, n
func (cpu *CPU_struct) op_0xe8(n uint8) {
	old_sp := cpu.Registers.SP
	cpu.Registers.SP = uint16(int16(cpu.Registers.SP) + int16(int8(n)))
	cpu.Registers.F = 0

	if (((int16(old_sp & 0x0f)) + (int16(int8(n)) & 0x0f)) & 0xf0) != 0 {
		cpu.Registers.F |= H_BIT
	}

	if (((int32(old_sp & 0x0ff)) + (int32(int8(n)) & 0xff)) & 0xf00) != 0 {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 16
}

// LDHL SP, n
func (cpu *CPU_struct) op_0xf8(n uint8) {
	old_sp := cpu.Registers.SP
	cpu.setHL(uint16(int16(cpu.Registers.SP) + int16(int8(n))))
	cpu.Registers.F = 0

	if (((int16(old_sp & 0x0f)) + (int16(int8(n)) & 0x0f)) & 0xf0) != 0 {
		cpu.Registers.F |= H_BIT
	}

	if (((int32(old_sp & 0x0ff)) + (int32(int8(n)) & 0xff)) & 0xf00) != 0 {
		cpu.Registers.F |= C_BIT
	}

	cpu.Cycle += 16
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
	// cpu.Registers.IF = 0

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

func (cpu *CPU_struct) rlc(n *uint8) {
	// var carry uint8 = cpu.Registers.F & C_BIT
	cpu.Registers.F = 0
	// bit 7 from A to carry
	cpu.Registers.F |= ((*n & 0x80) >> 3)

	*n = (*n << 1) | (*n >> 7)
	// if carry != 0 {
	// 	*n |= 0x01
	// }

	cpu.setZ(*n)

	cpu.Cycle += 8
}

// 0xcb 0x00 RLC B
func (cpu *CPU_struct) op_cb_0x00() {
	cpu.rlc(&cpu.Registers.B)
}

// 0xcb 0x01 RLC C
func (cpu *CPU_struct) op_cb_0x01() {
	cpu.rlc(&cpu.Registers.C)
}

// 0xcb 0x02 RLC D
func (cpu *CPU_struct) op_cb_0x02() {
	cpu.rlc(&cpu.Registers.D)
}

// 0xcb 0x03 RLC E
func (cpu *CPU_struct) op_cb_0x03() {
	cpu.rlc(&cpu.Registers.E)
}

// 0xcb 0x04 RLC H
func (cpu *CPU_struct) op_cb_0x04() {
	cpu.rlc(&cpu.Registers.H)
}

// 0xcb 0x05 RLC L
func (cpu *CPU_struct) op_cb_0x05() {
	cpu.rlc(&cpu.Registers.L)
}

// 0xcb 0x06 RLC C
func (cpu *CPU_struct) op_cb_0x06() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.rlc(&n)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0x07 RLC A
func (cpu *CPU_struct) op_cb_0x07() {
	cpu.rlc(&cpu.Registers.A)
}

// TODO: double check if carry will go to bit 7
func (cpu *CPU_struct) rrc(n *uint8) {
	// var carry uint8 = cpu.Registers.F & C_BIT
	cpu.Registers.F = 0
	// bit 0 from n to carry
	cpu.Registers.F |= ((*n & 0x01) << 4)

	*n = (*n >> 1) | (*n << 7)
	// if carry != 0 {
	// 	*n |= 0x80
	// }

	cpu.setZ(*n)

	cpu.Cycle += 8
}

// 0xcb 0x08 RRC B
func (cpu *CPU_struct) op_cb_0x08() {
	cpu.rrc(&cpu.Registers.B)
}

// 0xcb 0x09 RRC C
func (cpu *CPU_struct) op_cb_0x09() {
	cpu.rrc(&cpu.Registers.C)
}

// 0xcb 0x0a RRC D
func (cpu *CPU_struct) op_cb_0x0a() {
	cpu.rrc(&cpu.Registers.D)
}

// 0xcb 0x0b RRC E
func (cpu *CPU_struct) op_cb_0x0b() {
	cpu.rrc(&cpu.Registers.E)
}

// 0xcb 0x0c RRC H
func (cpu *CPU_struct) op_cb_0x0c() {
	cpu.rrc(&cpu.Registers.H)
}

// 0xcb 0x0d RRC L
func (cpu *CPU_struct) op_cb_0x0d() {
	cpu.rrc(&cpu.Registers.L)
}

// 0xcb 0x0e RRC (HL)
func (cpu *CPU_struct) op_cb_0x0e() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.rrc(&n)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0x0f RRC A
func (cpu *CPU_struct) op_cb_0x0f() {
	cpu.rrc(&cpu.Registers.A)
}

func (cpu *CPU_struct) rl(n *uint8) {
	var carry uint8 = cpu.Registers.F & C_BIT
	cpu.Registers.F = 0
	// bit 7 from A to carry
	cpu.Registers.F |= ((*n & 0x80) >> 3)

	*n = (*n << 1)
	if carry != 0 {
		*n |= 0x01
	}

	cpu.setZ(*n)

	cpu.Cycle += 8
}

// 0xcb 0x10 RL B
func (cpu *CPU_struct) op_cb_0x10() {
	cpu.rl(&cpu.Registers.B)
}

// 0xcb 0x11 RL C
func (cpu *CPU_struct) op_cb_0x11() {
	cpu.rl(&cpu.Registers.C)
}

// 0xcb 0x12 RL D
func (cpu *CPU_struct) op_cb_0x12() {
	cpu.rl(&cpu.Registers.D)
}

// 0xcb 0x13 RL E
func (cpu *CPU_struct) op_cb_0x13() {
	cpu.rl(&cpu.Registers.E)
}

// 0xcb 0x14 RL H
func (cpu *CPU_struct) op_cb_0x14() {
	cpu.rl(&cpu.Registers.H)
}

// 0xcb 0x15 RL L
func (cpu *CPU_struct) op_cb_0x15() {
	cpu.rl(&cpu.Registers.L)
}

// 0xcb 0x16 RL C
func (cpu *CPU_struct) op_cb_0x16() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.rl(&n)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0x17 RL A
func (cpu *CPU_struct) op_cb_0x17() {
	cpu.rl(&cpu.Registers.A)
}

// TODO: double check if carry will go to bit 7
func (cpu *CPU_struct) rr(n *uint8) {
	var carry uint8 = cpu.Registers.F & C_BIT
	cpu.Registers.F = 0
	// bit 0 from n to carry
	cpu.Registers.F |= ((*n & 0x01) << 4)

	*n = (*n >> 1)
	if carry != 0 {
		*n |= 0x80
	}

	cpu.setZ(*n)

	cpu.Cycle += 8
}

// 0xcb 0x18 RR B
func (cpu *CPU_struct) op_cb_0x18() {
	cpu.rr(&cpu.Registers.B)
}

// 0xcb 0x19 RR C
func (cpu *CPU_struct) op_cb_0x19() {
	cpu.rr(&cpu.Registers.C)
}

// 0xcb 0x1a RR D
func (cpu *CPU_struct) op_cb_0x1a() {
	cpu.rr(&cpu.Registers.D)
}

// 0xcb 0x1b RR E
func (cpu *CPU_struct) op_cb_0x1b() {
	cpu.rr(&cpu.Registers.E)
}

// 0xcb 0x1c RR H
func (cpu *CPU_struct) op_cb_0x1c() {
	cpu.rr(&cpu.Registers.H)
}

// 0xcb 0x1d RR L
func (cpu *CPU_struct) op_cb_0x1d() {
	cpu.rr(&cpu.Registers.L)
}

// 0xcb 0x1e RR (HL)
func (cpu *CPU_struct) op_cb_0x1e() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.rr(&n)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0x1f RR A
func (cpu *CPU_struct) op_cb_0x1f() {
	cpu.rr(&cpu.Registers.A)
}

func (cpu *CPU_struct) sla(n *uint8) {
	cpu.Registers.F = 0
	// bit 7 from A to carry
	cpu.Registers.F |= ((*n & 0x80) >> 3)

	*n = (*n << 1)

	cpu.setZ(*n)

	cpu.Cycle += 8
}

// 0xcb 0x20 SLA B
func (cpu *CPU_struct) op_cb_0x20() {
	cpu.sla(&cpu.Registers.B)
}

// 0xcb 0x21 SLA C
func (cpu *CPU_struct) op_cb_0x21() {
	cpu.sla(&cpu.Registers.C)
}

// 0xcb 0x22 SLA D
func (cpu *CPU_struct) op_cb_0x22() {
	cpu.sla(&cpu.Registers.D)
}

// 0xcb 0x23 SLA E
func (cpu *CPU_struct) op_cb_0x23() {
	cpu.sla(&cpu.Registers.E)
}

// 0xcb 0x24 SLA H
func (cpu *CPU_struct) op_cb_0x24() {
	cpu.sla(&cpu.Registers.H)
}

// 0xcb 0x25 SLA L
func (cpu *CPU_struct) op_cb_0x25() {
	cpu.sla(&cpu.Registers.L)
}

// 0xcb 0x26 SLA C
func (cpu *CPU_struct) op_cb_0x26() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.sla(&n)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0x27 SLA A
func (cpu *CPU_struct) op_cb_0x27() {
	cpu.sla(&cpu.Registers.A)
}

func (cpu *CPU_struct) sra(n *uint8) {
	cpu.Registers.F = 0
	// bit 0 from n to carry
	cpu.Registers.F |= ((*n & 0x01) << 4)

	*n = (*n >> 1) | (*n & 0x80)

	cpu.setZ(*n)

	cpu.Cycle += 8
}

// 0xcb 0x28 SRA B
func (cpu *CPU_struct) op_cb_0x28() {
	cpu.sra(&cpu.Registers.B)
}

// 0xcb 0x29 SRA C
func (cpu *CPU_struct) op_cb_0x29() {
	cpu.sra(&cpu.Registers.C)
}

// 0xcb 0x2a SRA D
func (cpu *CPU_struct) op_cb_0x2a() {
	cpu.sra(&cpu.Registers.D)
}

// 0xcb 0x2b SRA E
func (cpu *CPU_struct) op_cb_0x2b() {
	cpu.sra(&cpu.Registers.E)
}

// 0xcb 0x2c SRA H
func (cpu *CPU_struct) op_cb_0x2c() {
	cpu.sra(&cpu.Registers.H)
}

// 0xcb 0x2d SRA L
func (cpu *CPU_struct) op_cb_0x2d() {
	cpu.sra(&cpu.Registers.L)
}

// 0xcb 0x2e SRA (HL)
func (cpu *CPU_struct) op_cb_0x2e() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.sra(&n)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0x2f SRA A
func (cpu *CPU_struct) op_cb_0x2f() {
	cpu.sra(&cpu.Registers.A)
}

func (cpu *CPU_struct) swap(n *uint8) {
	*n = (*n << 4) | (*n >> 4)

	cpu.Registers.F = 0
	cpu.setZ(*n)

	cpu.Cycle += 8
}

// 0xcb 0x30 SWAP B
func (cpu *CPU_struct) op_cb_0x30() {
	cpu.swap(&cpu.Registers.B)
}

// 0xcb 0x31 SWAP C
func (cpu *CPU_struct) op_cb_0x31() {
	cpu.swap(&cpu.Registers.C)
}

// 0xcb 0x32 SWAP D
func (cpu *CPU_struct) op_cb_0x32() {
	cpu.swap(&cpu.Registers.D)
}

// 0xcb 0x33 SWAP E
func (cpu *CPU_struct) op_cb_0x33() {
	cpu.swap(&cpu.Registers.E)
}

// 0xcb 0x34 SWAP H
func (cpu *CPU_struct) op_cb_0x34() {
	cpu.swap(&cpu.Registers.H)
}

// 0xcb 0x35 SWAP L
func (cpu *CPU_struct) op_cb_0x35() {
	cpu.swap(&cpu.Registers.L)
}

// 0xcb 0x36 SWAP C
func (cpu *CPU_struct) op_cb_0x36() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.swap(&n)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0x37 SWAP A
func (cpu *CPU_struct) op_cb_0x37() {
	cpu.swap(&cpu.Registers.A)
}

func (cpu *CPU_struct) srl(n *uint8) {
	cpu.Registers.F = 0
	// bit 0 from n to carry
	cpu.Registers.F |= ((*n & 0x01) << 4)

	*n = (*n >> 1)

	cpu.setZ(*n)

	cpu.Cycle += 8
}

// 0xcb 0x38 SRL B
func (cpu *CPU_struct) op_cb_0x38() {
	cpu.srl(&cpu.Registers.B)
}

// 0xcb 0x39 SRL C
func (cpu *CPU_struct) op_cb_0x39() {
	cpu.srl(&cpu.Registers.C)
}

// 0xcb 0x3a SRL D
func (cpu *CPU_struct) op_cb_0x3a() {
	cpu.srl(&cpu.Registers.D)
}

// 0xcb 0x3b SRL E
func (cpu *CPU_struct) op_cb_0x3b() {
	cpu.srl(&cpu.Registers.E)
}

// 0xcb 0x3c SRL H
func (cpu *CPU_struct) op_cb_0x3c() {
	cpu.srl(&cpu.Registers.H)
}

// 0xcb 0x3d SRL L
func (cpu *CPU_struct) op_cb_0x3d() {
	cpu.srl(&cpu.Registers.L)
}

// 0xcb 0x3e SRL (HL)
func (cpu *CPU_struct) op_cb_0x3e() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.srl(&n)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0x3f SRL A
func (cpu *CPU_struct) op_cb_0x3f() {
	cpu.srl(&cpu.Registers.A)
}

func (cpu *CPU_struct) bit(n uint8, mask uint8) {
	cpu.Registers.F &= C_BIT
	cpu.Registers.F |= H_BIT

	cpu.setZ(n & mask)

	cpu.Cycle += 8
}

// 0xcb 0x40 BIT 0, B
func (cpu *CPU_struct) op_cb_0x40() {
	cpu.bit(cpu.Registers.B, 0b00000001)
}

// 0xcb 0x41 BIT 0, C
func (cpu *CPU_struct) op_cb_0x41() {
	cpu.bit(cpu.Registers.C, 0b00000001)
}

// 0xcb 0x42 BIT 0, D
func (cpu *CPU_struct) op_cb_0x42() {
	cpu.bit(cpu.Registers.D, 0b00000001)
}

// 0xcb 0x43 BIT 0, E
func (cpu *CPU_struct) op_cb_0x43() {
	cpu.bit(cpu.Registers.E, 0b00000001)
}

// 0xcb 0x44 BIT 0, H
func (cpu *CPU_struct) op_cb_0x44() {
	cpu.bit(cpu.Registers.H, 0b00000001)
}

// 0xcb 0x45 BIT 0, L
func (cpu *CPU_struct) op_cb_0x45() {
	cpu.bit(cpu.Registers.L, 0b00000001)
}

// 0xcb 0x46 BIT 0, (HL)
func (cpu *CPU_struct) op_cb_0x46() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.bit(n, 0b00000001)

	cpu.Cycle += 8
}

// 0xcb 0x47 BIT 0, A
func (cpu *CPU_struct) op_cb_0x47() {
	cpu.bit(cpu.Registers.A, 0b00000001)
}

// 0xcb 0x48 BIT 1, B
func (cpu *CPU_struct) op_cb_0x48() {
	cpu.bit(cpu.Registers.B, 0b00000010)
}

// 0xcb 0x49 BIT 1, C
func (cpu *CPU_struct) op_cb_0x49() {
	cpu.bit(cpu.Registers.C, 0b00000010)
}

// 0xcb 0x4a BIT 1, D
func (cpu *CPU_struct) op_cb_0x4a() {
	cpu.bit(cpu.Registers.D, 0b00000010)
}

// 0xcb 0x4b BIT 1, E
func (cpu *CPU_struct) op_cb_0x4b() {
	cpu.bit(cpu.Registers.E, 0b00000010)
}

// 0xcb 0x4c BIT 1, H
func (cpu *CPU_struct) op_cb_0x4c() {
	cpu.bit(cpu.Registers.H, 0b00000010)
}

// 0xcb 0x4d BIT 1, L
func (cpu *CPU_struct) op_cb_0x4d() {
	cpu.bit(cpu.Registers.L, 0b00000010)
}

// 0xcb 0x4e BIT 1, (HL)
func (cpu *CPU_struct) op_cb_0x4e() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.bit(n, 0b00000010)

	cpu.Cycle += 8
}

// 0xcb 0x4f BIT 1, A
func (cpu *CPU_struct) op_cb_0x4f() {
	cpu.bit(cpu.Registers.A, 0b00000010)
}

// 0xcb 0x50 BIT 2, B
func (cpu *CPU_struct) op_cb_0x50() {
	cpu.bit(cpu.Registers.B, 0b00000100)
}

// 0xcb 0x51 BIT 2, C
func (cpu *CPU_struct) op_cb_0x51() {
	cpu.bit(cpu.Registers.C, 0b00000100)
}

// 0xcb 0x52 BIT 2, D
func (cpu *CPU_struct) op_cb_0x52() {
	cpu.bit(cpu.Registers.D, 0b00000100)
}

// 0xcb 0x53 BIT 2, E
func (cpu *CPU_struct) op_cb_0x53() {
	cpu.bit(cpu.Registers.E, 0b00000100)
}

// 0xcb 0x54 BIT 2, H
func (cpu *CPU_struct) op_cb_0x54() {
	cpu.bit(cpu.Registers.H, 0b00000100)
}

// 0xcb 0x55 BIT 2, L
func (cpu *CPU_struct) op_cb_0x55() {
	cpu.bit(cpu.Registers.L, 0b00000100)
}

// 0xcb 0x56 BIT 2, (HL)
func (cpu *CPU_struct) op_cb_0x56() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.bit(n, 0b00000100)

	cpu.Cycle += 8
}

// 0xcb 0x57 BIT 2, A
func (cpu *CPU_struct) op_cb_0x57() {
	cpu.bit(cpu.Registers.A, 0b00000100)
}

// 0xcb 0x58 BIT 3, B
func (cpu *CPU_struct) op_cb_0x58() {
	cpu.bit(cpu.Registers.B, 0b00001000)
}

// 0xcb 0x59 BIT 3, C
func (cpu *CPU_struct) op_cb_0x59() {
	cpu.bit(cpu.Registers.C, 0b00001000)
}

// 0xcb 0x5a BIT 3, D
func (cpu *CPU_struct) op_cb_0x5a() {
	cpu.bit(cpu.Registers.D, 0b00001000)
}

// 0xcb 0x5b BIT 3, E
func (cpu *CPU_struct) op_cb_0x5b() {
	cpu.bit(cpu.Registers.E, 0b00001000)
}

// 0xcb 0x5c BIT 3, H
func (cpu *CPU_struct) op_cb_0x5c() {
	cpu.bit(cpu.Registers.H, 0b00001000)
}

// 0xcb 0x5d BIT 3, L
func (cpu *CPU_struct) op_cb_0x5d() {
	cpu.bit(cpu.Registers.L, 0b00001000)
}

// 0xcb 0x5e BIT 3, (HL)
func (cpu *CPU_struct) op_cb_0x5e() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.bit(n, 0b00001000)

	cpu.Cycle += 8
}

// 0xcb 0x5f BIT 3, A
func (cpu *CPU_struct) op_cb_0x5f() {
	cpu.bit(cpu.Registers.A, 0b00001000)
}

// 0xcb 0x60 BIT 4, B
func (cpu *CPU_struct) op_cb_0x60() {
	cpu.bit(cpu.Registers.B, 0b00010000)
}

// 0xcb 0x61 BIT 4, C
func (cpu *CPU_struct) op_cb_0x61() {
	cpu.bit(cpu.Registers.C, 0b00010000)
}

// 0xcb 0x62 BIT 4, D
func (cpu *CPU_struct) op_cb_0x62() {
	cpu.bit(cpu.Registers.D, 0b00010000)
}

// 0xcb 0x63 BIT 4, E
func (cpu *CPU_struct) op_cb_0x63() {
	cpu.bit(cpu.Registers.E, 0b00010000)
}

// 0xcb 0x64 BIT 4, H
func (cpu *CPU_struct) op_cb_0x64() {
	cpu.bit(cpu.Registers.H, 0b00010000)
}

// 0xcb 0x65 BIT 4, L
func (cpu *CPU_struct) op_cb_0x65() {
	cpu.bit(cpu.Registers.L, 0b00010000)
}

// 0xcb 0x66 BIT 4, (HL)
func (cpu *CPU_struct) op_cb_0x66() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.bit(n, 0b00010000)

	cpu.Cycle += 8
}

// 0xcb 0x67 BIT 4, A
func (cpu *CPU_struct) op_cb_0x67() {
	cpu.bit(cpu.Registers.A, 0b00010000)
}

// 0xcb 0x68 BIT 5, B
func (cpu *CPU_struct) op_cb_0x68() {
	cpu.bit(cpu.Registers.B, 0b00100000)
}

// 0xcb 0x69 BIT 5, C
func (cpu *CPU_struct) op_cb_0x69() {
	cpu.bit(cpu.Registers.C, 0b00100000)
}

// 0xcb 0x6a BIT 5, D
func (cpu *CPU_struct) op_cb_0x6a() {
	cpu.bit(cpu.Registers.D, 0b00100000)
}

// 0xcb 0x6b BIT 5, E
func (cpu *CPU_struct) op_cb_0x6b() {
	cpu.bit(cpu.Registers.E, 0b00100000)
}

// 0xcb 0x6c BIT 5, H
func (cpu *CPU_struct) op_cb_0x6c() {
	cpu.bit(cpu.Registers.H, 0b00100000)
}

// 0xcb 0x6d BIT 5, L
func (cpu *CPU_struct) op_cb_0x6d() {
	cpu.bit(cpu.Registers.L, 0b00100000)
}

// 0xcb 0x6e BIT 5, (HL)
func (cpu *CPU_struct) op_cb_0x6e() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.bit(n, 0b00100000)

	cpu.Cycle += 8
}

// 0xcb 0x6f BIT 5, A
func (cpu *CPU_struct) op_cb_0x6f() {
	cpu.bit(cpu.Registers.A, 0b00100000)
}

// 0xcb 0x70 BIT 6, B
func (cpu *CPU_struct) op_cb_0x70() {
	cpu.bit(cpu.Registers.B, 0b01000000)
}

// 0xcb 0x71 BIT 6, C
func (cpu *CPU_struct) op_cb_0x71() {
	cpu.bit(cpu.Registers.C, 0b01000000)
}

// 0xcb 0x72 BIT 6, D
func (cpu *CPU_struct) op_cb_0x72() {
	cpu.bit(cpu.Registers.D, 0b01000000)
}

// 0xcb 0x73 BIT 6, E
func (cpu *CPU_struct) op_cb_0x73() {
	cpu.bit(cpu.Registers.E, 0b01000000)
}

// 0xcb 0x74 BIT 6, H
func (cpu *CPU_struct) op_cb_0x74() {
	cpu.bit(cpu.Registers.H, 0b01000000)
}

// 0xcb 0x75 BIT 6, L
func (cpu *CPU_struct) op_cb_0x75() {
	cpu.bit(cpu.Registers.L, 0b01000000)
}

// 0xcb 0x76 BIT 6, (HL)
func (cpu *CPU_struct) op_cb_0x76() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.bit(n, 0b01000000)

	cpu.Cycle += 8
}

// 0xcb 0x77 BIT 6, A
func (cpu *CPU_struct) op_cb_0x77() {
	cpu.bit(cpu.Registers.A, 0b01000000)
}

// 0xcb 0x78 BIT 7, B
func (cpu *CPU_struct) op_cb_0x78() {
	cpu.bit(cpu.Registers.B, 0b10000000)
}

// 0xcb 0x79 BIT 7, C
func (cpu *CPU_struct) op_cb_0x79() {
	cpu.bit(cpu.Registers.C, 0b10000000)
}

// 0xcb 0x7a BIT 7, D
func (cpu *CPU_struct) op_cb_0x7a() {
	cpu.bit(cpu.Registers.D, 0b10000000)
}

// 0xcb 0x7b BIT 7, E
func (cpu *CPU_struct) op_cb_0x7b() {
	cpu.bit(cpu.Registers.E, 0b10000000)
}

// 0xcb 0x7c BIT 7, H
func (cpu *CPU_struct) op_cb_0x7c() {
	cpu.bit(cpu.Registers.H, 0b10000000)
}

// 0xcb 0x7d BIT 7, L
func (cpu *CPU_struct) op_cb_0x7d() {
	cpu.bit(cpu.Registers.L, 0b10000000)
}

// 0xcb 0x7e BIT 7, (HL)
func (cpu *CPU_struct) op_cb_0x7e() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.bit(n, 0b10000000)

	cpu.Cycle += 8
}

// 0xcb 0x7f BIT 7, A
func (cpu *CPU_struct) op_cb_0x7f() {
	cpu.bit(cpu.Registers.A, 0b10000000)
}

func (cpu *CPU_struct) res(n *uint8, mask uint8) {
	*n &= mask

	cpu.Cycle += 8
}

// 0xcb 0x80 RES 0, B
func (cpu *CPU_struct) op_cb_0x80() {
	cpu.res(&cpu.Registers.B, 0b11111110)
}

// 0xcb 0x81 RES 0, C
func (cpu *CPU_struct) op_cb_0x81() {
	cpu.res(&cpu.Registers.C, 0b11111110)
}

// 0xcb 0x82 RES 0, D
func (cpu *CPU_struct) op_cb_0x82() {
	cpu.res(&cpu.Registers.D, 0b11111110)
}

// 0xcb 0x83 RES 0, E
func (cpu *CPU_struct) op_cb_0x83() {
	cpu.res(&cpu.Registers.E, 0b11111110)
}

// 0xcb 0x84 RES 0, H
func (cpu *CPU_struct) op_cb_0x84() {
	cpu.res(&cpu.Registers.H, 0b11111110)
}

// 0xcb 0x85 RES 0, L
func (cpu *CPU_struct) op_cb_0x85() {
	cpu.res(&cpu.Registers.L, 0b11111110)
}

// 0xcb 0x86 RES 0, (HL)
func (cpu *CPU_struct) op_cb_0x86() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.res(&n, 0b11111110)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0x87 RES 0, A
func (cpu *CPU_struct) op_cb_0x87() {
	cpu.res(&cpu.Registers.A, 0b11111110)
}

// 0xcb 0x88 RES 1, B
func (cpu *CPU_struct) op_cb_0x88() {
	cpu.res(&cpu.Registers.B, 0b11111101)
}

// 0xcb 0x89 RES 1, C
func (cpu *CPU_struct) op_cb_0x89() {
	cpu.res(&cpu.Registers.C, 0b11111101)
}

// 0xcb 0x8a RES 1, D
func (cpu *CPU_struct) op_cb_0x8a() {
	cpu.res(&cpu.Registers.D, 0b11111101)
}

// 0xcb 0x8b RES 1, E
func (cpu *CPU_struct) op_cb_0x8b() {
	cpu.res(&cpu.Registers.E, 0b11111101)
}

// 0xcb 0x8c RES 1, H
func (cpu *CPU_struct) op_cb_0x8c() {
	cpu.res(&cpu.Registers.H, 0b11111101)
}

// 0xcb 0x8d RES 1, L
func (cpu *CPU_struct) op_cb_0x8d() {
	cpu.res(&cpu.Registers.L, 0b11111101)
}

// 0xcb 0x8e RES 1, (HL)
func (cpu *CPU_struct) op_cb_0x8e() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.res(&n, 0b11111101)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0x8f RES 1, A
func (cpu *CPU_struct) op_cb_0x8f() {
	cpu.res(&cpu.Registers.A, 0b11111101)
}

// 0xcb 0x90 RES 2, B
func (cpu *CPU_struct) op_cb_0x90() {
	cpu.res(&cpu.Registers.B, 0b11111011)
}

// 0xcb 0x91 RES 2, C
func (cpu *CPU_struct) op_cb_0x91() {
	cpu.res(&cpu.Registers.C, 0b11111011)
}

// 0xcb 0x92 RES 2, D
func (cpu *CPU_struct) op_cb_0x92() {
	cpu.res(&cpu.Registers.D, 0b11111011)
}

// 0xcb 0x93 RES 2, E
func (cpu *CPU_struct) op_cb_0x93() {
	cpu.res(&cpu.Registers.E, 0b11111011)
}

// 0xcb 0x94 RES 2, H
func (cpu *CPU_struct) op_cb_0x94() {
	cpu.res(&cpu.Registers.H, 0b11111011)
}

// 0xcb 0x95 RES 2, L
func (cpu *CPU_struct) op_cb_0x95() {
	cpu.res(&cpu.Registers.L, 0b11111011)
}

// 0xcb 0x96 RES 2, (HL)
func (cpu *CPU_struct) op_cb_0x96() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.res(&n, 0b11111011)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0x97 RES 2, A
func (cpu *CPU_struct) op_cb_0x97() {
	cpu.res(&cpu.Registers.A, 0b11111011)
}

// 0xcb 0x98 RES 3, B
func (cpu *CPU_struct) op_cb_0x98() {
	cpu.res(&cpu.Registers.B, 0b11110111)
}

// 0xcb 0x99 RES 3, C
func (cpu *CPU_struct) op_cb_0x99() {
	cpu.res(&cpu.Registers.C, 0b11110111)
}

// 0xcb 0x9a RES 3, D
func (cpu *CPU_struct) op_cb_0x9a() {
	cpu.res(&cpu.Registers.D, 0b11110111)
}

// 0xcb 0x9b RES 3, E
func (cpu *CPU_struct) op_cb_0x9b() {
	cpu.res(&cpu.Registers.E, 0b11110111)
}

// 0xcb 0x9c RES 3, H
func (cpu *CPU_struct) op_cb_0x9c() {
	cpu.res(&cpu.Registers.H, 0b11110111)
}

// 0xcb 0x9d RES 3, L
func (cpu *CPU_struct) op_cb_0x9d() {
	cpu.res(&cpu.Registers.L, 0b11110111)
}

// 0xcb 0x9e RES 3, (HL)
func (cpu *CPU_struct) op_cb_0x9e() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.res(&n, 0b11110111)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0x9f RES 3, A
func (cpu *CPU_struct) op_cb_0x9f() {
	cpu.res(&cpu.Registers.A, 0b11110111)
}

// 0xcb 0xa0 RES 4, B
func (cpu *CPU_struct) op_cb_0xa0() {
	cpu.res(&cpu.Registers.B, 0b11101111)
}

// 0xcb 0xa1 RES 4, C
func (cpu *CPU_struct) op_cb_0xa1() {
	cpu.res(&cpu.Registers.C, 0b11101111)
}

// 0xcb 0xa2 RES 4, D
func (cpu *CPU_struct) op_cb_0xa2() {
	cpu.res(&cpu.Registers.D, 0b11101111)
}

// 0xcb 0xa3 RES 4, E
func (cpu *CPU_struct) op_cb_0xa3() {
	cpu.res(&cpu.Registers.E, 0b11101111)
}

// 0xcb 0xa4 RES 4, H
func (cpu *CPU_struct) op_cb_0xa4() {
	cpu.res(&cpu.Registers.H, 0b11101111)
}

// 0xcb 0xa5 RES 4, L
func (cpu *CPU_struct) op_cb_0xa5() {
	cpu.res(&cpu.Registers.L, 0b11101111)
}

// 0xcb 0xa6 RES 4, (HL)
func (cpu *CPU_struct) op_cb_0xa6() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.res(&n, 0b11101111)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0xa7 RES 4, A
func (cpu *CPU_struct) op_cb_0xa7() {
	cpu.res(&cpu.Registers.A, 0b11101111)
}

// 0xcb 0xa8 RES 5, B
func (cpu *CPU_struct) op_cb_0xa8() {
	cpu.res(&cpu.Registers.B, 0b11011111)
}

// 0xcb 0xa9 RES 5, C
func (cpu *CPU_struct) op_cb_0xa9() {
	cpu.res(&cpu.Registers.C, 0b11011111)
}

// 0xcb 0xaa RES 5, D
func (cpu *CPU_struct) op_cb_0xaa() {
	cpu.res(&cpu.Registers.D, 0b11011111)
}

// 0xcb 0xab RES 5, E
func (cpu *CPU_struct) op_cb_0xab() {
	cpu.res(&cpu.Registers.E, 0b11011111)
}

// 0xcb 0xac RES 5, H
func (cpu *CPU_struct) op_cb_0xac() {
	cpu.res(&cpu.Registers.H, 0b11011111)
}

// 0xcb 0xad RES 5, L
func (cpu *CPU_struct) op_cb_0xad() {
	cpu.res(&cpu.Registers.L, 0b11011111)
}

// 0xcb 0xae RES 5, (HL)
func (cpu *CPU_struct) op_cb_0xae() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.res(&n, 0b11011111)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0xaf RES 5, A
func (cpu *CPU_struct) op_cb_0xaf() {
	cpu.res(&cpu.Registers.A, 0b11011111)
}

// 0xcb 0xb0 RES 6, B
func (cpu *CPU_struct) op_cb_0xb0() {
	cpu.res(&cpu.Registers.B, 0b10111111)
}

// 0xcb 0xb1 RES 6, C
func (cpu *CPU_struct) op_cb_0xb1() {
	cpu.res(&cpu.Registers.C, 0b10111111)
}

// 0xcb 0xb2 RES 6, D
func (cpu *CPU_struct) op_cb_0xb2() {
	cpu.res(&cpu.Registers.D, 0b10111111)
}

// 0xcb 0xb3 RES 6, E
func (cpu *CPU_struct) op_cb_0xb3() {
	cpu.res(&cpu.Registers.E, 0b10111111)
}

// 0xcb 0xb4 RES 6, H
func (cpu *CPU_struct) op_cb_0xb4() {
	cpu.res(&cpu.Registers.H, 0b10111111)
}

// 0xcb 0xb5 RES 6, L
func (cpu *CPU_struct) op_cb_0xb5() {
	cpu.res(&cpu.Registers.L, 0b10111111)
}

// 0xcb 0xb6 RES 6, (HL)
func (cpu *CPU_struct) op_cb_0xb6() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.res(&n, 0b10111111)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0xb7 RES 6, A
func (cpu *CPU_struct) op_cb_0xb7() {
	cpu.res(&cpu.Registers.A, 0b10111111)
}

// 0xcb 0xb8 RES 7, B
func (cpu *CPU_struct) op_cb_0xb8() {
	cpu.res(&cpu.Registers.B, 0b01111111)
}

// 0xcb 0xb9 RES 7, C
func (cpu *CPU_struct) op_cb_0xb9() {
	cpu.res(&cpu.Registers.C, 0b01111111)
}

// 0xcb 0xba RES 7, D
func (cpu *CPU_struct) op_cb_0xba() {
	cpu.res(&cpu.Registers.D, 0b01111111)
}

// 0xcb 0xbb RES 7, E
func (cpu *CPU_struct) op_cb_0xbb() {
	cpu.res(&cpu.Registers.E, 0b01111111)
}

// 0xcb 0xbc RES 7, H
func (cpu *CPU_struct) op_cb_0xbc() {
	cpu.res(&cpu.Registers.H, 0b01111111)
}

// 0xcb 0xbd RES 7, L
func (cpu *CPU_struct) op_cb_0xbd() {
	cpu.res(&cpu.Registers.L, 0b01111111)
}

// 0xcb 0xbe RES 7, (HL)
func (cpu *CPU_struct) op_cb_0xbe() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.res(&n, 0b01111111)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0xbf RES 7, A
func (cpu *CPU_struct) op_cb_0xbf() {
	cpu.res(&cpu.Registers.A, 0b01111111)
}

func (cpu *CPU_struct) set(n *uint8, mask uint8) {
	*n |= mask

	cpu.Cycle += 8
}

// 0xcb 0xc0 SET 0, B
func (cpu *CPU_struct) op_cb_0xc0() {
	cpu.set(&cpu.Registers.B, 0b00000001)
}

// 0xcb 0xc1 SET 0, C
func (cpu *CPU_struct) op_cb_0xc1() {
	cpu.set(&cpu.Registers.C, 0b00000001)
}

// 0xcb 0xc2 SET 0, D
func (cpu *CPU_struct) op_cb_0xc2() {
	cpu.set(&cpu.Registers.D, 0b00000001)
}

// 0xcb 0xc3 SET 0, E
func (cpu *CPU_struct) op_cb_0xc3() {
	cpu.set(&cpu.Registers.E, 0b00000001)
}

// 0xcb 0xc4 SET 0, H
func (cpu *CPU_struct) op_cb_0xc4() {
	cpu.set(&cpu.Registers.H, 0b00000001)
}

// 0xcb 0xc5 SET 0, L
func (cpu *CPU_struct) op_cb_0xc5() {
	cpu.set(&cpu.Registers.L, 0b00000001)
}

// 0xcb 0xc6 SET 0, (HL)
func (cpu *CPU_struct) op_cb_0xc6() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.set(&n, 0b00000001)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0xc7 SET 0, A
func (cpu *CPU_struct) op_cb_0xc7() {
	cpu.set(&cpu.Registers.A, 0b00000001)
}

// 0xcb 0xc8 SET 1, B
func (cpu *CPU_struct) op_cb_0xc8() {
	cpu.set(&cpu.Registers.B, 0b00000010)
}

// 0xcb 0xc9 SET 1, C
func (cpu *CPU_struct) op_cb_0xc9() {
	cpu.set(&cpu.Registers.C, 0b00000010)
}

// 0xcb 0xca SET 1, D
func (cpu *CPU_struct) op_cb_0xca() {
	cpu.set(&cpu.Registers.D, 0b00000010)
}

// 0xcb 0xcb SET 1, E
func (cpu *CPU_struct) op_cb_0xcb() {
	cpu.set(&cpu.Registers.E, 0b00000010)
}

// 0xcb 0xcc SET 1, H
func (cpu *CPU_struct) op_cb_0xcc() {
	cpu.set(&cpu.Registers.H, 0b00000010)
}

// 0xcb 0xcd SET 1, L
func (cpu *CPU_struct) op_cb_0xcd() {
	cpu.set(&cpu.Registers.L, 0b00000010)
}

// 0xcb 0xce SET 1, (HL)
func (cpu *CPU_struct) op_cb_0xce() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.set(&n, 0b00000010)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0xcf SET 1, A
func (cpu *CPU_struct) op_cb_0xcf() {
	cpu.set(&cpu.Registers.A, 0b00000010)
}

// 0xcb 0xd0 SET 2, B
func (cpu *CPU_struct) op_cb_0xd0() {
	cpu.set(&cpu.Registers.B, 0b00000100)
}

// 0xcb 0xd1 SET 2, C
func (cpu *CPU_struct) op_cb_0xd1() {
	cpu.set(&cpu.Registers.C, 0b00000100)
}

// 0xcb 0xd2 SET 2, D
func (cpu *CPU_struct) op_cb_0xd2() {
	cpu.set(&cpu.Registers.D, 0b00000100)
}

// 0xcb 0xd3 SET 2, E
func (cpu *CPU_struct) op_cb_0xd3() {
	cpu.set(&cpu.Registers.E, 0b00000100)
}

// 0xcb 0xd4 SET 2, H
func (cpu *CPU_struct) op_cb_0xd4() {
	cpu.set(&cpu.Registers.H, 0b00000100)
}

// 0xcb 0xd5 SET 2, L
func (cpu *CPU_struct) op_cb_0xd5() {
	cpu.set(&cpu.Registers.L, 0b00000100)
}

// 0xcb 0xd6 SET 2, (HL)
func (cpu *CPU_struct) op_cb_0xd6() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.set(&n, 0b00000100)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0xd7 SET 2, A
func (cpu *CPU_struct) op_cb_0xd7() {
	cpu.set(&cpu.Registers.A, 0b00000100)
}

// 0xcb 0xd8 SET 3, B
func (cpu *CPU_struct) op_cb_0xd8() {
	cpu.set(&cpu.Registers.B, 0b00001000)
}

// 0xcb 0xd9 SET 3, C
func (cpu *CPU_struct) op_cb_0xd9() {
	cpu.set(&cpu.Registers.C, 0b00001000)
}

// 0xcb 0xda SET 3, D
func (cpu *CPU_struct) op_cb_0xda() {
	cpu.set(&cpu.Registers.D, 0b00001000)
}

// 0xcb 0xdb SET 3, E
func (cpu *CPU_struct) op_cb_0xdb() {
	cpu.set(&cpu.Registers.E, 0b00001000)
}

// 0xcb 0xdc SET 3, H
func (cpu *CPU_struct) op_cb_0xdc() {
	cpu.set(&cpu.Registers.H, 0b00001000)
}

// 0xcb 0xdd SET 3, L
func (cpu *CPU_struct) op_cb_0xdd() {
	cpu.set(&cpu.Registers.L, 0b00001000)
}

// 0xcb 0xde SET 3, (HL)
func (cpu *CPU_struct) op_cb_0xde() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.set(&n, 0b00001000)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0xdf SET 3, A
func (cpu *CPU_struct) op_cb_0xdf() {
	cpu.set(&cpu.Registers.A, 0b00001000)
}

// 0xcb 0xe0 SET 4, B
func (cpu *CPU_struct) op_cb_0xe0() {
	cpu.set(&cpu.Registers.B, 0b00010000)
}

// 0xcb 0xe1 SET 4, C
func (cpu *CPU_struct) op_cb_0xe1() {
	cpu.set(&cpu.Registers.C, 0b00010000)
}

// 0xcb 0xe2 SET 4, D
func (cpu *CPU_struct) op_cb_0xe2() {
	cpu.set(&cpu.Registers.D, 0b00010000)
}

// 0xcb 0xe3 SET 4, E
func (cpu *CPU_struct) op_cb_0xe3() {
	cpu.set(&cpu.Registers.E, 0b00010000)
}

// 0xcb 0xe4 SET 4, H
func (cpu *CPU_struct) op_cb_0xe4() {
	cpu.set(&cpu.Registers.H, 0b00010000)
}

// 0xcb 0xe5 SET 4, L
func (cpu *CPU_struct) op_cb_0xe5() {
	cpu.set(&cpu.Registers.L, 0b00010000)
}

// 0xcb 0xe6 SET 4, (HL)
func (cpu *CPU_struct) op_cb_0xe6() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.set(&n, 0b00010000)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0xe7 SET 4, A
func (cpu *CPU_struct) op_cb_0xe7() {
	cpu.set(&cpu.Registers.A, 0b00010000)
}

// 0xcb 0xe8 SET 5, B
func (cpu *CPU_struct) op_cb_0xe8() {
	cpu.set(&cpu.Registers.B, 0b00100000)
}

// 0xcb 0xe9 SET 5, C
func (cpu *CPU_struct) op_cb_0xe9() {
	cpu.set(&cpu.Registers.C, 0b00100000)
}

// 0xcb 0xea SET 5, D
func (cpu *CPU_struct) op_cb_0xea() {
	cpu.set(&cpu.Registers.D, 0b00100000)
}

// 0xcb 0xeb SET 5, E
func (cpu *CPU_struct) op_cb_0xeb() {
	cpu.set(&cpu.Registers.E, 0b00100000)
}

// 0xcb 0xec SET 5, H
func (cpu *CPU_struct) op_cb_0xec() {
	cpu.set(&cpu.Registers.H, 0b00100000)
}

// 0xcb 0xed SET 5, L
func (cpu *CPU_struct) op_cb_0xed() {
	cpu.set(&cpu.Registers.L, 0b00100000)
}

// 0xcb 0xee SET 5, (HL)
func (cpu *CPU_struct) op_cb_0xee() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.set(&n, 0b00100000)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0xef SET 5, A
func (cpu *CPU_struct) op_cb_0xef() {
	cpu.set(&cpu.Registers.A, 0b00100000)
}

// 0xcb 0xf0 SET 6, B
func (cpu *CPU_struct) op_cb_0xf0() {
	cpu.set(&cpu.Registers.B, 0b01000000)
}

// 0xcb 0xf1 SET 6, C
func (cpu *CPU_struct) op_cb_0xf1() {
	cpu.set(&cpu.Registers.C, 0b01000000)
}

// 0xcb 0xf2 SET 6, D
func (cpu *CPU_struct) op_cb_0xf2() {
	cpu.set(&cpu.Registers.D, 0b01000000)
}

// 0xcb 0xf3 SET 6, E
func (cpu *CPU_struct) op_cb_0xf3() {
	cpu.set(&cpu.Registers.E, 0b01000000)
}

// 0xcb 0xf4 SET 6, H
func (cpu *CPU_struct) op_cb_0xf4() {
	cpu.set(&cpu.Registers.H, 0b01000000)
}

// 0xcb 0xf5 SET 6, L
func (cpu *CPU_struct) op_cb_0xf5() {
	cpu.set(&cpu.Registers.L, 0b01000000)
}

// 0xcb 0xf6 SET 6, (HL)
func (cpu *CPU_struct) op_cb_0xf6() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.set(&n, 0b01000000)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0xf7 SET 6, A
func (cpu *CPU_struct) op_cb_0xf7() {
	cpu.set(&cpu.Registers.A, 0b01000000)
}

// 0xcb 0xf8 SET 7, B
func (cpu *CPU_struct) op_cb_0xf8() {
	cpu.set(&cpu.Registers.B, 0b10000000)
}

// 0xcb 0xf9 SET 7, C
func (cpu *CPU_struct) op_cb_0xf9() {
	cpu.set(&cpu.Registers.C, 0b10000000)
}

// 0xcb 0xfa SET 7, D
func (cpu *CPU_struct) op_cb_0xfa() {
	cpu.set(&cpu.Registers.D, 0b10000000)
}

// 0xcb 0xfb SET 7, E
func (cpu *CPU_struct) op_cb_0xfb() {
	cpu.set(&cpu.Registers.E, 0b10000000)
}

// 0xcb 0xfc SET 7, H
func (cpu *CPU_struct) op_cb_0xfc() {
	cpu.set(&cpu.Registers.H, 0b10000000)
}

// 0xcb 0xfd SET 7, L
func (cpu *CPU_struct) op_cb_0xfd() {
	cpu.set(&cpu.Registers.L, 0b10000000)
}

// 0xcb 0xfe SET 7, (HL)
func (cpu *CPU_struct) op_cb_0xfe() {
	var address uint16 = cpu.getHL()
	var n uint8 = cpu.MMU.ReadByte(address)
	cpu.set(&n, 0b10000000)
	cpu.MMU.WriteByte(address, n)

	cpu.Cycle += 8
}

// 0xcb 0xff SET 7, A
func (cpu *CPU_struct) op_cb_0xff() {
	cpu.set(&cpu.Registers.A, 0b10000000)
}
