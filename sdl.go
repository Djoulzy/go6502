package main

import (
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

var screen []byte
var elapsedTime float32

func setPixel(index int, c Byte) {
	screen[index] = Colors[c].r
	screen[index+1] = Colors[c].g
	screen[index+2] = Colors[c].b
}

func draw8pixels(x, y int, fg_color, bg_color, value Byte) {
	// t0 := time.Now()
	index := (y*winWidth + x) * 3
	for i := 0; i < 8; i++ {
		base := index + (i * 3)
		if value&(0x1<<(7-i)) > 0 {
			setPixel(base, fg_color)
		} else {
			setPixel(base, bg_color)
		}
	}
	// t1 := time.Now()
	// dur := int64(cpuCycle) - t1.Sub(t0).Milliseconds()
	// log.Printf("The call took %v to run.\n", t1.Sub(t0))
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
	window.SetResizable(true)
	window.SetSize(winWidth*2, winHeight*2)

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_RGB24, sdl.TEXTUREACCESS_STATIC, int32(winWidth), int32(winHeight))
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
