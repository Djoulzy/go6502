package cpu

import (
	"fmt"
)

func (C *CPU) op_ADC_IM() {
	value := C.fetchByte()
	result := uint16(C.A) + uint16(value) + uint16(C.getC())
	C.setC(result > 0x0FF)
	C.setV(C.A, value, byte(result))
	C.A = byte(result)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("ADC #$%02X", value)
		C.debug = fmt.Sprintf(" = %02X", result)
	}
}

func (C *CPU) op_ADC_ZP() {
	zpAddress := C.fetchByte()
	value := C.readByte(uint16(zpAddress))
	result := uint16(C.A) + uint16(value) + uint16(C.getC())
	C.setC(result > 0x0FF)
	C.setV(C.A, value, byte(result))
	C.A = byte(result)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = "ADC ZP"
	}
}

func (C *CPU) op_ADC_ZPX() {
	zpAddress := C.fetchByte() + C.X
	value := C.readByte(uint16(zpAddress))
	result := uint16(C.A) + uint16(value) + uint16(C.getC())
	C.setC(result > 0x0FF)
	C.setV(C.A, value, byte(result))
	C.A = byte(result)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = "ADC ZP,X"
	}
}

func (C *CPU) op_ADC_ABS() {
	absAddress := C.fetchWord()
	value := C.readByte(absAddress)
	result := uint16(C.A) + uint16(value) + uint16(C.getC())
	C.setC(result > 0x0FF)
	C.setV(C.A, value, byte(result))
	C.A = byte(result)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = "ADC Abs"
	}
}

func (C *CPU) op_ADC_ABX() {
	absAddress := C.fetchWord() + uint16(C.X)
	value := C.readByte(absAddress)
	result := uint16(C.A) + uint16(value) + uint16(C.getC())
	C.setC(result > 0x0FF)
	C.setV(C.A, value, byte(result))
	C.A = byte(result)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = "ADC Abs,X"
	}
}

func (C *CPU) op_ADC_ABY() {
	absAddress := C.fetchWord() + uint16(C.Y)
	value := C.readByte(absAddress)
	result := uint16(C.A) + uint16(value) + uint16(C.getC())
	C.setC(result > 0x0FF)
	C.setV(C.A, value, byte(result))
	C.A = byte(result)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = "ADC Abs,Y"
	}
}

func (C *CPU) op_ADC_INX() {
	zpAddr := C.fetchByte()
	wordZP := C.Indexed_indirect_X(zpAddr, C.X)
	value := C.readByte(wordZP)
	result := uint16(C.A) + uint16(value) + uint16(C.getC())
	C.setC(result > 0x0FF)
	C.setV(C.A, value, byte(result))
	C.A = byte(result)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = "ADC (ZP,X)"
	}
}

func (C *CPU) op_ADC_INY() {
	zpAddr := C.fetchByte()
	wordZP := C.Indirect_index_Y(zpAddr, C.Y)
	value := C.readByte(wordZP)
	result := uint16(C.A) + uint16(value) + uint16(C.getC())
	C.setC(result > 0x0FF)
	C.setV(C.A, value, byte(result))
	C.A = byte(result)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = "ADC (ZP),Y"
	}
}

func (C *CPU) op_SBC_IM() {
	addr := C.fetchByte()
	value := ^addr
	result := uint16(C.A) + uint16(value) + uint16(C.getC())
	C.setC(result > 0x0FF)
	C.setV(C.A, value, byte(result))
	C.A = byte(result)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("SBC #$%02X", addr)
		C.debug = fmt.Sprintf("A - %02X = %02X", addr, result)
	}
}

func (C *CPU) op_SBC_ZP() {
	zpAddress := C.fetchByte()
	content := C.readByte(uint16(zpAddress))
	value := ^content
	result := uint16(C.A) + uint16(value) + uint16(C.getC())
	C.setC(result > 0x0FF)
	C.setV(C.A, value, byte(result))
	C.A = byte(result)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("SBC $%02X", zpAddress)
		C.debug = fmt.Sprintf("A - %02X = %02X", content, result)
	}
}

func (C *CPU) op_SBC_ZPX() { C.opName = "ToDO" }
func (C *CPU) op_SBC_ABS() { C.opName = "ToDO" }
func (C *CPU) op_SBC_ABX() { C.opName = "ToDO" }
func (C *CPU) op_SBC_ABY() { C.opName = "ToDO" }
func (C *CPU) op_SBC_INX() { C.opName = "ToDO" }
func (C *CPU) op_SBC_INY() { C.opName = "ToDO" }

func (C *CPU) op_CMP_IM() {
	value := C.fetchByte()
	C.setC(C.A >= value)
	res := C.A - value
	C.setNZStatus(res)

	if C.Display {
		C.opName = fmt.Sprintf("CMP #$%02X", value)
	}
}

