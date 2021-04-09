package cpu

import (
	"fmt"
	"go6502/mem"
)

func (C *CPU) getRelativeAddr(dist byte) uint16 {
	signedDist := int(int8(dist))
	newAddr := int(C.PC) + signedDist
	return uint16(newAddr)
}

func (C *CPU) op_BCC_REL(mem *mem.Memory) {
	relative := C.fetchByte(mem)
	val := C.getRelativeAddr(relative)
	C.opName = fmt.Sprintf("BCC %02X -> $%04X", relative, val)
	if !C.testC() {
		C.PC = val
	}
}

func (C *CPU) op_BCS_REL(mem *mem.Memory) {
	relative := C.fetchByte(mem)
	val := C.getRelativeAddr(relative)
	C.opName = fmt.Sprintf("BCS %02X -> $%04X", relative, val)
	if C.testC() {
		C.PC = val
	}
}

func (C *CPU) op_BEQ_REL(mem *mem.Memory) {
	relative := C.fetchByte(mem)
	val := C.getRelativeAddr(relative)
	C.opName = fmt.Sprintf("BEQ %02X -> $%04X", relative, val)
	if C.testZ() {
		C.PC = val
	}
}

func (C *CPU) op_BMI_REL(mem *mem.Memory) {
	relative := C.fetchByte(mem)
	val := C.getRelativeAddr(relative)
	C.opName = fmt.Sprintf("BMI %02X -> $%04X", relative, val)
	if C.testN() {
		C.PC = val
	}
}

func (C *CPU) op_BNE_REL(mem *mem.Memory) {
	relative := C.fetchByte(mem)
	val := C.getRelativeAddr(relative)
	C.opName = fmt.Sprintf("BNE %02X -> $%04X", relative, val)
	if !C.testZ() {
		C.PC = val
	}
}

func (C *CPU) op_BPL_REL(mem *mem.Memory) {
	relative := C.fetchByte(mem)
	val := C.getRelativeAddr(relative)
	C.opName = fmt.Sprintf("BPL %02X -> $%04X", relative, val)
	if !C.testN() {
		C.PC = val
	}
}
func (C *CPU) op_BVC_REL(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_BVS_REL(mem *mem.Memory) { C.opName = "ToDO" }
