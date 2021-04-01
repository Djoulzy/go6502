package graphic

import (
	"go6502/globals"
)

type Driver interface {
	Init(int, int)
	DrawPixel(int, int, globals.RGB)
	DisplayFrame()
	CloseAll()
}
