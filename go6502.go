package main

import (
	"os"
	"time"
)

func (C *CPU) reset(mem *Memory) {
	mem.Init()
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

func (C *CPU) pushWordStack(mem *Memory, val Word) {
	low := Byte(val)
	hi := Byte(val >> 8)
	C.pushByteStack(mem, hi)
	C.pushByteStack(mem, low)
}

func (C *CPU) fetchWordStack(mem *Memory) Word {
	low := C.pullByteStack(mem)
	hi := Word(C.pullByteStack(mem)) << 8
	return hi + Word(low)
}

func (C *CPU) pushByteStack(mem *Memory, val Byte) {
	mem.Stack[C.SP] = val
	C.SP--
	if C.SP < 0 {
		panic("Stack overflow")
	}
}

func (C *CPU) pullByteStack(mem *Memory) Byte {
	C.SP++
	if C.SP > 0xFF {
		panic("Stack overflow")
	}
	return mem.Stack[C.SP]
}

//////////////////////////////////
/////// Memory Operations ////////
//////////////////////////////////

func (C *CPU) fetchWord(mem *Memory) Word {
	low := C.fetchByte(mem)
	value := Word(C.fetchByte(mem)) << 8
	value += Word(low)
	return value
}

func (C *CPU) fetchByte(mem *Memory) Byte {
	time.Sleep(time.Second)
	value := mem.Data[C.PC]
	C.PC++
	return value
}

func (C *CPU) exec(mem *Memory) {
	if C.exit {
		time.Sleep(time.Second)
		os.Exit(1)
	}
	opCode := C.fetchByte(mem)
	Nemonic[opCode](mem)
}

func (C *CPU) initLanguage() {
	Nemonic = make(map[Byte]func(*Memory))

	Nemonic[NOP] = C.op_NOP
	Nemonic[BRK] = C.op_BRK

	Nemonic[LDA_IM] = C.op_LDA_IM
	Nemonic[LDA_ZP] = C.op_LDA_ZP
	Nemonic[LDA_ZPX] = C.op_LDA_ZPX
	Nemonic[LDA_INX] = C.op_LDA_INX
	Nemonic[LDA_INY] = C.op_LDA_INY
	Nemonic[LDX_IM] = C.op_LDX_IM
	Nemonic[LDX_ZP] = C.op_LDX_ZP
	Nemonic[LDX_ZPY] = C.op_LDX_ZPY
	Nemonic[LDY_IM] = C.op_LDY_IM
	Nemonic[LDY_ZP] = C.op_LDY_ZP
	Nemonic[LDY_ZPX] = C.op_LDY_ZPX

	Nemonic[AND_IM] = C.op_AND_IM
	Nemonic[AND_ZP] = C.op_AND_ZP
	Nemonic[AND_ZPX] = C.op_AND_ZPX
	Nemonic[AND_ABS] = C.op_AND_ABS
	Nemonic[AND_ABX] = C.op_AND_ABX
	Nemonic[AND_ABY] = C.op_AND_ABY
	Nemonic[AND_INX] = C.op_AND_INX
	Nemonic[AND_INY] = C.op_AND_INY

	Nemonic[TXS] = C.op_TXS
	Nemonic[PHA] = C.op_PHA
	Nemonic[PLA] = C.op_PLA

	Nemonic[JMP_ABS] = C.op_JMP_ABS
	Nemonic[JMP_IND] = C.op_JMP_IND
	Nemonic[JSR] = C.op_JSR
	Nemonic[RTS] = C.op_RTS
}

func main() {
	mem := Memory{}
	cpu := CPU{}
	cpu.initLanguage()

	go cpu.output(&mem)

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
