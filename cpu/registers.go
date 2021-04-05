package cpu

import "go6502/mem"

func (C *CPU) op_TAX(mem *mem.Memory) {
	C.opName = "\tTAX"
	C.X = C.A
	C.setNZStatus(C.X)
	C.dbus.WaitBusLow()
}

func (C *CPU) op_TAY(mem *mem.Memory) {
	C.opName = "\tTAY"
	C.Y = C.A
	C.setNZStatus(C.Y)
	C.dbus.WaitBusLow()
}

func (C *CPU) op_TXA(mem *mem.Memory) {
	C.opName = "\tTXA"
	C.A = C.X
	C.setNZStatus(C.A)
	C.dbus.WaitBusLow()
}

func (C *CPU) op_TYA(mem *mem.Memory) {
	C.opName = "\tTYA"
	C.A = C.Y
	C.setNZStatus(C.A)
	C.dbus.WaitBusLow()
}
