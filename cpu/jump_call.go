package cpu

import (
	"go6502/mem"
)

func (C *CPU) op_JMP_ABS(mem *mem.Memory) {
	C.opName = "JMP Abs"
	address := C.fetchWord(mem)
	C.PC = address
}

func (C *CPU) op_JMP_IND(mem *mem.Memory) {
	C.opName = "JMP Ind"
	address := C.fetchWord(mem)
	C.PC = C.readWord(address)
}

func (C *CPU) op_JSR(mem *mem.Memory) {
	C.opName = "JSR"
	C.pushWordStack(mem, C.PC+1)
	address := C.fetchWord(mem)
	C.PC = address
}

func (C *CPU) op_RTS(mem *mem.Memory) {
	C.opName = "RTS"
	C.PC = C.fetchWordStack(mem) + 1
}
