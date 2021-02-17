// http://www.obelisk.me.uk/6502/

package main

import (
	"go6502/cpu"
	"go6502/mem"
	"go6502/vic"
	"runtime"
	"time"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	mem := mem.Memory{}
	mem.Init()

	// mem.dumpChar(0x2F)
	// os.Exit(1)
	cpu := cpu.CPU{}
	cpu.Init(&mem)

	vic := vic.VIC{}
	vic.Init(&mem, cpu.Cycle)

	go cpu.Run()

	if cpu.Display {
		for {
			cpu.Cycle <- true
			time.Sleep(time.Second)
		}
	}
	vic.Run()
}
