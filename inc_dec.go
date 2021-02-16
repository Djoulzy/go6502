package main

//////////////////////////////////
///////////// INC ////////////////
//////////////////////////////////

func (C *CPU) op_INC_ZP(mem *Memory)  {}
func (C *CPU) op_INC_ZPX(mem *Memory) {}
func (C *CPU) op_INC_ABS(mem *Memory) {}
func (C *CPU) op_INC_ABX(mem *Memory) {}

//////////////////////////////////
///////////// INX ////////////////
//////////////////////////////////

// op_INX : Increment X
func (C *CPU) op_INX(mem *Memory) {
	C.opName = "INX"
	C.X += 1
	C.setNZStatus(C.X)
}

//////////////////////////////////
///////////// INY ////////////////
//////////////////////////////////

func (C *CPU) op_INY(mem *Memory) {}

//////////////////////////////////
///////////// DEC ////////////////
//////////////////////////////////

func (C *CPU) op_DEC_ZP(mem *Memory)  {}
func (C *CPU) op_DEC_ZPX(mem *Memory) {}
func (C *CPU) op_DEC_ABS(mem *Memory) {}
func (C *CPU) op_DEC_ABX(mem *Memory) {}
func (C *CPU) op_DEX(mem *Memory)     {}
func (C *CPU) op_DEY(mem *Memory)     {}
