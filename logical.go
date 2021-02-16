package main

func (C *CPU) op_AND_IM(mem *Memory) {
	C.opName = "AND Imm"
	C.A &= C.fetchByte(mem)
	C.setNZStatus(C.A)
}

func (C *CPU) op_AND_ZP(mem *Memory) {
	C.opName = "AND ZP"
	zpAddress := C.fetchByte(mem)
	C.A &= mem.Data[zpAddress]
	C.setNZStatus(C.A)
}

func (C *CPU) op_AND_ZPX(mem *Memory) {
	C.opName = "ToDO"
}

func (C *CPU) op_AND_ABS(mem *Memory) {
	C.opName = "ToDO"
}

func (C *CPU) op_AND_ABX(mem *Memory) {
	C.opName = "ToDO"
}

func (C *CPU) op_AND_ABY(mem *Memory) {
	C.opName = "ToDO"
}

func (C *CPU) op_AND_INX(mem *Memory) {
	C.opName = "ToDO"
}

func (C *CPU) op_AND_INY(mem *Memory) {
	C.opName = "ToDO"
}