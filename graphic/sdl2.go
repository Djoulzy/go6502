package graphic

import (
	"go6502/globals"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

type SDL2Driver struct {
	winHeight int
	winWidth  int
	window    *sdl.Window
	renderer  *sdl.Renderer
	texture   *sdl.Texture
	screen    []byte
}

func (S *SDL2Driver) SetPixel(index int, c globals.RGB) {
}

func (S *SDL2Driver) Draw8pixels(x, y int, fg_color, bg_color globals.RGB, value globals.Byte) {
	for i := 0; i < 8; i++ {
		if value&(0x1<<(7-i)) > 0 {
			S.renderer.SetDrawColor(uint8(fg_color.R), uint8(fg_color.G), uint8(fg_color.B), 255)
		} else {
			S.renderer.SetDrawColor(uint8(bg_color.R), uint8(bg_color.G), uint8(bg_color.B), 255)
		}
		S.renderer.DrawPoint(int32(x+i), int32(y))
	}
}

func (S *SDL2Driver) CloseAll() {
	S.window.Destroy()
	S.renderer.Destroy()
	S.texture.Destroy()
	sdl.Quit()
}

func (S *SDL2Driver) Init(winWidth, winHeight int) {
	S.winHeight = winHeight
	S.winWidth = winWidth

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}

	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")

	S.window, err = sdl.CreateWindow("VIC-II", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(S.winWidth), int32(S.winHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	S.window.SetResizable(true)
	S.window.SetSize(int32(S.winWidth*2), int32(S.winHeight*2))

	S.renderer, err = sdl.CreateRenderer(S.window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	// S.texture, err = S.renderer.CreateTexture(sdl.PIXELFORMAT_RGB24, sdl.TEXTUREACCESS_STATIC, int32(S.winWidth), int32(S.winHeight))
	// if err != nil {
	// 	panic(err)
	// }

	S.screen = make([]byte, S.winWidth*S.winHeight*3)
}

func (S *SDL2Driver) DisplayFrame() {
	S.renderer.Present()

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			os.Exit(1)
		}
	}
}
