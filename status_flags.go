package main

func (C *CPU) setN(register Byte) {
	if register&0b10000000 > 0 {
		C.S |= ^N_mask
	} else {
		C.S &= N_mask
	}
}

func (C *CPU) setZ(register Byte) {
	if register == 0 {
		C.S |= ^Z_mask
	} else {
		C.S &= Z_mask
	}
}

func (C *CPU) setC(register Byte, value Byte) {
	if register >= value {
		C.S |= ^C_mask
	} else {
		C.S &= C_mask
	}
}

func (C *CPU) setNZStatus(register Byte) {
	C.setN(register)
	C.setZ(register)
}

func (C *CPU) testZ() bool {
	return C.S & ^Z_mask > 0
}
