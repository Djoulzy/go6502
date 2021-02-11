package main

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

	visibleFirstLine = 51
	visibleLastLine  = 251
	visibleFirstCol  = 11
	visibleLastCol   = 51
)

func (V *VIC) isVisibleArea(x, y int) bool {
	if (y > visibleFirstLine) && (y < visibleLastLine) {
		if (x >= visibleFirstCol) && (x <= visibleLastCol) {
			return true
		}
	}
	return false
}

func (V *VIC) drawByte(beamX, beamY int) {
	if V.isVisibleArea(beamX, beamY) {
		for i := 0; i < 8; i++ {
			setPixel(beamX*8+i, beamY, Blue)
		}
	} else {
		for i := 0; i < 8; i++ {
			setPixel(beamX*8+i, beamY, Lightblue)
		}
	}
}

func (V *VIC) run(mem *Memory) {
	win, rend, tex := initSDL()
	defer closeAll(win, rend, tex)

	var codeA Word
	codeA = 0x0041
	for i := 0; i < 40; i++ {
		V.Buffer[i] = codeA
	}

	for {
		HBlank := true
		VBlank := true
		for beamY := 0; beamY < screenHeightPAL; beamY++ {
			if beamY > 15 && beamY < 300 {
				VBlank = false
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
					V.drawByte(beamX, beamY)
				}
			}
		}
		//setPixel(beamX, beamY, Red)
		displayFrame(rend, tex)
	}
}
