package cpu

import (
	"fmt"
	"go6502/mem"
)

//////////////////////////////////
///////////// LDA ////////////////
//////////////////////////////////

// op_LDA_IM : LDA Immediate
func (C *CPU) op_LDA_IM(mem *mem.Memory) {
	C.A = C.fetchByte(mem)
	C.opName = fmt.Sprintf("LDA #$%02X", C.A)
	C.setNZStatus(C.A)
}

// op_LDA_ZP : LDA Zero Page
func (C *CPU) op_LDA_ZP(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	C.A = mem.Read(zpAddress)
	C.opName = fmt.Sprintf("LDA $%02X", zpAddress)
	C.setNZStatus(C.A)
}

// op_LDA_ZPX : LDA Zero Page,X
func (C *CPU) op_LDA_ZPX(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	dest := zpAddress + C.X
	C.A = mem.Read(dest)
	C.opName = fmt.Sprintf("LDA $%02X,X", zpAddress)
	C.setNZStatus(C.A)
}

func (C *CPU) op_LDA_ABS(mem *mem.Memory) {
	C.opName = "LDA Abs"
	absAddress := C.fetchWord(mem)
	C.A = mem.Read(absAddress)
	C.opName = fmt.Sprintf("LDA $%04X", absAddress)
	C.setNZStatus(C.A)
}

func (C *CPU) op_LDA_ABX(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	C.A = mem.Read(absAddress + uint16(C.X))
	C.opName = fmt.Sprintf("LDA $%04X,X", absAddress)
	C.setNZStatus(C.A)
}

func (C *CPU) op_LDA_ABY(mem *mem.Memory) {
	C.opName = "LDA Abs,Y"
	absAddress := C.fetchWord(mem) + uint16(C.Y)
	C.A = mem.Read(absAddress)
	C.setNZStatus(C.A)
}

func (C *CPU) op_LDA_INX(mem *mem.Memory) {
	zpAddr := C.fetchByte(mem)
	C.opName = fmt.Sprintf("LDA ($%02X,X)", zpAddr)
	wordZP := C.Indexed_indirect_X(zpAddr, C.X)
	C.A = mem.Read(wordZP)
	C.setNZStatus(C.A)
}

func (C *CPU) op_LDA_INY(mem *mem.Memory) {
	zpAddr := C.fetchByte(mem)
	C.opName = fmt.Sprintf("LDA ($%02X),Y", zpAddr)
	wordZP := C.Indirect_index_Y(zpAddr, C.Y)
	C.A = mem.Read(wordZP)
	C.setNZStatus(C.A)
}

//////////////////////////////////
///////////// LDX ////////////////
//////////////////////////////////

// op_LDX_IM : LDA Immediate
func (C *CPU) op_LDX_IM(mem *mem.Memory) {
	C.X = C.fetchByte(mem)
	C.opName = fmt.Sprintf("LDX #$%02X", C.X)
	C.setNZStatus(C.X)
}

// op_LDX_ZP : LDA Zero Page
func (C *CPU) op_LDX_ZP(mem *mem.Memory) {
	C.opName = "LDX ZP"
	zpAddress := C.fetchByte(mem)
	C.X = mem.Read(zpAddress)
	C.setNZStatus(C.X)
}

// op_LDX_ZPY : LDA Zero Page,Y
func (C *CPU) op_LDX_ZPY(mem *mem.Memory) {
	C.opName = "LDX ZP,Y"
	zpAddress := C.fetchByte(mem) + C.Y
	C.X = mem.Read(zpAddress)
	C.setNZStatus(C.X)
}

func (C *CPU) op_LDX_ABS(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	C.X = mem.Read(absAddress)
	C.opName = fmt.Sprintf("LDX $%04X", absAddress)
	C.setNZStatus(C.X)
}

func (C *CPU) op_LDX_ABY(mem *mem.Memory) {
	C.opName = "LDX Abs,X"
	C.opName = "ToDO"
	C.setNZStatus(C.A)
}

//////////////////////////////////
///////////// LDY ////////////////
//////////////////////////////////

// op_LDY_IM : LDA Immediate
func (C *CPU) op_LDY_IM(mem *mem.Memory) {
	C.Y = C.fetchByte(mem)
	C.opName = fmt.Sprintf("LDY #$%02X", C.Y)
	C.setNZStatus(C.Y)
}

// op_LDY_ZP : LDA Zero Page
func (C *CPU) op_LDY_ZP(mem *mem.Memory) {
	C.opName = "LDY ZP"
	zpAddress := C.fetchByte(mem)
	C.Y = mem.Read(zpAddress)
	C.setNZStatus(C.Y)
}

