package cpu

import (
	"fmt"
	"go6502/mem"
)

func (C *CPU) op_JMP_ABS(mem *mem.Memory) {
	address := C.fetchWord(mem)
	C.opName = fmt.Sprintf("JMP $%04X", address)
	C.PC = address
}

func (C *CPU) op_JMP_IND(mem *mem.Memory) {
	address := C.fetchWord(mem)
	C.PC = C.readWord(address)
	C.opName = fmt.Sprintf("JMP ($%04X)", address)
	C.debug = fmt.Sprintf("target: $%04X", C.PC)
}

func (C *CPU) op_JSR(mem *mem.Memory) {
	// originPC := C.PC - 1
	C.pushWordStack(C.PC + 1)
	// stack := C.PC + 1
	address := C.fetchWord(mem)
	C.opName = fmt.Sprintf("JSR $%04X", address)
	C.PC = address
	// clog.File("go6502", "JSR", "PC:%04X -> %04X - Push: %04X - Return to: %04X", originPC, address, stack, stack+1)
}

func (C *CPU) op_RTS(mem *mem.Memory) {
	// originPC := C.PC - 1
	C.opName = "RTS"
	dest := C.pullWordStack()
	C.debug = fmt.Sprintf("target: $%04X", dest)
	C.PC = dest + 1
	// clog.File("go6502", "RTS", "PC:%04X -> %04X - Pull: %04X - Call from: %04X", originPC, C.PC, dest, dest-2)
}

func (C *CPU) op_RTI(mem *mem.Memory) {
	C.S = C.pullByteStack()
	C.PC = C.pullWordStack()
}
