package main

func (C *CPU) op_TXS(mem *Memory) {
	C.opName = "TXS     "
	// mem.Stack[C.SP] = C.X
	// C.SP++
}

func (C *CPU) op_PHA(mem *Memory) {
	C.opName = "PHA     "
	C.pushByteStack(mem, C.A)
}

func (C *CPU) op_PLA(mem *Memory) {
	C.opName = "PLA     "
	C.A = C.pullByteStack(mem)
	C.setLDAStatus(C.A)
}
