package cpu

type Register struct {
	val int16
}

func NewRegister() *Register {
	return &Register{0}
}

func (reg *Register) Get() int16 {
	return reg.val
}

func (reg *Register) Set(val int16) {
	reg.val = val
}
