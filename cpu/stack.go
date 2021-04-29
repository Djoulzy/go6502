package cpu

import "go6502/mem"

func (C *CPU) op_TXS(mem *mem.Memory) {
	C.SP = C.X
	C.dbus.Release()

	if C.Display {
		C.opName = "TXS"
	}
}

func (C *CPU) op_TSX(mem *mem.Memory) {
	C.X = C.SP
	C.setNZStatus(C.X)
	C.dbus.Release()

	if C.Display {
		C.opName = "TSX"
	}
}

func (C *CPU) op_PHA(mem *mem.Memory) {
	C.pushByteStack(C.A)

	if C.Display {
		C.opName = "PHA"
	}
}

func (C *CPU) op_PLA(mem *mem.Memory) {
	C.A = C.pullByteStack()
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = "PLA"
	}
}

func (C *CPU) op_PHP(mem *mem.Memory) {
	C.pushByteStack(C.S)
	C.dbus.Release()

	if C.Display {
		C.opName = "PHP"
	}
}

func (C *CPU) op_PLP(mem *mem.Memory) {
	C.dbus.Release()
	C.S = C.pullByteStack()
	C.dbus.Release()

	if C.Display {
		C.opName = "PLP"
	}
}
