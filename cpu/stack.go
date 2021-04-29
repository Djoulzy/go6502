package cpu

func (C *CPU) op_TXS() {
	C.SP = C.X
	C.dbus.Release()

	if C.Display {
		C.opName = "TXS"
	}
}

func (C *CPU) op_TSX() {
	C.X = C.SP
	C.setNZStatus(C.X)
	C.dbus.Release()

	if C.Display {
		C.opName = "TSX"
	}
}

func (C *CPU) op_PHA() {
	C.pushByteStack(C.A)

	if C.Display {
		C.opName = "PHA"
	}
}

func (C *CPU) op_PLA() {
	C.A = C.pullByteStack()
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = "PLA"
	}
}

func (C *CPU) op_PHP() {
	C.pushByteStack(C.S)
	C.dbus.Release()

	if C.Display {
		C.opName = "PHP"
	}
}

func (C *CPU) op_PLP() {
	C.dbus.Release()
	C.S = C.pullByteStack()
	C.dbus.Release()

	if C.Display {
		C.opName = "PLP"
	}
}
