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

	visibleFirstLine = 56
	visibleLastLine  = 255
	visibleFirstCol  = 11
	visibleLastCol   = 50
	visibleFirstRow  = 7
)

func (V *VIC) readScreenData(mem *mem.Memory, y int) {
	if (y >= visibleFirstLine) && (y <= visibleLastLine) {
		start := globals.Word(V.RowCounter-visibleFirstRow) * 40
		// log.Printf("Y: %d", V.RowCounter-visibleFirstRow)
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
	if (y >= visibleFirstLine) && (y <= visibleLastLine) {
		if (x >= visibleFirstCol) && (x <= visibleLastCol) {
			return true
		}
	}
	return false
}

func (V *VIC) drawgByte(beamX, beamY int) {
	if V.isVisibleArea(beamX, beamY) {
		charColor := globals.Byte(V.Buffer[beamX-visibleFirstCol] >> 8)
		charAddr := globals.Byte(V.Buffer[beamX-visibleFirstCol])
		charRomAddr := V.ram.CharGen[globals.Word(charAddr)<<3+globals.Word(V.BadLineCounter)]
		V.graph.Draw8pixels(beamX*8, beamY, charColor, Blue, charRomAddr)
	} else {
		V.graph.Draw8pixels(beamX*8, beamY, Lightblue, Lightblue, globals.Byte(0xFF))
	}
}

func (V *VIC) Init(mem *mem.Memory, cpuCycle chan bool) {
	V.cpuCycle = cpuCycle
	V.ram = mem
}

func (V *VIC) Run() {
	V.graph = graphic.SDLDriver{}

	V.graph.Init(winWidth, winHeight)
	defer func() {
		V.ram.Dump(0x0590)
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
			select {
			case <-ticker.C:
				if beamY > 15 && beamY < 300 {
					VBlank = false
					if V.BadLineCounter == 0 {
						V.readScreenData(V.ram, beamY)
					}
				} else {
					VBlank = true
				}

				for beamX := 0; beamX < cyclesPerLine; beamX++ {
					if beamX > 5 && beamX < 57 {
						HBlank = false
					} else {
						HBlank = true
					}

					if VBlank || HBlank {
						V.graph.Draw8pixels(beamX*8, beamY, Black, Red, globals.Byte(0xFF))
					} else {
						V.drawgByte(beamX, beamY)
					}
					V.cpuCycle <- true
				}
				V.BadLineCounter++
				if V.BadLineCounter == 8 {
					V.BadLineCounter = 0
					V.RowCounter++
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
