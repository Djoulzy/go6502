package cpu

import "go6502/mem"

func (C *CPU) op_BCC(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_BCS(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_BEQ(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_BMI(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_BNE(mem *mem.Memory) {
	C.opName = "BNE"
	address := C.fetchWord(mem)
	if !C.testZ() {
		C.PC = address
	}
}

func (C *CPU) op_BPL(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_BVC(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_BVS(mem *mem.Memory) { C.opName = "ToDO" }
