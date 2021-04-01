package cpu

import (
	"go6502/globals"
	"go6502/mem"
)

//////////////////////////////////
///////////// INC ////////////////
//////////////////////////////////

func (C *CPU) op_INC_ZP(mem *mem.Memory) {
	C.opName = "INC ZP"
	zpAddr := C.fetchByte(mem)
	mem.Data[zpAddr] += 1
	C.setNZStatus(mem.Data[zpAddr])
}

func (C *CPU) op_INC_ZPX(mem *mem.Memory) {
	C.opName = "INC ZP,X"
	zpAddr := C.fetchByte(mem) + C.X
	mem.Data[zpAddr] += 1
	C.setNZStatus(mem.Data[zpAddr])
}

func (C *CPU) op_INC_ABS(mem *mem.Memory) {
	C.opName = "INC Abs"
	address := C.fetchWord(mem)
	C.dbus.WaitBusLow()
	val := mem.Data[address]
	C.dbus.WaitBusLow()
	val += 1
	C.dbus.WaitBusLow()
	mem.Data[address] = val
	C.setNZStatus(mem.Data[address])
}

func (C *CPU) op_INC_ABX(mem *mem.Memory) {
	C.opName = "INC Abs,X"
	absAddress := C.fetchWord(mem) + globals.Word(C.X)
	mem.Data[absAddress] += 1
	C.setNZStatus(mem.Data[absAddress])
}

//////////////////////////////////
///////////// DEC ////////////////
//////////////////////////////////

func (C *CPU) op_DEC_ZP(mem *mem.Memory) {
	C.opName = "DEC ZP"
	zpAddr := C.fetchByte(mem)
	mem.Data[zpAddr] -= 1
	C.setNZStatus(mem.Data[zpAddr])
}

func (C *CPU) op_DEC_ZPX(mem *mem.Memory) {
	C.opName = "DEC ZP,X"
	zpAddr := C.fetchByte(mem) + C.X
	mem.Data[zpAddr] -= 1
	C.setNZStatus(mem.Data[zpAddr])
}

func (C *CPU) op_DEC_ABS(mem *mem.Memory) {
	C.opName = "DEC Abs"
	address := C.fetchWord(mem)
	mem.Data[address] -= 1
	C.setNZStatus(mem.Data[address])
}

func (C *CPU) op_DEC_ABX(mem *mem.Memory) {
	C.opName = "DEC Abs,X"
	absAddress := C.fetchWord(mem) + globals.Word(C.X)
	mem.Data[absAddress] -= 1
	C.setNZStatus(mem.Data[absAddress])
}

func (C *CPU) op_DEX(mem *mem.Memory) {
	C.opName = "DEX"
	C.X -= 1
	C.setNZStatus(C.X)
}

func (C *CPU) op_DEY(mem *mem.Memory) {}

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
