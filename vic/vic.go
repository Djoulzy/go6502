package vic

import (
	"fmt"
	"go6502/globals"
	"go6502/graphic"
	"go6502/mem"
	"time"
)

const (
	cpuClock        = 985248                         // Mesure en Hz
	cpuCycle        = (1 / float32(cpuClock)) * 1000 // 1 cycle en ms
	screenWidthPAL  = 504
	screenHeightPAL = 312
	rasterWidthPAL  = 403
	rasterHeightPAL = 284
	cyclesPerLine   = 63

	rasterTime = 1                  // Nb of cycle to put 1 globals.Byte on a line
	rasterLine = rasterWidthPAL / 8 // Nb of cycle to draw a full line
	fullRaster = rasterLine * rasterHeightPAL

	lineRefresh   = cyclesPerLine * cpuCycle                   // Time for a line in ms
	screenRefresh = screenHeightPAL * cyclesPerLine * cpuCycle // Time for a full screen display in ms
	fps           = 1 / screenRefresh

	winWidth      = screenWidthPAL
	winHeight     = screenHeightPAL
	visibleWidth  = 320
	visibleHeight = 200

	firstVBlankLine  = 300
	lastVBlankLine   = 15
	visibleFirstLine = 51
	visibleLastLine  = visibleFirstLine + visibleHeight

	firstHBlankCol  = 453
	lastHBlankCol   = 50
	visibleFirstCol = 92
	visibleLastCol  = 412
)

func (V *VIC) Init(mem *mem.Memory, cpuCycle chan bool) {
	V.cpuCycle = cpuCycle
	V.ram = mem
	V.ram.Data[REG_EC] = Lightblue
	V.ram.Data[REG_B0C] = Blue
}

func (V *VIC) readScreenData(mem *mem.Memory, y int) {
	if (y >= visibleFirstLine) && (y < visibleLastLine) {
		start := globals.Word(V.RowCounter) * 40
		for i := 0; i < 40; i++ {
			V.Buffer[i] = globals.Word(mem.Color[int(start)+i]) << 8
			V.Buffer[i] |= globals.Word(mem.Screen[int(start)+i])
		}
	}
}

func (V *VIC) getPixelColor(beamX int) globals.Byte {
	origin := beamX - visibleFirstCol
	col := origin >> 3
	bufferValue := V.Buffer[col]

	bit := globals.Byte(0b10000000 >> (origin % 8))
	charAddr := globals.Byte(bufferValue)
	charRomAddr := V.ram.CharGen[globals.Word(charAddr)<<3+globals.Word(V.BadLineCounter)]
	if charRomAddr&bit > 0 {
		return globals.Byte(bufferValue>>8) & 0b00001111
	}
	return V.ram.Data[REG_B0C] & 0b00001111
}

func (V *VIC) saveRasterPos(val int) {
	V.ram.Data[REG_RASTER] = globals.Byte(val)
	if (globals.Byte(globals.Word(val) >> 8)) == 0x1 {
		V.ram.Data[REG_RST8] |= 0b10000000
	} else {
		V.ram.Data[REG_RST8] &= 0b01111111
	}
	// fmt.Printf("val: %d - RST8: %08b - RASTER: %08b\n", val, V.ram.Data[REG_RST8], V.ram.Data[REG_RASTER])
}

func (V *VIC) Run() {
	var VBlank, HBlank, VisibleArea bool
	var pixelColor globals.RGB

	V.graph = &graphic.SDLDriver{}
	// V.graph = &graphic.SDL2Driver{}

	V.graph.Init(winWidth, winHeight)
	defer func() {
		V.graph.CloseAll()
	}()

	// cpuTimer, _ := time.ParseDuration(fmt.Sprintf("%fms", lineRefresh))
	cpuTimer, _ := time.ParseDuration(fmt.Sprintf("%fms", 0.05))

	// fmt.Printf("cpuTimer %v.\n", cpuTimer)
	ticker := time.NewTicker(cpuTimer)
	defer func() {
		ticker.Stop()
	}()

	for {
		HBlank = true
		VBlank = true
		VisibleArea = false
		V.BadLineCounter = 0
		V.RowCounter = 0

		t0 := time.Now()
		for beamY := 0; beamY < screenHeightPAL; beamY++ {

			V.saveRasterPos(beamY)
			// // fmt.Printf("raster: %d - BadLineCounter: %d - RowCounter: %d\n", beamY, V.BadLineCounter, V.RowCounter)
			select {
			case <-ticker.C:
			if beamY > lastVBlankLine && beamY < firstVBlankLine {
				VBlank = false
				if beamY >= visibleFirstLine && beamY < visibleLastLine {
					VisibleArea = true
					if V.BadLineCounter == 0 {
						V.readScreenData(V.ram, beamY)
					}
				} else {
					VisibleArea = false
				}
			} else {
				VBlank = true
			}

			beamX := 0
			for cycle := 0; cycle < cyclesPerLine; cycle++ {
				if beamX >= lastHBlankCol && beamX < firstHBlankCol {
					HBlank = false
				} else {
					HBlank = true
				}
				for column := 0; column < 8; column++ {
					if VBlank || HBlank {
						pixelColor = Colors[Black]
					} else {
						if beamX >= visibleFirstCol && beamX < visibleLastCol && VisibleArea {
							pixelColor = Colors[V.getPixelColor(beamX)]
						} else {
							pixelColor = Colors[V.ram.Data[REG_EC]&0b00001111]
						}
					}
					V.graph.DrawPixel(beamX, beamY, pixelColor)
					beamX++
				}
				// V.cpuCycle <- true
				V.ram.WaitFor(false)
			}
			if beamY >= visibleFirstLine && beamY < visibleLastLine {
				V.BadLineCounter++
				if V.BadLineCounter == 8 {
					V.BadLineCounter = 0
					V.RowCounter++
				}
			}

			}
		}
		// setPixel(visibleFirstCol*8, visibleFirstLine, White)
		// setPixel(visibleLastCol*8, visibleLastLine, White)
		V.graph.DisplayFrame()
		// os.Exit(1)
		t1 := time.Now()
		fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
	}
}
