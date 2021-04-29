package cpu

import "go6502/mem"

func (C *CPU) op_TAX(mem *mem.Memory) {
	C.X = C.A
	C.setNZStatus(C.X)
	C.dbus.Release()

	if C.Display {
		C.opName = "TAX"
	}
}

func (C *CPU) op_TAY(mem *mem.Memory) {
	C.Y = C.A
	C.setNZStatus(C.Y)
	C.dbus.Release()

	if C.Display {
		C.opName = "TAY"
	}
}

func (C *CPU) op_TXA(mem *mem.Memory) {
	C.A = C.X
	C.setNZStatus(C.A)
	C.dbus.Release()

	if C.Display {
		C.opName = "TXA"
	}
}

func (C *CPU) op_TYA(mem *mem.Memory) {
	C.A = C.Y
	C.setNZStatus(C.A)
	C.dbus.Release()

	if C.Display {
		C.opName = "TYA"
	}
}
