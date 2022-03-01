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
}

type Instructions_maps struct {
	Instructions_map_uint8  map[uint8]func(uint8)
	Instructions_map_uint16 map[uint8]func(uint16)
	Instructions_map_       map[uint8]func()
	Instructions_map_CB     map[uint8]func()
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
		var next_instruction uint8 = cpu.MMU.ReadByte(cpu.Registers.PC)

		if cpu.Registers.PC == 0x00FE {
			fmt.Printf("%X, %X, %X, %d %d\n", cpu.Registers.PC, next_instruction, cpu.Registers.SP, cpu.Cycle, cpu.Instructions_count)
			// cpu.MMU.DumpMemory()
			// fmt.Scanln()
		}

		if next_instruction == 0xCB {
			cpu.Registers.PC += 1
			next_instruction = cpu.MMU.ReadByte(cpu.Registers.PC)
			// TODO: check if can simply skip cycle
			cpu.Cycle += 4
			if op_func, ok := cpu.Instructions_maps.Instructions_map_CB[next_instruction]; ok {
				op_func()
				cpu.Registers.PC += 1
			} else {
				fmt.Printf("Instruction not implemented 0xCB 0x%X\n", next_instruction)
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

		// if cpu.Instructions_count%500000 == 0 { // && cpu.Registers.PC != 0x64 && cpu.Registers.PC != 0x66 && cpu.Registers.PC != 0x68 {
		if cpu.Registers.PC != 0x64 && cpu.Registers.PC != 0x66 && cpu.Registers.PC != 0x68 {
			fmt.Printf("%X, %X, %X, %d %d\n", cpu.Registers.PC, next_instruction, cpu.Registers.SP, cpu.Cycle, cpu.Instructions_count)
		}

	}
}

func (cpu *CPU_struct) interruption() {
	if cpu.Registers.IME == 0 {
		return
	}

	if cpu.Registers.IE != 0 {
		cpu.Registers.IME = 0
		if cpu.Registers.IF&0x01 != 0 {
			// "call" 0x40
			cpu.op_0xcd(0x40)
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
