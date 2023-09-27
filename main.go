package main

import (
	"computer/cpu"
	"computer/memory"
)

func main() {
	ram := memory.NewRAM()
	computer := cpu.NewCPU(ram)
	for i := 0; i < 100; i++ {
		computer.Update()
	}
	ram.Print()
}
