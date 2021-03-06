package vic

import (
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
	{R: 0, G: 0, B: 0},
	{R: 255, G: 255, B: 255},
	{R: 137, G: 78, B: 67},
	{R: 170, G: 255, B: 238},
	{R: 204, G: 68, B: 204},
	{R: 0, G: 204, B: 85},
	{R: 67, G: 60, B: 165},
	{R: 238, G: 238, B: 119},
	{R: 221, G: 136, B: 85},
	{R: 102, G: 68, B: 0},
	{R: 255, G: 119, B: 119},
	{R: 51, G: 51, B: 51},
	{R: 119, G: 119, B: 119},
	{R: 170, G: 255, B: 102},
	{R: 132, G: 126, B: 216},
	{R: 187, G: 187, B: 187},
}

// VIC :
type VIC struct {
	Buffer         [40]globals.Word
	BadLineCounter globals.Byte
	RowCounter     globals.Byte
	graph          graphic.Driver
	cpuCycle       chan bool
	ram            *mem.Memory
}

const (
	REG_EC     = globals.Word(0xD020) // Border Color
	REG_RASTER = globals.Word(0xD012) // Raster 8 first bits
	REG_RST8   = globals.Word(0xD011) // Raster 9eme bit
)
