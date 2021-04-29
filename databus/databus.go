package databus

import (
	"fmt"
	"go6502/vic"
	"sync"
	"time"
)

const (
	cpuClock = 985248                  // Mesure en Hz
	cpuCycle = (1 / float32(cpuClock)) // 1 cycle en ms
)

var start time.Time

type Bus struct {
	// CPU    sync.Mutex
	// VIC    sync.Mutex
	Access sync.Mutex
	level  bool // True: CPU / False: VIC
	vic    *vic.VIC
	Cycles int
	Timer  uint16
}

func (B *Bus) Init(vic *vic.VIC) {
	B.level = false
	B.vic = vic
	B.Cycles = 0
	B.Timer = 0
}

func (B *Bus) Get() {
	start = time.Now()
	B.Cycles = 0
}

func (B *Bus) Release() {
	// KEEPBUS:
	elapsed := time.Since(start)
	start = time.Now()
	fmt.Printf("%d - %v\n", B.Cycles, elapsed)
	B.vic.Run()
	B.Cycles++
	B.Timer++
	// if !B.vic.BA {
	// 	goto KEEPBUS
	// }
}
