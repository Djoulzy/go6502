package cpu

import (
	"go6502/globals"
	"go6502/mem"
)

func (C *CPU) getRelativeAddr(dist globals.Byte) globals.Word {
	signedDist := int(int8(dist))
	newAddr := int(C.PC) + signedDist
	return globals.Word(newAddr)
}

func (C *CPU) op_BCC_REL(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_BCS_REL(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_BEQ_REL(mem *mem.Memory) {
	C.opName = "BEQ"
	if C.testZ() {
		C.PC = C.getRelativeAddr(C.fetchByte(mem))
	}
}

func (C *CPU) op_BMI_REL(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_BNE_REL(mem *mem.Memory) {
	C.opName = "BNE"
	if !C.testZ() {
		C.PC = C.getRelativeAddr(C.fetchByte(mem))
	}
}

func (C *CPU) op_BPL_REL(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_BVC_REL(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_BVS_REL(mem *mem.Memory) { C.opName = "ToDO" }
