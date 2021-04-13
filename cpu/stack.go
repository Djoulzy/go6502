package cpu

import "go6502/mem"

func (C *CPU) op_TXS(mem *mem.Memory) {
	C.opName = "TXS"
	C.SP = C.X
}

func (C *CPU) op_TSX(mem *mem.Memory) {
	C.opName = "TSX"
	C.X = C.SP
	C.setNZStatus(C.X)
}

func (C *CPU) op_PHA(mem *mem.Memory) {
	C.opName = "PHA"
	C.pushByteStack(mem, C.A)
}

func (C *CPU) op_PLA(mem *mem.Memory) {
	C.opName = "PLA"
	C.A = C.pullByteStack(mem)
	C.setNZStatus(C.A)
}

func (C *CPU) op_PHP(mem *mem.Memory) {
	C.opName = "PHP"
	C.pushByteStack(mem, C.S)
	C.dbus.Release()
}

func (C *CPU) op_PLP(mem *mem.Memory) {
	C.opName = "PLP"
	C.dbus.Release()
	C.S = C.pullByteStack(mem)
	C.dbus.Release()
}
