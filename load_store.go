package main

//////////////////////////////////
///////////// LDA ////////////////
//////////////////////////////////

// op_LDA_IM : LDA Immediate
func (C *CPU) op_LDA_IM(mem *Memory) {
	C.opName = "LDA Imm"
	C.A = C.fetchByte(mem)
	C.setNZStatus(C.A)
}

// op_LDA_ZP : LDA Zero Page
func (C *CPU) op_LDA_ZP(mem *Memory) {
	C.opName = "LDA ZP"
	zpAddress := C.fetchByte(mem)
	C.A = mem.Data[zpAddress]
	C.setNZStatus(C.A)
}

// op_LDA_ZPX : LDA Zero Page,X
func (C *CPU) op_LDA_ZPX(mem *Memory) {
	C.opName = "LDA ZP,X"
	zpAddress := C.fetchByte(mem)
	zpAddress += C.X
	C.A = mem.Data[zpAddress]
	C.setNZStatus(C.A)
}

func (C *CPU) op_LDA_ABS(mem *Memory) {
	C.opName = "LDA Abs"
	absAddress := C.fetchWord(mem)
	C.A = mem.Data[absAddress]
	C.setNZStatus(C.A)
}

func (C *CPU) op_LDA_ABX(mem *Memory) {
	C.opName = "LDA Abs,X"
	absAddress := C.fetchWord(mem) + Word(C.X)
	C.A = mem.Data[absAddress]
	C.setNZStatus(C.A)
}

func (C *CPU) op_LDA_ABY(mem *Memory) {
	C.opName = "LDA Abs,Y"
	absAddress := C.fetchWord(mem) + Word(C.Y)
	C.A = mem.Data[absAddress]
	C.setNZStatus(C.A)
}

func (C *CPU) op_LDA_INX(mem *Memory) {
	C.opName = "LDA (ZP,X)"
	zpAddress := C.fetchByte(mem) + C.X
	C.A = mem.Data[zpAddress]
	C.setNZStatus(C.A)
}

func (C *CPU) op_LDA_INY(mem *Memory) {
	C.opName = "LDA (ZP),Y"
	zpAddress := C.fetchByte(mem)
	C.A = mem.Data[zpAddress] + C.Y
	C.setNZStatus(C.A)
}

//////////////////////////////////
///////////// LDX ////////////////
//////////////////////////////////

// op_LDX_IM : LDA Immediate
func (C *CPU) op_LDX_IM(mem *Memory) {
	C.opName = "LDX Imm"
	C.X = C.fetchByte(mem)
	C.setNZStatus(C.X)
}

// op_LDX_ZP : LDA Zero Page
func (C *CPU) op_LDX_ZP(mem *Memory) {
	C.opName = "LDX ZP"
	zpAddress := C.fetchByte(mem)
	C.X = mem.Data[zpAddress]
	C.setNZStatus(C.X)
}

// op_LDX_ZPY : LDA Zero Page,Y
func (C *CPU) op_LDX_ZPY(mem *Memory) {
	C.opName = "LDX ZP,Y"
	zpAddress := C.fetchByte(mem)
	zpAddress += C.Y
	C.X = mem.Data[zpAddress]
	C.setNZStatus(C.X)
}

func (C *CPU) op_LDX_ABS(mem *Memory) {
	C.opName = "LDA Abs"

	C.setNZStatus(C.A)
}

func (C *CPU) op_LDX_ABY(mem *Memory) {
	C.opName = "LDA Abs,X"

	C.setNZStatus(C.A)
}

//////////////////////////////////
///////////// LDY ////////////////
//////////////////////////////////

// op_LDY_IM : LDA Immediate
func (C *CPU) op_LDY_IM(mem *Memory) {
	C.opName = "LDY Imm"
	C.Y = C.fetchByte(mem)
	C.setNZStatus(C.Y)
}

// op_LDY_ZP : LDA Zero Page
func (C *CPU) op_LDY_ZP(mem *Memory) {
	C.opName = "LDY ZP"
	zpAddress := C.fetchByte(mem)
	C.Y = mem.Data[zpAddress]
	C.setNZStatus(C.Y)
}

// op_LDY_ZPX : LDA Zero Page,X
func (C *CPU) op_LDY_ZPX(mem *Memory) {
	C.opName = "LDY ZP,X"
	zpAddress := C.fetchByte(mem)
	zpAddress += C.X
	C.Y = mem.Data[zpAddress]
	C.setNZStatus(C.Y)
}

func (C *CPU) op_LDY_ABS(mem *Memory) {
	C.opName = "LDA Abs"

	C.setNZStatus(C.A)
}

func (C *CPU) op_LDY_ABX(mem *Memory) {
	C.opName = "LDA Abs,X"

	C.setNZStatus(C.A)
}

//////////////////////////////////
///////////// STA ////////////////
//////////////////////////////////

func (C *CPU) op_STA_ZP(mem *Memory) {
	C.opName = "STA ZP"
	zpAddress := C.fetchByte(mem)
	mem.Data[zpAddress] = C.A
}

func (C *CPU) op_STA_ZPX(mem *Memory) {
	C.opName = "STA ZP,X"
	zpAddress := C.fetchByte(mem) + C.X
	mem.Data[zpAddress] = C.A
}

func (C *CPU) op_STA_ABS(mem *Memory) {
	C.opName = "STA Abs"
	absAddress := C.fetchWord(mem)
	mem.Data[absAddress] = C.A
}

func (C *CPU) op_STA_ABX(mem *Memory) {}
func (C *CPU) op_STA_ABY(mem *Memory) {}
func (C *CPU) op_STA_INX(mem *Memory) {}
func (C *CPU) op_STA_INY(mem *Memory) {}

//////////////////////////////////
///////////// STX ////////////////
//////////////////////////////////

func (C *CPU) op_STX_ZP(mem *Memory) {
	C.opName = "STX ZP"
	zpAddress := C.fetchByte(mem)
	mem.Data[zpAddress] = C.X
}

func (C *CPU) op_STX_ZPY(mem *Memory) {
	C.opName = "STX ZP,Y"
	zpAddress := C.fetchByte(mem) + C.Y
	mem.Data[zpAddress] = C.X
}

func (C *CPU) op_STX_ABS(mem *Memory) {}

//////////////////////////////////
///////////// STY ////////////////
//////////////////////////////////

func (C *CPU) op_STY_ZP(mem *Memory)  {}
func (C *CPU) op_STY_ZPX(mem *Memory) {}
func (C *CPU) op_STY_ABS(mem *Memory) {}
