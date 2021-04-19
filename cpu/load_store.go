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
	C.A = C.readByte(uint16(zpAddress))
	C.opName = fmt.Sprintf("LDA $%02X", zpAddress)
	C.setNZStatus(C.A)
}

// op_LDA_ZPX : LDA Zero Page,X
func (C *CPU) op_LDA_ZPX(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	dest := zpAddress + C.X
	C.A = C.readByte(uint16(dest))
	C.opName = fmt.Sprintf("LDA $%02X,X", zpAddress)
	C.debug = fmt.Sprintf("-> $%02X", dest)
	C.setNZStatus(C.A)
}

func (C *CPU) op_LDA_ABS(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	C.A = C.readByte(absAddress)
	C.opName = fmt.Sprintf("LDA $%04X", absAddress)
	C.setNZStatus(C.A)
}

func (C *CPU) op_LDA_ABX(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	C.A = C.readByte(absAddress + uint16(C.X))
	C.opName = fmt.Sprintf("LDA $%04X,X", absAddress)
	C.setNZStatus(C.A)
}

func (C *CPU) op_LDA_ABY(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	dest := absAddress + uint16(C.Y)
	C.opName = fmt.Sprintf("LDA $%04X,Y", absAddress)
	C.debug = fmt.Sprintf("-> $%04X", dest)
	C.A = C.readByte(dest)
	C.setNZStatus(C.A)
}

func (C *CPU) op_LDA_INX(mem *mem.Memory) {
	zpAddr := C.fetchByte(mem)
	C.opName = fmt.Sprintf("LDA ($%02X,X)", zpAddr)
	wordZP := C.Indexed_indirect_X(zpAddr, C.X)
	C.A = C.readByte(wordZP)
	C.setNZStatus(C.A)
}

func (C *CPU) op_LDA_INY(mem *mem.Memory) {
	zpAddr := C.fetchByte(mem)
	C.opName = fmt.Sprintf("LDA ($%02X),Y", zpAddr)
	wordZP := C.Indirect_index_Y(zpAddr, C.Y)
	C.A = C.readByte(wordZP)
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
	zpAddress := C.fetchByte(mem)
	C.opName = fmt.Sprintf("LDX $%02X", zpAddress)
	C.X = C.readByte(uint16(zpAddress))
	C.setNZStatus(C.X)
}

// op_LDX_ZPY : LDA Zero Page,Y
func (C *CPU) op_LDX_ZPY(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	dest := zpAddress + C.Y
	C.opName = fmt.Sprintf("LDX $%02X,Y", zpAddress)
	C.debug = fmt.Sprintf("-> $%02X", dest)
	C.X = C.readByte(uint16(dest))
	C.setNZStatus(C.X)
}

func (C *CPU) op_LDX_ABS(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	C.X = C.readByte(absAddress)
	C.opName = fmt.Sprintf("LDX $%04X", absAddress)
	C.setNZStatus(C.X)
}

func (C *CPU) op_LDX_ABY(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	dest := absAddress + uint16(C.Y)
	C.opName = fmt.Sprintf("LDX $%04X,Y", absAddress)
	C.debug = fmt.Sprintf("-> $%04X", dest)
	C.X = C.readByte(dest)
	C.setNZStatus(C.X)
}

//////////////////////////////////
///////////// LDY ////////////////
//////////////////////////////////

func (C *CPU) op_LDY_IM(mem *mem.Memory) {
	C.Y = C.fetchByte(mem)
	C.opName = fmt.Sprintf("LDY #$%02X", C.Y)
	C.setNZStatus(C.Y)
}

func (C *CPU) op_LDY_ZP(mem *mem.Memory) {
	C.opName = "LDY ZP"
	zpAddress := C.fetchByte(mem)
	C.Y = C.readByte(uint16(zpAddress))
	C.setNZStatus(C.Y)
}

func (C *CPU) op_LDY_ZPX(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	dest := zpAddress + C.X
	C.Y = C.readByte(uint16(dest))
	C.opName = fmt.Sprintf("LDY $%02X,X", zpAddress)
	C.debug = fmt.Sprintf("-> $%02X", dest)
	C.setNZStatus(C.Y)
}

