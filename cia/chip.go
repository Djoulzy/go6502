package cia

import "go6502/mem"

type CIA struct {
	mem []mem.Cell
	PRA byte
}

func (C *CIA) setValue(addr byte, value byte) {
	for i := 0; i < 16; i++ {
		zone := addr + byte(16*i)
		C.mem[zone].Ram = value
	}
}

func (C *CIA) Init(mem []mem.Cell) {
	C.mem = mem
	C.PRA = 0
	C.setValue(C.PRA, 0xC7)
	// C.PRA.Rom = 0x47
}
