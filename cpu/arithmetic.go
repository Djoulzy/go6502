package cpu

import (
	"fmt"
	"go6502/mem"
)

func (C *CPU) op_ADC_IM(mem *mem.Memory) {
	C.opName = "ADC Im"
	value := C.fetchByte(mem)
	result := uint16(C.A) + uint16(value) + uint16(C.getC())
	C.setC(result > 0x0FF)
	C.setV(C.A, value, byte(result))
	C.A = byte(result)
	C.setNZStatus(C.A)
}

func (C *CPU) op_ADC_ZP(mem *mem.Memory) {
	C.opName = "ADC ZP"
	zpAddress := C.fetchByte(mem)
	value := mem.Read(zpAddress)
	result := uint16(C.A) + uint16(value) + uint16(C.getC())
	C.setC(result > 0x0FF)
	C.setV(C.A, value, byte(result))
	C.A = byte(result)
	C.setNZStatus(C.A)
}

func (C *CPU) op_ADC_ZPX(mem *mem.Memory) {
	C.opName = "ADC ZP,X"
	zpAddress := C.fetchByte(mem) + C.X
	value := mem.Read(zpAddress)
	result := uint16(C.A) + uint16(value) + uint16(C.getC())
	C.setC(result > 0x0FF)
	C.setV(C.A, value, byte(result))
	C.A = byte(result)
	C.setNZStatus(C.A)
}

func (C *CPU) op_ADC_ABS(mem *mem.Memory) {
	C.opName = "ADC Abs"
	absAddress := C.fetchWord(mem)
	value := mem.Read(absAddress)
	result := uint16(C.A) + uint16(value) + uint16(C.getC())
	C.setC(result > 0x0FF)
	C.setV(C.A, value, byte(result))
	C.A = byte(result)
	C.setNZStatus(C.A)
}

func (C *CPU) op_ADC_ABX(mem *mem.Memory) {
	C.opName = "ADC Abs,X"
	absAddress := C.fetchWord(mem) + uint16(C.X)
	value := mem.Read(absAddress)
	result := uint16(C.A) + uint16(value) + uint16(C.getC())
	C.setC(result > 0x0FF)
	C.setV(C.A, value, byte(result))
	C.A = byte(result)
	C.setNZStatus(C.A)
}

func (C *CPU) op_ADC_ABY(mem *mem.Memory) {
	C.opName = "ADC Abs,Y"
	absAddress := C.fetchWord(mem) + uint16(C.Y)
	value := mem.Read(absAddress)
	result := uint16(C.A) + uint16(value) + uint16(C.getC())
	C.setC(result > 0x0FF)
	C.setV(C.A, value, byte(result))
	C.A = byte(result)
	C.setNZStatus(C.A)
}

func (C *CPU) op_ADC_INX(mem *mem.Memory) {
	C.opName = "ADC (ZP,X)"
	zpAddr := C.fetchByte(mem)
	wordZP := C.Indexed_indirect_X(zpAddr, C.X)
	value := mem.Read(wordZP)
	result := uint16(C.A) + uint16(value) + uint16(C.getC())
	C.setC(result > 0x0FF)
	C.setV(C.A, value, byte(result))
	C.A = byte(result)
	C.setNZStatus(C.A)
}

func (C *CPU) op_ADC_INY(mem *mem.Memory) {
	C.opName = "ADC (ZP),Y"
	zpAddr := C.fetchByte(mem)
	wordZP := C.Indirect_index_Y(zpAddr, C.Y)
	value := mem.Read(wordZP)
	result := uint16(C.A) + uint16(value) + uint16(C.getC())
	C.setC(result > 0x0FF)
	C.setV(C.A, value, byte(result))
	C.A = byte(result)
	C.setNZStatus(C.A)
}

func (C *CPU) op_SBC_IM(mem *mem.Memory) {
	C.opName = "SBC Im"
	value := ^C.fetchByte(mem)
	result := uint16(C.A) + uint16(value) + uint16(C.getC())
	C.setC(result > 0x0FF)
	C.setV(C.A, value, byte(result))
	C.A = byte(result)
	C.setNZStatus(C.A)
}

