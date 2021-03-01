// http://www.obelisk.me.uk/6502/

package main

import (
	"go6502/assembler"
	"go6502/cpu"
	"go6502/mem"
	"go6502/vic"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	args := os.Args

	// test := 0xF4
	// fmt.Printf("%d\n", int8(test))
	// os.Exit(1)

	mem := mem.Memory{}
	mem.Init()

	cpu := cpu.CPU{}
	cpu.Init(&mem, false)

	if len(args) > 1 {
		ass := assembler.Assembler{}
		ass.Init()

		switch filepath.Ext(args[1]) {
		case ".asm":
			code := ass.Assemble(args[1])
			cpu.PC, _ = assembler.LoadHex(&mem, code)
		case ".hex":
			cpu.PC, _ = assembler.LoadFile(&mem, args[1])
		}
	}

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
