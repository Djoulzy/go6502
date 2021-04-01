package databus

import (
	"sync"
)

type Databus struct {
	// CPU    sync.Mutex
	// VIC    sync.Mutex
	Access sync.Mutex
	level  bool // True: CPU / False: VIC
}

func (D *Databus) Init() {
	D.level = false
}

// WaitBusLow : CPU wait
func (D *Databus) WaitBusLow() {
	D.Access.Lock()
	D.level = true
	D.Access.Unlock()
	for D.level == true {
	}
}

// WaitBusLow : VIC wait
func (D *Databus) WaitBusHigh() {
	D.Access.Lock()
	D.level = false
	D.Access.Unlock()
	for D.level == false {
	}
}
