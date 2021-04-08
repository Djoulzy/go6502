package cpu

import (
	"fmt"
	"go6502/databus"
	"go6502/mem"
	"os"
	"time"
)

func (C *CPU) reset(mem *mem.Memory) {
	C.A = 0xAA
	C.X = 0
	C.Y = 0
	C.S = 0b00000000

	C.PC = 0xFF00
	C.SP = 0xFF

	C.exit = false
}

//////////////////////////////////
//////// Stack Operations ////////
//////////////////////////////////

// Word
func (C *CPU) pushWordStack(mem *mem.Memory, val uint16) {
	low := byte(val)
	hi := byte(val >> 8)
	C.pushByteStack(mem, hi)
	C.pushByteStack(mem, low)
}

func (C *CPU) pullWordStack(mem *mem.Memory) uint16 {
	low := C.pullByteStack(mem)
	hi := uint16(C.pullByteStack(mem)) << 8
	return hi + uint16(low)
}

// Byte
func (C *CPU) pushByteStack(mem *mem.Memory, val byte) {
	mem.Stack[C.SP].Ram = val
	C.SP--
	if C.SP < 0 {
		panic("Stack overflow")
	}
	C.dbus.WaitBusLow()
}

func (C *CPU) pullByteStack(mem *mem.Memory) byte {
	C.SP++
	if C.SP > 0xFF {
		panic("Stack overflow")
	}
	C.dbus.WaitBusLow()
	return mem.Stack[C.SP].Ram
}

//////////////////////////////////
////// Addressage Indirect ///////
//////////////////////////////////

// https://stackoverflow.com/questions/46262435/indirect-y-indexed-addressing-mode-in-mos-6502
// http://www.emulator101.com/6502-addressing-modes.html

func (C *CPU) Indirect_index_Y(addr byte, y byte) uint16 {
	zpAddr := uint16(addr)
	wordZP := C.readWord(zpAddr) + uint16(y)
	return wordZP
}

func (C *CPU) Indexed_indirect_X(addr byte, x byte) uint16 {
	zpAddr := uint16(addr + x)
	wordZP := C.readWord(zpAddr)
	return wordZP
}

//////////////////////////////////
/////// Memory Operations ////////
//////////////////////////////////

func (C *CPU) readWord(addr uint16) uint16 {
	low := C.ram.Read(addr)
	value := uint16(C.ram.Read(addr+1)) << 8
	C.dbus.WaitBusLow()
	value += uint16(low)
	C.dbus.WaitBusLow()
	return value
}

func (C *CPU) fetchWord(mem *mem.Memory) uint16 {
	low := C.fetchByte(mem)
	value := uint16(C.fetchByte(mem)) << 8
	value += uint16(low)
	return value
}

func (C *CPU) fetchByte(mem *mem.Memory) byte {
	// if C.Display {
	// 	C.refreshScreen(mem)
	// }
	value := mem.Mem[C.PC].Ram
	fmt.Printf(" %02X", value)
	C.PC++
	C.dbus.WaitBusLow()
	return value
}

func (C *CPU) exec(mem *mem.Memory) {
	if C.exit || C.PC == C.BP {
		time.Sleep(time.Second)
		os.Exit(1)
	}
	if C.Display {
		fmt.Printf("\n%08b - A:%02X X:%02X Y:%02X - %04X:", C.SP, C.A, C.X, C.Y, C.PC)
	}
	opCode := C.fetchByte(mem)
	if C.Display {
		fmt.Printf(" %02X", opCode)
	}
	Mnemonic[opCode](mem)
	if C.Display {
		fmt.Printf("\t%s", C.opName)
	}
}

func (C *CPU) SetBreakpoint(bp uint16) {
	C.BP = bp
}

//////////////////////////////////
//////////// Language ////////////
//////////////////////////////////

// func (C *CPU) CheckMnemonic(code string) {
// 	test := Mnemonic[code]
// }

