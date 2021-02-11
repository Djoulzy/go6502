package main

import (
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

type rgb struct {
	r byte
	g byte
	b byte
}

var (
	Black      = rgb{r: 0, g: 0, b: 0}
	White      = rgb{r: 255, g: 255, b: 255}
	Red        = rgb{r: 137, g: 78, b: 67}
	Cyan       = rgb{r: 170, g: 255, b: 238}
	Violet     = rgb{r: 204, g: 68, b: 204}
	Green      = rgb{r: 0, g: 204, b: 85}
	Blue       = rgb{r: 67, g: 60, b: 165}
	Yellow     = rgb{r: 238, g: 238, b: 119}
	Orange     = rgb{r: 221, g: 136, b: 85}
	Brown      = rgb{r: 102, g: 68, b: 0}
	Lightred   = rgb{r: 255, g: 119, b: 119}
	Darkgrey   = rgb{r: 51, g: 51, b: 51}
	Grey       = rgb{r: 119, g: 119, b: 119}
	Lightgreen = rgb{r: 170, g: 255, b: 102}
	Lightblue  = rgb{r: 132, g: 126, b: 216}
	Lightgrey  = rgb{r: 187, g: 187, b: 187}
)

var screen []byte
var elapsedTime float32

func setPixel(x, y int, c rgb) {
	index := (y*winWidth + x) * 3

	if index < len(screen)-3 && index >= 0 {
		screen[index] = c.r
		screen[index+1] = c.g
		screen[index+2] = c.b
	}
}

func closeAll(win *sdl.Window, rend *sdl.Renderer, tex *sdl.Texture) {
	win.Destroy()
	rend.Destroy()
	tex.Destroy()
	sdl.Quit()
}

func initSDL() (*sdl.Window, *sdl.Renderer, *sdl.Texture) {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}

	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")

	window, err := sdl.CreateWindow("VIC-II", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_RGB24, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	if err != nil {
		panic(err)
	}

	screen = make([]byte, winWidth*winHeight*3)

	return window, renderer, tex
}

func displayFrame(rend *sdl.Renderer, tex *sdl.Texture) {

	tex.Update(nil, screen, winWidth*3)
	rend.Copy(tex, nil, nil)
	rend.Present()

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			os.Exit(1)
		}
	}

}
