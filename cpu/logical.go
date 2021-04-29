package cpu

import (
	"fmt"
	"go6502/mem"
)

func (C *CPU) op_AND_IM(mem *mem.Memory) {
	val := C.fetchByte(mem)
	C.A &= val
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("AND #$%02X", val)
	}
}

func (C *CPU) op_AND_ZP(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	C.A &= C.readByte(uint16(zpAddress))
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = "AND ZP"
	}
}

func (C *CPU) op_AND_ZPX(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem) + C.X
	C.A &= C.readByte(uint16(zpAddress))
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = "AND ZP,X"
	}
}

func (C *CPU) op_AND_ABS(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	C.A &= C.readByte(absAddress)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = "AND Abs"
	}
}

func (C *CPU) op_AND_ABX(mem *mem.Memory) {
	absAddress := C.fetchWord(mem) + uint16(C.X)
	C.A &= C.readByte(absAddress)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = "AND Abs,X"
	}
}

func (C *CPU) op_AND_ABY(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	dest := absAddress + uint16(C.Y)
	C.A &= C.readByte(dest)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("AND $%04X,Y", absAddress)
	}
}

func (C *CPU) op_AND_INX(mem *mem.Memory) {
	zpAddr := C.fetchByte(mem)
	wordZP := C.Indexed_indirect_X(zpAddr, C.X)
	C.A &= C.readByte(wordZP)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("AND ($%02X,X)", zpAddr)
	}
}

func (C *CPU) op_AND_INY(mem *mem.Memory) {
	zpAddr := C.fetchByte(mem)
	wordZP := C.Indirect_index_Y(zpAddr, C.Y)
	C.A &= C.readByte(wordZP)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("AND ($%02X),Y", zpAddr)
	}
}

func (C *CPU) op_EOR_IM(mem *mem.Memory) {
	val := C.fetchByte(mem)
	C.A ^= val
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("EOR #$%02X", val)
	}
}

func (C *CPU) op_EOR_ZP(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	C.A ^= C.readByte(uint16(zpAddress))
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("EOR $%02X", zpAddress)
	}
}

func (C *CPU) op_EOR_ZPX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_EOR_ABS(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_EOR_ABX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_EOR_ABY(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_EOR_INX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_EOR_INY(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_ORA_IM(mem *mem.Memory) {
	val := C.fetchByte(mem)
	C.A |= val
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("ORA #$%02X", val)
	}
}

func (C *CPU) op_ORA_ZP(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	C.A |= C.readByte(uint16(zpAddress))
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("ORA $%02X", zpAddress)
	}
}

func (C *CPU) op_ORA_ZPX(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_ORA_ABS(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	C.A |= C.readByte(absAddress)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("ORA $%04X", absAddress)
	}
}

func (C *CPU) op_ORA_ABX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ORA_ABY(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ORA_INX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ORA_INY(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_BIT_ZP(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	val := C.readByte(uint16(zpAddress))
	if val&0b01000000 != 0 {
		C.S |= ^V_mask
	} else {
		C.S &= V_mask
	}
	result := val & C.A
	C.setNZStatus(result)

	if C.Display {
		C.opName = fmt.Sprintf("BIT $%02X", zpAddress)
	}
}

func (C *CPU) op_BIT_ABS(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	val := C.readByte(absAddress)
	if val&0b01000000 != 0 {
		C.S |= ^V_mask
	} else {
		C.S &= V_mask
	}
	result := val & C.A
	C.setNZStatus(result)

	if C.Display {
		C.opName = fmt.Sprintf("BIT $%04X", absAddress)
	}
}
