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

func (B *Bus) Release() {
KEEPBUS:
	B.vic.Run()
	if !B.vic.BA {
		goto KEEPBUS
	}
}
