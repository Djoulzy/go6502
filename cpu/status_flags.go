package cpu

import "go6502/globals"

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

func (C *CPU) setC(register globals.Byte, value globals.Byte) {
	if register >= value {
		C.S |= ^C_mask
	} else {
		C.S &= C_mask
	}
}

func (C *CPU) setNZStatus(register globals.Byte) {
	C.setN(register)
	C.setZ(register)
}

func (C *CPU) testZ() bool {
	return C.S & ^Z_mask > 0
}
