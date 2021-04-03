package cpu

import (
	"fmt"
	"go6502/mem"
)

func (C *CPU) op_NOP(mem *mem.Memory) {
	C.opName = "NOP"
}

func (C *CPU) op_SHW(mem *mem.Memory) {
	fmt.Printf("Code      |  PC  | SP | A  | X  | Y  | NV-BDIZC\n")
	fmt.Printf("%-10s| %04X | %02X | %02X | %02X | %02X | %08b\n", C.opName, C.PC, C.SP, C.A, C.X, C.Y, C.S)
}

func (C *CPU) op_BRK(mem *mem.Memory) {
	// C.op_SHW(mem)
	C.opName = "BRK"
	C.pushWordStack(mem, C.PC)
	C.pushByteStack(mem, C.S)
	address := C.readWord(0xFFFE)
	C.PC = address
	C.dbus.WaitBusLow()
	C.setB(true)
}

func (C *CPU) op_DMP(mem *mem.Memory) {
	C.op_SHW(mem)
	C.opName = "DMP"
	absAddress := C.fetchWord(mem)
	mem.Dump(absAddress)
	C.exit = true
}
