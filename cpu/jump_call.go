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
	C.opName = "JMP Ind"
	address := C.fetchWord(mem)
	C.PC = C.readWord(address)
	C.opName = fmt.Sprintf("JMP ($%04X) -> %04X", address, C.PC)
}

func (C *CPU) op_JSR(mem *mem.Memory) {
	C.pushWordStack(mem, C.PC + 1)
	address := C.fetchWord(mem)
	C.opName = fmt.Sprintf("JSR $%04X", address)
	C.PC = address
}

func (C *CPU) op_RTS(mem *mem.Memory) {
	C.opName = "RTS"
	dest := C.pullWordStack(mem)
	C.opName = fmt.Sprintf("RTS -> $%04X", dest)
	C.PC = dest + 1
}
