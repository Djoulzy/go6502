package mem

import "go6502/globals"

const (
	memorySize  = 65536
	stackStart  = 0x0100
	stackEnd    = 0x01FF
	screenStart = 0x0400
	screenEnd   = 0x07FF
	colorStart  = 0xD800
	colorEnd    = 0xDBFF
	intAddr     = 0xFFFA
	resetAddr   = 0xFFFC
	brkAddr     = 0xFFFE

	KernalStart = 0xE000
	KernalEnd   = 0xFFFF
	BasicStart  = 0xA000
	BasicEnd    = 0xC000
)

// Memory :
type Memory struct {
	Data    [memorySize]globals.Byte
	Kernal  []globals.Byte
	Basic   []globals.Byte
	CharGen []globals.Byte
	Stack   []globals.Byte
	Screen  []globals.Byte
	Color   []globals.Byte
	Vic     [4][]globals.Byte
}
