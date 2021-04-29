package cpu

import (
	"fmt"
	"go6502/mem"
)

func (C *CPU) op_ASL_IM(mem *mem.Memory) {
	result := uint16(C.A) << 1
	C.setC(result > 0x00FF)
	C.A <<= 1
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = "ASL"
	}
}

func (C *CPU) op_ASL_ZP(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	value := C.readByte(uint16(zpAddress))
	result := uint16(value) << 1
	C.setC(result > 0x00FF)
	res8 := byte(result)
	C.setNZStatus(res8)
	C.writeByte(uint16(zpAddress), res8)

	if C.Display {
		C.opName = fmt.Sprintf("ASL $%02X", zpAddress)
	}
}

func (C *CPU) op_ASL_ZPX(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	dest := zpAddress + C.X
	value := C.readByte(uint16(dest))
	result := uint16(value) << 1
	C.setC(result > 0x00FF)
	res8 := byte(result)
	C.setNZStatus(res8)
	C.writeByte(uint16(dest), res8)

	if C.Display {
		C.opName = fmt.Sprintf("ASL $%02X,X", zpAddress)
	}
}

func (C *CPU) op_ASL_ABS(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	value := C.readByte(uint16(absAddress))
	result := uint16(value) << 1
	C.setC(result > 0x00FF)
	res8 := byte(result)
	C.setNZStatus(res8)
	C.writeByte(uint16(absAddress), res8)

	if C.Display {
		C.opName = fmt.Sprintf("ASL $%04X", absAddress)
	}
}

func (C *CPU) op_ASL_ABX(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_LSR_IM(mem *mem.Memory) {
	C.setC(C.A&0x01 == 0x01)
	C.A >>= 1
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = "LSR"
	}
}

func (C *CPU) op_LSR_ZP(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	value := C.readByte(uint16(zpAddress))
	C.setC(value&0x01 == 0x01)
	result := value >> 1
	C.setNZStatus(result)
	C.writeByte(uint16(zpAddress), result)

	if C.Display {
		C.opName = fmt.Sprintf("LSR $%02X", zpAddress)
	}
}

func (C *CPU) op_LSR_ZPX(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	dest := zpAddress + C.X
	value := C.readByte(uint16(dest))
	C.setC(value&0x01 == 0x01)
	result := value >> 1
	C.setNZStatus(result)
	C.writeByte(uint16(dest), result)

	if C.Display {
		C.opName = fmt.Sprintf("LSR $%02X,X", zpAddress)
		C.debug = fmt.Sprintf("-> $%02X", dest)
	}
}

func (C *CPU) op_LSR_ABS(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	value := C.readByte(uint16(absAddress))
	C.setC(value&0x01 == 0x01)
	result := value >> 1
	C.setNZStatus(result)
	C.writeByte(uint16(absAddress), result)

	if C.Display {
		C.opName = fmt.Sprintf("LSR $%04X", absAddress)
	}
}

func (C *CPU) op_LSR_ABX(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_ROL_IM(mem *mem.Memory) {
	result := uint16(C.A) << 1
	if C.testC() {
		result++
	}
	C.setC(result > 0x00FF)
	C.A = byte(result)
	C.setNZStatus(C.A)
	C.dbus.Release()

	if C.Display {
		C.opName = "ROL"
	}
}

func (C *CPU) op_ROL_ZP(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	value := C.readByte(uint16(zpAddress))
	result := uint16(value) << 1
	C.setC(result > 0x00FF)
	res8 := byte(result)
	C.setNZStatus(res8)
	C.writeByte(uint16(zpAddress), res8)

	if C.Display {
		C.opName = fmt.Sprintf("ROL $%02X", zpAddress)
	}
}

func (C *CPU) op_ROL_ZPX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ROL_ABS(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ROL_ABX(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_ROR_IM(mem *mem.Memory) {
	carry := C.A&0b00000001 > 0
	C.A >>= 1
	if C.testC() {
		C.A |= 0b10000000
	}
	C.setC(carry)
	C.setNZStatus(C.A)
	C.dbus.Release()

	if C.Display {
		C.opName = "ROR"
	}
}

func (C *CPU) op_ROR_ZP(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	value := C.readByte(uint16(zpAddress))
	carry := value&0b00000001 > 0
	value >>= 1
	if C.testC() {
		value |= 0b10000000
	}
	C.setC(carry)
	C.setNZStatus(value)
	C.writeByte(uint16(zpAddress), value)
	C.dbus.Release()

	if C.Display {
		C.opName = fmt.Sprintf("ROR $%02X", zpAddress)
	}
}

func (C *CPU) op_ROR_ZPX(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	dest := zpAddress + C.X
	value := C.readByte(uint16(dest))
	carry := value&0b00000001 > 0
	C.dbus.Release()
	value >>= 1
	if C.testC() {
		value |= 0b10000000
	}
	C.setC(carry)
	C.setNZStatus(value)
	C.writeByte(uint16(dest), value)
	C.dbus.Release()

	if C.Display {
		C.opName = fmt.Sprintf("ROR $%02X,X", zpAddress)
		C.debug = fmt.Sprintf("#%02X -> $%02X", value, dest)
	}
}

func (C *CPU) op_ROR_ABS(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ROR_ABX(mem *mem.Memory) { C.opName = "ToDO" }
