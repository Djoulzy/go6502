package main

const (
	memorySize  = 0xFFFF
	stackStart  = 0x0100
	stackEnd    = 0x01FF
	screenStart = 0x0400
	screenEnd   = 0x07FF
	colorStart  = 0xD800
	colorEnd    = 0xDBFF
	intAddr     = 0xFFFA
	resetAddr   = 0xFFFC
	brkAddr     = 0xFFFE
)

// Byte :
type Byte uint8

// Word :
type Word uint16

// Memory :
type Memory struct {
	Data    [memorySize]Byte
	CharGen [4096]Byte
	Stack   []Byte
	Screen  []Byte
	Color   []Byte
	Vic     [4][]Byte
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

type rgb struct {
	r byte
	g byte
	b byte
}

var (
	Black      Byte = 0
	White      Byte = 1
	Red        Byte = 2
	Cyan       Byte = 3
	Violet     Byte = 4
	Green      Byte = 5
	Blue       Byte = 6
	Yellow     Byte = 7
	Orange     Byte = 8
	Brown      Byte = 9
	Lightred   Byte = 10
	Darkgrey   Byte = 11
	Grey       Byte = 12
	Lightgreen Byte = 13
	Lightblue  Byte = 14
	Lightgrey  Byte = 15
)

var Colors [16]rgb = [16]rgb{
	{r: 0, g: 0, b: 0},
	{r: 255, g: 255, b: 255},
	{r: 137, g: 78, b: 67},
	{r: 170, g: 255, b: 238},
	{r: 204, g: 68, b: 204},
	{r: 0, g: 204, b: 85},
	{r: 67, g: 60, b: 165},
	{r: 238, g: 238, b: 119},
	{r: 221, g: 136, b: 85},
	{r: 102, g: 68, b: 0},
	{r: 255, g: 119, b: 119},
	{r: 51, g: 51, b: 51},
	{r: 119, g: 119, b: 119},
	{r: 170, g: 255, b: 102},
	{r: 132, g: 126, b: 216},
	{r: 187, g: 187, b: 187},
}

// CPU :
type CPU struct {
	PC Word
	SP Byte
	A  Byte
	X  Byte
	Y  Byte
	S  Byte

	opName  string
	exit    bool
	cycle   chan bool
	display bool
}

// VIC :
type VIC struct {
	Buffer         [40]Word
	BadLineCounter Byte
	RowCounter     Byte
}

// Nemonic :
var Nemonic map[Byte]func(*Memory)

// :
const (
	BRK = 0x00
	NOP = 0xEA

	LDA_IM  = 0xA9
	LDA_ZP  = 0xA5
	LDA_ZPX = 0xB5
	LDA_ABS = 0xAD
	LDA_ABX = 0xBD
	LDA_ABY = 0xB9
	LDA_INX = 0xA1
	LDA_INY = 0xB1

	LDX_IM  = 0xA2
	LDX_ZP  = 0xA6
	LDX_ZPY = 0xB6
	LDX_ABS = 0xAE
	LDX_ABY = 0xBE

	LDY_IM  = 0xA0
	LDY_ZP  = 0xA4
	LDY_ZPX = 0xB4
	LDY_ABS = 0xAC
	LDY_ABX = 0xBC

	STA_ZP  = 0x85
	STA_ZPX = 0x95
	STA_ABS = 0x8D
	STA_ABX = 0x9D
	STA_ABY = 0x99
	STA_INX = 0x81
	STA_INY = 0x91

	STX_ZP  = 0x86
	STX_ZPY = 0x96
	STX_ABS = 0x8E

	STY_ZP  = 0x84
	STY_ZPX = 0x94
	STY_ABS = 0x8

	AND_IM  = 0x29
	AND_ZP  = 0x25
	AND_ZPX = 0x35
	AND_ABS = 0x2D
	AND_ABX = 0x3D
	AND_ABY = 0x39
	AND_INX = 0x21
	AND_INY = 0x31

	TXS = 0x9A
	PHA = 0x48
	PLA = 0x68

	JMP_ABS = 0x4C
	JMP_IND = 0x6C
	JSR     = 0x20
	RTS     = 0x60
)
