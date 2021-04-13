package cpu

import (
	"go6502/mem"
)

func (C *CPU) op_ASL_IM(mem *mem.Memory)  { C.opName = "ToDO" }
func (C *CPU) op_ASL_ZP(mem *mem.Memory)  { C.opName = "ToDO" }
func (C *CPU) op_ASL_ZPX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ASL_ABS(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ASL_ABX(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_LSR_IM(mem *mem.Memory)  { C.opName = "ToDO" }
func (C *CPU) op_LSR_ZP(mem *mem.Memory)  { C.opName = "ToDO" }
func (C *CPU) op_LSR_ZPX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_LSR_ABS(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_LSR_ABX(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_ROL_IM(mem *mem.Memory) {
	C.opName = "ROL"
	result := uint16(C.A << 1)
	C.setC(result > 0x0FF)
	C.A = byte(result)
	C.dbus.Release()
	C.setNZStatus(C.Y)
}

func (C *CPU) op_ROL_ZP(mem *mem.Memory)  { C.opName = "ToDO" }
func (C *CPU) op_ROL_ZPX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ROL_ABS(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ROL_ABX(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_ROR_IM(mem *mem.Memory)  { C.opName = "ToDO" }
func (C *CPU) op_ROR_ZP(mem *mem.Memory)  { C.opName = "ToDO" }
func (C *CPU) op_ROR_ZPX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ROR_ABS(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ROR_ABX(mem *mem.Memory) { C.opName = "ToDO" }
