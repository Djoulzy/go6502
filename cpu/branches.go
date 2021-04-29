package cpu

import (
	"fmt"
)

func (C *CPU) getRelativeAddr(dist byte) uint16 {
	signedDist := int(int8(dist))
	newAddr := int(C.PC) + signedDist
	return uint16(newAddr)
}

func (C *CPU) op_BCC_REL() {
	relative := C.fetchByte()
	val := C.getRelativeAddr(relative)
	if !C.testC() {
		C.PC = val
	}

	if C.Display {
		C.opName = fmt.Sprintf("BCC $%04X", val)
	}
}

func (C *CPU) op_BCS_REL() {
	relative := C.fetchByte()
	val := C.getRelativeAddr(relative)
	if C.testC() {
		C.PC = val
	}

	if C.Display {
		C.opName = fmt.Sprintf("BCS $%04X", val)
	}
}

func (C *CPU) op_BEQ_REL() {
	relative := C.fetchByte()
	val := C.getRelativeAddr(relative)
	if C.testZ() {
		C.PC = val
	}

	if C.Display {
		C.opName = fmt.Sprintf("BEQ $%04X", val)
	}
}

func (C *CPU) op_BMI_REL() {
	relative := C.fetchByte()
	val := C.getRelativeAddr(relative)
	if C.testN() {
		C.PC = val
	}

	if C.Display {
		C.opName = fmt.Sprintf("BMI $%04X", val)
	}
}

func (C *CPU) op_BNE_REL() {
	relative := C.fetchByte()
	val := C.getRelativeAddr(relative)
	if !C.testZ() {
		C.PC = val
	}

	if C.Display {
		C.opName = fmt.Sprintf("BNE $%04X", val)
	}
}

func (C *CPU) op_BPL_REL() {
	relative := C.fetchByte()
	val := C.getRelativeAddr(relative)
	if !C.testN() {
		C.PC = val
	}

	if C.Display {
		C.opName = fmt.Sprintf("BPL $%04X", val)
	}
}
func (C *CPU) op_BVC_REL() { C.opName = "ToDO" }

func (C *CPU) op_BVS_REL() {
	relative := C.fetchByte()
	val := C.getRelativeAddr(relative)
	if C.testV() {
		C.PC = val
	}

	if C.Display {
		C.opName = fmt.Sprintf("BVS $%04X", val)
	}
}
