package graphic

import (
	"go6502/globals"
)

type Driver interface {
	Init(int, int)
	SetPixel(int, globals.RGB)
	Draw8pixels(int, int, globals.RGB, globals.RGB, globals.Byte)
	DisplayFrame()
	CloseAll()
}
