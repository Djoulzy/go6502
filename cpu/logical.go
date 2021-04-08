package cpu

import (
	"fmt"
	"go6502/mem"
)

func (C *CPU) op_AND_IM(mem *mem.Memory) {
	C.opName = "AND #$"
	val := C.fetchByte(mem)
	C.A &= val
	C.opName = fmt.Sprintf("%s%02X", C.opName, val)
	C.setNZStatus(C.A)
}

func (C *CPU) op_AND_ZP(mem *mem.Memory) {
	C.opName = "AND ZP"
	zpAddress := C.fetchByte(mem)
	C.A &= mem.Read(zpAddress)
	C.setNZStatus(C.A)
}

func (C *CPU) op_AND_ZPX(mem *mem.Memory) {
	C.opName = "AND ZP,X"
	zpAddress := C.fetchByte(mem) + C.X
	C.A &= mem.Read(zpAddress)
	C.setNZStatus(C.A)
}

func (C *CPU) op_AND_ABS(mem *mem.Memory) {
	C.opName = "AND Abs"
	absAddress := C.fetchWord(mem)
	C.A &= mem.Read(absAddress)
	C.setNZStatus(C.A)
}

func (C *CPU) op_AND_ABX(mem *mem.Memory) {
	C.opName = "AND Abs,X"
	absAddress := C.fetchWord(mem) + uint16(C.X)
	C.A &= mem.Read(absAddress)
	C.setNZStatus(C.A)
}

func (C *CPU) op_AND_ABY(mem *mem.Memory) {
	C.opName = "AND Abs,Y"
	absAddress := C.fetchWord(mem) + uint16(C.Y)
	C.A &= mem.Read(absAddress)
	C.setNZStatus(C.A)
}

func (C *CPU) op_AND_INX(mem *mem.Memory) {
	C.opName = "AND (ZP,X)"
	zpAddr := C.fetchByte(mem)
	wordZP := C.Indexed_indirect_X(zpAddr, C.X)
	C.A &= mem.Read(wordZP)
	C.setNZStatus(C.A)
}

func (C *CPU) op_AND_INY(mem *mem.Memory) {
	C.opName = "AND (ZP),Y"
	zpAddr := C.fetchByte(mem)
	wordZP := C.Indirect_index_Y(zpAddr, C.Y)
	C.A &= mem.Read(wordZP)
	C.setNZStatus(C.A)
}

func (C *CPU) op_EOR_IM(mem *mem.Memory)  { C.opName = "ToDO" }
func (C *CPU) op_EOR_ZP(mem *mem.Memory)  { C.opName = "ToDO" }
func (C *CPU) op_EOR_ZPX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_EOR_ABS(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_EOR_ABX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_EOR_ABY(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_EOR_INX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_EOR_INY(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_ORA_IM(mem *mem.Memory) {
	C.opName = "ORA #$"
	val := C.fetchByte(mem)
	C.A |= val
	C.opName = fmt.Sprintf("%s%02X", C.opName, val)
	C.setNZStatus(C.A)
}

func (C *CPU) op_ORA_ZP(mem *mem.Memory)  { C.opName = "ToDO" }
func (C *CPU) op_ORA_ZPX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ORA_ABS(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ORA_ABX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ORA_ABY(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ORA_INX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ORA_INY(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_BIT_ZP(mem *mem.Memory)  { C.opName = "ToDO" }
func (C *CPU) op_BIT_ABS(mem *mem.Memory) { C.opName = "ToDO" }
