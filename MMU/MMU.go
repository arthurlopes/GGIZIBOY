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
	SetScroll_x(uint8)
	SetScroll_y(uint8)
	GetLYC() uint8
	SetLYC(uint8)
	GetLCDC() uint8
	GetSTAT() uint8
	SetLCDC(uint8)
	SetSTAT(uint8)
	GetBGP() uint8
	SetBGP(uint8)
}

type ICPU interface {
	GetIE() uint8
	GetIF() uint8
	GetTIMA() uint8
	GetTMA() uint8
	GetTAC() uint8
	GetDIV() uint8
	SetIE(uint8)
	SetIF(uint8)
	SetTIMA(uint8)
	SetTMA(uint8)
	SetTAC(uint8)
	SetDIV(uint8)
}

type MMU_struct struct {
	Memory            [0xffff + 1]uint8
	Rom               []uint8
	Bootstrap_path    string
	ROM_path          string
	Bootstrap_enabled bool
	cartridge_type    uint8
	VRAM_modified     bool
	memory_bank       uint8

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
	mmu.Bootstrap_enabled = true
	mmu.Bootstrap_path = "data/bootstrap/dmg_boot.bin"
	// mmu.ROM_path = "data/ROM/Pokemon - Blue Version (USA, Europe) (SGB Enhanced).gb"
	mmu.ROM_path = "data/ROM/Super Mario Land (JUE) (V1.1) [!].gb"
	// mmu.ROM_path = "data/ROM/Dr. Mario (World).gb"
	// mmu.ROM_path = "data/ROM/cpu_instrs.gb"
	// mmu.ROM_path = "data/ROM/Tetris.gb"
	// mmu.ROM_path = "data/ROM/01-special.gb" // OK
	// mmu.ROM_path = "data/ROM/02-interrupts.gb" // OK
	// mmu.ROM_path = "data/ROM/03-op sp,hl.gb" // OK
	// mmu.ROM_path = "data/ROM/04-op r,imm.gb" // OK
	// mmu.ROM_path = "data/ROM/05-op rp.gb" // OK
	// mmu.ROM_path = "data/ROM/06-ld r,r.gb" // OK
	// mmu.ROM_path = "data/ROM/07-jr,jp,call,ret,rst.gb" // OK
	// mmu.ROM_path = "data/ROM/08-misc instrs.gb" // OK
	// mmu.ROM_path = "data/ROM/09-op r,r.gb" // OK
	// mmu.ROM_path = "data/ROM/10-bit ops.gb" // OK
	// mmu.ROM_path = "data/ROM/11-op a,(hl).gb" // OK
	// mmu.ROM_path = "data/ROM/naughtyemu.gb"

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
	if err != nil {
		log.Fatalln(err)
	}

	// get cartridge type
	mmu.cartridge_type = mmu.Rom[0x147]
	fmt.Printf("Cartridge type %d\n", mmu.cartridge_type)
}

func (mmu *MMU_struct) ReadByte(address uint16) uint8 {
	if (address >= 0x0100 || !mmu.Bootstrap_enabled) && address < 0x8000 {
		if address >= 0x4000 {
			return mmu.Rom[address+uint16(mmu.memory_bank-1)*0x4000]
		} else {
			return mmu.Rom[address]
		}
	}

	if address == 0xff47 {
		return mmu.GPU.GetBGP()
	}
	if address == 0xff45 {
		return mmu.GPU.GetLYC()
	}
	if address == 0xff44 {
		return mmu.GPU.GetLine()
	}
	if address == 0xff43 {
		return mmu.GPU.GetScroll_x()
	}
	if address == 0xff42 {
		return mmu.GPU.GetScroll_y()
	}
	if address == 0xff41 {
		return mmu.GPU.GetSTAT()
	}
	if address == 0xff40 {
		return mmu.GPU.GetLCDC()
	}
	if address == 0xFFFF {
		return mmu.CPU.GetIE()
	}
	if address == 0xFF0F {
		return mmu.CPU.GetIF()
	}
	if address == 0xFF04 {
		return mmu.CPU.GetDIV()
	}
	if address == 0xFF05 {
		return mmu.CPU.GetTIMA()
	}
	if address == 0xFF06 {
		return mmu.CPU.GetTMA()
	}
	if address == 0xFF07 {
		return mmu.CPU.GetTAC()
	}
	if address == 0xFF00 {
		return 0xFF
	}

	return mmu.Memory[address]
}

func (mmu *MMU_struct) WriteByte(address uint16, n uint8) {
	if address == 0xff50 && n > 0 {
		mmu.Memory[address] = n
		mmu.Bootstrap_enabled = false
		return
	}

	// Tiles
	if (address >= 0x8000) && (address <= 0x9000) {
		mmu.Memory[address] = n
		mmu.VRAM_modified = true
		return
	}

	// Bg
	if (address >= 0x9000) && (address <= 0x9bff) {
		mmu.Memory[address] = n
		mmu.VRAM_modified = true
		return
	}

	// OAM
	if (address >= 0xFE00) && (address <= 0xFE9F) {
		mmu.Memory[address] = n
		mmu.VRAM_modified = true
		return
	}

	// RAM
	if (address >= 0xC000) && (address <= 0xDE00) {
		mmu.Memory[address] = n
		// Write same value on ghost RAM
		mmu.Memory[(address + 0x2000)] = n
		return
	}

	if address == 0xff47 {
		mmu.GPU.SetBGP(n)
		return
	}
	if address == 0xff46 {
		mmu.Memory[address] = n
		// DMA
		DMA_startAddress := uint16(n) << 8
		OAM_startAdress := uint16(0xFE00)
		for offset := uint16(0); offset < 280; offset++ {
			mmu.Memory[OAM_startAdress+offset] = mmu.Memory[DMA_startAddress+offset]
		}
		return
	}
	if address == 0xff45 {
		mmu.GPU.SetLYC(n)
		return
	}
	if address == 0xff41 {
		mmu.GPU.SetSTAT(n)
		return
	}
	if address == 0xff40 {
		mmu.GPU.SetLCDC(n)
		return
	}
	if address == 0xff43 {
		mmu.GPU.SetScroll_x(n)
		return
	}
	if address == 0xff42 {
		mmu.GPU.SetScroll_y(n)
		return
	}
	if address == 0xFFFF {
		mmu.CPU.SetIE(n)
		return
	}
	if address == 0xFF0F {
		mmu.CPU.SetIF(n)
		return
	}
	if address == 0xFF04 {
		mmu.CPU.SetDIV(n)
		return
	}
	if address == 0xFF05 {
		mmu.CPU.SetTIMA(n)
		return
	}
	if address == 0xFF06 {
		mmu.CPU.SetTMA(n)
		return
	}
	if address == 0xFF07 {
		mmu.CPU.SetTAC(n)
		return
	}

	if (address >= 0x2000) && (address < 0x4000) {
		mmu.memory_bank = n & 0b00011111
	}

	mmu.Memory[address] = n
}

func (mmu *MMU_struct) ReadWord(address uint16) uint16 {
	word := (uint16(mmu.ReadByte(address+1)) << 8) | uint16(mmu.ReadByte(address))
	return word
}

func (mmu *MMU_struct) WriteWord(address uint16, nn uint16) {
	mmu.WriteByte(address, uint8(nn&0xff))
	mmu.WriteByte(address+1, uint8(nn>>8))
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

func (mmu *MMU_struct) Get_VRAM_modified() bool {
	return mmu.VRAM_modified
}

func (mmu *MMU_struct) Set_VRAM_modified(b bool) {
	mmu.VRAM_modified = b
}
