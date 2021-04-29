package cpu

func (C *CPU) op_TAX() {
	C.X = C.A
	C.setNZStatus(C.X)
	C.dbus.Release()

	if C.Display {
		C.opName = "TAX"
	}
}

func (C *CPU) op_TAY() {
	C.Y = C.A
	C.setNZStatus(C.Y)
	C.dbus.Release()

	if C.Display {
		C.opName = "TAY"
	}
}

func (C *CPU) op_TXA() {
	C.A = C.X
	C.setNZStatus(C.A)
	C.dbus.Release()

	if C.Display {
		C.opName = "TXA"
	}
}

func (C *CPU) op_TYA() {
	C.A = C.Y
	C.setNZStatus(C.A)
	C.dbus.Release()

	if C.Display {
		C.opName = "TYA"
	}
}
