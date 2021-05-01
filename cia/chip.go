package cia

import (
	"go6502/mem"
	"log"
)

type CIA struct {
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

func (C *CIA) TimerA() {
	defer func() {
		C.timerArunning = false
	}()
	C.timerArunning = true

	last := *C.systemCycle
	for {
		select {
		case <-C.timerAcom:
			return
		default:
			if *C.systemCycle != last {
				C.timerAlatch--
				if C.timerAlatch < 0 {
					if C.mem[CRA].Zone[mem.IO]&0b00001000 > 0 {
						return
					} else {
						C.timerAlatch = int32(C.mem[TAHI].Zone[mem.RAM])<<8 + int32(C.mem[TALO].Zone[mem.RAM])
					}
					if (C.mem[ICR].Zone[mem.RAM]&0b00000001 > 0) && (C.mem[ICR].Zone[mem.IO]&0b1000000 == 0) {

					}
				}
				log.Printf("timerA: %d\n", C.timerAlatch)
				last = *C.systemCycle
			}
		}
	}
}

func (C *CIA) Init(mem []mem.Cell, timer *uint16) {
	C.mem = mem
	C.systemCycle = timer

	C.timerArunning = false
	C.timerAlatch = 0
	C.timerAcom = make(chan int)
	C.timerBrunning = false
	C.timerBlatch = 0
	C.timerBcom = make(chan int)
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
			C.timerAcom <- 1
		}
	}
	if C.mem[CRB].IsWrite {
		C.mem[CRB].Zone[mem.IO] = C.mem[CRA].Zone[mem.RAM]
		C.mem[CRB].IsWrite = false
	}
}

func (C *CIA) Run() {
	C.updateStates()
	// C.execTimerA()
}
