package main

import "fmt"

func (C *CPU) op_NOP(mem *Memory) {
	C.opName = "NOP"
}

func (C *CPU) op_BRK(mem *Memory) {
	fmt.Printf("%-10s| PC:%04X | SP:%02X | A:%02X | X:%02X | Y:%02X | S:%08b\n", C.opName, C.PC, C.SP, C.A, C.X, C.Y, C.S)
	C.opName = "BRK"
	C.exit = true
}

func (C *CPU) op_DMP(mem *Memory) {
	fmt.Printf("%-10s| PC:%04X | SP:%02X | A:%02X | X:%02X | Y:%02X | S:%08b\n", C.opName, C.PC, C.SP, C.A, C.X, C.Y, C.S)
	C.opName = "DMP"
	absAddress := C.fetchWord(mem)
	mem.dump(absAddress)
	C.exit = true
}
