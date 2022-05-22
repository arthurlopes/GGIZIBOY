package CPU

import "fmt"

type IGPU interface {
	Update_clock(int, int)
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

	DIV      uint8
	DIV_aux  float32
	TIMA     uint8
	TIMA_aux float32
	TMA      uint8
	TAC      uint8

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
			cpu.GPU.Update_clock(4, cpu.Cycle)
			cpu.GPU.Step()
			cpu.timer(4, cpu.Registers.TAC)
			cpu.interruption(cpu.Registers.IME)
			continue
		}

		var old_cycle int = cpu.Cycle

		var next_instruction uint8 = cpu.MMU.ReadByte(cpu.Registers.PC)
		var next_instruction_str string
		if next_instruction == 0xCB {
			next_instruction_str = cpu.Instructions_maps.Instructions_map_CB_str[cpu.MMU.ReadByte(cpu.Registers.PC+1)]
		} else {
			next_instruction_str = cpu.Instructions_maps.Instructions_map_str[next_instruction]
		}

		if cpu.Registers.PC == 0x100 {
			fmt.Printf("%s, %X, %X, %X, %d %d\n", next_instruction_str, cpu.Registers.PC, next_instruction, cpu.Registers.SP, cpu.Cycle, cpu.Instructions_count)
			// cpu.MMU.DumpMemory()
			// fmt.Scanln()
		}

		// if cpu.Registers.PC == 0xfe {
		// 	fmt.Printf("%s, %X, %X, %X, %d %d\n", next_instruction_str, cpu.Registers.PC, next_instruction, cpu.Registers.SP, cpu.Cycle, cpu.Instructions_count)
		// 	// cpu.MMU.DumpMemory()
		// 	// fmt.Scanln()
		// }

		// if cpu.Registers.PC == 0x2BA {
		// 	fmt.Printf("%s, %X, %X, %X, %d %d\n", next_instruction_str, cpu.Registers.PC, next_instruction, cpu.Registers.SP, cpu.Cycle, cpu.Instructions_count)
		// 	fmt.Printf("\n")
		// 	// cpu.MMU.DumpMemory()
		// 	// fmt.Scanln()
		// }

		// fmt.Printf("0x%x 0x%x 0x%x 0x%x 0x%x 0x%x 0x%x 0x%x 0x%x\n", cpu.Registers.PC, next_instruction, cpu.Registers.A, cpu.Registers.B, cpu.Registers.C, cpu.Registers.D, cpu.Registers.E, cpu.getHL(), cpu.Registers.F) //, cpu.MMU.ReadByte(0xff44))

		var oldIME = cpu.Registers.IME
		var oldTAC uint8 = cpu.Registers.TAC
		if next_instruction == 0xCB {
			cpu.Registers.PC += 1
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
		cpu.GPU.Update_clock(cpu.Cycle-old_cycle, cpu.Cycle)
		cpu.GPU.Step()

		cpu.timer(cpu.Cycle-old_cycle, oldTAC)
		cpu.interruption(oldIME)

		cpu.Instructions_count += 1
		if limit > 0 && cpu.Instructions_count == limit {
			break
		}
	}
}

func (cpu *CPU_struct) timer(cycles int, TAC uint8) {
	cpu.Registers.DIV_aux += float32(cycles) / 4 / 4 / 16
	cpu.Registers.DIV = uint8(cpu.Registers.DIV_aux)

	if (TAC & 0x04) != 0 {
		var speed uint8 = TAC & 0x03
		if speed == 0 {
			speed = 64
		} else if speed == 1 {
			speed = 1
		} else if speed == 2 {
			speed = 4
		} else {
			speed = 16
		}

		var old_TIMA uint8 = cpu.Registers.TIMA
		cpu.Registers.TIMA_aux += float32(cycles) / 4 / 4 / float32(speed)
		cpu.Registers.TIMA = uint8(cpu.Registers.TIMA_aux)

		if cpu.Registers.TIMA < old_TIMA {
			cpu.Registers.IF |= 0x04
			cpu.Registers.TIMA = cpu.Registers.TMA
		}
	}

}

func (cpu *CPU_struct) interruption(IME uint8) {
	if (IME == 0 && cpu.Registers.Halt == 0) || ((cpu.Registers.IE & cpu.Registers.IF) == 0) {
		return
	}

	if cpu.Registers.IE != 0 {
		cpu.Registers.IME = 0
		// V BLANK
		if (cpu.Registers.IF & 0x01) != 0 {
			// "call" 0x40
			// fmt.Printf("V BLANK\n")
			cpu.call(0x40)
			if cpu.Registers.Halt == 0 {
				cpu.Registers.IF &= 0b11111110
			}
			cpu.Registers.Halt = 0
			// LCD status triggers
		} else if (cpu.Registers.IF & 0x02) != 0 {
			// "call" 0x48
			// fmt.Printf("LCD status triggers\n")
			cpu.call(0x48)
			if cpu.Registers.Halt == 0 {
				cpu.Registers.IF &= 0b11111101
			}
			cpu.Registers.Halt = 0
			// Timer overflow
		} else if (cpu.Registers.IF & 0x04) != 0 {
			// fmt.Printf("Timer overflow\n")
			cpu.call(0x50)
			if cpu.Registers.Halt == 0 {
				cpu.Registers.IF &= 0b11111011
			}
			cpu.Registers.Halt = 0
			// Serial link
		} else if (cpu.Registers.IF & 0x08) != 0 {
			// fmt.Printf("Serial link\n")
			cpu.call(0x58)
			if cpu.Registers.Halt == 0 {
				cpu.Registers.IF &= 0b11110111
			}
			cpu.Registers.Halt = 0
			// Joypad press
		} else if (cpu.Registers.IF & 0x10) != 0 {
			// fmt.Printf("Joypad press\n")
			cpu.call(0x60)
			if cpu.Registers.Halt == 0 {
				cpu.Registers.IF &= 0b11101111
			}
			cpu.Registers.Halt = 0
		}
	}
}

func (cpu *CPU_struct) GetIE() uint8 {
	return cpu.Registers.IE
}

func (cpu *CPU_struct) SetIE(IE uint8) {
	cpu.Registers.IE = IE
}

func (cpu CPU_struct) GetIF() uint8 {
	return cpu.Registers.IF
}

func (cpu CPU_struct) SetInterrupt(mask uint8) {
	cpu.Registers.IF |= mask
}

func (cpu *CPU_struct) SetIF(IF uint8) {
	cpu.Registers.IF = IF
}

func (cpu *CPU_struct) GetTIMA() uint8 {
	return cpu.Registers.TIMA
}

func (cpu *CPU_struct) SetTIMA(TIMA uint8) {
	cpu.Registers.TIMA = TIMA
	cpu.Registers.TIMA_aux = float32(TIMA)
}

func (cpu CPU_struct) GetTMA() uint8 {
	return cpu.Registers.TMA
}

func (cpu *CPU_struct) SetTMA(TMA uint8) {
	cpu.Registers.TMA = TMA
}

func (cpu CPU_struct) GetTAC() uint8 {
	return cpu.Registers.TAC
}

func (cpu *CPU_struct) SetTAC(TAC uint8) {
	cpu.Registers.TAC = TAC
}

func (cpu CPU_struct) GetDIV() uint8 {
	return cpu.Registers.DIV
}

func (cpu *CPU_struct) SetDIV(DIV uint8) {
	cpu.Registers.DIV = DIV
	cpu.Registers.DIV_aux = float32(DIV)
}
