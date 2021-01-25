package main

func (C *CPU) op_JMP_ABS(mem *Memory) {
	C.opName = "JMP Abs "
	address := Word(C.fetchByte(mem)) << 8
	address += Word(C.fetchByte(mem))
	C.PC = address
}

func (C *CPU) op_JMP_IND(mem *Memory) {

}
