package cpu

import "go6502/mem"

func (C *CPU) op_AND_IM(mem *mem.Memory) {
	C.opName = "AND Imm"
	C.A &= C.fetchByte(mem)
	C.setNZStatus(C.A)
}

func (C *CPU) op_AND_ZP(mem *mem.Memory) {
	C.opName = "AND ZP"
	zpAddress := C.fetchByte(mem)
	C.A &= mem.Data[zpAddress]
	C.setNZStatus(C.A)
}

func (C *CPU) op_AND_ZPX(mem *mem.Memory) {
	C.opName = "ToDO"
}

func (C *CPU) op_AND_ABS(mem *mem.Memory) {
	C.opName = "ToDO"
}

func (C *CPU) op_AND_ABX(mem *mem.Memory) {
	C.opName = "ToDO"
}

func (C *CPU) op_AND_ABY(mem *mem.Memory) {
	C.opName = "ToDO"
}

func (C *CPU) op_AND_INX(mem *mem.Memory) {
	C.opName = "ToDO"
}

func (C *CPU) op_AND_INY(mem *mem.Memory) {
	C.opName = "ToDO"
}
