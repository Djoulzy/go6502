package cpu

import (
	"go6502/globals"
	"go6502/mem"
)

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
	C.opName = "AND ZP,X"
	zpAddress := C.fetchByte(mem) + C.X
	C.A &= mem.Data[zpAddress]
	C.setNZStatus(C.A)
}

func (C *CPU) op_AND_ABS(mem *mem.Memory) {
	C.opName = "AND Abs"
	absAddress := C.fetchWord(mem)
	C.A &= mem.Data[absAddress]
	C.setNZStatus(C.A)
}

func (C *CPU) op_AND_ABX(mem *mem.Memory) {
	C.opName = "AND Abs,X"
	absAddress := C.fetchWord(mem) + globals.Word(C.X)
	C.A &= mem.Data[absAddress]
	C.setNZStatus(C.A)
}

func (C *CPU) op_AND_ABY(mem *mem.Memory) {
	C.opName = "AND Abs,Y"
	absAddress := C.fetchWord(mem) + globals.Word(C.Y)
	C.A &= mem.Data[absAddress]
	C.setNZStatus(C.A)
}

func (C *CPU) op_AND_INX(mem *mem.Memory) {
	C.opName = "AND (ZP,X)"
	zpAddr := C.fetchByte(mem)
	wordZP := C.Indexed_indirect_X(zpAddr, C.X)
	C.A &= mem.Data[wordZP]
	C.setNZStatus(C.A)
}

func (C *CPU) op_AND_INY(mem *mem.Memory) {
	C.opName = "AND (ZP),Y"
	zpAddr := C.fetchByte(mem)
	wordZP := C.Indirect_index_Y(zpAddr, C.Y)
	C.A &= mem.Data[wordZP]
	C.setNZStatus(C.A)
}
