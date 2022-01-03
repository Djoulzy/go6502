package cia

import (
	"go6502/mem"
)

func (C *CIA) TimerA() {
	last := *C.systemCycle
	for {
		select {
		case <-C.timerAcom:
			return
		default:
			if *C.systemCycle != last {
				C.timerAlatch--
				C.mem[TAHI].Zone[mem.IO] = byte(C.timerAlatch >> 8)
				C.mem[TALO].Zone[mem.IO] = byte(C.timerAlatch)
				if C.timerAlatch < 0 {
					// log.Println("underflow timer A")
					if (C.mem[ICR].Zone[mem.RAM]&0b00000001 > 0) && (C.mem[ICR].Zone[mem.IO]&0b1000000 == 0) {
						C.mem[ICR].Zone[mem.IO] |= 0b10000001
						// log.Printf("%s: Int timer A\n", C.name)
						*C.Signal_Pin = 1
					}
					if C.mem[CRA].Zone[mem.IO]&0b00001000 > 0 {
						return
					} else {
						C.timerAlatch = int32(C.mem[TAHI].Zone[mem.RAM])<<8 + int32(C.mem[TALO].Zone[mem.RAM])
					}
				}
				// log.Printf("timerA: %d\n", C.timerAlatch)
				last = *C.systemCycle
			}
		}
	}
}

func (C *CIA) TimerB() {
	last := *C.systemCycle
	for {
		select {
		case <-C.timerBcom:
			return
		default:
			if *C.systemCycle != last {
				C.timerAlatch--
				C.mem[TBHI].Zone[mem.IO] = byte(C.timerBlatch >> 8)
				C.mem[TBLO].Zone[mem.IO] = byte(C.timerBlatch)
				if C.timerAlatch < 0 {
					// log.Println("underflow timer B")
					if (C.mem[ICR].Zone[mem.RAM]&0b00000010 > 0) && (C.mem[ICR].Zone[mem.IO]&0b1000000 == 0) {
						C.mem[ICR].Zone[mem.IO] |= 0b10000010
						// log.Printf("%s: Int timer B\n", C.name)
						*C.Signal_Pin = 1
					}
					if C.mem[CRB].Zone[mem.IO]&0b00001000 > 0 {
						return
					} else {
						C.timerBlatch = int32(C.mem[TBHI].Zone[mem.RAM])<<8 + int32(C.mem[TBLO].Zone[mem.RAM])
					}
				}
				// log.Printf("timerA: %d\n", C.timerAlatch)
				last = *C.systemCycle
			}
		}
	}
}
