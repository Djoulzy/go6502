package cpu

import (
	"fmt"
	"go6502/mem"
)

//////////////////////////////////
///////////// INC ////////////////
//////////////////////////////////

func (C *CPU) op_INC_ZP(mem *mem.Memory) {
	zpAddr := C.fetchByte(mem)
	C.opName = fmt.Sprintf("INC $%02X", zpAddr)
	val := mem.Read(zpAddr)
	val += 1
	C.setNZStatus(val)
	mem.Write(zpAddr, val)
}

func (C *CPU) op_INC_ZPX(mem *mem.Memory) {
	C.opName = "INC ZP,X"
	zpAddr := C.fetchByte(mem) + C.X
	val := mem.Read(zpAddr)
	val += 1
	C.setNZStatus(val)
	mem.Write(zpAddr, val)
}

func (C *CPU) op_INC_ABS(mem *mem.Memory) {
	C.opName = "INC Abs"
	address := C.fetchWord(mem)
	C.dbus.Release()
	val := mem.Read(address)
	C.dbus.Release()
	val += 1
	C.dbus.Release()
	C.setNZStatus(val)
	mem.Write(address, val)
}

func (C *CPU) op_INC_ABX(mem *mem.Memory) {
	C.opName = "INC Abs,X"
	absAddress := C.fetchWord(mem) + uint16(C.X)
	val := mem.Read(absAddress)
	val += 1
	C.setNZStatus(val)
	mem.Write(absAddress, val)
}

//////////////////////////////////
///////////// DEC ////////////////
//////////////////////////////////

func (C *CPU) op_DEC_ZP(mem *mem.Memory) {
	C.opName = "DEC ZP"
	zpAddr := C.fetchByte(mem)
	val := mem.Read(zpAddr)
	val -= 1
	C.setNZStatus(val)
	mem.Write(zpAddr, val)
}

func (C *CPU) op_DEC_ZPX(mem *mem.Memory) {
	C.opName = "DEC ZP,X"
	zpAddr := C.fetchByte(mem) + C.X
	val := mem.Read(zpAddr)
	val -= 1
	C.setNZStatus(val)
	mem.Write(zpAddr, val)
}

func (C *CPU) op_DEC_ABS(mem *mem.Memory) {
	C.opName = "DEC Abs"
	address := C.fetchWord(mem)
	val := mem.Read(address)
	val -= 1
	C.setNZStatus(val)
	mem.Write(address, val)
}

func (C *CPU) op_DEC_ABX(mem *mem.Memory) {
	C.opName = "DEC Abs,X"
	absAddress := C.fetchWord(mem) + uint16(C.X)
	val := mem.Read(absAddress)
	val -= 1
	C.setNZStatus(val)
	mem.Write(absAddress, val)
}

func (C *CPU) op_DEX(mem *mem.Memory) {
	C.opName = "DEX"
	C.X -= 1
	C.setNZStatus(C.X)
}

func (C *CPU) op_DEY(mem *mem.Memory) {
	C.opName = "DEY"
	C.Y -= 1
	C.setNZStatus(C.Y)
}

//////////////////////////////////
///////////// INX ////////////////
//////////////////////////////////

// op_INX : Increment X
func (C *CPU) op_INX(mem *mem.Memory) {
	C.opName = "INX"
	C.X += 1
	C.setNZStatus(C.X)
}

//////////////////////////////////
///////////// INY ////////////////
//////////////////////////////////

func (C *CPU) op_INY(mem *mem.Memory) {
	C.opName = "INY"
	C.Y += 1
	C.setNZStatus(C.Y)
}
