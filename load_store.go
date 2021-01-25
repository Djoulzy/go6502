package main

func (C *CPU) setLDAStatus(register Byte) {
	if register == 0 {
		C.S |= ^Z_mask
	} else {
		C.S &= Z_mask
	}
	if register&0b10000000 > 0 {
		C.S |= ^N_mask
	} else {
		C.S &= N_mask
	}
}

// op_LDA_IM : LDA Immediate
func (C *CPU) op_LDA_IM(mem *Memory) {
	C.opName = "LDA Imm "
	C.A = C.fetchByte(mem)
	C.setLDAStatus(C.A)
}

// op_LDA_ZP : LDA Zero Page
func (C *CPU) op_LDA_ZP(mem *Memory) {
	C.opName = "LDA ZP  "
	zpAddress := C.fetchByte(mem)
	C.A = mem.Data[zpAddress]
	C.setLDAStatus(C.A)
}

// op_LDA_ZPX : LDA Zero Page,X
func (C *CPU) op_LDA_ZPX(mem *Memory) {
	C.opName = "LDA ZP,X"
	zpAddress := C.fetchByte(mem)
	zpAddress += C.X
	C.A = mem.Data[zpAddress]
	C.setLDAStatus(C.A)
}

// op_LDX_IM : LDA Immediate
func (C *CPU) op_LDX_IM(mem *Memory) {
	C.opName = "LDX Imm "
	C.X = C.fetchByte(mem)
	C.setLDAStatus(C.X)
}

// op_LDX_ZP : LDA Zero Page
func (C *CPU) op_LDX_ZP(mem *Memory) {
	C.opName = "LDX ZP  "
	zpAddress := C.fetchByte(mem)
	C.X = mem.Data[zpAddress]
	C.setLDAStatus(C.X)
}

// op_LDX_ZPY : LDA Zero Page,Y
func (C *CPU) op_LDX_ZPY(mem *Memory) {
	C.opName = "LDX ZP,Y"
	zpAddress := C.fetchByte(mem)
	zpAddress += C.Y
	C.X = mem.Data[zpAddress]
	C.setLDAStatus(C.X)
}

// op_LDY_IM : LDA Immediate
func (C *CPU) op_LDY_IM(mem *Memory) {
	C.opName = "LDY Imm "
	C.Y = C.fetchByte(mem)
	C.setLDAStatus(C.Y)
}

// op_LDY_ZP : LDA Zero Page
func (C *CPU) op_LDY_ZP(mem *Memory) {
	C.opName = "LDY ZP  "
	zpAddress := C.fetchByte(mem)
	C.Y = mem.Data[zpAddress]
	C.setLDAStatus(C.Y)
}

// op_LDY_ZPX : LDA Zero Page,X
func (C *CPU) op_LDY_ZPX(mem *Memory) {
	C.opName = "LDY ZP,X"
	zpAddress := C.fetchByte(mem)
	zpAddress += C.X
	C.Y = mem.Data[zpAddress]
	C.setLDAStatus(C.Y)
}
