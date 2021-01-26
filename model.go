package main

const (
	memorySize = 0xFFFF
	stackStart = 0x0100
	stackEnd   = 0x01FF
	intAddr    = 0xFFFA
	resetAddr  = 0xFFFC
	brkAddr    = 0xFFFE
)

// Byte :
type Byte uint8

// Word :
type Word uint16

// Memory :
type Memory struct {
	Data  [memorySize]Byte
	Stack []Byte
}

//
const (
	C_mask Byte = 0b11111110
	Z_mask Byte = 0b11111101
	I_mask Byte = 0b11111011
	D_mask Byte = 0b11110111
	B_mask Byte = 0b11101111

	V_mask Byte = 0b10111111
	N_mask Byte = 0b01111111
)

// CPU :
type CPU struct {
	PC Word
	SP Byte
	A  Byte
	X  Byte
	Y  Byte
	S  Byte

	opName string
}

// Mnemonic :
var Mnemonic map[Byte]func(*Memory)

// :
const (
	BRK = 0x00
	NOP = 0xEA

	LDA_IM  = 0x0A
	LDA_ZP  = 0xA5
	LDA_ZPX = 0xB5
	LDX_IM  = 0xA2
	LDX_ZP  = 0xA6
	LDX_ZPY = 0xB6
	LDY_IM  = 0xA0
	LDY_ZP  = 0xA4
	LDY_ZPX = 0xB4

	TXS = 0x9A
	PHA = 0x48
	PLA = 0x68

	JMP_ABS = 0x4C
	JMP_IND = 0x6C
	JSR     = 0x20
	RTS     = 0x60
)
