package cpu

import (
	"go6502/databus"
	"go6502/mem"

	"github.com/mattn/go-tty"
)

//
const (
	C_mask byte = 0b11111110
	Z_mask byte = 0b11111101
	I_mask byte = 0b11111011
	D_mask byte = 0b11110111
	B_mask byte = 0b11101111

	V_mask byte = 0b10111111
	N_mask byte = 0b01111111
)

// CPU :
type CPU struct {
	PC uint16
	SP byte
	A  byte
	X  byte
	Y  byte
	S  byte

	IRQ     int
	opName  string
	debug   string
	exit    bool
	Display bool
	ram     *mem.Memory
	dbus    *databus.Bus
	BP      uint16
	Step    bool
	Dump    uint16
	Zone    int
	tty     *tty.TTY
}

// Mnemonic :
var Mnemonic map[byte]func(*mem.Memory)
var CodeAddr = map[string]byte{
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
	"DEC_ZP":  0xC6,
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

	"BCC_REL": 0x90,
	"BCS_REL": 0xB0,
	"BEQ_REL": 0xF0,
	"BMI_REL": 0x30,
	"BNE_REL": 0xD0,
	"BPL_REL": 0x10,
	"BVC_REL": 0x50,
	"BVS_REL": 0x70,

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

	"STY_ZP":  0x84,
	"STY_ZPX": 0x94,
	"STY_ABS": 0x8C,

	"AND_IM":  0x29,
	"AND_ZP":  0x25,
	"AND_ZPX": 0x35,
	"AND_ABS": 0x2D,
	"AND_ABX": 0x3D,
	"AND_ABY": 0x39,
	"AND_INX": 0x21,
	"AND_INY": 0x31,

	"EOR_IM":  0x49,
	"EOR_ZP":  0x45,
	"EOR_ZPX": 0x55,
	"EOR_ABS": 0x4D,
	"EOR_ABX": 0x5D,
	"EOR_ABY": 0x59,
	"EOR_INX": 0x41,
	"EOR_INY": 0x51,

	"ORA_IM":  0x09,
	"ORA_ZP":  0x05,
	"ORA_ZPX": 0x15,
	"ORA_ABS": 0x0D,
	"ORA_ABX": 0x1D,
	"ORA_ABY": 0x19,
	"ORA_INX": 0x01,
	"ORA_INY": 0x11,

	"BIT_ZP":  0x24,
	"BIT_ABS": 0x2C,

	"TXS": 0x9A,
	"TSX": 0xBA,
	"PHA": 0x48,
	"PHP": 0x08,
	"PLA": 0x68,
	"PLP": 0x28,

	"TAX": 0xAA,
	"TAY": 0xA8,
	"TXA": 0x8A,
	"TYA": 0x98,

	"JMP_ABS": 0x4C,
	"JMP_IND": 0x6C,
	"JSR":     0x20,
	"RTS":     0x60,
	"RTI":     0x40,

	"CLC": 0x18,
	"CLD": 0xD8,
	"CLI": 0x58,
	"CLV": 0xB8,
	"SEC": 0x38,
	"SED": 0xF8,
	"SEI": 0x78,

	"ASL_IM":  0x0A,
	"ASL_ZP":  0x06,
	"ASL_ZPX": 0x16,
	"ASL_ABS": 0x0E,
	"ASL_ABX": 0x1E,

	"LSR_IM":  0x4A,
	"LSR_ZP":  0x46,
	"LSR_ZPX": 0x56,
	"LSR_ABS": 0x4E,
	"LSR_ABX": 0x5E,

	"ROL_IM":  0x2A,
	"ROL_ZP":  0x26,
	"ROL_ZPX": 0x36,
	"ROL_ABS": 0x2E,
	"ROL_ABX": 0x3E,

	"ROR_IM":  0x6A,
	"ROR_ZP":  0x66,
	"ROR_ZPX": 0x76,
	"ROR_ABS": 0x6E,
	"ROR_ABX": 0x7E,
}
