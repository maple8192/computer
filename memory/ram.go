package memory

import "fmt"

type RAM struct {
	mem []uint8
}

func NewRAM() *RAM {
	return &RAM{
		[]uint8{
			0x70, 0x00, 0x05, 0x71, 0x00, 0x08, 0x10, 0x90,
			0x80, 0x00, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		},
	}
}

func (ram *RAM) Get(addr int16) uint8 {
	if addr >= int16(len(ram.mem)) {
		return 0
	}
	return ram.mem[addr]
}

func (ram *RAM) Set(addr int16, val uint8) {
	ram.mem[addr] = val
}

func (ram *RAM) Print() {
	fmt.Println(ram.mem)
}
