package main

const (
	cpuClock        = 985248
	cpuCycle        = 1 / float32(cpuClock)
	screenWidthPAL  = 504
	screenHeightPAL = 312
	rasterWidthPAL  = 403
	rasterHeightPAL = 284

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
	visibleLastLine  = 250
	visibleFirstCol  = 50 + 41
	visibleLastCol   = visibleFirstCol + 320
)

func isVisibleArea(x, y int) bool {
	if (y >= visibleFirstLine) && (y <= visibleLastLine) {
		if (x >= visibleFirstCol) && (x <= visibleLastCol) {
			return true
		}
	}
	return false
}

func main() {
	win, rend, tex := initSDL()
	defer closeAll(win, rend, tex)

	for {
		HBlank := true
		VBlank := true
		for beamY := 0; beamY < screenHeightPAL; beamY++ {
			if beamY > 15 && beamY < 300 {
				VBlank = false
			} else {
				VBlank = true
			}
			charCpt := 0
			for beamX := 0; beamX < screenWidthPAL; beamX++ {
				if beamX > 50 && beamX < 453 {
					HBlank = false
				} else {
					HBlank = true
				}
				if VBlank || HBlank {
					setPixel(beamX, beamY, Black)
				} else {
					if isVisibleArea(beamX, beamY) {
						if charCpt == 7 {
							setPixel(beamX, beamY, Black)
						} else {
							setPixel(beamX, beamY, Blue)
						}
						charCpt++
					} else {
						setPixel(beamX, beamY, Lightblue)
					}

					if charCpt == 8 {
						charCpt = 0
					}
				}
			}
		}
		//setPixel(beamX, beamY, Red)
		displayFrame(rend, tex)
	}
}
