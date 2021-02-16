// http://www.obelisk.me.uk/6502/

package main

import (
	"runtime"
	"time"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	mem := Memory{}
	mem.Init()

	// mem.dumpChar(0x2F)
	// os.Exit(1)
	cpu := CPU{}

	cpu.cycle = make(chan bool, 1)
	cpu.display = false

	vic := VIC{}

	go cpu.run(&mem)

	if cpu.display {
		for {
			cpu.cycle <- true
			time.Sleep(time.Second)
		}
	}
	vic.run(&mem, cpu.cycle)
}
