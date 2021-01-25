package main

func (C *CPU) op_TXS(mem *Memory) {
	C.opName = "TXS     "
	// mem.Stack[C.SP] = C.X
	// C.SP++
}

func (C *CPU) op_PHA(mem *Memory) {
	C.opName = "PHA     "
	mem.Stack[C.SP] = C.A
	C.SP++
}

func (C *CPU) op_PLA(mem *Memory) {
	C.opName = "PLA     "
	C.SP--
	C.A = mem.Stack[C.SP]
}
