package cpu

type ALU struct {
	result int16
}

func NewALU() *ALU {
	return &ALU{0}
}

func (alu *ALU) GetResult() int16 {
	return alu.result
}

func (alu *ALU) Add(a, b int16) {
	alu.result = a + b
}

func (alu *ALU) Sub(a, b int16) {
	alu.result = a - b
}

func (alu *ALU) Mul(a, b int16) {
	alu.result = a * b
}

func (alu *ALU) Div(a, b int16) {
	alu.result = a / b
}

func (alu *ALU) Mod(a, b int16) {
	alu.result = a % b
}

func (alu *ALU) RSft(a, b int16) {
	alu.result = a >> b
}

func (alu *ALU) LSft(a, b int16) {
	alu.result = a << b
}
