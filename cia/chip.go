package cia

import "go6502/mem"

type CIA struct {
	mem []mem.Cell
}

const (
	PRA byte = iota
	PRB
	DDRA
	DDRB
	TALO
	TAHI
	TBLO
	TBHI
	TOD10THS
	TODSEC
	TODMIN
	TODHR
	SRD
	ICR
	CRA
	CRB
)

func (C *CIA) SetValue(port byte, value byte) {
	for i := 0; i < 16; i++ {
		zone := port + byte(16*i)
		C.mem[zone].Zone[mem.IO] = value
	}
}

func (C *CIA) Init(mem []mem.Cell) {
	C.mem = mem
	// C.PRA = 0
	// C.setValue(C.PRA, 0xC7)
	// // C.PRA.Rom = 0x47
}
