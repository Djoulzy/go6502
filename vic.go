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

	visibleFirstLine = 55
	visibleLastLine  = 255
	visibleFirstCol  = 11
	visibleLastCol   = 50
)

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
		xPos := beamX - visibleFirstCol
		// log.Printf("%v", beamX)
		charRomAddr := int(V.Buffer[xPos])*8 + V.LineCounter
		data := byte(mem.CharGen[charRomAddr])

		for i := 0; i < 8; i++ {
			shift := data & byte(0x1<<i)
			if shift > 0 {
				setPixel(beamX*8+i, beamY, Black)
			} else {
				setPixel(beamX*8+i, beamY, Blue)
			}
		}
		// setPixel(beamX*8, beamY, Black)
		// setPixel(beamX*8+7, beamY, Black)
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
	codeA = 0x0001
	for i := 0; i < 40; i++ {
		V.Buffer[i] = codeA
	}

	for {
		HBlank := true
		VBlank := true
		V.LineCounter = 0
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
					V.drawByte(mem, beamX, beamY)
				}
			}
			V.LineCounter++
			if V.LineCounter == 8 {
				V.LineCounter = 0
			}
		}
		setPixel(visibleFirstCol*8, visibleFirstLine, White)
		setPixel(visibleLastCol*8, visibleLastLine, White)
		displayFrame(rend, tex)
	}
}
