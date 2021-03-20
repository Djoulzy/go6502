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
	firstDisplayLine = 51
	lastDisplayLine  = 250

	firstHBlankCol  = 453
	lastHBlankCol   = 50
	visibleFirstCol = 92
	visibleLastCol  = 412
)

func (V *VIC) Init(mem *mem.Memory, cpuCycle chan bool) {
	V.graph = &graphic.SDLDriver{}
	V.graph.Init(winWidth, winHeight)

	V.cpuCycle = cpuCycle
	V.ram = mem
	V.ram.Data[REG_EC] = Lightblue
	V.ram.Data[REG_B0C] = Blue
	V.BA = true
}

// func (V *VIC) readScreenData(mem *mem.Memory, y int) {
// 	if (y >= firstDisplayLine) && (y <= lastDisplayLine) {
// 		start := globals.Word(V.RowCounter) * 40
// 		for i := 0; i < 40; i++ {
// 			V.Buffer[i] = globals.Word(mem.Color[int(start)+i]) << 8
// 			V.Buffer[i] |= globals.Word(mem.Screen[int(start)+i])
// 		}
// 	}
// }

// func (V *VIC) readCharGen() {
// 	if V.displayArea {
// 		V.VC++
// 		V.VMLI++
// 	}
// }

// func (V *VIC) getPixelColor(beamX int) globals.Byte {
// 	origin := beamX - visibleFirstCol
// 	col := origin >> 3
// 	bufferValue := V.Buffer[col]

// 	bit := globals.Byte(0b10000000 >> (origin % 8))
// 	charAddr := globals.Byte(bufferValue)
// 	charRomAddr := V.ram.CharGen[globals.Word(charAddr)<<3+globals.Word(V.BadLineCounter)]
// 	if charRomAddr&bit > 0 {
// 		return globals.Byte(bufferValue>>8) & 0b00001111
// 	}
// 	return V.ram.Data[REG_B0C] & 0b00001111
// }

func (V *VIC) saveRasterPos(val int) {
	V.ram.Data[REG_RASTER] = globals.Byte(val)
	if (globals.Byte(globals.Word(val) >> 8)) == 0x1 {
		V.ram.Data[REG_RST8] |= 0b10000000
	} else {
		V.ram.Data[REG_RST8] &= 0b01111111
	}
	// fmt.Printf("val: %d - RST8: %08b - RASTER: %08b\n", val, V.ram.Data[REG_RST8], V.ram.Data[REG_RASTER])
}

func (V *VIC) readVideoMatrix() {
	if !V.BA {
		V.ColorBuffer[V.VMLI] = V.ram.Color[V.VC]
		V.CharBuffer[V.VMLI] = V.ram.Screen[V.VC]
	}
}

func (V *VIC) drawChar(X int, Y int) {
	if V.drawArea {
		charAddr := (globals.Word(V.CharBuffer[V.VMLI]) << 3) + globals.Word(V.RC)
		charData := V.ram.CharGen[charAddr]
		// fmt.Printf("SC: %02X - RC: %d - %04X - %02X = %08b\n", V.CharBuffer[V.VMLI], V.RC, charAddr, charData, charData)
		// if V.CharBuffer[V.VMLI] == 0 {
		// fmt.Printf("Raster: %d - Cycle: %d - BA: %t - VMLI: %d - VCBASE/VC: %d/%d - RC: %d - Char: %02X\n", Y, X, V.BA, V.VMLI, V.VCBASE, V.VC, V.RC, V.CharBuffer[V.VMLI])
		// }
		for column := 0; column < 8; column++ {
			bit := globals.Byte(0b10000000 >> column)
			if charData&bit > 0 {
				V.graph.DrawPixel(X+column, Y, Colors[V.ColorBuffer[V.VMLI]])
			} else {
				V.graph.DrawPixel(X+column, Y, Colors[V.ram.Data[REG_B0C]&0b00001111])
			}
		}
		V.VMLI++
		V.VC++
	} else if V.visibleArea {
		for column := 0; column < 8; column++ {
			V.graph.DrawPixel(X+column, Y, Colors[V.ram.Data[REG_EC]&0b00001111])
		}
	}
}

