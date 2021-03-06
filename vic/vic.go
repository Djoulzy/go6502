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
	visibleFirstLine = 51  // 56
	visibleLastLine  = 251 // 255
	// visibleFirstCol  = 11
	// visibleLastCol   = 50
	firstXcoord = 24
	lastXcoord  = 344

	visibleFirstCol = firstXcoord / 8
	visibleLastCol  = lastXcoord / 8
)

func (V *VIC) readScreenData(mem *mem.Memory, y int) {
	if (y >= visibleFirstLine) && (y < visibleLastLine) {
		start := globals.Word(V.RowCounter) * 40
		for i := 0; i < 40; i++ {
			// log.Printf("X: %d Y: %d", i, start)
			V.Buffer[i] = globals.Word(mem.Color[int(start)+i]) << 8
			V.Buffer[i] |= globals.Word(mem.Screen[int(start)+i])
			// log.Printf("Mem Color: %d, Value: %x", start+i, mem.Color[start+i])
			// log.Printf("Mem Screen: %d, Value: %x", start+i, mem.Screen[start+i])
			// log.Printf("Buffer: %d, Value: %x", i, V.Buffer[i])
		}
	}
}

func (V *VIC) isVisibleArea(x, y int) bool {
	if (y >= visibleFirstLine) && (y < visibleLastLine) {
		if (x >= visibleFirstCol) && (x <= visibleLastCol) {
			return true
		}
	}
	return false
}

func (V *VIC) drawByte(beamX, beamY int) {
	if V.isVisibleArea(beamX, beamY) {
		charColor := globals.Byte(V.Buffer[beamX-visibleFirstCol] >> 8)
		charAddr := globals.Byte(V.Buffer[beamX-visibleFirstCol])
		charRomAddr := V.ram.CharGen[globals.Word(charAddr)<<3+globals.Word(V.BadLineCounter)]
		V.graph.Draw8pixels(beamX*8, beamY, Colors[charColor], Colors[Blue], charRomAddr)
	} else {
		color := V.ram.Data[REG_EC]
		V.graph.Draw8pixels(beamX*8, beamY, Colors[color], Colors[color], globals.Byte(0xFF))
	}
}

func (V *VIC) Init(mem *mem.Memory, cpuCycle chan bool) {
	V.cpuCycle = cpuCycle
	V.ram = mem
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
	V.graph = &graphic.SDLDriver{}
	// V.graph = &graphic.SDL2Driver{}

	V.graph.Init(winWidth, winHeight)
	defer func() {
		V.graph.CloseAll()
	}()

	cpuTimer, _ := time.ParseDuration(fmt.Sprintf("%fms", lineRefresh))
	// fmt.Printf("cpuTimer %v.\n", cpuTimer)
	ticker := time.NewTicker(cpuTimer)
	defer func() {
		ticker.Stop()
	}()

	for {
		HBlank := true
		VBlank := true
		V.BadLineCounter = 0
		V.RowCounter = 0

		// t0 := time.Now()
		for beamY := 0; beamY < screenHeightPAL; beamY++ {
			V.saveRasterPos(beamY)
			// fmt.Printf("raster: %d - BadLineCounter: %d - RowCounter: %d\n", beamY, V.BadLineCounter, V.RowCounter)
			select {
			case <-ticker.C:
				if beamY > lastVBlankLine && beamY < firstVBlankLine {
					VBlank = false
					if beamY >= visibleFirstLine && beamY < visibleLastLine {
						if V.BadLineCounter == 0 {
							V.readScreenData(V.ram, beamY)
						}
					}
				} else {
					VBlank = true
				}

				for beamX := 0; beamX < cyclesPerLine; beamX++ {
					if beamX >= visibleFirstCol && beamX < visibleLastCol {
						HBlank = false
					} else {
						HBlank = true
					}

					if VBlank || HBlank {
						V.graph.Draw8pixels(beamX*8, beamY, Colors[Black], Colors[Black], globals.Byte(0xFF))
					} else {
						V.drawByte(beamX, beamY)
					}
					V.cpuCycle <- true
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
		// t1 := time.Now()
		// fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
		// setPixel(visibleFirstCol*8, visibleFirstLine, White)
		// setPixel(visibleLastCol*8, visibleLastLine, White)
		V.graph.DisplayFrame()
		// os.Exit(1)
	}
}