// op_LDY_ZPX : LDA Zero Page,X
func (C *CPU) op_LDY_ZPX(mem *mem.Memory) {
	C.opName = "LDY ZP,X"
	zpAddress := C.fetchByte(mem) + C.X
	C.Y = mem.Read(zpAddress)
	C.setNZStatus(C.Y)
}

func (C *CPU) op_LDY_ABS(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	C.Y = mem.Read(absAddress)
	C.opName = fmt.Sprintf("LDY $%04X", absAddress)
	C.setNZStatus(C.Y)
}

func (C *CPU) op_LDY_ABX(mem *mem.Memory) {
	C.opName = "LDY Abs,X"
	C.opName = "ToDO"
	C.setNZStatus(C.A)
}

//////////////////////////////////
///////////// STA ////////////////
//////////////////////////////////

func (C *CPU) op_STA_ZP(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	C.opName = fmt.Sprintf("STA $%02X", zpAddress)
	mem.Write(zpAddress, C.A)
	val := mem.Read(zpAddress)
	C.debug = fmt.Sprintf("%02X -> $%02X", val, zpAddress)
}

func (C *CPU) op_STA_ZPX(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	dest := zpAddress + C.X
	C.opName = fmt.Sprintf("STA $%02X,X", zpAddress)
	mem.Write(dest, C.A)
}

func (C *CPU) op_STA_ABS(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	C.opName = fmt.Sprintf("STA $%04X", absAddress)
	if !mem.Write(absAddress, C.A) {
		C.debug = "Write to ROM"
	}
}

func (C *CPU) op_STA_ABX(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	dest := absAddress + uint16(C.X)
	C.dbus.Release()
	C.opName = fmt.Sprintf("STA $%04X,X", absAddress)
	C.debug = fmt.Sprintf("%02X -> $%04X", C.A, dest)
	mem.Write(dest, C.A)
	C.dbus.Release()
}

func (C *CPU) op_STA_ABY(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	dest := absAddress + uint16(C.Y)
	C.opName = fmt.Sprintf("STA $%04X,Y", absAddress)
	C.debug = fmt.Sprintf("%02X -> $%04X", C.A, dest)
	C.dbus.Release()
	mem.Write(dest, C.A)
	C.dbus.Release()
}

func (C *CPU) op_STA_INX(mem *mem.Memory) {
	zpAddr := C.fetchByte(mem)
	C.opName = fmt.Sprintf("STA ($%02X,X)", zpAddr)
	wordZP := C.Indexed_indirect_X(zpAddr, C.X)
	C.debug = fmt.Sprintf("%02X -> $%04X", C.A, wordZP)
	mem.Write(wordZP, C.A)
}

func (C *CPU) op_STA_INY(mem *mem.Memory) {
	zpAddr := C.fetchByte(mem)
	wordZP := C.Indirect_index_Y(zpAddr, C.Y)
	C.opName = fmt.Sprintf("STA ($%02X),Y", zpAddr)
	// mem.Dump(uint16(zpAddr))
	mem.Write(wordZP, C.A)
	C.debug = fmt.Sprintf("%02X -> $%04X", C.A, wordZP)
}

//////////////////////////////////
///////////// STX ////////////////
//////////////////////////////////

func (C *CPU) op_STX_ZP(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	C.opName = fmt.Sprintf("STX $%02X", zpAddress)
	mem.Write(zpAddress, C.X)
	C.dbus.Release()
}

func (C *CPU) op_STX_ZPY(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem) + C.Y
	C.opName = fmt.Sprintf("STX $%02X,Y", zpAddress)
	mem.Write(zpAddress, C.X)
}

func (C *CPU) op_STX_ABS(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	C.opName = fmt.Sprintf("STX $%04X", absAddress)
	mem.Write(absAddress, C.X)
}

//////////////////////////////////
///////////// STY ////////////////
//////////////////////////////////

func (C *CPU) op_STY_ZP(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	C.opName = fmt.Sprintf("STY $%02X", zpAddress)
	mem.Write(zpAddress, C.Y)
	C.dbus.Release()
}

func (C *CPU) op_STY_ZPX(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	dest := zpAddress + C.X
	C.opName = fmt.Sprintf("STY $%02X,X", zpAddress)
	mem.Write(dest, C.Y)
}

func (C *CPU) op_STY_ABS(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	C.opName = fmt.Sprintf("STY $%04X", absAddress)
	mem.Write(absAddress, C.Y)
}
