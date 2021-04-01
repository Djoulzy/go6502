package vic

import (
	"go6502/databus"
	"go6502/globals"
	"go6502/graphic"
	"go6502/mem"
)

var (
	Black      globals.Byte = 0
	White      globals.Byte = 1
	Red        globals.Byte = 2
	Cyan       globals.Byte = 3
	Violet     globals.Byte = 4
	Green      globals.Byte = 5
	Blue       globals.Byte = 6
	Yellow     globals.Byte = 7
	Orange     globals.Byte = 8
	Brown      globals.Byte = 9
	Lightred   globals.Byte = 10
	Darkgrey   globals.Byte = 11
	Grey       globals.Byte = 12
	Lightgreen globals.Byte = 13
	Lightblue  globals.Byte = 14
	Lightgrey  globals.Byte = 15
)

var Colors [16]globals.RGB = [16]globals.RGB{
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
	VML    [40]globals.Word // Video Matrix Line
	VMLI   globals.Byte     // Video Matrix Line Indexer
	VC     globals.Word     // Vide Counter
	VCBASE globals.Word     // Video Counter Base
	RC     globals.Byte     // Row counter
	BA     bool             // High: normal / Low: BadLine

	visibleArea bool
	displayArea bool
	drawArea    bool

	ColorBuffer [40]globals.Byte
	CharBuffer  [40]globals.Byte

	graph     graphic.Driver
	ram       *mem.Memory
	dbus      *databus.Databus
}

const (
	REG_RST8   = globals.Word(0xD011) // Raster 9eme bit
	REG_RASTER = globals.Word(0xD012) // Raster 8 first bits
	REG_EC     = globals.Word(0xD020) // Border Color
	REG_B0C    = globals.Word(0xD021) // Background color 0
)
