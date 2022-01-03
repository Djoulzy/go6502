package cpu

import (
	"fmt"
)

func (C *CPU) op_JMP_ABS() {
	C.PC = C.fetchWord()

	if C.Display {
		C.opName = fmt.Sprintf("JMP $%04X", C.PC)
	}
}

func (C *CPU) op_JMP_IND() {
	address := C.fetchWord()
	C.PC = C.readWord(address)

	if C.Display {
		C.opName = fmt.Sprintf("JMP ($%04X)", address)
		C.debug = fmt.Sprintf("target: $%04X", C.PC)
	}
}

func (C *CPU) op_JSR() {
	C.pushWordStack(C.PC + 1)
	C.PC = C.fetchWord()

	if C.Display {
		C.opName = fmt.Sprintf("JSR $%04X", C.PC)
	}
}

func (C *CPU) op_RTS() {
	C.PC = C.pullWordStack() + 1

	if C.Display {
		C.opName = "RTS"
		C.debug = fmt.Sprintf("target: $%04X", C.PC)
	}
}

func (C *CPU) op_RTI() {
	C.S = C.pullByteStack()
	C.PC = C.pullWordStack()

	if C.Display {
		C.opName = "RTI"
	}
}
