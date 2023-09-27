package cpu

import (
	"computer/memory"
)

type CPU struct {
	state   State
	pc      *Counter
	bp      *Register
	sp      *Register
	acc     *Register
	data    *Register
	fetcher *Fetcher
	decoder *Decoder
	alu     *ALU
	ram     *memory.RAM
}

func NewCPU(ram *memory.RAM) *CPU {
	return &CPU{
		Fetch,
		NewCounter(),
		NewRegister(),
		NewRegister(),
		NewRegister(),
		NewRegister(),
		NewFetcher(ram),
		NewDecoder(),
		NewALU(),
		ram,
	}
}

func (cpu *CPU) Update() {
	switch cpu.state {
	case Fetch:
		cpu.fetcher.Fetch(cpu.pc.Get())
		cpu.state = Decode
	case Decode:
		cpu.decoder.Decode(cpu.fetcher.GetInstruction())
		cpu.state = Execute
	case Execute:
		switch cpu.decoder.GetInstruction() {
		case Nop:
		case Add:
			cpu.alu.Add(cpu.acc.Get(), cpu.data.Get())
			cpu.pc.Inc()
		case Sub:
			cpu.alu.Sub(cpu.acc.Get(), cpu.data.Get())
			cpu.pc.Inc()
		case Mul:
			cpu.alu.Mul(cpu.acc.Get(), cpu.data.Get())
			cpu.pc.Inc()
		case Div:
			cpu.alu.Div(cpu.acc.Get(), cpu.data.Get())
			cpu.pc.Inc()
		case Mod:
			cpu.alu.Mod(cpu.acc.Get(), cpu.data.Get())
			cpu.pc.Inc()
		case MovA:
			cpu.acc.Set(int16(uint16(cpu.ram.Get(cpu.fetcher.GetOperand()))<<8 | uint16(cpu.ram.Get(cpu.fetcher.GetOperand()+1))))
			cpu.pc.Inc()
			cpu.pc.Inc()
			cpu.pc.Inc()
		case MovD:
			cpu.data.Set(int16(uint16(cpu.ram.Get(cpu.fetcher.GetOperand()))<<8 | uint16(cpu.ram.Get(cpu.fetcher.GetOperand()+1))))
			cpu.pc.Inc()
			cpu.pc.Inc()
			cpu.pc.Inc()
		case MovP:
			cpu.pc.Set(int16(uint16(cpu.ram.Get(cpu.fetcher.GetOperand()))<<8 | uint16(cpu.ram.Get(cpu.fetcher.GetOperand()+1))))
		case MovB:
			cpu.bp.Set(int16(uint16(cpu.ram.Get(cpu.fetcher.GetOperand()))<<8 | uint16(cpu.ram.Get(cpu.fetcher.GetOperand()+1))))
			cpu.pc.Inc()
			cpu.pc.Inc()
			cpu.pc.Inc()
		case MovS:
			cpu.sp.Set(int16(uint16(cpu.ram.Get(cpu.fetcher.GetOperand()))<<8 | uint16(cpu.ram.Get(cpu.fetcher.GetOperand()+1))))
			cpu.pc.Inc()
			cpu.pc.Inc()
			cpu.pc.Inc()
		case SetA:
			cpu.acc.Set(cpu.fetcher.GetOperand())
			cpu.pc.Inc()
			cpu.pc.Inc()
			cpu.pc.Inc()
		case SetD:
			cpu.data.Set(cpu.fetcher.GetOperand())
			cpu.pc.Inc()
			cpu.pc.Inc()
			cpu.pc.Inc()
		case SetP:
			cpu.pc.Set(cpu.fetcher.GetOperand())
		case SetB:
			cpu.bp.Set(cpu.fetcher.GetOperand())
			cpu.pc.Inc()
			cpu.pc.Inc()
			cpu.pc.Inc()
		case SetS:
			cpu.sp.Set(cpu.fetcher.GetOperand())
			cpu.pc.Inc()
			cpu.pc.Inc()
			cpu.pc.Inc()
		case StoreA:
			cpu.ram.Set(cpu.fetcher.GetOperand(), uint8(cpu.acc.Get()>>8))
			cpu.ram.Set(cpu.fetcher.GetOperand()+1, uint8(cpu.acc.Get()))
			cpu.pc.Inc()
			cpu.pc.Inc()
			cpu.pc.Inc()
		case StoreD:
			cpu.ram.Set(cpu.fetcher.GetOperand(), uint8(cpu.data.Get()>>8))
			cpu.ram.Set(cpu.fetcher.GetOperand()+1, uint8(cpu.data.Get()))
			cpu.pc.Inc()
			cpu.pc.Inc()
			cpu.pc.Inc()
		case StoreP:
			cpu.ram.Set(cpu.fetcher.GetOperand(), uint8(cpu.pc.Get()>>8))
			cpu.ram.Set(cpu.fetcher.GetOperand()+1, uint8(cpu.pc.Get()))
			cpu.pc.Inc()
			cpu.pc.Inc()
			cpu.pc.Inc()
		case StoreB:
			cpu.ram.Set(cpu.fetcher.GetOperand(), uint8(cpu.bp.Get()>>8))
			cpu.ram.Set(cpu.fetcher.GetOperand()+1, uint8(cpu.bp.Get()))
			cpu.pc.Inc()
			cpu.pc.Inc()
			cpu.pc.Inc()
		case StoreS:
			cpu.ram.Set(cpu.fetcher.GetOperand(), uint8(cpu.sp.Get()>>8))
			cpu.ram.Set(cpu.fetcher.GetOperand()+1, uint8(cpu.sp.Get()))
			cpu.pc.Inc()
			cpu.pc.Inc()
			cpu.pc.Inc()
		case ResA:
			cpu.acc.Set(cpu.alu.GetResult())
			cpu.pc.Inc()
		case ResD:
			cpu.data.Set(cpu.alu.GetResult())
			cpu.pc.Inc()
		case ResP:
			cpu.pc.Set(cpu.alu.GetResult())
		case ResB:
			cpu.bp.Set(cpu.alu.GetResult())
			cpu.pc.Inc()
		case ResS:
			cpu.sp.Set(cpu.alu.GetResult())
			cpu.pc.Inc()
		}
		cpu.state = Fetch
	}
}
