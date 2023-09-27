package cpu

import "computer/memory"

type Fetcher struct {
	ram         *memory.RAM
	instruction uint8
	operand     [2]uint8
}

func NewFetcher(ram *memory.RAM) *Fetcher {
	return &Fetcher{ram, 0, [2]uint8{}}
}

func (ftc *Fetcher) Fetch(addr int16) {
	ftc.instruction = ftc.ram.Get(addr)
	ftc.operand[0] = ftc.ram.Get(addr + 1)
	ftc.operand[1] = ftc.ram.Get(addr + 2)
}

func (ftc *Fetcher) GetInstruction() uint8 {
	return ftc.instruction
}

func (ftc *Fetcher) GetOperand() int16 {
	return int16(uint16(ftc.operand[0])<<8 | uint16(ftc.operand[1]))
}
