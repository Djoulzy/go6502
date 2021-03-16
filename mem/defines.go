package mem

import (
	"go6502/globals"
	"sync"
)

const (
	memorySize  = globals.Word(0xFFFF)
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

// Memory :
type Memory struct {
	Data    [memorySize]globals.Byte
	CharGen [4096]globals.Byte
	Stack   []globals.Byte
	Screen  []globals.Byte
	Color   []globals.Byte
	Vic     [4][]globals.Byte
	Access  bool
	mu      sync.RWMutex
}
