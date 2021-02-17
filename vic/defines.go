package vic

import (
	"go6502/globals"
	"go6502/graphic"
	"go6502/mem"
)

type rgb struct {
	r globals.Byte
	g globals.Byte
	b globals.Byte
}

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

// VIC :
type VIC struct {
	Buffer         [40]globals.Word
	BadLineCounter globals.Byte
	RowCounter     globals.Byte
	graph          graphic.Driver
	cpuCycle       chan bool
	ram            *mem.Memory
}