func (C *CPU) op_SBC_ZP(mem *mem.Memory)  { C.opName = "ToDO" }
func (C *CPU) op_SBC_ZPX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_SBC_ABS(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_SBC_ABX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_SBC_ABY(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_SBC_INX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_SBC_INY(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_CMP_IM(mem *mem.Memory) {
	C.opName = "CMP Im"
	value := C.fetchByte(mem)
	C.setC(C.A >= value)
	res := C.A - value
	C.setZ(res)
	C.setN(res)
}

func (C *CPU) op_CMP_ZP(mem *mem.Memory) {
	C.opName = "CMP ZP"
	zpAddress := C.fetchByte(mem)
	value := mem.Read(zpAddress)
	C.setC(C.A >= value)
	res := C.A - value
	C.setZ(res)
	C.setN(res)
}

func (C *CPU) op_CMP_ZPX(mem *mem.Memory) {
	C.opName = "CMP ZP,X"
	zpAddress := C.fetchByte(mem) + C.X
	value := mem.Read(zpAddress)
	C.setC(C.A >= value)
	res := C.A - value
	C.setZ(res)
	C.setN(res)
}

func (C *CPU) op_CMP_ABS(mem *mem.Memory) {
	C.opName = "CMP Abs"
	absAddress := C.fetchWord(mem)
	value := mem.Read(absAddress)
	C.setC(C.A >= value)
	res := C.A - value
	C.setZ(res)
	C.setN(res)
}

func (C *CPU) op_CMP_ABX(mem *mem.Memory) {
	absAddress := C.fetchWord(mem)
	dest := absAddress + uint16(C.X)
	value := mem.Read(dest)
	C.opName = fmt.Sprintf("CMP $%04X,X -> %02X", absAddress, value)
	C.setC(C.A >= value)
	res := C.A - value
	C.setZ(res)
	C.setN(res)
}

func (C *CPU) op_CMP_ABY(mem *mem.Memory) {
	C.opName = "CMP Abs,Y"
	absAddress := C.fetchWord(mem) + uint16(C.Y)
	value := mem.Read(absAddress)
	C.setC(C.A >= value)
	res := C.A - value
	C.setZ(res)
	C.setN(res)
}

func (C *CPU) op_CMP_INX(mem *mem.Memory) {
	C.opName = "CMP (ZP,X)"
	zpAddr := C.fetchByte(mem)
	wordZP := C.Indexed_indirect_X(zpAddr, C.X)
	value := mem.Read(wordZP)
	C.setC(C.A >= value)
	res := C.A - value
	C.setZ(res)
	C.setN(res)
}

func (C *CPU) op_CMP_INY(mem *mem.Memory) {
	zpAddr := C.fetchByte(mem)
	wordZP := C.Indirect_index_Y(zpAddr, C.Y)
	value := mem.Read(wordZP)
	C.opName = fmt.Sprintf("CMP ($%02X),Y -> %02X", zpAddr, value)
	C.setC(C.A >= value)
	res := C.A - value
	C.setZ(res)
	C.setN(res)
}

func (C *CPU) op_CPX_IM(mem *mem.Memory) {
	C.opName = "CPX Im"
	value := C.fetchByte(mem)
	C.setC(C.X >= value)
	res := C.X - value
	C.setZ(res)
	C.setN(res)
}

func (C *CPU) op_CPX_ZP(mem *mem.Memory) {
	C.opName = "CPX ZP"
	zpAddress := C.fetchByte(mem)
	value := mem.Read(zpAddress)
	C.setC(C.X >= value)
	res := C.X - value
	C.setZ(res)
	C.setN(res)
}

func (C *CPU) op_CPX_ABS(mem *mem.Memory) {
	C.opName = "CPX Abs"
	absAddress := C.fetchWord(mem)
	value := mem.Read(absAddress)
	C.setC(C.X >= value)
	res := C.X - value
	C.setZ(res)
	C.setN(res)
}

func (C *CPU) op_CPY_IM(mem *mem.Memory) {
	C.opName = "CPY Im"
	value := C.fetchByte(mem)
	C.setC(C.Y >= value)
	res := C.Y - value
	C.setZ(res)
	C.setN(res)
}

func (C *CPU) op_CPY_ZP(mem *mem.Memory) {
	C.opName = "CPY ZP"
	zpAddress := C.fetchByte(mem)
	value := mem.Read(zpAddress)
	C.setC(C.Y >= value)
	res := C.Y - value
	C.setZ(res)
	C.setN(res)
}

func (C *CPU) op_CPY_ABS(mem *mem.Memory) {
	C.opName = "CPY Abs"
	absAddress := C.fetchWord(mem)
	value := mem.Read(absAddress)
	C.setC(C.Y >= value)
	res := C.Y - value
	C.setZ(res)
	C.setN(res)
}
