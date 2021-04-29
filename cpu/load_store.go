package cpu

import (
	"fmt"
)

//////////////////////////////////
///////////// LDA ////////////////
//////////////////////////////////

// op_LDA_IM : LDA Immediate
func (C *CPU) op_LDA_IM() {
	C.A = C.fetchByte()
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("LDA #$%02X", C.A)
	}
}

// op_LDA_ZP : LDA Zero Page
func (C *CPU) op_LDA_ZP() {
	zpAddress := C.fetchByte()
	C.A = C.readByte(uint16(zpAddress))
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("LDA $%02X", zpAddress)
	}
}

// op_LDA_ZPX : LDA Zero Page,X
func (C *CPU) op_LDA_ZPX() {
	zpAddress := C.fetchByte()
	dest := zpAddress + C.X
	C.A = C.readByte(uint16(dest))
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("LDA $%02X,X", zpAddress)
		C.debug = fmt.Sprintf("-> $%02X", dest)
	}
}

func (C *CPU) op_LDA_ABS() {
	absAddress := C.fetchWord()
	C.A = C.readByte(absAddress)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("LDA $%04X", absAddress)
	}
}

func (C *CPU) op_LDA_ABX() {
	absAddress := C.fetchWord()
	C.A = C.readByte(absAddress + uint16(C.X))
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("LDA $%04X,X", absAddress)
	}
}

func (C *CPU) op_LDA_ABY() {
	absAddress := C.fetchWord()
	dest := absAddress + uint16(C.Y)
	C.A = C.readByte(dest)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("LDA $%04X,Y", absAddress)
	}
	C.debug = fmt.Sprintf("-> $%04X", dest)
}

func (C *CPU) op_LDA_INX() {
	zpAddr := C.fetchByte()
	wordZP := C.Indexed_indirect_X(zpAddr, C.X)
	C.A = C.readByte(wordZP)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("LDA ($%02X,X)", zpAddr)
	}
}

func (C *CPU) op_LDA_INY() {
	zpAddr := C.fetchByte()
	wordZP := C.Indirect_index_Y(zpAddr, C.Y)
	C.A = C.readByte(wordZP)
	C.setNZStatus(C.A)

	if C.Display {
		C.opName = fmt.Sprintf("LDA ($%02X),Y", zpAddr)
	}
}

//////////////////////////////////
///////////// LDX ////////////////
//////////////////////////////////

// op_LDX_IM : LDA Immediate
func (C *CPU) op_LDX_IM() {
	C.X = C.fetchByte()
	C.setNZStatus(C.X)

	if C.Display {
		C.opName = fmt.Sprintf("LDX #$%02X", C.X)
	}
}

// op_LDX_ZP : LDA Zero Page
func (C *CPU) op_LDX_ZP() {
	zpAddress := C.fetchByte()
	C.X = C.readByte(uint16(zpAddress))
	C.setNZStatus(C.X)

	if C.Display {
		C.opName = fmt.Sprintf("LDX $%02X", zpAddress)
	}
}

// op_LDX_ZPY : LDA Zero Page,Y
func (C *CPU) op_LDX_ZPY() {
	zpAddress := C.fetchByte()
	dest := zpAddress + C.Y
	C.X = C.readByte(uint16(dest))
	C.setNZStatus(C.X)

	if C.Display {
		C.opName = fmt.Sprintf("LDX $%02X,Y", zpAddress)
		C.debug = fmt.Sprintf("-> $%02X", dest)
	}
}

func (C *CPU) op_LDX_ABS() {
	absAddress := C.fetchWord()
	C.X = C.readByte(absAddress)
	C.setNZStatus(C.X)

	if C.Display {
		C.opName = fmt.Sprintf("LDX $%04X", absAddress)
	}
}

func (C *CPU) op_LDX_ABY() {
	absAddress := C.fetchWord()
	dest := absAddress + uint16(C.Y)
	C.X = C.readByte(dest)
	C.setNZStatus(C.X)

	if C.Display {
		C.opName = fmt.Sprintf("LDX $%04X,Y", absAddress)
		C.debug = fmt.Sprintf("-> $%04X", dest)
	}
}

//////////////////////////////////
///////////// LDY ////////////////
//////////////////////////////////

func (C *CPU) op_LDY_IM() {
	C.Y = C.fetchByte()
	C.setNZStatus(C.Y)

	if C.Display {
		C.opName = fmt.Sprintf("LDY #$%02X", C.Y)
	}
}

func (C *CPU) op_LDY_ZP() {
	zpAddress := C.fetchByte()
	C.Y = C.readByte(uint16(zpAddress))
	C.setNZStatus(C.Y)

	if C.Display {
		C.opName = "LDY ZP"
	}
}

func (C *CPU) op_LDY_ZPX() {
	zpAddress := C.fetchByte()
	dest := zpAddress + C.X
	C.Y = C.readByte(uint16(dest))
	C.setNZStatus(C.Y)

	if C.Display {
		C.opName = fmt.Sprintf("LDY $%02X,X", zpAddress)
		C.debug = fmt.Sprintf("-> $%02X", dest)
	}
}

