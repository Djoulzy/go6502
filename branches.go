package main

func (C *CPU) op_BCC(mem *Memory) { C.opName = "ToDO" }
func (C *CPU) op_BCS(mem *Memory) { C.opName = "ToDO" }
func (C *CPU) op_BEQ(mem *Memory) { C.opName = "ToDO" }
func (C *CPU) op_BMI(mem *Memory) { C.opName = "ToDO" }

func (C *CPU) op_BNE(mem *Memory) {
	C.opName = "BNE"
	address := C.fetchWord(mem)
	if !C.testZ() {
		C.PC = address
	}
}

func (C *CPU) op_BPL(mem *Memory) { C.opName = "ToDO" }
func (C *CPU) op_BVC(mem *Memory) { C.opName = "ToDO" }
func (C *CPU) op_BVS(mem *Memory) { C.opName = "ToDO" }
