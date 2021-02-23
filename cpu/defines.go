package cpu

import (
	"go6502/globals"
	"go6502/mem"
)

//
const (
	C_mask globals.Byte = 0b11111110
	Z_mask globals.Byte = 0b11111101
	I_mask globals.Byte = 0b11111011
	D_mask globals.Byte = 0b11110111
	B_mask globals.Byte = 0b11101111

	V_mask globals.Byte = 0b10111111
	N_mask globals.Byte = 0b01111111
)

// CPU :
type CPU struct {
	PC globals.Word
	SP globals.Byte
	A  globals.Byte
	X  globals.Byte
	Y  globals.Byte
	S  globals.Byte

	opName  string
	exit    bool
	Cycle   chan bool
	Display bool
	ram     *mem.Memory
}

// Mnemonic :
var Mnemonic map[globals.Byte]func(*mem.Memory)
var CodeAddr = map[string]globals.Byte{
	"SHW": 0xEF,
	"DMP": 0xFF,
	"BRK": 0x00,
	"NOP": 0xEA,

	"INC_ZP":  0xE6,
	"INC_ZPX": 0xF6,
	"INC_ABS": 0xEE,
	"INC_ABX": 0xFE,
	"INX":     0xE8,
	"INY":     0xC8,
	"DEC_ZP ": 0xC6,
	"DEC_ZPX": 0xD6,
	"DEC_ABS": 0xCE,
	"DEC_ABX": 0xDE,
	"DEX":     0xCA,
	"DEY":     0x88,

	"ADC_IM":  0x69,
	"ADC_ZP":  0x65,
	"ADC_ZPX": 0x75,
	"ADC_ABS": 0x6D,
	"ADC_ABX": 0x7D,
	"ADC_ABY": 0x79,
	"ADC_INX": 0x61,
	"ADC_INY": 0x71,

	"SBC_IM":  0xE9,
	"SBC_ZP":  0xE5,
	"SBC_ZPX": 0xF5,
	"SBC_ABS": 0xED,
	"SBC_ABX": 0xFD,
	"SBC_ABY": 0xF9,
	"SBC_INX": 0xE1,
	"SBC_INY": 0xF1,

	"CMP_IM":  0xC9,
	"CMP_ZP":  0xC5,
	"CMP_ZPX": 0xD5,
	"CMP_ABS": 0xCD,
	"CMP_ABX": 0xDD,
	"CMP_ABY": 0xD9,
	"CMP_INX": 0xC1,
	"CMP_INY": 0xD1,

	"CPX_IM":  0xE0,
	"CPX_ZP":  0xE4,
	"CPX_ABS": 0xEC,

	"CPY_IM":  0xC0,
	"CPY_ZP":  0xC4,
	"CPY_ABS": 0xCC,

	"BCC": 0x90,
	"BCS": 0xB0,
	"BEQ": 0xF0,
	"BMI": 0x30,
	"BNE": 0xD0,
	"BPL": 0x10,
	"BVC": 0x50,
	"BVS": 0x70,

	"LDA_IM":  0xA9,
	"LDA_ZP":  0xA5,
	"LDA_ZPX": 0xB5,
	"LDA_ABS": 0xAD,
	"LDA_ABX": 0xBD,
	"LDA_ABY": 0xB9,
	"LDA_INX": 0xA1,
	"LDA_INY": 0xB1,

	"LDX_IM":  0xA2,
	"LDX_ZP":  0xA6,
	"LDX_ZPY": 0xB6,
	"LDX_ABS": 0xAE,
	"LDX_ABY": 0xBE,

	"LDY_IM":  0xA0,
	"LDY_ZP":  0xA4,
	"LDY_ZPX": 0xB4,
	"LDY_ABS": 0xAC,
	"LDY_ABX": 0xBC,

	"STA_ZP":  0x85,
	"STA_ZPX": 0x95,
	"STA_ABS": 0x8D,
	"STA_ABX": 0x9D,
	"STA_ABY": 0x99,
	"STA_INX": 0x81,
	"STA_INY": 0x91,

	"STX_ZP":  0x86,
	"STX_ZPY": 0x96,
	"STX_ABS": 0x8E,

	"STY_ZP ": 0x84,
	"STY_ZPX": 0x94,
	"STY_ABS": 0x8C,

	"AND_IM ": 0x29,
	"AND_ZP ": 0x25,
	"AND_ZPX": 0x35,
	"AND_ABS": 0x2D,
	"AND_ABX": 0x3D,
	"AND_ABY": 0x39,
	"AND_INX": 0x21,
	"AND_INY": 0x31,

	"TXS": 0x9A,
	"PHA": 0x48,
	"PLA": 0x68,

	"TAX": 0xAA,
	"TAY": 0xA8,
	"TXA": 0x8A,
	"TYA": 0x98,

	"JMP_ABS": 0x4C,
	"JMP_IND": 0x6C,
	"JSR":     0x20,
	"RTS":     0x60,

	"CLC": 0x18,
	"CLD": 0xD8,
	"CLI": 0x58,
	"CLV": 0xB8,
	"SEC": 0x38,
	"SED": 0xF8,
	"SEI": 0x78,
}
