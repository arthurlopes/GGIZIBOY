package gameboy

import (
	"crypto/md5"
	"testing"
)

func TestLoadBoostrap(t *testing.T) {
	var gb = GameboyFactory()

	md5 := md5.Sum(gb.mmu.Memory[:0x100])

	var bootstrap_md5 [16]byte = [16]byte{0x32, 0xfb, 0xbd, 0x84, 0x16, 0x8d, 0x34, 0x82, 0x95, 0x6e, 0xb3, 0xc5, 0x05, 0x16, 0x37, 0xf5}

	if md5 != bootstrap_md5 {
		t.Errorf("Incorrect bootstrap md5")
	}
}

func TestWriteRead(t *testing.T) {
	var gb = GameboyFactory()

	gb.mmu.WriteByte(0x10, 0x66)
	if gb.mmu.ReadByte(0x10) != 0x66 {
		t.Errorf("Incorrect value")
	}
}
