package vic

import (
	"go6502/graphic"
	"go6502/mem"
)

var (
	Black      byte = 0
	White      byte = 1
	Red        byte = 2
	Cyan       byte = 3
	Violet     byte = 4
	Green      byte = 5
	Blue       byte = 6
	Yellow     byte = 7
	Orange     byte = 8
	Brown      byte = 9
	Lightred   byte = 10
	Darkgrey   byte = 11
	Grey       byte = 12
	Lightgreen byte = 13
	Lightblue  byte = 14
	Lightgrey  byte = 15
)

var Colors [16]graphic.RGB = [16]graphic.RGB{
	{R: 0, G: 0, B: 0},       // Black
	{R: 255, G: 255, B: 255}, // White
	{R: 137, G: 78, B: 67},   // Red
	{R: 146, G: 195, B: 203}, // Cyan
	{R: 138, G: 87, B: 176},  // Violet
	{R: 128, G: 174, B: 89},  // Green
	{R: 68, G: 63, B: 164},   // Blue
	{R: 215, G: 221, B: 137}, // Yellow
	{R: 146, G: 106, B: 56},  // Orange
	{R: 100, G: 82, B: 23},   // Brown
	{R: 184, G: 132, B: 122}, // Lightred
	{R: 96, G: 96, B: 96},    // Darkgrey
	{R: 138, G: 138, B: 138}, // Grey
	{R: 191, G: 233, B: 155}, // Lightgreen
	{R: 131, G: 125, B: 216}, // Lightblue
	{R: 179, G: 179, B: 179}, // Lightgrey
}

// VIC :
type VIC struct {
	VML    [40]uint16 // Video Matrix Line
	VMLI   byte       // Video Matrix Line Indexer
	VC     uint16     // Vide Counter
	VCBASE uint16     // Video Counter Base
	RC     byte       // Row counter
	BA     bool       // High: normal / Low: BadLine

	beamX int
	beamY int
	cycle int

	visibleArea bool
	displayArea bool
	drawArea    bool

	ColorBuffer [40]byte
	CharBuffer  [40]byte

	graph graphic.Driver
	ram   *mem.Memory
}

const (
	REG_RST8        = 0xD011 // Raster 9eme bit of Raster (0b10000000)
	REG_SCR_CONTROL = 0xD011 // Screen control (0b01111111)
	REG_RASTER      = 0xD012 // Raster 8 first bits
	REG_EC          = 0xD020 // Border Color
	REG_B0C         = 0xD021 // Background color 0
	PALNTSC         = 0x02A6
)
