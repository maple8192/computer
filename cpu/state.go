package cpu

type State int

const (
	Fetch State = iota
	Decode
	Execute
)
