package cpu

import (
	"fmt"
	"go6502/mem"
)

func (C *CPU) op_JMP_ABS(mem *mem.Memory) {
	C.PC = C.fetchWord(mem)

	if C.Display {
		C.opName = fmt.Sprintf("JMP $%04X", C.PC)
	}
}

func (C *CPU) op_JMP_IND(mem *mem.Memory) {
	address := C.fetchWord(mem)
	C.PC = C.readWord(address)

	if C.Display {
		C.opName = fmt.Sprintf("JMP ($%04X)", address)
		C.debug = fmt.Sprintf("target: $%04X", C.PC)
	}
}

func (C *CPU) op_JSR(mem *mem.Memory) {
	C.pushWordStack(C.PC + 1)
	C.PC = C.fetchWord(mem)

	if C.Display {
		C.opName = fmt.Sprintf("JSR $%04X", C.PC)
	}
}

func (C *CPU) op_RTS(mem *mem.Memory) {
	C.PC = C.pullWordStack() + 1

	if C.Display {
		C.opName = "RTS"
		C.debug = fmt.Sprintf("target: $%04X", C.PC)
	}
}

func (C *CPU) op_RTI(mem *mem.Memory) {
	C.S = C.pullByteStack()
	C.PC = C.pullWordStack()

	if C.Display {
		C.opName = "RTI"
	}
}
