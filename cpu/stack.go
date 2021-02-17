package cpu

import "go6502/mem"

func (C *CPU) op_TXS(mem *mem.Memory) {
	C.opName = "TXS"
	// mem.Stack[C.SP] = C.X
	// C.SP++
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
