package main

func (C *CPU) op_JMP_ABS(mem *Memory) {
	C.opName = "JMP Abs "
	address := C.fetchWord(mem)
	C.PC = address
}

func (C *CPU) op_JMP_IND(mem *Memory) {

}

func (C *CPU) op_JSR(mem *Memory) {

}

func (C *CPU) op_RTS(mem *Memory) {

}