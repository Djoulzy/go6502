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

func (C *CPU) setC(on bool) {
	if on {
		C.S |= ^C_mask
	} else {
		C.S &= C_mask
	}
}

func (C *CPU) testC() bool {
	return C.S & ^C_mask > 0
}

func (C *CPU) setNZStatus(register globals.Byte) {
	C.setN(register)
	C.setZ(register)
}

func (C *CPU) testZ() bool {
	return C.S & ^Z_mask > 0
}

func (C *CPU) setV(m, n globals.Byte) {
	// c6 := (m & 0b01000000) & (n & 0b01000000)
	// m7 := m & 0b10000000
	// n7 := n & 0b10000000
	// if ((^m7 & ^n7 & c6) | (m7 & n7 & ^c6)) == 0 {
	// 	C.S &= V_mask
	// } else {
	// 	C.S |= ^V_mask
	// }
	result := m + n
	if (m^result)&(n^result)&0x80 != 0 {
		C.S |= ^V_mask
	} else {
		C.S &= V_mask
	}
}
