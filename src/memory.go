package main

type Memory struct {
	ram []uint8
}

func (mem *Memory) Read(addr uint16) uint8 {
	return mem.ram[addr]
}
