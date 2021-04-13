package vic

import (
	"go6502/graphic"
	"go6502/mem"
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

	firstVBlankLine  = 300
	lastVBlankLine   = 15
	firstDisplayLine = 51
	lastDisplayLine  = 250

	// firstHBlankCol  = 453
	// lastHBlankCol   = 50
	// visibleFirstCol = 92
	// visibleLastCol  = 412
)

func (V *VIC) Init(mem *mem.Memory) {
	V.graph = &graphic.SDLDriver{}
	V.graph.Init(winWidth, winHeight)

	V.ram = mem
	V.ram.Mem[REG_EC].Rom = Lightblue
	V.ram.Mem[REG_B0C].Rom = Blue
	V.BA = true
	V.ram.Mem[PALNTSC].Ram = 0x01 // PAL

	V.VCBASE = 0
	V.beamX = 0
	V.beamY = 0
	V.cycle = 1
}

func (V *VIC) saveRasterPos(val int) {
	V.ram.Mem[REG_RASTER].Ram = byte(val)
	if (byte(uint16(val) >> 8)) == 0x1 {
		V.ram.Mem[REG_RST8].Rom |= 0b10000000
	} else {
		V.ram.Mem[REG_RST8].Rom &= 0b01111111
	}
	// fmt.Printf("val: %d - RST8: %08b - RASTER: %08b\n", val, V.ram.Data[REG_RST8], V.ram.Data[REG_RASTER])
}

func (V *VIC) readVideoMatrix() {
	if !V.BA {
		V.ColorBuffer[V.VMLI] = V.ram.Color[V.VC].Ram
		V.CharBuffer[V.VMLI] = V.ram.Screen[V.VC].Ram
	}
}

func (V *VIC) drawChar(X int, Y int) {
	if V.drawArea {
		charAddr := (uint16(V.CharBuffer[V.VMLI]) << 3) + uint16(V.RC)
		charData := V.ram.CharGen[charAddr].Rom
		// fmt.Printf("SC: %02X - RC: %d - %04X - %02X = %08b\n", V.CharBuffer[V.VMLI], V.RC, charAddr, charData, charData)
		// if V.CharBuffer[V.VMLI] == 0 {
		// fmt.Printf("Raster: %d - Cycle: %d - BA: %t - VMLI: %d - VCBASE/VC: %d/%d - RC: %d - Char: %02X\n", Y, X, V.BA, V.VMLI, V.VCBASE, V.VC, V.RC, V.CharBuffer[V.VMLI])
		// }
		for column := 0; column < 8; column++ {
			bit := byte(0b10000000 >> column)
			if charData&bit > 0 {
				V.graph.DrawPixel(X+column, Y, Colors[V.ColorBuffer[V.VMLI]])
			} else {
				V.graph.DrawPixel(X+column, Y, Colors[V.ram.Mem[REG_B0C].Rom&0b00001111])
			}
		}
		V.VMLI++
		V.VC++
	} else if V.visibleArea {
		for column := 0; column < 8; column++ {
			V.graph.DrawPixel(X+column, Y, Colors[V.ram.Mem[REG_EC].Rom&0b00001111])
		}
	}
}

func (V *VIC) Run() {
	V.saveRasterPos(V.beamY)
	V.visibleArea = (V.beamY > lastVBlankLine) && (V.beamY < firstVBlankLine)
	V.displayArea = (V.beamY >= firstDisplayLine) && (V.beamY <= lastDisplayLine) && V.visibleArea
	V.BA = !(((V.beamY-firstDisplayLine)%8 == 0) && V.displayArea)
	V.beamX = (V.cycle - 1) * 8
	V.drawArea = ((V.cycle > 15) && (V.cycle < 56)) && V.displayArea
	// if V.drawArea {
	// 	fmt.Printf("Raster: %d - Cycle: %d - BA: %t - VMLI: %d - VCBASE/VC: %d/%d - RC: %d - Char: %02X\n", V.beamY, V.cycle, V.BA, V.VMLI, V.VCBASE, V.VC, V.RC, V.CharBuffer[V.VMLI])
	// }

	switch V.cycle {
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
		V.drawChar(V.beamX, V.beamY)
	case 12:
		V.drawChar(V.beamX, V.beamY)
	case 13:
		V.drawChar(V.beamX, V.beamY)
	case 14:
		V.VC = V.VCBASE
		V.VMLI = 0
		if !V.BA {
			V.RC = 0
		}
		V.drawChar(V.beamX, V.beamY)
	case 15: // Debut de la lecture de la memoire video en mode BadLine
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 16: // Debut de la zone d'affichage
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 17:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 18:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 19:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 20:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 21:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 22:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 23:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 24:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 25:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 26:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 27:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 28:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 29:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 30:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 31:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 32:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 33:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 34:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 35:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 36:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 37:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 38:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 39:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 40:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 41:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 42:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 43:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 44:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 45:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 46:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 47:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 48:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 49:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 50:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 51:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 52:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 53:
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 54: // Dernier lecture de la matrice video ram
		V.drawChar(V.beamX, V.beamY)
		V.readVideoMatrix()
	case 55: // Fin de la zone de display
		V.drawChar(V.beamX, V.beamY)
	case 56: // Debut de la zone visible
		V.drawChar(V.beamX, V.beamY)
	case 57:
		V.drawChar(V.beamX, V.beamY)
	case 58:
		if V.RC == 7 {
			V.VCBASE = V.VC
		}
		if V.displayArea {
			V.RC++
		}
		V.drawChar(V.beamX, V.beamY)
	case 59:
		V.drawChar(V.beamX, V.beamY)
	case 60:
	case 61:
	case 62:
	case 63:
	}
	// V.beamX += 8
	V.cycle++
	if V.cycle > cyclesPerLine {
		V.cycle = 1
		V.beamY++
		if V.beamY >= screenHeightPAL {
			V.beamY = 0
			V.VCBASE = 0
			V.graph.DisplayFrame()
		}
	}

}
