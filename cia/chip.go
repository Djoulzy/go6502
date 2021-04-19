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
	ICR // Interrupt control register
	CRA // Timer A Control
	CRB // Timer B Control
)

func (C *CIA) Init(mem []mem.Cell) {
	C.mem = mem
	// C.PRA = 0
	// C.setValue(C.PRA, 0xC7)
	// // C.PRA.Rom = 0x47
}

func (C *CIA) SetValue(port byte, value byte) {
	for i := 0; i < 16; i++ {
		zone := port + byte(16*i)
		C.mem[zone].Zone[mem.IO] = value
	}
}

func (C *CIA) updateStates() {
	order := C.mem[ICR].Zone[mem.RAM]
	mask := order & 0b00001111
	if mask > 0 {
		if order&0b10000000 > 0 { // 7eme bit = 1 -> mask set
			C.mem[ICR].Zone[mem.IO] |= mask
		} else {
			C.mem[ICR].Zone[mem.IO] &= ^mask
		}
	}
	C.mem[CRA].Zone[mem.IO] = C.mem[CRA].Zone[mem.RAM]
	C.mem[CRB].Zone[mem.IO] = C.mem[CRB].Zone[mem.RAM]
}

func (C *CIA) Run() {
	C.updateStates()
}
