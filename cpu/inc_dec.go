package cpu

import (
	"fmt"
)

//////////////////////////////////
///////////// INC ////////////////
//////////////////////////////////

func (C *CPU) op_INC_ZP() {
	zpAddress := C.fetchByte()
	val := C.readByte(uint16(zpAddress))
	val += 1
	C.setNZStatus(val)
	C.writeByte(uint16(zpAddress), val)

	if C.Display {
		C.opName = fmt.Sprintf("INC $%02X", zpAddress)
	}
}

func (C *CPU) op_INC_ZPX() {
	zpAddress := C.fetchByte() + C.X
	val := C.readByte(uint16(zpAddress))
	val += 1
	C.setNZStatus(val)
	C.writeByte(uint16(zpAddress), val)

	if C.Display {
		C.opName = "INC ZP,X"
	}
}

func (C *CPU) op_INC_ABS() {
	address := C.fetchWord()
	C.dbus.Release()
	val := C.readByte(address)
	C.dbus.Release()
	val += 1
	C.dbus.Release()
	C.setNZStatus(val)
	C.writeByte(address, val)

	if C.Display {
		C.opName = "INC Abs"
	}
}

func (C *CPU) op_INC_ABX() {
	absAddress := C.fetchWord() + uint16(C.X)
	val := C.readByte(absAddress)
	val += 1
	C.setNZStatus(val)
	C.writeByte(absAddress, val)

	if C.Display {
		C.opName = "INC Abs,X"
	}
}

//////////////////////////////////
///////////// DEC ////////////////
//////////////////////////////////

func (C *CPU) op_DEC_ZP() {
	zpAddress := C.fetchByte()
	val := C.readByte(uint16(zpAddress))
	val -= 1
	C.setNZStatus(val)
	C.dbus.Release()
	C.writeByte(uint16(zpAddress), val)

	if C.Display {
		C.opName = "DEC ZP"
	}
}

func (C *CPU) op_DEC_ZPX() {
	zpAddress := C.fetchByte() + C.X
	val := C.readByte(uint16(zpAddress))
	val -= 1
	C.dbus.Release()
	C.setNZStatus(val)
	C.writeByte(uint16(zpAddress), val)

	if C.Display {
		C.opName = "DEC ZP,X"
	}
}

func (C *CPU) op_DEC_ABS() {
	address := C.fetchWord()
	val := C.readByte(address)
	val -= 1
	C.dbus.Release()
	C.setNZStatus(val)
	C.writeByte(address, val)

	if C.Display {
		C.opName = "DEC Abs"
	}
}

func (C *CPU) op_DEC_ABX() {
	absAddress := C.fetchWord() + uint16(C.X)
	C.dbus.Release()
	val := C.readByte(absAddress)
	val -= 1
	C.dbus.Release()
	C.setNZStatus(val)
	C.writeByte(absAddress, val)

	if C.Display {
		C.opName = "DEC Abs,X"
	}
}

func (C *CPU) op_DEX() {
	C.X -= 1
	C.setNZStatus(C.X)
	C.dbus.Release()

	if C.Display {
		C.opName = "DEX"
	}
}

func (C *CPU) op_DEY() {
	C.Y -= 1
	C.setNZStatus(C.Y)
	C.dbus.Release()

	if C.Display {
		C.opName = "DEY"
	}
}

//////////////////////////////////
///////////// INX ////////////////
//////////////////////////////////

// op_INX : Increment X
func (C *CPU) op_INX() {
	C.X += 1
	C.setNZStatus(C.X)
	C.dbus.Release()

	if C.Display {
		C.opName = "INX"
	}
}

//////////////////////////////////
///////////// INY ////////////////
//////////////////////////////////

func (C *CPU) op_INY() {
	C.Y += 1
	C.setNZStatus(C.Y)
	C.dbus.Release()

	if C.Display {
		C.opName = "INY"
	}
}
