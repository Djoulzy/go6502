package cpu

import "go6502/mem"

func (C *CPU) op_TAX(mem *mem.Memory) {
	C.opName = "TAX"
	C.X = C.A
	C.setNZStatus(C.X)
	C.dbus.Release()
}

func (C *CPU) op_TAY(mem *mem.Memory) {
	C.opName = "TAY"
	C.Y = C.A
	C.setNZStatus(C.Y)
	C.dbus.Release()
}

func (C *CPU) op_TXA(mem *mem.Memory) {
	C.opName = "TXA"
	C.A = C.X
	C.setNZStatus(C.A)
	C.dbus.Release()
}

func (C *CPU) op_TYA(mem *mem.Memory) {
	C.opName = "TYA"
	C.A = C.Y
	C.setNZStatus(C.A)
	C.dbus.Release()
}
