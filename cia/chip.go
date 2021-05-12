package cia

import (
	"go6502/mem"
)

type CIA struct {
	name        string
	mem         []mem.Cell
	IRQ_Pin     *int
	systemCycle *uint16

	timerArunning bool
	timerAlatch   int32
	timerAcom     chan int

	timerBrunning bool
	timerBlatch   int32
	timerBcom     chan int
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

func (C *CIA) Init(name string, memCells []mem.Cell, timer *uint16) {
	C.name = name
	C.mem = memCells
	C.systemCycle = timer

	C.timerArunning = false
	C.timerAlatch = 0
	C.timerAcom = make(chan int)
	C.timerBrunning = false
	C.timerBlatch = 0
	C.timerBcom = make(chan int)

	C.SetValue(PRA, 0x81)
	C.SetValue(PRB, 0xFF)
	C.SetValue(DDRA, 0)
	C.SetValue(DDRB, 0)
	C.SetValue(TALO, 0xFF)
	C.SetValue(TAHI, 0xFF)
	C.SetValue(TBLO, 0xFF)
	C.SetValue(TBHI, 0xFF)
	C.SetValue(TOD10THS, 0)
	C.SetValue(TODSEC, 0)
	C.SetValue(TODMIN, 0)
	C.SetValue(TODHR, 0x01)
	C.SetValue(SRD, 0)
	C.SetValue(ICR, 0)
	C.mem[ICR].Zone[mem.RAM] = 0b00001111
	C.SetValue(CRA, 0)
	C.SetValue(CRB, 0)
}

func (C *CIA) SetValue(port byte, value byte) {
	for i := 0; i < 16; i++ {
		zone := port + byte(16*i)
		C.mem[zone].Zone[mem.IO] = value
		C.mem[zone].Zone[mem.RAM] = value
	}
}

// func (C *CIA) execTimerA() {
// 	reg := C.mem[CRA].Zone[mem.IO]
// 	if reg&0x00000001 != 0 {

// 	}
// }

func (C *CIA) updateStates() {
	if C.mem[ICR].IsWrite {
		order := C.mem[ICR].Zone[mem.RAM]
		mask := order & 0b00001111
		if mask > 0 {
			if order&0b10000000 > 0 { // 7eme bit = 1 -> mask set
				C.mem[ICR].Zone[mem.IO] |= mask
			} else {
				C.mem[ICR].Zone[mem.IO] &= ^mask
			}
		}
		C.mem[ICR].IsWrite = false
	}

	if C.mem[CRA].IsWrite {
		C.mem[CRA].Zone[mem.IO] = C.mem[CRA].Zone[mem.RAM]
		C.mem[CRA].IsWrite = false
		if C.mem[CRA].Zone[mem.IO]&0b00010000 > 0 {
			C.timerAlatch = int32(C.mem[TAHI].Zone[mem.RAM])<<8 + int32(C.mem[TALO].Zone[mem.RAM])
		}
		if C.mem[CRA].Zone[mem.IO]&0b00000001 == 1 && !C.timerArunning {
			go C.TimerA()
		} else {
			if C.timerArunning {
				C.timerAcom <- 1
			}
		}
	}

	if C.mem[CRB].IsWrite {
		C.mem[CRB].Zone[mem.IO] = C.mem[CRB].Zone[mem.RAM]
		C.mem[CRB].IsWrite = false
		if C.mem[CRB].Zone[mem.IO]&0b00010000 > 0 {
			C.timerAlatch = int32(C.mem[TBHI].Zone[mem.RAM])<<8 + int32(C.mem[TBLO].Zone[mem.RAM])
		}
		if C.mem[CRB].Zone[mem.IO]&0b00000001 == 1 && !C.timerBrunning {
			go C.TimerB()
		} else {
			if C.timerBrunning {
				C.timerBcom <- 1
			}
		}
	}

	if C.mem[TALO].IsWrite {
		C.mem[TALO].Zone[mem.IO] = C.mem[TALO].Zone[mem.RAM]
		C.mem[TALO].IsWrite = false
		C.timerAlatch = int32(C.mem[TAHI].Zone[mem.RAM])<<8 + int32(C.mem[TALO].Zone[mem.RAM])
	}
	if C.mem[TAHI].IsWrite {
		C.mem[TAHI].Zone[mem.IO] = C.mem[TAHI].Zone[mem.RAM]
		C.mem[TAHI].IsWrite = false
		C.timerAlatch = int32(C.mem[TAHI].Zone[mem.RAM])<<8 + int32(C.mem[TALO].Zone[mem.RAM])
	}
	if C.mem[TBLO].IsWrite {
		C.mem[TBLO].Zone[mem.IO] = C.mem[TBLO].Zone[mem.RAM]
		C.mem[TBLO].IsWrite = false
		C.timerBlatch = int32(C.mem[TBHI].Zone[mem.RAM])<<8 + int32(C.mem[TBLO].Zone[mem.RAM])
	}
	if C.mem[TBHI].IsWrite {
		C.mem[TBHI].Zone[mem.IO] = C.mem[TBHI].Zone[mem.RAM]
		C.mem[TBHI].IsWrite = false
		C.timerBlatch = int32(C.mem[TBHI].Zone[mem.RAM])<<8 + int32(C.mem[TBLO].Zone[mem.RAM])
	}
}

func (C *CIA) Run() {
	C.updateStates()
}
