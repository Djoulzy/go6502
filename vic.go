package main

import (
	"log"
	"os"
)

const (
	cpuClock        = 985248
	cpuCycle        = 1 / float32(cpuClock)
	screenWidthPAL  = 504
	screenHeightPAL = 312
	rasterWidthPAL  = 403
	rasterHeightPAL = 284
	cyclesPerLine   = 63

	rasterTime = 1                  // Nb of cycle to put 1 byte on a line
	rasterLine = rasterWidthPAL / 8 // Nb of cycle to draw a full line
	fullRaster = rasterLine * rasterHeightPAL

	screenRefresh = fullRaster * cpuCycle // Time for a full screen display in ms
	fps           = 1 / screenRefresh

	winWidth      = screenWidthPAL
	winHeight     = screenHeightPAL
	visibleWidth  = 320
	visibleHeight = 200

	visibleFirstLine = 55
	visibleLastLine  = 255
	visibleFirstCol  = 11
	visibleLastCol   = 50
)

func (V *VIC) readScreenData(mem *Memory) {
	start := int(V.RowCounter) * 40
	for i := 0; i < 40; i++ {
		V.Buffer[i] = Word(mem.Color[start+i]) << 8
		V.Buffer[i] &= Word(mem.Screen[start+i])
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
		charRomAddr := V.Buffer[beamX-visibleFirstCol]<<3 + V.BadLineCounter
		for i := 0; i < 8; i++ {
			if mem.CharGen[charRomAddr]&(0x1<<i) > 0 {
				setPixel(beamX*8+i, beamY, Black)
			} else {
				setPixel(beamX*8+i, beamY, Blue)
			}
		}
	} else {
		for i := 0; i < 8; i++ {
			setPixel(beamX*8+i, beamY, Lightblue)
		}
	}
}

func (V *VIC) CheckForBadLines(y int) {

	if (y >= visibleFirstLine) && (y <= visibleLastLine) {
		log.Printf("Line : %d", V.BadLineCounter)
		if V.BadLineCounter == 8 {
			V.BadLineCounter = 0
			V.RowCounter++
		}

		if V.BadLineCounter == 0 {
			log.Printf("Bad Line")
		}

	}
}

func (V *VIC) run(mem *Memory) {
	win, rend, tex := initSDL()
	defer closeAll(win, rend, tex)

	var codeA Word
	codeA = 0x0043
	for i := 0; i < 40; i++ {
		V.Buffer[i] = codeA
	}

	for {
		HBlank := true
		VBlank := true
		V.BadLineCounter = 0
		V.RowCounter = 0
		for beamY := 0; beamY < screenHeightPAL; beamY++ {
			// log.Printf("Line : %d", V.BadLineCounter)
			if beamY > 15 && beamY < 300 {
				VBlank = false
				V.CheckForBadLines(beamY)
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
					for i := 0; i < 8; i++ {
						setPixel(beamX*8+i, beamY, Black)
					}
				} else {
					V.drawByte(mem, beamX, beamY)
				}
			}
			if VBlank == false {
				V.BadLineCounter++
			}
		}
		// setPixel(visibleFirstCol*8, visibleFirstLine, White)
		// setPixel(visibleLastCol*8, visibleLastLine, White)
		displayFrame(rend, tex)
		os.Exit(1)
	}

}
