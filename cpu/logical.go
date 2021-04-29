package cpu

import (
	"fmt"
)

func (C *CPU) op_AND_IM() {
	val := C.fetchByte()
	C.A &= val
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("AND #$%02X", val)
	}
}

func (C *CPU) op_AND_ZP() {
	zpAddress := C.fetchByte()
	C.A &= C.readByte(uint16(zpAddress))
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = "AND ZP"
	}
}

func (C *CPU) op_AND_ZPX() {
	zpAddress := C.fetchByte() + C.X
	C.A &= C.readByte(uint16(zpAddress))
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = "AND ZP,X"
	}
}

func (C *CPU) op_AND_ABS() {
	absAddress := C.fetchWord()
	C.A &= C.readByte(absAddress)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = "AND Abs"
	}
}

func (C *CPU) op_AND_ABX() {
	absAddress := C.fetchWord() + uint16(C.X)
	C.A &= C.readByte(absAddress)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = "AND Abs,X"
	}
}

func (C *CPU) op_AND_ABY() {
	absAddress := C.fetchWord()
	dest := absAddress + uint16(C.Y)
	C.A &= C.readByte(dest)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("AND $%04X,Y", absAddress)
	}
}

func (C *CPU) op_AND_INX() {
	zpAddr := C.fetchByte()
	wordZP := C.Indexed_indirect_X(zpAddr, C.X)
	C.A &= C.readByte(wordZP)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("AND ($%02X,X)", zpAddr)
	}
}

func (C *CPU) op_AND_INY() {
	zpAddr := C.fetchByte()
	wordZP := C.Indirect_index_Y(zpAddr, C.Y)
	C.A &= C.readByte(wordZP)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("AND ($%02X),Y", zpAddr)
	}
}

func (C *CPU) op_EOR_IM() {
	val := C.fetchByte()
	C.A ^= val
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("EOR #$%02X", val)
	}
}

func (C *CPU) op_EOR_ZP() {
	zpAddress := C.fetchByte()
	C.A ^= C.readByte(uint16(zpAddress))
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("EOR $%02X", zpAddress)
	}
}

func (C *CPU) op_EOR_ZPX() { C.opName = "ToDO" }
func (C *CPU) op_EOR_ABS() { C.opName = "ToDO" }
func (C *CPU) op_EOR_ABX() { C.opName = "ToDO" }
func (C *CPU) op_EOR_ABY() { C.opName = "ToDO" }
func (C *CPU) op_EOR_INX() { C.opName = "ToDO" }
func (C *CPU) op_EOR_INY() { C.opName = "ToDO" }

func (C *CPU) op_ORA_IM() {
	val := C.fetchByte()
	C.A |= val
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("ORA #$%02X", val)
	}
}

func (C *CPU) op_ORA_ZP() {
	zpAddress := C.fetchByte()
	C.A |= C.readByte(uint16(zpAddress))
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("ORA $%02X", zpAddress)
	}
}

func (C *CPU) op_ORA_ZPX() { C.opName = "ToDO" }

func (C *CPU) op_ORA_ABS() {
	absAddress := C.fetchWord()
	C.A |= C.readByte(absAddress)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("ORA $%04X", absAddress)
	}
}

func (C *CPU) op_ORA_ABX() { C.opName = "ToDO" }
func (C *CPU) op_ORA_ABY() { C.opName = "ToDO" }
func (C *CPU) op_ORA_INX() { C.opName = "ToDO" }
func (C *CPU) op_ORA_INY() { C.opName = "ToDO" }

func (C *CPU) op_BIT_ZP() {
	zpAddress := C.fetchByte()
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

func (C *CPU) op_BIT_ABS() {
	absAddress := C.fetchWord()
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
