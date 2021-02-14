package main

import (
	"fmt"
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

	rasterTime = 1                  // Nb of cycle to put 1 byte on a line
	rasterLine = rasterWidthPAL / 8 // Nb of cycle to draw a full line
	fullRaster = rasterLine * rasterHeightPAL

	lineRefresh   = cyclesPerLine * cpuCycle                   // Time for a line in ms
	screenRefresh = screenHeightPAL * cyclesPerLine * cpuCycle // Time for a full screen display in ms
	fps           = 1 / screenRefresh

	winWidth      = screenWidthPAL
	winHeight     = screenHeightPAL
	visibleWidth  = 320
	visibleHeight = 200

	visibleFirstLine = 55
	visibleLastLine  = 254
	visibleFirstCol  = 11
	visibleLastCol   = 50
)

func (V *VIC) readScreenData(mem *Memory) {
	start := int(V.RowCounter) * 40
	for i := 0; i < 40; i++ {
		V.Buffer[i] = Word(mem.Color[start+i]) << 8
		V.Buffer[i] |= Word(mem.Screen[start+i])
		// log.Printf("Mem Color: %d, Value: %x", start+i, mem.Color[start+i])
		// log.Printf("Mem Screen: %d, Value: %x", start+i, mem.Screen[start+i])
		// log.Printf("Buffer: %d, Value: %x", i, V.Buffer[i])
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

func (V *VIC) drawByte(mem *Memory, beamX, beamY int) {
	if V.isVisibleArea(beamX, beamY) {
		charColor := Byte(V.Buffer[beamX-visibleFirstCol] >> 8)
		charAddr := Byte(V.Buffer[beamX-visibleFirstCol])
		charRomAddr := mem.CharGen[charAddr<<3+V.BadLineCounter]
		draw8pixels(beamX*8, beamY, charColor, Blue, charRomAddr)
	} else {
		draw8pixels(beamX*8, beamY, Lightblue, Lightblue, Byte(0xFF))
	}
}

func (V *VIC) CheckForBadLines(y int, mem *Memory) {
	if (y >= visibleFirstLine) && (y <= visibleLastLine) {
		V.BadLineCounter++
		if V.BadLineCounter == 0 {
			V.readScreenData(mem)
		}

		if V.BadLineCounter == 8 {
			V.BadLineCounter = 0
			V.RowCounter++
		}
	}
	// log.Printf("Line : %d BadLineCounter %d RowCounter : %d", y, V.BadLineCounter, V.RowCounter)
}

func (V *VIC) run(mem *Memory, cpuCycle chan bool) {
	win, rend, tex := initSDL()
	defer closeAll(win, rend, tex)

	cpuTimer, _ := time.ParseDuration(fmt.Sprintf("%fms", lineRefresh))
	// fmt.Printf("cpuTimer %v.\n", cpuTimer)
	ticker := time.NewTicker(cpuTimer)
	defer func() {
		ticker.Stop()
	}()

	for {
		HBlank := true
		VBlank := true
		V.BadLineCounter = 0xFF
		V.RowCounter = 0

		// t0 := time.Now()
		for beamY := 0; beamY < screenHeightPAL; beamY++ {
			select {
			case <-ticker.C:
				// log.Printf("Line : %d", V.BadLineCounter)
				if beamY > 15 && beamY < 300 {
					VBlank = false
					V.CheckForBadLines(beamY, mem)
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
						draw8pixels(beamX*8, beamY, Black, Red, Byte(0xFF))
					} else {
						V.drawByte(mem, beamX, beamY)
					}
					cpuCycle <- true
				}
			}
		}
		// t1 := time.Now()
		// fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
		// setPixel(visibleFirstCol*8, visibleFirstLine, White)
		// setPixel(visibleLastCol*8, visibleLastLine, White)
		displayFrame(rend, tex)
		// os.Exit(1)
	}

}
