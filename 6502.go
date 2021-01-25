package main

import "time"

func (C *CPU) reset(mem *Memory) {
	mem.Init()
	C.A = 0
	C.X = 0
	C.Y = 0
	C.S = 0b00000000

	C.PC = 0xFF00
	C.SP = 0x00
}

func (C *CPU) fetchByte(mem *Memory) Byte {
	time.Sleep(time.Second)
	opCode := mem.Data[C.PC]
	C.PC++
	return opCode
}

func (C *CPU) exec(mem *Memory) {
	opCode := C.fetchByte(mem)
	Mnemonic[opCode](mem)
}

func (C *CPU) initLanguage() {
	Mnemonic = make(map[Byte]func(*Memory))

	Mnemonic[NOP] = C.op_NOP

	Mnemonic[LDA_IM] = C.op_LDA_IM
	Mnemonic[LDA_ZP] = C.op_LDA_ZP
	Mnemonic[LDA_ZPX] = C.op_LDA_ZPX
	Mnemonic[LDX_IM] = C.op_LDX_IM
	Mnemonic[LDX_ZP] = C.op_LDX_ZP
	Mnemonic[LDX_ZPY] = C.op_LDX_ZPY
	Mnemonic[LDY_IM] = C.op_LDY_IM
	Mnemonic[LDY_ZP] = C.op_LDY_ZP
	Mnemonic[LDY_ZPX] = C.op_LDY_ZPX

	Mnemonic[TXS] = C.op_TXS
	Mnemonic[PHA] = C.op_PHA
	Mnemonic[PLA] = C.op_PLA

	Mnemonic[JMP_ABS] = C.op_JMP_ABS
	Mnemonic[JMP_IND] = C.op_JMP_IND
}

func main() {
	mem := Memory{}
	cpu := CPU{}
	cpu.initLanguage()

	go cpu.output()

	cpu.reset(&mem)
	mem.load()

	// cpu.exec(&mem)
	// cpu.exec(&mem)
	// cpu.exec(&mem)

	// for {
	// }
	for {
		cpu.exec(&mem)
	}
}
