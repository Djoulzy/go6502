package databus

import (
	"go6502/vic"
	"sync"
)

type Bus struct {
	// CPU    sync.Mutex
	// VIC    sync.Mutex
	Access sync.Mutex
	level  bool // True: CPU / False: VIC
	vic    *vic.VIC
}

func (B *Bus) Init(vic *vic.VIC) {
	B.level = false
	B.vic = vic
}

// WaitBusLow : CPU wait
// func (D *Databus) WaitBusLow() {
// 	D.Access.Lock()
// 	D.level = true
// 	D.Access.Unlock()
// 	for D.level == true {
// 	}
// }

// WaitBusLow : VIC wait
// func (D *Databus) WaitBusHigh() {
// 	D.Access.Lock()
// 	D.level = false
// 	D.Access.Unlock()
// 	for D.level == false {
// 	}
// }

func (B *Bus) Release() {
KEEPBUS:
	B.vic.Run()
	if !B.vic.BA {
		goto KEEPBUS
	}
}
