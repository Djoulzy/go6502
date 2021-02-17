package graphic

import (
	"go6502/globals"

	"github.com/veandco/go-sdl2/sdl"
)

type Driver interface {
	Init(int, int)
	SetPixel(int, globals.RGB)
	Draw8pixels(int, int, globals.RGB, globals.RGB, globals.Byte)
	DisplayFrame()
	CloseAll()
}

type SDLDriver struct {
	winHeight int
	winWidth  int
	window    *sdl.Window
	renderer  *sdl.Renderer
	texture   *sdl.Texture
	screen    []byte
}
