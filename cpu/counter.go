package cpu

type Counter struct {
	val int16
}

func NewCounter() *Counter {
	return &Counter{}
}

func (cnt *Counter) Get() int16 {
	return cnt.val
}

func (cnt *Counter) Set(val int16) {
	cnt.val = val
}

func (cnt *Counter) Inc() {
	cnt.val++
}
