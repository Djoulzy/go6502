package cpu

import (
	"go6502/mem"
)

func (C *CPU) setN(register byte) {
	if register&0b10000000 > 0 {
		C.S |= ^N_mask
	} else {
		C.S &= N_mask
	}
}

func (C *CPU) testN() bool {
	return C.S & ^N_mask > 0
}

func (C *CPU) setZ(register byte) {
	if register == 0 {
		C.S |= ^Z_mask
	} else {
		C.S &= Z_mask
	}
}

func (C *CPU) setD(on bool) {
	if on {
		C.S |= ^D_mask
	} else {
		C.S &= D_mask
	}
}

func (C *CPU) setI(on bool) {
	if on {
		C.S |= ^I_mask
	} else {
		C.S &= I_mask
	}
}

func (C *CPU) setB(on bool) {
	if on {
		C.S |= ^B_mask
	} else {
		C.S &= B_mask
	}
}

func (C *CPU) setC(on bool) {
	if on {
		C.S |= ^C_mask
	} else {
		C.S &= C_mask
	}
}

func (C *CPU) getC() byte {
	if C.S & ^C_mask > 0 {
		return 0x01
	}
	return 0x00
}

func (C *CPU) testC() bool {
	return C.S & ^C_mask > 0
}

func (C *CPU) setNZStatus(register byte) {
	C.setN(register)
	C.setZ(register)
}

func (C *CPU) testZ() bool {
	return C.S & ^Z_mask > 0
}

func (C *CPU) setV(m, n, result byte) {
	if (m^result)&(n^result)&0x80 != 0 {
		C.S |= ^V_mask
	} else {
		C.S &= V_mask
	}
}

func (C *CPU) op_CLC(mem *mem.Memory) {
	C.opName = "CLC"
	C.setC(false)
	C.dbus.WaitBusLow()
}

func (C *CPU) op_CLD(mem *mem.Memory) {
	C.opName = "CLD"
	C.setD(false)
	C.dbus.WaitBusLow()
}

func (C *CPU) op_CLI(mem *mem.Memory) {
	C.opName = "CLI"
	C.setI(false)
	C.dbus.WaitBusLow()
}

func (C *CPU) op_CLV(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_SEC(mem *mem.Memory) {
	C.opName = "SEC"
	C.setC(true)
	C.dbus.WaitBusLow()
}

func (C *CPU) op_SED(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_SEI(mem *mem.Memory) {
	C.opName = "SEI"
	C.setI(true)
	C.dbus.WaitBusLow()
}
