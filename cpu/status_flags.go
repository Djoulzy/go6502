package cpu

import (
	"go6502/globals"
	"go6502/mem"
)

func (C *CPU) setN(register globals.Byte) {
	if register&0b10000000 > 0 {
		C.S |= ^N_mask
	} else {
		C.S &= N_mask
	}
}

func (C *CPU) setZ(register globals.Byte) {
	if register == 0 {
		C.S |= ^Z_mask
	} else {
		C.S &= Z_mask
	}
}

func (C *CPU) setC(on bool) {
	if on {
		C.S |= ^C_mask
	} else {
		C.S &= C_mask
	}
}

func (C *CPU) testC() globals.Byte {
	if C.S & ^C_mask > 0 {
		return 0x01
	}
	return 0x00
}

func (C *CPU) setNZStatus(register globals.Byte) {
	C.setN(register)
	C.setZ(register)
}

func (C *CPU) testZ() bool {
	return C.S & ^Z_mask > 0
}

func (C *CPU) setV(m, n, result globals.Byte) {
	if (m^result)&(n^result)&0x80 != 0 {
		C.S |= ^V_mask
	} else {
		C.S &= V_mask
	}
}

func (C *CPU) op_CLC(mem *mem.Memory) {
	C.opName = "CLC"
	C.setC(false)
}

func (C *CPU) op_CLD(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_CLI(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_CLV(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_SEC(mem *mem.Memory) {
	C.opName = "SEC"
	C.setC(true)
}

func (C *CPU) op_SED(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_SEI(mem *mem.Memory) { C.opName = "ToDO" }
