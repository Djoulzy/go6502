package cpu

func (C *CPU) op_NOP() {
	C.dbus.Release()

	if C.Display {
		C.opName = "NOP"
	}
}

func (C *CPU) op_SHW() {
	// fmt.Printf("\nCode      |  PC  | SP | A  | X  | Y  | NV-BDIZC\n")
	// fmt.Printf("%-10s| %04X | %02X | %02X | %02X | %02X | %08b\n", C.opName, C.PC, C.SP, C.A, C.X, C.Y, C.S)
}

func (C *CPU) op_BRK() {
	C.pushWordStack(C.PC)
	C.pushByteStack(C.S)
	address := C.readWord(0xFFFE)
	C.PC = address
	C.dbus.Release()
	C.setB(true)

	if C.Display {
		C.opName = "\tBRK"
	}
}

func (C *CPU) op_DMP() {
	// C.op_SHW(C.ram)
	// C.opName = "DMP"
	// absAddress := C.fetchWord(C.ram)
	// mem.Dump(absAddress)
	// C.exit = true
}
