package cpu

import "go6502/mem"

func (C *CPU) op_BCC_REL(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_BCS_REL(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_BEQ_REL(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_BMI_REL(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_BNE_REL(mem *mem.Memory) {
	C.opName = "BNE"
	address := C.fetchWord(mem)
	if !C.testZ() {
		C.PC = address
	}
}

func (C *CPU) op_BPL_REL(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_BVC_REL(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_BVS_REL(mem *mem.Memory) { C.opName = "ToDO" }
