// http://www.obelisk.me.uk/6502/

package main

import (
	"runtime"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	mem := Memory{}
	cpu := CPU{}
	vic := VIC{}

	cpu.cycle = make(chan bool, 1)
	cpu.display = true

	go cpu.run(&mem)

	// for {
	// 	cpu.cycle <- true
	// 	time.Sleep(time.Second)
	// }
	vic.run(&mem)
}
