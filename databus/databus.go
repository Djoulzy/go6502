package databus

import (
	"fmt"
	"go6502/vic"
	"os"
	"sync"
	"time"
)

const (
	cpuClock = 985248                            // Mesure en Hz
	cpuCycle = (1 / float32(cpuClock)) * 1000000 // 1 cycle en µs
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
	log    chan time.Duration
}

func logTime(log chan time.Duration) {
	var c time.Duration

	fileDesc, _ := os.Create("log/time.log")

	for {
		c = <-log
		fileDesc.Write([]byte(fmt.Sprintf("%v\n", c)))
	}
}

func (B *Bus) Init(vic *vic.VIC) {
	B.level = false
	B.vic = vic
	B.Cycles = 0
	B.Timer = 0

	B.log = make(chan time.Duration)

	// go logTime(B.log)
}

func (B *Bus) Get() {
	// start = time.Now()
	B.Cycles = 0
}

func (B *Bus) Release() {
	KEEPBUS:
	// time.Since(start)
	// B.log <- elapsed
	// start = time.Now()
	B.vic.Run()
	B.Cycles++
	B.Timer++
	if !B.vic.BA {
		goto KEEPBUS
	}
}
