package cpu

import (
	"fmt"
	"go6502/mem"
)

//////////////////////////////////
///////////// INC ////////////////
//////////////////////////////////

func (C *CPU) op_INC_ZP(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	val := C.readByte(uint16(zpAddress))
	val += 1
	C.setNZStatus(val)
	C.writeByte(uint16(zpAddress), val)

	if C.Display {
		C.opName = fmt.Sprintf("INC $%02X", zpAddress)
	}
}

func (C *CPU) op_INC_ZPX(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem) + C.X
	val := C.readByte(uint16(zpAddress))
	val += 1
	C.setNZStatus(val)
	C.writeByte(uint16(zpAddress), val)

	if C.Display {
		C.opName = "INC ZP,X"
	}
}

func (C *CPU) op_INC_ABS(mem *mem.Memory) {
	address := C.fetchWord(mem)
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

func (C *CPU) op_INC_ABX(mem *mem.Memory) {
	absAddress := C.fetchWord(mem) + uint16(C.X)
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

func (C *CPU) op_DEC_ZP(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	val := C.readByte(uint16(zpAddress))
	val -= 1
	C.setNZStatus(val)
	C.dbus.Release()
	C.writeByte(uint16(zpAddress), val)

	if C.Display {
		C.opName = "DEC ZP"
	}
}

func (C *CPU) op_DEC_ZPX(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem) + C.X
	val := C.readByte(uint16(zpAddress))
	val -= 1
	C.dbus.Release()
	C.setNZStatus(val)
	C.writeByte(uint16(zpAddress), val)

	if C.Display {
		C.opName = "DEC ZP,X"
	}
}

func (C *CPU) op_DEC_ABS(mem *mem.Memory) {
	address := C.fetchWord(mem)
	val := C.readByte(address)
	val -= 1
	C.dbus.Release()
	C.setNZStatus(val)
	C.writeByte(address, val)

	if C.Display {
		C.opName = "DEC Abs"
	}
}

func (C *CPU) op_DEC_ABX(mem *mem.Memory) {
	absAddress := C.fetchWord(mem) + uint16(C.X)
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

func (C *CPU) op_DEX(mem *mem.Memory) {
	C.X -= 1
	C.setNZStatus(C.X)
	C.dbus.Release()

	if C.Display {
		C.opName = "DEX"
	}
}

func (C *CPU) op_DEY(mem *mem.Memory) {
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
func (C *CPU) op_INX(mem *mem.Memory) {
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

func (C *CPU) op_INY(mem *mem.Memory) {
	C.Y += 1
	C.setNZStatus(C.Y)
	C.dbus.Release()

	if C.Display {
		C.opName = "INY"
	}
}
