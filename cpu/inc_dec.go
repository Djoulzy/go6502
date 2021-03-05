package cpu

import "go6502/mem"

//////////////////////////////////
///////////// INC ////////////////
//////////////////////////////////

func (C *CPU) op_INC_ZP(mem *mem.Memory)  {}
func (C *CPU) op_INC_ZPX(mem *mem.Memory) {}
func (C *CPU) op_INC_ABS(mem *mem.Memory) {}
func (C *CPU) op_INC_ABX(mem *mem.Memory) {}

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

func (C *CPU) op_INY(mem *mem.Memory) {}

//////////////////////////////////
///////////// DEC ////////////////
//////////////////////////////////

func (C *CPU) op_DEC_ZP(mem *mem.Memory)  {}
func (C *CPU) op_DEC_ZPX(mem *mem.Memory) {}
func (C *CPU) op_DEC_ABS(mem *mem.Memory) {}
func (C *CPU) op_DEC_ABX(mem *mem.Memory) {}

func (C *CPU) op_DEX(mem *mem.Memory)     {
	C.opName = "DEX"
	C.X -= 1
	C.setNZStatus(C.X)
}

func (C *CPU) op_DEY(mem *mem.Memory)     {}
