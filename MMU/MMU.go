package MMU

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type IGPU interface {
	GetLine() uint8
	GetScroll_x() uint8
	GetScroll_y() uint8
}

type ICPU interface {
	GetIE() uint8
	GetIF() uint8
}

type MMU_struct struct {
	Memory            [0xffff + 1]uint8
	Rom               []uint8
	Bootstrap_path    string
	ROM_path          string
	bootstrap_enabled bool

	GPU IGPU
	CPU ICPU
}

func MMUFactory() *MMU_struct {
	var mmu = MMU_struct{}
	mmu.Innit()
	return &mmu
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (mmu *MMU_struct) Innit() {
	mmu.bootstrap_enabled = true
	mmu.Bootstrap_path = "data/bootstrap/dmg_boot.bin"
	// mmu.ROM_path = "data/ROM/Pokemon - Blue Version (USA, Europe) (SGB Enhanced).gb"
	// mmu.ROM_path = "data/ROM/Super Mario Land (JUE) (V1.1) [!].gb"
	mmu.ROM_path = "data/ROM/ttt.gb"

	if !fileExists(mmu.ROM_path) {
		mmu.ROM_path = "../" + mmu.ROM_path
	}

	if !fileExists(mmu.Bootstrap_path) {
		mmu.Bootstrap_path = "../" + mmu.Bootstrap_path
	}

	//read bootstrap
	bootstrap_file, err := os.Open(mmu.Bootstrap_path)
	if err != nil {
		log.Fatalln(err)
	}
	defer bootstrap_file.Close()

	memory_bootstrap := mmu.Memory[:0x100]
	err = binary.Read(bootstrap_file, binary.LittleEndian, &memory_bootstrap)
	if err != nil {
		log.Fatalln(err)
	}

	// read ROM
	// rom_file, err := os.Open(mmu.ROM_path)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer rom_file.Close()

	// memory_rom := mmu.Rom[:]
	// err = binary.Read(rom_file, binary.LittleEndian, &memory_rom)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	file, err := os.Open(mmu.ROM_path)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}

	// calculate the bytes size
	var size int64 = info.Size()
	mmu.Rom = make([]uint8, size)

	// read into buffer
	buffer := bufio.NewReader(file)
	_, err = buffer.Read(mmu.Rom[:])

}

func (mmu *MMU_struct) ReadByte(address uint16) uint8 {
	if (address >= 0x0100 || !mmu.bootstrap_enabled) && address < 0x8000 {
		return mmu.Rom[address]
	}

	if address == 0xff44 {
		return mmu.GPU.GetLine()
	}

	// if address >= 0xff00 {
	// 	fmt.Println(address)
	// }

	if address == 0xff43 {
		return mmu.GPU.GetScroll_x()
	}
	if address == 0xff42 {
		return mmu.GPU.GetScroll_y()
	}
	// if address == 0xff40 {
	// 	return (mmu.GPU.Switchbg) +
	// 		(mmu.GPU.Bgmap << 4) +
	// 		(mmu.GPU.Bgtile << 5) +
	// 		(mmu.GPU.Switchlcd << 8)
	// }
	if address == 0xFFFF {
		return mmu.CPU.GetIE()
	}
	if address == 0xFF0F {
		return mmu.CPU.GetIF()
	}

	return mmu.Memory[address]
}

func (mmu *MMU_struct) WriteByte(address uint16, n uint8) {
	if address == 0xff50 && n > 0 {
		mmu.bootstrap_enabled = false
		return
	}

	// if address >= 0xff00 {
	// 	fmt.Println(address, n)
	// }

	if address >= 0x8000 && address <= 0x9000 {
		mmu.Memory[address] = n
		return
	}

	if address >= 0x9000 && address <= 0x9bff {
		mmu.Memory[address] = n
		return
	}

	if address >= 0xFE00 && address <= 0xFE9F {
		// fmt.Printf("address 0x%X, value 0x%X\n", address, n)
		mmu.Memory[address] = n
		return
	}

	mmu.Memory[address] = n
}

func (mmu *MMU_struct) ReadWord(address uint16) uint16 {
	var word uint16
	if (address >= 0x0100 || !mmu.bootstrap_enabled) && address < 0x8000 {
		word = (uint16(mmu.Rom[address+1]) << 8) + uint16(mmu.Rom[address])
	} else {
		word = (uint16(mmu.Memory[address+1]) << 8) + uint16(mmu.Memory[address])
	}

	return word
}

func (mmu *MMU_struct) WriteWord(address uint16, nn uint16) {
	mmu.Memory[address] = uint8(nn & 0xff)
	mmu.Memory[address+1] = uint8(nn >> 8)
}

func (mmu *MMU_struct) DumpMemory() {
	var byteArray []uint8 = make([]uint8, 0xffff)
	copy(byteArray[:], mmu.Memory[:])
	err := ioutil.WriteFile("data/dumps/dump.bin", byteArray, 0222)
	if err != nil {
		fmt.Println("Issue dumping memory")
	}
}

func (mmu *MMU_struct) LoadDump() {
	var dumpPath = "data/dumps/dump_after_boot.bin"

	if !fileExists(dumpPath) {
		dumpPath = "../" + dumpPath
	}

	//read dump
	dumpFile, err := os.Open(dumpPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer dumpFile.Close()

	memory_bootstrap := mmu.Memory[:0xffff]
	err = binary.Read(dumpFile, binary.LittleEndian, &memory_bootstrap)
	if err != nil {
		log.Fatalln(err)
	}

}
