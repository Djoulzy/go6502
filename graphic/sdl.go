package graphic

import (
	"go6502/globals"
	"go6502/vic"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

func (S *SDLDriver) SetPixel(index int, c globals.Byte) {
	S.screen[index] = vic.Colors[c].r
	S.screen[index+1] = vic.Colors[c].g
	S.screen[index+2] = vic.Colors[c].b
}

func (S *SDLDriver) Draw8pixels(x, y int, fg_color, bg_color, value globals.Byte) {
	// t0 := time.Now()
	index := (y*S.winWidth + x) * 3
	for i := 0; i < 8; i++ {
		base := index + (i * 3)
		if value&(0x1<<(7-i)) > 0 {
			S.SetPixel(base, fg_color)
		} else {
			S.SetPixel(base, bg_color)
		}
	}
	// t1 := time.Now()
	// dur := int64(cpuCycle) - t1.Sub(t0).Milliseconds()
	// log.Printf("The call took %v to run.\n", t1.Sub(t0))
}

func (S *SDLDriver) CloseAll() {
	S.window.Destroy()
	S.renderer.Destroy()
	S.texture.Destroy()
	sdl.Quit()
}

func (S *SDLDriver) Init(winWidth, winHeight int) {
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

	S.texture, err = S.renderer.CreateTexture(sdl.PIXELFORMAT_RGB24, sdl.TEXTUREACCESS_STATIC, int32(S.winWidth), int32(S.winHeight))
	if err != nil {
		panic(err)
	}

	S.screen = make([]byte, S.winWidth*S.winHeight*3)
}

func (S *SDLDriver) DisplayFrame() {

	S.texture.Update(nil, S.screen, S.winWidth*3)
	S.renderer.Copy(S.texture, nil, nil)
	S.renderer.Present()

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			os.Exit(1)
		}
	}

}
