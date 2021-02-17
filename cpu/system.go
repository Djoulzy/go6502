package cpu

import (
	"fmt"
	"go6502/mem"
)

func (C *CPU) op_NOP(mem *mem.Memory) {
	C.opName = "NOP"
}

func (C *CPU) op_BRK(mem *mem.Memory) {
	fmt.Printf("%-10s| PC:%04X | SP:%02X | A:%02X | X:%02X | Y:%02X | S:%08b\n", C.opName, C.PC, C.SP, C.A, C.X, C.Y, C.S)
	C.opName = "BRK"
	C.exit = true
}

func (C *CPU) op_DMP(mem *mem.Memory) {
	fmt.Printf("%-10s| PC:%04X | SP:%02X | A:%02X | X:%02X | Y:%02X | S:%08b\n", C.opName, C.PC, C.SP, C.A, C.X, C.Y, C.S)
	C.opName = "DMP"
	absAddress := C.fetchWord(mem)
	mem.Dump(absAddress)
	C.exit = true
}
