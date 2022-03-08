package CPU

import "fmt"

type IGPU interface {
	Update_clock(int)
	Step()
}

type IMMU interface {
	ReadByte(uint16) uint8
	ReadWord(uint16) uint16
	WriteByte(uint16, uint8)
	WriteWord(uint16, uint16)
	DumpMemory()
}

type Registers_struct struct {
	A uint8
	B uint8
	C uint8
	D uint8
	E uint8
	F uint8
	H uint8
	L uint8

	SP uint16
	PC uint16

	IME uint8
	IE  uint8
	IF  uint8

	Halt uint8
}

type Instructions_maps struct {
	Instructions_map_uint8  map[uint8]func(uint8)
	Instructions_map_uint16 map[uint8]func(uint16)
	Instructions_map_       map[uint8]func()
	Instructions_map_CB     map[uint8]func()
	Instructions_map_str    map[uint8]string
	Instructions_map_CB_str map[uint8]string
}

type CPU_struct struct {
	Registers         *Registers_struct
	Instructions_maps *Instructions_maps

	Cycle              int
	Instructions_count int

	MMU IMMU
	GPU IGPU
}

func CPUFactory() *CPU_struct {
	var cpu = CPU_struct{}
	cpu.Innit()
	return &cpu
}

func (cpu *CPU_struct) Innit() {
	cpu.Instructions_maps = new(Instructions_maps)
	cpu.Innit_Instruction_maps()
	cpu.Registers = new(Registers_struct)
}

func (cpu *CPU_struct) Run(limit int) {
	cpu.Instructions_count = 0
	for {
		if cpu.Registers.Halt != 0 {
			cpu.Cycle += 4
			cpu.GPU.Update_clock(cpu.Cycle)
			cpu.GPU.Step()
			cpu.interruption()
			continue
		}

		var next_instruction uint8 = cpu.MMU.ReadByte(cpu.Registers.PC)
		var next_instruction_str string
		if next_instruction == 0xCB {
			next_instruction_str = cpu.Instructions_maps.Instructions_map_CB_str[cpu.MMU.ReadByte(cpu.Registers.PC+1)]
		} else {
			next_instruction_str = cpu.Instructions_maps.Instructions_map_str[next_instruction]
		}

		if cpu.Registers.PC == 0xC325 {
			fmt.Printf("%s, %X, %X, %X, %d %d\n", next_instruction_str, cpu.Registers.PC, next_instruction, cpu.Registers.SP, cpu.Cycle, cpu.Instructions_count)
			// cpu.MMU.DumpMemory()
			// fmt.Scanln()
		}

		if cpu.Registers.PC == 0xC063 {
			fmt.Printf("%s, %X, %X, %X, %d %d\n", next_instruction_str, cpu.Registers.PC, next_instruction, cpu.Registers.SP, cpu.Cycle, cpu.Instructions_count)
			// cpu.MMU.DumpMemory()
			// fmt.Scanln()
		}

		if next_instruction == 0x55 {
			fmt.Printf("%s, %X, %X, %X, %d %d\n", next_instruction_str, cpu.Registers.PC, next_instruction, cpu.Registers.SP, cpu.Cycle, cpu.Instructions_count)
			// cpu.MMU.DumpMemory()
			// fmt.Scanln()
		}

		if next_instruction == 0xCB {
			cpu.Registers.PC += 1
			// fmt.Printf("Instruction 0xCB\n")
			next_instruction = cpu.MMU.ReadByte(cpu.Registers.PC)
			// TODO: check if can simply skip cycle
			// cpu.Cycle += 4
			if op_func, ok := cpu.Instructions_maps.Instructions_map_CB[next_instruction]; ok {
				op_func()
				cpu.Registers.PC += 1
			} else {
				fmt.Printf("Instruction not implemented 0xCB 0x%X\n", next_instruction)
				fmt.Scan()
				panic("CB Instruction not implemented")
			}
		} else {
			if op_func, ok := cpu.Instructions_maps.Instructions_map_uint8[next_instruction]; ok {
				cpu.Registers.PC += 1
				var n uint8 = cpu.MMU.ReadByte(cpu.Registers.PC)
				cpu.Registers.PC += 1

				op_func(n)
			} else if op_func, ok := cpu.Instructions_maps.Instructions_map_uint16[next_instruction]; ok {
				cpu.Registers.PC += 1
				var nn uint16 = cpu.MMU.ReadWord(cpu.Registers.PC)
				cpu.Registers.PC += 2

				op_func(nn)
			} else if op_func, ok := cpu.Instructions_maps.Instructions_map_[next_instruction]; ok {
				cpu.Registers.PC += 1

				op_func()
			} else {
				fmt.Printf("Instruction not implemented 0x%X\n", next_instruction)
				fmt.Scan()
				panic("Instruction not implemented")
			}
		}
		cpu.GPU.Update_clock(cpu.Cycle)
		cpu.GPU.Step()

		cpu.interruption()

		cpu.Instructions_count += 1
		if limit > 0 && cpu.Instructions_count == limit {
			break
		}

		// if cpu.Registers.PC != 0x64 && cpu.Registers.PC != 0x66 && cpu.Registers.PC != 0x68 {
		// 	fmt.Printf("%s, %X, %X, %X, %d %d\n", next_instruction_str, cpu.Registers.PC, next_instruction, cpu.Registers.SP, cpu.Cycle, cpu.Instructions_count)
		// }

	}
}

func (cpu *CPU_struct) interruption() {
	if (cpu.Registers.IME & cpu.Registers.IE & cpu.Registers.IF) == 0 {
		return
	}

	if cpu.Registers.IE != 0 {
		cpu.Registers.IME = 0
		cpu.Registers.Halt = 0
		// V BLANK
		if (cpu.Registers.IF & 0x01) != 0 {
			// "call" 0x40
			fmt.Printf("V BLANK\n")
			cpu.op_0xcd(0x40)
			cpu.Registers.IF &= 0b11111110
			// LCD status triggers
		} else if (cpu.Registers.IF & 0x02) != 0 {
			// "call" 0x48
			fmt.Printf("LCD status triggers\n")
			cpu.op_0xcd(0x48)
			cpu.Registers.IF &= 0b11111101
			// Timer overflow
		} else if (cpu.Registers.IF & 0x04) != 0 {
			// "call" 0x50
			fmt.Printf("Timer overflow\n")
			cpu.op_0xcd(0x50)
			cpu.Registers.IF &= 0b11111011
			// Serial link
		} else if (cpu.Registers.IF & 0x08) != 0 {
			// "call" 0x58
			fmt.Printf("Serial link\n")
			cpu.op_0xcd(0x58)
			cpu.Registers.IF &= 0b11110111
			// Joypad press
		} else if (cpu.Registers.IF & 0x10) != 0 {
			// "call" 0x60
			fmt.Printf("Joypad press\n")
			cpu.op_0xcd(0x60)
			cpu.Registers.IF &= 0b11101111
		}
	}
}

func (cpu *CPU_struct) GetIE() uint8 {
	return cpu.Registers.IE
}

func (cpu *CPU_struct) SetIE(IE uint8) {
	cpu.Registers.IE |= IE
}

func (cpu CPU_struct) GetIF() uint8 {
	return cpu.Registers.IF
}

func (cpu *CPU_struct) SetIF(IF uint8) {
	cpu.Registers.IF |= IF
}