func (C *CPU) op_LDY_ABS(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	C.Y = C.readByte(absAddress)
	C.opName = fmt.Sprintf("LDY $%04X", absAddress)
	C.setNZStatus(C.Y)
}

func (C *CPU) op_LDY_ABX(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	dest := absAddress + uint16(C.X)
	C.opName = fmt.Sprintf("LDY $%04X,X", absAddress)
	C.Y = C.readByte(dest)
	C.setNZStatus(C.Y)
}

//////////////////////////////////
///////////// STA ////////////////
//////////////////////////////////

func (C *CPU) op_STA_ZP(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	C.opName = fmt.Sprintf("STA $%02X", zpAddress)
	C.writeByte(uint16(zpAddress), C.A)
	val := C.readByte(uint16(zpAddress))
	C.debug = fmt.Sprintf("%02X -> $%02X", val, zpAddress)
}

func (C *CPU) op_STA_ZPX(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	dest := zpAddress + C.X
	C.opName = fmt.Sprintf("STA $%02X,X", zpAddress)
	C.debug = fmt.Sprintf("-> $%02X", dest)
	C.writeByte(uint16(dest), C.A)
}

func (C *CPU) op_STA_ABS(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	C.opName = fmt.Sprintf("STA $%04X", absAddress)
	C.writeByte(absAddress, C.A)
}

func (C *CPU) op_STA_ABX(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	dest := absAddress + uint16(C.X)
	C.dbus.Release()
	C.opName = fmt.Sprintf("STA $%04X,X", absAddress)
	C.debug = fmt.Sprintf("%02X -> $%04X", C.A, dest)
	C.writeByte(dest, C.A)
}

func (C *CPU) op_STA_ABY(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	dest := absAddress + uint16(C.Y)
	C.dbus.Release()
	C.opName = fmt.Sprintf("STA $%04X,Y", absAddress)
	C.debug = fmt.Sprintf("%02X -> $%04X", C.A, dest)
	C.writeByte(dest, C.A)
}

func (C *CPU) op_STA_INX(mem *mem.Memory) {
	zpAddr := C.fetchByte(mem)
	C.opName = fmt.Sprintf("STA ($%02X,X)", zpAddr)
	wordZP := C.Indexed_indirect_X(zpAddr, C.X)
	C.debug = fmt.Sprintf("%02X -> $%04X", C.A, wordZP)
	C.dbus.Release()
	C.writeByte(wordZP, C.A)
}

func (C *CPU) op_STA_INY(mem *mem.Memory) {
	zpAddr := C.fetchByte(mem)
	wordZP := C.Indirect_index_Y(zpAddr, C.Y)
	C.opName = fmt.Sprintf("STA ($%02X),Y", zpAddr)
	C.debug = fmt.Sprintf("%02X -> $%04X", C.A, wordZP)
	C.dbus.Release()
	C.writeByte(wordZP, C.A)
}

//////////////////////////////////
///////////// STX ////////////////
//////////////////////////////////

func (C *CPU) op_STX_ZP(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	C.opName = fmt.Sprintf("STX $%02X", zpAddress)
	C.writeByte(uint16(zpAddress), C.X)
}

func (C *CPU) op_STX_ZPY(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem) + C.Y
	C.opName = fmt.Sprintf("STX $%02X,Y", zpAddress)
	C.writeByte(uint16(zpAddress), C.X)
}

func (C *CPU) op_STX_ABS(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	C.opName = fmt.Sprintf("STX $%04X", absAddress)
	C.writeByte(absAddress, C.X)
}

//////////////////////////////////
///////////// STY ////////////////
//////////////////////////////////

func (C *CPU) op_STY_ZP(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	C.opName = fmt.Sprintf("STY $%02X", zpAddress)
	C.writeByte(uint16(zpAddress), C.Y)
	C.dbus.Release()
}

func (C *CPU) op_STY_ZPX(mem *mem.Memory) {
	zpAddress := C.fetchByte(mem)
	dest := zpAddress + C.X
	C.opName = fmt.Sprintf("STY $%02X,X", zpAddress)
	C.writeByte(uint16(dest), C.Y)
}

func (C *CPU) op_STY_ABS(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	C.opName = fmt.Sprintf("STY $%04X", absAddress)
	C.writeByte(absAddress, C.Y)
}
