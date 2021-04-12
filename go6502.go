// http://www.obelisk.me.uk/6502/

package main

import (
	"go6502/assembler"
	"go6502/cia"
	"go6502/clog"
	"go6502/confload"
	"go6502/cpu"
	"go6502/databus"
	"go6502/mem"
	"go6502/vic"
	"os"
	"path/filepath"
	"runtime"
)

var conf = &confload.ConfigData{}

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	args := os.Args
	confload.Load("config.ini", conf)

	clog.LogLevel = conf.LogLevel
	clog.StartLogging = conf.StartLogging
	if conf.FileLog != "" {
		clog.EnableFileLog(conf.FileLog)
	}

	dbus := databus.Databus{}
	dbus.Init()

	mem := mem.Memory{}
	mem.Init()

	cia2 := cia.CIA{}
	cia2.Init(mem.Mem[0xDD00:])

	cpu := cpu.CPU{}
	cpu.Init(&dbus, &mem, conf)

	if len(args) > 1 {
		ass := assembler.Assembler{}
		ass.Init()

		switch filepath.Ext(args[1]) {
		case ".asm":
			code := ass.Assemble(args[1])
			cpu.PC, _ = assembler.LoadHex(&mem, code)
		case ".hex":
			cpu.PC, _ = assembler.LoadFile(&mem, args[1])
		case ".prg":
			cpu.PC, _ = assembler.LoadPRG(&mem, args[1])
		default:
			cpu.PC = 0xFCE2 // Reset call
		}
	}
	cpu.PC = 0xFCE2

	vic := vic.VIC{}
	vic.Init(&dbus, &mem)

	go cpu.Run()
	vic.Run()
}