func (V *VIC) Run() {
	defer func() {
		V.graph.CloseAll()
	}()

	cpuTimer, _ := time.ParseDuration(fmt.Sprintf("%fms", lineRefresh))
	// fmt.Printf("cpuTimer %v.\n", cpuTimer)
	ticker := time.NewTicker(cpuTimer)
	defer func() {
		ticker.Stop()
	}()

	beamY := 0
	V.VCBASE = 0
	for {
		select {
		case <-ticker.C:
			V.saveRasterPos(beamY)
			V.visibleArea = (beamY > lastVBlankLine) && (beamY < firstVBlankLine)
			V.displayArea = (beamY >= firstDisplayLine) && (beamY <= lastDisplayLine) && V.visibleArea
			V.BA = !(((beamY-firstDisplayLine)%8 == 0) && V.displayArea)
			beamX := 0
			for cycle := 1; cycle <= cyclesPerLine; cycle++ {
				V.drawArea = ((cycle > 15) && (cycle < 56)) && V.displayArea
				// if V.drawArea {
				// 	fmt.Printf("Raster: %d - Cycle: %d - BA: %t - VMLI: %d - VCBASE/VC: %d/%d - RC: %d - Char: %02X\n", beamY, cycle, V.BA, V.VMLI, V.VCBASE, V.VC, V.RC, V.CharBuffer[V.VMLI])
				// }
				switch cycle {
				case 1:
				case 2:
				case 3:
				case 4:
				case 5:
				case 6:
				case 7:
				case 8:
				case 9:
				case 10:
				case 11: // Debut de la zone visible
					V.drawChar(beamX, beamY)
				case 12:
					V.drawChar(beamX, beamY)
				case 13:
					V.drawChar(beamX, beamY)
				case 14:
					V.VC = V.VCBASE
					V.VMLI = 0
					if !V.BA {
						V.RC = 0
					}
					V.drawChar(beamX, beamY)
				case 15: // Debut de la lecture de la memoire video en mode BadLine
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 16: // Debut de la zone d'affichage
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 17:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 18:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 19:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 20:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 21:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 22:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 23:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 24:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 25:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 26:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 27:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 28:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 29:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 30:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 31:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 32:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 33:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 34:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 35:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 36:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 37:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 38:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 39:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 40:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 41:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 42:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 43:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 44:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 45:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 46:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 47:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 48:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 49:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 50:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 51:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 52:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 53:
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 54: // Dernier lecture de la matrice video ram
					V.drawChar(beamX, beamY)
					V.readVideoMatrix()
				case 55: // Fin de la zone de display
					V.drawChar(beamX, beamY)

				case 56: // Debut de la zone visible
					V.drawChar(beamX, beamY)
				case 57:
					V.drawChar(beamX, beamY)
				case 58:
					if V.RC == 7 {
						V.VCBASE = V.VC
					}
					if V.displayArea {
						V.RC++
					}
					V.drawChar(beamX, beamY)
				case 59:
					V.drawChar(beamX, beamY)
				case 60:
				case 61:
				case 62:
				case 63:
				}
				beamX += 8
			}
		}
		beamY++
		// if beamY == firstDisplayLine+9 {
		// 	os.Exit(1)
		// }
		if beamY >= screenHeightPAL {
			beamY = 0
			V.VCBASE = 0
			V.graph.DisplayFrame()
		}
	}
}

// func (V *VIC) Run() {
// 	var VBlank, HBlank, VisibleArea bool
// 	var pixelColor globals.RGB

// 	V.graph = &graphic.SDLDriver{}
// 	// V.graph = &graphic.SDL2Driver{}

// 	V.graph.Init(winWidth, winHeight)
// 	defer func() {
// 		V.graph.CloseAll()
// 	}()

// 	// cpuTimer, _ := time.ParseDuration(fmt.Sprintf("%fms", lineRefresh))
// 	cpuTimer, _ := time.ParseDuration(fmt.Sprintf("%fms", 0.05))

// 	// fmt.Printf("cpuTimer %v.\n", cpuTimer)
// 	ticker := time.NewTicker(cpuTimer)
// 	defer func() {
// 		ticker.Stop()
// 	}()

// 	for {
// 		HBlank = true
// 		VBlank = true
// 		VisibleArea = false
// 		V.BadLineCounter = 0
// 		V.RowCounter = 0

// 		t0 := time.Now()
// 		for beamY := 0; beamY < screenHeightPAL; beamY++ {

// 			V.saveRasterPos(beamY)
// 			// // fmt.Printf("raster: %d - BadLineCounter: %d - RowCounter: %d\n", beamY, V.BadLineCounter, V.RowCounter)
// 			select {
// 			case <-ticker.C:
// 			if beamY > lastVBlankLine && beamY < firstVBlankLine {
// 				VBlank = false
// 				if beamY >= visibleFirstLine && beamY < visibleLastLine {
// 					VisibleArea = true
// 					if V.BadLineCounter == 0 {
// 						V.readScreenData(V.ram, beamY)
// 					}
// 				} else {
// 					VisibleArea = false
// 				}
// 			} else {
// 				VBlank = true
// 			}

// 			beamX := 0
// 			for cycle := 0; cycle < cyclesPerLine; cycle++ {
// 				if beamX >= lastHBlankCol && beamX < firstHBlankCol {
// 					HBlank = false
// 				} else {
// 					HBlank = true
// 				}
// 				for column := 0; column < 8; column++ {
// 					if VBlank || HBlank {
// 						pixelColor = Colors[Black]
// 					} else {
// 						if beamX >= visibleFirstCol && beamX < visibleLastCol && VisibleArea {
// 							pixelColor = Colors[V.getPixelColor(beamX)]
// 						} else {
// 							pixelColor = Colors[V.ram.Data[REG_EC]&0b00001111]
// 						}
// 					}
// 					V.graph.DrawPixel(beamX, beamY, pixelColor)
// 					beamX++
// 				}
// 				// V.cpuCycle <- true
// 				V.ram.WaitFor(false)
// 			}
// 			if beamY >= visibleFirstLine && beamY < visibleLastLine {
// 				V.BadLineCounter++
// 				if V.BadLineCounter == 8 {
// 					V.BadLineCounter = 0
// 					V.RowCounter++
// 				}
// 			}

// 			}
// 		}
// 		// setPixel(visibleFirstCol*8, visibleFirstLine, White)
// 		// setPixel(visibleLastCol*8, visibleLastLine, White)
// 		V.graph.DisplayFrame()
// 		// os.Exit(1)
// 		t1 := time.Now()
// 		fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
// 	}
// }
