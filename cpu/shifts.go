package cpu

import (
	"fmt"
	"go6502/mem"
)

func (C *CPU) op_ASL_IM(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_ASL_ZP(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	C.opName = fmt.Sprintf("ASL $%02X", zpAddress)
	value := C.readByte(uint16(zpAddress))
	result := uint16(value << 1)
	C.setC(result > 0x0FF)
	res8 := byte(result)
	C.setNZStatus(res8)
	C.writeByte(uint16(zpAddress), res8)
}

func (C *CPU) op_ASL_ZPX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ASL_ABS(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ASL_ABX(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_LSR_IM(mem *mem.Memory)  { C.opName = "ToDO" }

func (C *CPU) op_LSR_ZP(mem *mem.Memory) {
	 zpAddress := C.fetchByte(mem)
	 C.opName = fmt.Sprintf("LSR $%02X", zpAddress)
	 value := C.readByte(uint16(zpAddress))
	 C.setC(value & 0x01 == 0x01)
	 result := value >> 1
	 C.setNZStatus(result)
	 C.writeByte(uint16(zpAddress), result)
}

func (C *CPU) op_LSR_ZPX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_LSR_ABS(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_LSR_ABX(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_ROL_IM(mem *mem.Memory) {
	C.opName = "ROL"
	result := uint16(C.A << 1)
	C.setC(result > 0x0FF)
	C.A = byte(result)
	C.setNZStatus(C.Y)
	C.dbus.Release()
}

func (C *CPU) op_ROL_ZP(mem *mem.Memory)  { 
	zpAddress := C.fetchByte(mem)
	C.opName = fmt.Sprintf("ROL $%02X", zpAddress)
	value := C.readByte(uint16(zpAddress))
	result := uint16(value << 1)
	C.setC(result > 0x0FF)
	res8 := byte(result)
	C.setNZStatus(res8)
	C.writeByte(uint16(zpAddress), res8)
}

func (C *CPU) op_ROL_ZPX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ROL_ABS(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ROL_ABX(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_ROR_IM(mem *mem.Memory)  { C.opName = "ToDO" }
func (C *CPU) op_ROR_ZP(mem *mem.Memory)  { C.opName = "ToDO" }
func (C *CPU) op_ROR_ZPX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ROR_ABS(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ROR_ABX(mem *mem.Memory) { C.opName = "ToDO" }
