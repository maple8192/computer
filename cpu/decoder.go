package cpu

type Instruction int

const (
	Nop Instruction = iota
	Add
	Sub
	Mul
	Div
	Mod
	MovA
	MovD
	MovP
	MovB
	MovS
	SetA
	SetD
	SetP
	SetB
	SetS
	StoreA
	StoreD
	StoreP
	StoreB
	StoreS
	ResA
	ResD
	ResP
	ResB
	ResS
)

type Decoder struct {
	instruction Instruction
}

func NewDecoder() *Decoder {
	return &Decoder{Nop}
}

func (dec *Decoder) GetInstruction() Instruction {
	return dec.instruction
}

func (dec *Decoder) Decode(inst uint8) {
	switch inst {
	case 0x00:
		dec.instruction = Nop
	case 0x10:
		dec.instruction = Add
	case 0x20:
		dec.instruction = Sub
	case 0x30:
		dec.instruction = Mul
	case 0x40:
		dec.instruction = Div
	case 0x50:
		dec.instruction = Mod
	case 0x60:
		dec.instruction = MovA
	case 0x61:
		dec.instruction = MovD
	case 0x62:
		dec.instruction = MovP
	case 0x63:
		dec.instruction = MovB
	case 0x64:
		dec.instruction = MovS
	case 0x70:
		dec.instruction = SetA
	case 0x71:
		dec.instruction = SetD
	case 0x72:
		dec.instruction = SetP
	case 0x73:
		dec.instruction = SetB
	case 0x74:
		dec.instruction = SetS
	case 0x80:
		dec.instruction = StoreA
	case 0x81:
		dec.instruction = StoreD
	case 0x82:
		dec.instruction = StoreP
	case 0x83:
		dec.instruction = StoreB
	case 0x84:
		dec.instruction = StoreS
	case 0x90:
		dec.instruction = ResA
	case 0x91:
		dec.instruction = ResD
	case 0x92:
		dec.instruction = ResP
	case 0x93:
		dec.instruction = ResB
	case 0x94:
		dec.instruction = ResS
	}
}