func (C *CPU) op_LDY_ABS() {
	absAddress := C.fetchWord()
	C.Y = C.readByte(absAddress)
	C.setNZStatus(C.Y)

	if C.Display {
		C.opName = fmt.Sprintf("LDY $%04X", absAddress)
	}
}

func (C *CPU) op_LDY_ABX() {
	absAddress := C.fetchWord()
	dest := absAddress + uint16(C.X)
	C.Y = C.readByte(dest)
	C.setNZStatus(C.Y)

	if C.Display {
		C.opName = fmt.Sprintf("LDY $%04X,X", absAddress)
	}
}

//////////////////////////////////
///////////// STA ////////////////
//////////////////////////////////

func (C *CPU) op_STA_ZP() {
	zpAddress := C.fetchByte()
	C.writeByte(uint16(zpAddress), C.A)
	val := C.readByte(uint16(zpAddress))

	if C.Display {
		C.opName = fmt.Sprintf("STA $%02X", zpAddress)
		C.debug = fmt.Sprintf("%02X -> $%02X", val, zpAddress)
	}
}

func (C *CPU) op_STA_ZPX() {
	zpAddress := C.fetchByte()
	dest := zpAddress + C.X
	C.writeByte(uint16(dest), C.A)

	if C.Display {
		C.opName = fmt.Sprintf("STA $%02X,X", zpAddress)
	}
	C.debug = fmt.Sprintf("-> $%02X", dest)
}

func (C *CPU) op_STA_ABS() {
	absAddress := C.fetchWord()
	C.writeByte(absAddress, C.A)

	if C.Display {
		C.opName = fmt.Sprintf("STA $%04X", absAddress)
	}
}

func (C *CPU) op_STA_ABX() {
	absAddress := C.fetchWord()
	dest := absAddress + uint16(C.X)
	C.dbus.Release()
	C.writeByte(dest, C.A)

	if C.Display {
		C.opName = fmt.Sprintf("STA $%04X,X", absAddress)
		C.debug = fmt.Sprintf("%02X -> $%04X", C.A, dest)
	}
}

func (C *CPU) op_STA_ABY() {
	absAddress := C.fetchWord()
	dest := absAddress + uint16(C.Y)
	C.dbus.Release()
	C.writeByte(dest, C.A)

	if C.Display {
		C.opName = fmt.Sprintf("STA $%04X,Y", absAddress)
		C.debug = fmt.Sprintf("%02X -> $%04X", C.A, dest)
	}
}

func (C *CPU) op_STA_INX() {
	zpAddr := C.fetchByte()
	wordZP := C.Indexed_indirect_X(zpAddr, C.X)
	C.dbus.Release()
	C.writeByte(wordZP, C.A)

	if C.Display {
		C.opName = fmt.Sprintf("STA ($%02X,X)", zpAddr)
		C.debug = fmt.Sprintf("%02X -> $%04X", C.A, wordZP)
	}
}

func (C *CPU) op_STA_INY() {
	zpAddr := C.fetchByte()
	wordZP := C.Indirect_index_Y(zpAddr, C.Y)
	C.dbus.Release()
	C.writeByte(wordZP, C.A)

	if C.Display {
		C.opName = fmt.Sprintf("STA ($%02X),Y", zpAddr)
		C.debug = fmt.Sprintf("%02X -> $%04X", C.A, wordZP)
	}
}

//////////////////////////////////
///////////// STX ////////////////
//////////////////////////////////

func (C *CPU) op_STX_ZP() {
	zpAddress := C.fetchByte()
	C.writeByte(uint16(zpAddress), C.X)

	if C.Display {
		C.opName = fmt.Sprintf("STX $%02X", zpAddress)
	}
}

func (C *CPU) op_STX_ZPY() {
	zpAddress := C.fetchByte() + C.Y
	C.writeByte(uint16(zpAddress), C.X)

	if C.Display {
		C.opName = fmt.Sprintf("STX $%02X,Y", zpAddress)
	}
}

func (C *CPU) op_STX_ABS() {
	absAddress := C.fetchWord()
	C.writeByte(absAddress, C.X)

	if C.Display {
		C.opName = fmt.Sprintf("STX $%04X", absAddress)
	}
}

//////////////////////////////////
///////////// STY ////////////////
//////////////////////////////////

func (C *CPU) op_STY_ZP() {
	zpAddress := C.fetchByte()
	C.writeByte(uint16(zpAddress), C.Y)
	C.dbus.Release()

	if C.Display {
		C.opName = fmt.Sprintf("STY $%02X", zpAddress)
	}
}

func (C *CPU) op_STY_ZPX() {
	zpAddress := C.fetchByte()
	dest := zpAddress + C.X
	C.writeByte(uint16(dest), C.Y)

	if C.Display {
		C.opName = fmt.Sprintf("STY $%02X,X", zpAddress)
	}
}

func (C *CPU) op_STY_ABS() {
	absAddress := C.fetchWord()
	C.writeByte(absAddress, C.Y)

	if C.Display {
		C.opName = fmt.Sprintf("STY $%04X", absAddress)
	}
}
