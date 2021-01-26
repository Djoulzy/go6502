package main

func (C *CPU) op_JMP_ABS(mem *Memory) {
	C.opName = "JMP Abs "
	address := C.fetchWord(mem)
	C.PC = address
}

func (C *CPU) op_JMP_IND(mem *Memory) {

}

func (C *CPU) op_JSR(mem *Memory) {
	C.opName = "JSR     "
	address := C.fetchWord(mem)
	C.pushWordStack(mem, C.PC)
	C.PC = address
}

func (C *CPU) op_RTS(mem *Memory) {
	C.PC = C.fetchWordStack(mem)
}
