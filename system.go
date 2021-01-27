package main

func (C *CPU) op_NOP(mem *Memory) {
	C.opName = "NOP     "
}

func (C *CPU) op_BRK(mem *Memory) {
	C.opName = "BRK     "
	C.exit = true
}