func (C *CPU) op_CMP_ZP() {
	zpAddress := C.fetchByte()
	value := C.readByte(uint16(zpAddress))
	C.setC(C.A >= value)
	res := C.A - value
	C.setNZStatus(res)

	if C.Display {
		C.opName = fmt.Sprintf("CMP $%02X", zpAddress)
	}
}

func (C *CPU) op_CMP_ZPX() {
	zpAddress := C.fetchByte() + C.X
	dest := zpAddress + C.X
	value := C.readByte(uint16(dest))
	C.setC(C.A >= value)
	res := C.A - value
	C.setNZStatus(res)

	if C.Display {
		C.opName = "CMP ZP,X"
	}
}

func (C *CPU) op_CMP_ABS() {
	absAddress := C.fetchWord()
	value := C.readByte(absAddress)
	C.setC(C.A >= value)
	res := C.A - value
	C.setNZStatus(res)

	if C.Display {
		C.opName = fmt.Sprintf("CMP $%04X", absAddress)
	}
}

func (C *CPU) op_CMP_ABX() {
	absAddress := C.fetchWord()
	dest := absAddress + uint16(C.X)
	value := C.readByte(dest)
	C.setC(C.A >= value)
	res := C.A - value
	C.setNZStatus(res)

	if C.Display {
		C.opName = fmt.Sprintf("CMP $%04X,X", absAddress)
		C.debug = fmt.Sprintf("($%04X) = %02X vs %02X", dest, value, C.A)
	}
}

func (C *CPU) op_CMP_ABY() {
	absAddress := C.fetchWord() + uint16(C.Y)
	value := C.readByte(absAddress)
	C.setC(C.A >= value)
	res := C.A - value
	C.setNZStatus(res)

	if C.Display {
		C.opName = fmt.Sprintf("CMP $%04X,Y", absAddress)
	}
}

func (C *CPU) op_CMP_INX() {
	zpAddr := C.fetchByte()
	wordZP := C.Indexed_indirect_X(zpAddr, C.X)
	value := C.readByte(wordZP)
	C.setC(C.A >= value)
	res := C.A - value
	C.setNZStatus(res)

	if C.Display {
		C.opName = fmt.Sprintf("CMP ($%02X,X)", zpAddr)
		C.debug = fmt.Sprintf("%02X vs %02X", value, C.A)
	}
}

func (C *CPU) op_CMP_INY() {
	zpAddr := C.fetchByte()
	wordZP := C.Indirect_index_Y(zpAddr, C.Y)
	value := C.readByte(wordZP)
	C.setC(C.A >= value)
	res := C.A - value
	C.setNZStatus(res)

	if C.Display {
		C.opName = fmt.Sprintf("CMP ($%02X),Y", zpAddr)
		C.debug = fmt.Sprintf("%02X vs %02X", value, C.A)
	}
}

func (C *CPU) op_CPX_IM() {
	value := C.fetchByte()
	C.setC(C.X >= value)
	res := C.X - value
	C.setNZStatus(res)

	if C.Display {
		C.opName = fmt.Sprintf("CPX #$%02X", value)
	}
}

func (C *CPU) op_CPX_ZP() {
	zpAddress := C.fetchByte()
	value := C.readByte(uint16(zpAddress))
	C.setC(C.X >= value)
	res := C.X - value
	C.setNZStatus(res)

	if C.Display {
		C.opName = fmt.Sprintf("CPX $%02X", zpAddress)
		C.debug = fmt.Sprintf("-> %02X", value)
	}
}

func (C *CPU) op_CPX_ABS() {
	absAddress := C.fetchWord()
	value := C.readByte(absAddress)
	C.setC(C.X >= value)
	res := C.X - value
	C.setNZStatus(res)

	if C.Display {
		C.opName = fmt.Sprintf("CPX $%04X", absAddress)
	}
}

func (C *CPU) op_CPY_IM() {
	value := C.fetchByte()
	C.setC(C.Y >= value)
	res := C.Y - value
	C.setNZStatus(res)

	if C.Display {
		C.opName = fmt.Sprintf("CPY #$%02X", value)
	}
}

func (C *CPU) op_CPY_ZP() {
	zpAddress := C.fetchByte()
	value := C.readByte(uint16(zpAddress))
	C.setC(C.Y >= value)
	res := C.Y - value
	C.setNZStatus(res)

	if C.Display {
		C.opName = fmt.Sprintf("CPY $%02X", zpAddress)
		C.debug = fmt.Sprintf("-> %02X", value)
	}
}

func (C *CPU) op_CPY_ABS() {
	absAddress := C.fetchWord()
	value := C.readByte(absAddress)
	C.setC(C.Y >= value)
	res := C.Y - value
	C.setNZStatus(res)

	if C.Display {
		C.opName = fmt.Sprintf("CPY $%04X", absAddress)
	}
}
