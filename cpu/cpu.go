package cpu

import (
	"go6502/globals"
	"go6502/mem"
	"os"
	"time"
)

func (C *CPU) reset(mem *mem.Memory) {
	C.A = 0
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

func (C *CPU) pushWordStack(mem *mem.Memory, val globals.Word) {
	low := globals.Byte(val)
	hi := globals.Byte(val >> 8)
	C.pushByteStack(mem, hi)
	C.pushByteStack(mem, low)
}

func (C *CPU) fetchWordStack(mem *mem.Memory) globals.Word {
	low := C.pullByteStack(mem)
	hi := globals.Word(C.pullByteStack(mem)) << 8
	return hi + globals.Word(low)
}

func (C *CPU) pushByteStack(mem *mem.Memory, val globals.Byte) {
	mem.Stack[C.SP] = val
	C.SP--
	if C.SP < 0 {
		panic("Stack overflow")
	}
}

func (C *CPU) pullByteStack(mem *mem.Memory) globals.Byte {
	C.SP++
	if C.SP > 0xFF {
		panic("Stack overflow")
	}
	return mem.Stack[C.SP]
}

//////////////////////////////////
/////// mem.Memory Operations ////////
//////////////////////////////////

func (C *CPU) fetchWord(mem *mem.Memory) globals.Word {
	low := C.fetchByte(mem)
	value := globals.Word(C.fetchByte(mem)) << 8
	value += globals.Word(low)
	return value
}

func (C *CPU) fetchByte(mem *mem.Memory) globals.Byte {
	if C.display {
		C.refreshScreen(mem)
	}
	<-C.cycle
	value := mem.Data[C.PC]
	C.PC++
	return value
}

func (C *CPU) exec(mem *mem.Memory) {
	if C.exit {
		time.Sleep(time.Second)
		os.Exit(1)
	}
	opCode := C.fetchByte(mem)
	Mnemonic[opCode](mem)
}

func (C *CPU) initLanguage() {
	Mnemonic = make(map[globals.Byte]func(*mem.Memory))

	Mnemonic[DMP] = C.op_DMP

	Mnemonic[NOP] = C.op_NOP
	Mnemonic[BRK] = C.op_BRK

	Mnemonic[INC_ZP] = C.op_INC_ZP
	Mnemonic[INC_ZPX] = C.op_INC_ZPX
	Mnemonic[INC_ABS] = C.op_INC_ABS
	Mnemonic[INC_ABX] = C.op_INC_ABX
	Mnemonic[INX] = C.op_INX
	Mnemonic[INY] = C.op_INY
	Mnemonic[DEC_ZP] = C.op_DEC_ZP
	Mnemonic[DEC_ZPX] = C.op_DEC_ZPX
	Mnemonic[DEC_ABS] = C.op_DEC_ABS
	Mnemonic[DEC_ABX] = C.op_DEC_ABX
	Mnemonic[DEX] = C.op_DEX
	Mnemonic[DEY] = C.op_DEY

	Mnemonic[ADC_IM] = C.op_ADC_IM
	Mnemonic[ADC_ZP] = C.op_ADC_ZP
	Mnemonic[ADC_ZPX] = C.op_ADC_ZPX
	Mnemonic[ADC_ABS] = C.op_ADC_ABS
	Mnemonic[ADC_ABX] = C.op_ADC_ABX
	Mnemonic[ADC_ABY] = C.op_ADC_ABY
	Mnemonic[ADC_INX] = C.op_ADC_INX
	Mnemonic[ADC_INY] = C.op_ADC_INY

	Mnemonic[SBC_IM] = C.op_SBC_IM
	Mnemonic[SBC_ZP] = C.op_SBC_ZP
	Mnemonic[SBC_ZPX] = C.op_SBC_ZPX
	Mnemonic[SBC_ABS] = C.op_SBC_ABS
	Mnemonic[SBC_ABX] = C.op_SBC_ABX
	Mnemonic[SBC_ABY] = C.op_SBC_ABY
	Mnemonic[SBC_INX] = C.op_SBC_INX
	Mnemonic[SBC_INY] = C.op_SBC_INY

	Mnemonic[CMP_IM] = C.op_CMP_IM
	Mnemonic[CMP_ZP] = C.op_CMP_ZP
	Mnemonic[CMP_ZPX] = C.op_CMP_ZPX
	Mnemonic[CMP_ABS] = C.op_CMP_ABS
	Mnemonic[CMP_ABX] = C.op_CMP_ABX
	Mnemonic[CMP_ABY] = C.op_CMP_ABY
	Mnemonic[CMP_INX] = C.op_CMP_INX
	Mnemonic[CMP_INY] = C.op_CMP_INY

	Mnemonic[CPX_IM] = C.op_CPX_IM
	Mnemonic[CPX_ZP] = C.op_CPX_ZP
	Mnemonic[CPX_ABS] = C.op_CPX_ABS

	Mnemonic[CPY_IM] = C.op_CPY_IM
	Mnemonic[CPY_ZP] = C.op_CPY_ZP
	Mnemonic[CPY_ABS] = C.op_CPY_ABS

	Mnemonic[BCC] = C.op_BCC
	Mnemonic[BCS] = C.op_BCS
	Mnemonic[BEQ] = C.op_BEQ
	Mnemonic[BMI] = C.op_BMI
	Mnemonic[BNE] = C.op_BNE
	Mnemonic[BPL] = C.op_BPL
	Mnemonic[BVC] = C.op_BVC
	Mnemonic[BVS] = C.op_BVS

	Mnemonic[LDA_IM] = C.op_LDA_IM
	Mnemonic[LDA_ZP] = C.op_LDA_ZP
	Mnemonic[LDA_ZPX] = C.op_LDA_ZPX
	Mnemonic[LDA_INX] = C.op_LDA_INX
	Mnemonic[LDA_INY] = C.op_LDA_INY
	Mnemonic[LDA_ABS] = C.op_LDA_ABS
	Mnemonic[LDA_ABX] = C.op_LDA_ABX
	Mnemonic[LDA_ABY] = C.op_LDA_ABY

	Mnemonic[LDX_IM] = C.op_LDX_IM
	Mnemonic[LDX_ZP] = C.op_LDX_ZP
	Mnemonic[LDX_ZPY] = C.op_LDX_ZPY
	Mnemonic[LDX_ABS] = C.op_LDX_ABS
	Mnemonic[LDX_ABY] = C.op_LDX_ABY

	Mnemonic[LDY_IM] = C.op_LDY_IM
	Mnemonic[LDY_ZP] = C.op_LDY_ZP
	Mnemonic[LDY_ZPX] = C.op_LDY_ZPX
	Mnemonic[LDY_ABS] = C.op_LDY_ABS
	Mnemonic[LDY_ABX] = C.op_LDY_ABX

	Mnemonic[STA_ZP] = C.op_STA_ZP
	Mnemonic[STA_ZPX] = C.op_STA_ZPX
	Mnemonic[STA_INX] = C.op_STA_INX
	Mnemonic[STA_INY] = C.op_STA_INY
	Mnemonic[STA_ABS] = C.op_STA_ABS
	Mnemonic[STA_ABX] = C.op_STA_ABX
	Mnemonic[STA_ABY] = C.op_STA_ABY

	Mnemonic[STX_ZP] = C.op_STX_ZP
	Mnemonic[STX_ZPY] = C.op_STX_ZPY
	Mnemonic[STX_ABS] = C.op_STX_ABS

	Mnemonic[STY_ZP] = C.op_STY_ZP
	Mnemonic[STY_ZPX] = C.op_STY_ZPX
	Mnemonic[STY_ABS] = C.op_STY_ABS

	Mnemonic[AND_IM] = C.op_AND_IM
	Mnemonic[AND_ZP] = C.op_AND_ZP
	Mnemonic[AND_ZPX] = C.op_AND_ZPX
	Mnemonic[AND_ABS] = C.op_AND_ABS
	Mnemonic[AND_ABX] = C.op_AND_ABX
	Mnemonic[AND_ABY] = C.op_AND_ABY
	Mnemonic[AND_INX] = C.op_AND_INX
	Mnemonic[AND_INY] = C.op_AND_INY

	Mnemonic[TXS] = C.op_TXS
	Mnemonic[PHA] = C.op_PHA
	Mnemonic[PLA] = C.op_PLA

	Mnemonic[JMP_ABS] = C.op_JMP_ABS
	Mnemonic[JMP_IND] = C.op_JMP_IND
	Mnemonic[JSR] = C.op_JSR
	Mnemonic[RTS] = C.op_RTS
}

func (C *CPU) Init(mem *mem.Memory) {
	C.Cycle = make(chan bool, 1)
	C.Display = false
	C.ram = mem
}

func (C *CPU) Run() {
	C.initLanguage()
	if C.Display {
		C.initOutput(C.ram)
	}

	C.reset(C.ram)
	C.load0(C.ram)

	// for i := range mem.Screen {
	// 	mem.Screen[i] = 0x39
	// }
	for {
		C.exec(C.ram)
	}
}