func (C *CPU) initLanguage() {
	Mnemonic = make(map[byte]func(*mem.Memory))

	Mnemonic[CodeAddr["SHW"]] = C.op_SHW
	Mnemonic[CodeAddr["DMP"]] = C.op_DMP

	Mnemonic[CodeAddr["NOP"]] = C.op_NOP
	Mnemonic[CodeAddr["BRK"]] = C.op_BRK

	Mnemonic[CodeAddr["INC_ZP"]] = C.op_INC_ZP
	Mnemonic[CodeAddr["INC_ZPX"]] = C.op_INC_ZPX
	Mnemonic[CodeAddr["INC_ABS"]] = C.op_INC_ABS
	Mnemonic[CodeAddr["INC_ABX"]] = C.op_INC_ABX
	Mnemonic[CodeAddr["INX"]] = C.op_INX
	Mnemonic[CodeAddr["INY"]] = C.op_INY
	Mnemonic[CodeAddr["DEC_ZP"]] = C.op_DEC_ZP
	Mnemonic[CodeAddr["DEC_ZPX"]] = C.op_DEC_ZPX
	Mnemonic[CodeAddr["DEC_ABS"]] = C.op_DEC_ABS
	Mnemonic[CodeAddr["DEC_ABX"]] = C.op_DEC_ABX
	Mnemonic[CodeAddr["DEX"]] = C.op_DEX
	Mnemonic[CodeAddr["DEY"]] = C.op_DEY

	Mnemonic[CodeAddr["ADC_IM"]] = C.op_ADC_IM
	Mnemonic[CodeAddr["ADC_ZP"]] = C.op_ADC_ZP
	Mnemonic[CodeAddr["ADC_ZPX"]] = C.op_ADC_ZPX
	Mnemonic[CodeAddr["ADC_ABS"]] = C.op_ADC_ABS
	Mnemonic[CodeAddr["ADC_ABX"]] = C.op_ADC_ABX
	Mnemonic[CodeAddr["ADC_ABY"]] = C.op_ADC_ABY
	Mnemonic[CodeAddr["ADC_INX"]] = C.op_ADC_INX
	Mnemonic[CodeAddr["ADC_INY"]] = C.op_ADC_INY

	Mnemonic[CodeAddr["SBC_IM"]] = C.op_SBC_IM
	Mnemonic[CodeAddr["SBC_ZP"]] = C.op_SBC_ZP
	Mnemonic[CodeAddr["SBC_ZPX"]] = C.op_SBC_ZPX
	Mnemonic[CodeAddr["SBC_ABS"]] = C.op_SBC_ABS
	Mnemonic[CodeAddr["SBC_ABX"]] = C.op_SBC_ABX
	Mnemonic[CodeAddr["SBC_ABY"]] = C.op_SBC_ABY
	Mnemonic[CodeAddr["SBC_INX"]] = C.op_SBC_INX
	Mnemonic[CodeAddr["SBC_INY"]] = C.op_SBC_INY

	Mnemonic[CodeAddr["CMP_IM"]] = C.op_CMP_IM
	Mnemonic[CodeAddr["CMP_ZP"]] = C.op_CMP_ZP
	Mnemonic[CodeAddr["CMP_ZPX"]] = C.op_CMP_ZPX
	Mnemonic[CodeAddr["CMP_ABS"]] = C.op_CMP_ABS
	Mnemonic[CodeAddr["CMP_ABX"]] = C.op_CMP_ABX
	Mnemonic[CodeAddr["CMP_ABY"]] = C.op_CMP_ABY
	Mnemonic[CodeAddr["CMP_INX"]] = C.op_CMP_INX
	Mnemonic[CodeAddr["CMP_INY"]] = C.op_CMP_INY

	Mnemonic[CodeAddr["CPX_IM"]] = C.op_CPX_IM
	Mnemonic[CodeAddr["CPX_ZP"]] = C.op_CPX_ZP
	Mnemonic[CodeAddr["CPX_ABS"]] = C.op_CPX_ABS

	Mnemonic[CodeAddr["CPY_IM"]] = C.op_CPY_IM
	Mnemonic[CodeAddr["CPY_ZP"]] = C.op_CPY_ZP
	Mnemonic[CodeAddr["CPY_ABS"]] = C.op_CPY_ABS

	Mnemonic[CodeAddr["BCC_REL"]] = C.op_BCC_REL
	Mnemonic[CodeAddr["BCS_REL"]] = C.op_BCS_REL
	Mnemonic[CodeAddr["BEQ_REL"]] = C.op_BEQ_REL
	Mnemonic[CodeAddr["BMI_REL"]] = C.op_BMI_REL
	Mnemonic[CodeAddr["BNE_REL"]] = C.op_BNE_REL
	Mnemonic[CodeAddr["BPL_REL"]] = C.op_BPL_REL
	Mnemonic[CodeAddr["BVC_REL"]] = C.op_BVC_REL
	Mnemonic[CodeAddr["BVS_REL"]] = C.op_BVS_REL

	Mnemonic[CodeAddr["LDA_IM"]] = C.op_LDA_IM
	Mnemonic[CodeAddr["LDA_ZP"]] = C.op_LDA_ZP
	Mnemonic[CodeAddr["LDA_ZPX"]] = C.op_LDA_ZPX
	Mnemonic[CodeAddr["LDA_INX"]] = C.op_LDA_INX
	Mnemonic[CodeAddr["LDA_INY"]] = C.op_LDA_INY
	Mnemonic[CodeAddr["LDA_ABS"]] = C.op_LDA_ABS
	Mnemonic[CodeAddr["LDA_ABX"]] = C.op_LDA_ABX
	Mnemonic[CodeAddr["LDA_ABY"]] = C.op_LDA_ABY

	Mnemonic[CodeAddr["LDX_IM"]] = C.op_LDX_IM
	Mnemonic[CodeAddr["LDX_ZP"]] = C.op_LDX_ZP
	Mnemonic[CodeAddr["LDX_ZPY"]] = C.op_LDX_ZPY
	Mnemonic[CodeAddr["LDX_ABS"]] = C.op_LDX_ABS
	Mnemonic[CodeAddr["LDX_ABY"]] = C.op_LDX_ABY

	Mnemonic[CodeAddr["LDY_IM"]] = C.op_LDY_IM
	Mnemonic[CodeAddr["LDY_ZP"]] = C.op_LDY_ZP
	Mnemonic[CodeAddr["LDY_ZPX"]] = C.op_LDY_ZPX
	Mnemonic[CodeAddr["LDY_ABS"]] = C.op_LDY_ABS
	Mnemonic[CodeAddr["LDY_ABX"]] = C.op_LDY_ABX

	Mnemonic[CodeAddr["STA_ZP"]] = C.op_STA_ZP
	Mnemonic[CodeAddr["STA_ZPX"]] = C.op_STA_ZPX
	Mnemonic[CodeAddr["STA_INX"]] = C.op_STA_INX
	Mnemonic[CodeAddr["STA_INY"]] = C.op_STA_INY
	Mnemonic[CodeAddr["STA_ABS"]] = C.op_STA_ABS
	Mnemonic[CodeAddr["STA_ABX"]] = C.op_STA_ABX
	Mnemonic[CodeAddr["STA_ABY"]] = C.op_STA_ABY

	Mnemonic[CodeAddr["STX_ZP"]] = C.op_STX_ZP
	Mnemonic[CodeAddr["STX_ZPY"]] = C.op_STX_ZPY
	Mnemonic[CodeAddr["STX_ABS"]] = C.op_STX_ABS

	Mnemonic[CodeAddr["STY_ZP"]] = C.op_STY_ZP
	Mnemonic[CodeAddr["STY_ZPX"]] = C.op_STY_ZPX
	Mnemonic[CodeAddr["STY_ABS"]] = C.op_STY_ABS

	Mnemonic[CodeAddr["AND_IM"]] = C.op_AND_IM
	Mnemonic[CodeAddr["AND_ZP"]] = C.op_AND_ZP
	Mnemonic[CodeAddr["AND_ZPX"]] = C.op_AND_ZPX
	Mnemonic[CodeAddr["AND_ABS"]] = C.op_AND_ABS
	Mnemonic[CodeAddr["AND_ABX"]] = C.op_AND_ABX
	Mnemonic[CodeAddr["AND_ABY"]] = C.op_AND_ABY
	Mnemonic[CodeAddr["AND_INX"]] = C.op_AND_INX
	Mnemonic[CodeAddr["AND_INY"]] = C.op_AND_INY

	Mnemonic[CodeAddr["EOR_IM"]] = C.op_EOR_IM
	Mnemonic[CodeAddr["EOR_ZP"]] = C.op_EOR_ZP
	Mnemonic[CodeAddr["EOR_ZPX"]] = C.op_EOR_ZPX
	Mnemonic[CodeAddr["EOR_ABS"]] = C.op_EOR_ABS
	Mnemonic[CodeAddr["EOR_ABX"]] = C.op_EOR_ABX
	Mnemonic[CodeAddr["EOR_ABY"]] = C.op_EOR_ABY
	Mnemonic[CodeAddr["EOR_INX"]] = C.op_EOR_INX
	Mnemonic[CodeAddr["EOR_INY"]] = C.op_EOR_INY

	Mnemonic[CodeAddr["ORA_IM"]] = C.op_ORA_IM
	Mnemonic[CodeAddr["ORA_ZP"]] = C.op_ORA_ZP
	Mnemonic[CodeAddr["ORA_ZPX"]] = C.op_ORA_ZPX
	Mnemonic[CodeAddr["ORA_ABS"]] = C.op_ORA_ABS
	Mnemonic[CodeAddr["ORA_ABX"]] = C.op_ORA_ABX
	Mnemonic[CodeAddr["ORA_ABY"]] = C.op_ORA_ABY
	Mnemonic[CodeAddr["ORA_INX"]] = C.op_ORA_INX
	Mnemonic[CodeAddr["ORA_INY"]] = C.op_ORA_INY

	Mnemonic[CodeAddr["BIT_ZP"]] = C.op_BIT_ZP
	Mnemonic[CodeAddr["BIT_ABS"]] = C.op_BIT_ABS

	Mnemonic[CodeAddr["TXS"]] = C.op_TXS
	Mnemonic[CodeAddr["PHA"]] = C.op_PHA
	Mnemonic[CodeAddr["PLA"]] = C.op_PLA
	Mnemonic[CodeAddr["TSX"]] = C.op_TSX
	Mnemonic[CodeAddr["PHP"]] = C.op_PHP
	Mnemonic[CodeAddr["PLP"]] = C.op_PLP

	Mnemonic[CodeAddr["TAX"]] = C.op_TAX
	Mnemonic[CodeAddr["TAY"]] = C.op_TAY
	Mnemonic[CodeAddr["TXA"]] = C.op_TXA
	Mnemonic[CodeAddr["TYA"]] = C.op_TYA

	Mnemonic[CodeAddr["JMP_ABS"]] = C.op_JMP_ABS
	Mnemonic[CodeAddr["JMP_IND"]] = C.op_JMP_IND
	Mnemonic[CodeAddr["JSR"]] = C.op_JSR
	Mnemonic[CodeAddr["RTS"]] = C.op_RTS

	Mnemonic[CodeAddr["CLC"]] = C.op_CLC
	Mnemonic[CodeAddr["CLD"]] = C.op_CLD
	Mnemonic[CodeAddr["CLI"]] = C.op_CLI
	Mnemonic[CodeAddr["CLV"]] = C.op_CLV
	Mnemonic[CodeAddr["SEC"]] = C.op_SEC
	Mnemonic[CodeAddr["SED"]] = C.op_SED
	Mnemonic[CodeAddr["SEI"]] = C.op_SEI

	Mnemonic[CodeAddr["ASL_IM"]] = C.op_ASL_IM
	Mnemonic[CodeAddr["ASL_ZP"]] = C.op_ASL_ZP
	Mnemonic[CodeAddr["ASL_ZPX"]] = C.op_ASL_ZPX
	Mnemonic[CodeAddr["ASL_ABS"]] = C.op_ASL_ABS
	Mnemonic[CodeAddr["ASL_ABX"]] = C.op_ASL_ABX

	Mnemonic[CodeAddr["LSR_IM"]] = C.op_LSR_IM
	Mnemonic[CodeAddr["LSR_ZP"]] = C.op_LSR_ZP
	Mnemonic[CodeAddr["LSR_ZPX"]] = C.op_LSR_ZPX
	Mnemonic[CodeAddr["LSR_ABS"]] = C.op_LSR_ABS
	Mnemonic[CodeAddr["LSR_ABX"]] = C.op_LSR_ABX

	Mnemonic[CodeAddr["ROL_IM"]] = C.op_ROL_IM
	Mnemonic[CodeAddr["ROL_ZP"]] = C.op_ROL_ZP
	Mnemonic[CodeAddr["ROL_ZPX"]] = C.op_ROL_ZPX
	Mnemonic[CodeAddr["ROL_ABS"]] = C.op_ROL_ABS
	Mnemonic[CodeAddr["ROL_ABX"]] = C.op_ROL_ABX

	Mnemonic[CodeAddr["ROR_IM"]] = C.op_ROR_IM
	Mnemonic[CodeAddr["ROR_ZP"]] = C.op_ROR_ZP
	Mnemonic[CodeAddr["ROR_ZPX"]] = C.op_ROR_ZPX
	Mnemonic[CodeAddr["ROR_ABS"]] = C.op_ROR_ABS
	Mnemonic[CodeAddr["ROR_ABX"]] = C.op_ROR_ABX
}

func (C *CPU) Init(dbus *databus.Databus, mem *mem.Memory, disp bool) {
	C.Display = disp
	C.ram = mem
	C.dbus = dbus
	C.BP = 0

	C.initLanguage()
	// if C.Display {
	// 	C.initOutput(C.ram)
	// }

	C.reset(C.ram)
	// NMI
	C.ram.Mem[0xFFFA].Ram = 0x43
	C.ram.Mem[0xFFFB].Ram = 0xFE
	// Cold Start
	C.ram.Mem[0xFFFC].Ram = 0xE2
	C.ram.Mem[0xFFFD].Ram = 0xFC
	// IRQ
	C.ram.Mem[0xFFFE].Ram = 0x48
	C.ram.Mem[0xFFFF].Ram = 0xFF
}

func (C *CPU) Run() {
	for {
		C.exec(C.ram)
	}
}
