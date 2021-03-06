package vic

import (
	"fmt"
	"go6502/graphic"
	"go6502/mem"
)

const (
	screenWidthPAL  = 504
	screenHeightPAL = 312
	rasterWidthPAL  = 403
	rasterHeightPAL = 284
	cyclesPerLine   = 63

	rasterTime = 1                  // Nb of cycle to put 1 byte on a line
	rasterLine = rasterWidthPAL / 8 // Nb of cycle to draw a full line
	fullRaster = rasterLine * rasterHeightPAL

	// lineRefresh   = cyclesPerLine * cpuCycle                   // Time for a line in ms
	// screenRefresh = screenHeightPAL * cyclesPerLine * cpuCycle // Time for a full screen display in ms
	// fps           = 1 / screenRefresh

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

func (V *VIC) Init(memory *mem.Memory, video graphic.Driver) {
	V.graph = video
	V.graph.Init(winWidth, winHeight)

	V.ram = memory
	V.ram.Mem[REG_EC].Read = &V.ram.PLA.Ram
	V.ram.Mem[REG_EC].Zone[mem.RAM] = 0xFE // Border Color : Lightblue
	V.ram.Mem[REG_B0C].Read = &V.ram.PLA.Ram
	V.ram.Mem[REG_B0C].Zone[mem.RAM] = 0xF6 // Background Color : Blue
	V.ram.Mem[REG_CTRL1].Zone[mem.IO] = 0b10011011
	V.ram.Mem[REG_CTRL1].Zone[mem.RAM] = 0b00000000
	V.ram.Mem[REG_RASTER].Zone[mem.RAM] = 0b00000000
	V.ram.Mem[REG_CTRL2].Zone[mem.IO] = 0b00001000
	V.ram.Mem[PALNTSC].Zone[mem.RAM] = 0x01 // PAL
	V.ram.Mem[REG_IRQ].Zone[mem.IO] = 0b00001111
	V.ram.Mem[REG_SETIRQ].Zone[mem.IO] = 0b00000000
	V.ram.Mem[REG_SETIRQ].Zone[mem.RAM] = 0b00000000

	V.BA = true
	V.VCBASE = 0
	V.beamX = 0
	V.beamY = 0
	V.cycle = 1
	V.RasterIRQ = 0xFFFF
}

func (V *VIC) saveRasterPos(val int) {
	V.ram.Mem[REG_RASTER].Zone[mem.IO] = byte(val)
	if (byte(uint16(val) >> 8)) == 0x1 {
		V.ram.Mem[REG_CTRL1].Zone[mem.IO] |= RST8
	} else {
		V.ram.Mem[REG_CTRL1].Zone[mem.IO] &= ^RST8
	}
	// fmt.Printf("val: %d - RST8: %08b - RASTER: %08b\n", val, V.ram.Data[REG_RST8], V.ram.Data[REG_RASTER])
}

func (V *VIC) readVideoMatrix() {
	if !V.BA {
		V.ColorBuffer[V.VMLI] = V.ram.Color[V.VC].Zone[mem.RAM] & 0b00001111
		V.CharBuffer[V.VMLI] = V.ram.Screen[V.VC].Zone[mem.RAM]
	}
}

func (V *VIC) drawChar(X int, Y int) {
	if V.drawArea && (V.ram.Mem[REG_CTRL1].Zone[mem.IO]&DEN > 0) {
		charAddr := (uint16(V.CharBuffer[V.VMLI]) << 3) + uint16(V.RC)
		charData := V.ram.CharGen[charAddr].Zone[mem.CHAR]
		// fmt.Printf("SC: %02X - RC: %d - %04X - %02X = %08b\n", V.CharBuffer[V.VMLI], V.RC, charAddr, charData, charData)
		// if V.CharBuffer[V.VMLI] == 0 {
		// fmt.Printf("Raster: %d - Cycle: %d - BA: %t - VMLI: %d - VCBASE/VC: %d/%d - RC: %d - Char: %02X\n", Y, X, V.BA, V.VMLI, V.VCBASE, V.VC, V.RC, V.CharBuffer[V.VMLI])
		// }
		for column := 0; column < 8; column++ {
			bit := byte(0b10000000 >> column)
			if charData&bit > 0 {
				V.graph.DrawPixel(X+column, Y, Colors[V.ColorBuffer[V.VMLI]])
			} else {
				V.graph.DrawPixel(X+column, Y, Colors[V.ram.Mem[REG_B0C].Zone[mem.RAM]&0b00001111])
			}
		}
		V.VMLI++
		V.VC++
	} else if V.visibleArea {
		for column := 0; column < 8; column++ {
			V.graph.DrawPixel(X+column, Y, Colors[V.ram.Mem[REG_EC].Zone[mem.RAM]&0b00001111])
		}
	}
}

func (V *VIC) registersManagement() {
	V.saveRasterPos(V.beamY)

	V.ram.Mem[REG_SETIRQ].Zone[mem.IO] = V.ram.Mem[REG_SETIRQ].Zone[mem.RAM]

	if V.ram.Mem[REG_CTRL1].IsWrite || V.ram.Mem[REG_RASTER].IsWrite {
		V.RasterIRQ = uint16(V.ram.Mem[REG_CTRL1].Zone[mem.RAM]&0b10000000) << 8
		V.RasterIRQ += uint16(V.ram.Mem[REG_RASTER].Zone[mem.RAM])
		V.ram.Mem[REG_CTRL1].IsWrite = false
		V.ram.Mem[REG_RASTER].IsWrite = false
	}

	if V.ram.Mem[REG_IRQ].IsWrite {
		V.ram.Mem[REG_IRQ].Zone[mem.IO] = V.ram.Mem[REG_IRQ].Zone[mem.RAM]
		V.ram.Mem[REG_IRQ].Zone[mem.IO] &= 0b01111111
		// *V.IRQ_Pin = 0
		V.ram.Mem[REG_IRQ].IsWrite = false
	}
}

func (V *VIC) Run() {
	V.registersManagement()

	V.visibleArea = (V.beamY > lastVBlankLine) && (V.beamY < firstVBlankLine)
	// V.displayArea = (V.beamY >= firstDisplayLine) && (V.beamY <= lastDisplayLine) && V.visibleArea
	V.displayArea = (V.beamY >= firstDisplayLine) && (V.beamY <= lastDisplayLine)
	V.beamX = (V.cycle - 1) * 8
	V.drawArea = ((V.cycle > 15) && (V.cycle < 56)) && V.displayArea

	V.BA = !(((V.beamY-firstDisplayLine)%8 == 0) && V.displayArea && (V.cycle > 11) && (V.cycle < 55))

	// if V.drawArea {
	// 	fmt.Printf("Raster: %d - Cycle: %d - BA: %t - VMLI: %d - VCBASE/VC: %d/%d - RC: %d - Char: %02X\n", V.beamY, V.cycle, V.BA, V.VMLI, V.VCBASE, V.VC, V.RC, V.CharBuffer[V.VMLI])
	// }

	switch V.cycle {
	case 1:
		if V.ram.Mem[REG_SETIRQ].Zone[mem.IO]&IRQ_RASTER > 0 {
			if V.RasterIRQ == uint16(V.beamY) {
				//fmt.Printf("\nIRQ: %04X - %04X", V.RasterIRQ, uint16(V.beamY))
				fmt.Println("Rastrer Interrupt")
				*V.IRQ_Pin = 1
				V.ram.Mem[REG_IRQ].Zone[mem.IO] |= 0b10000001
			}
		}
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
			V.graph.UpdateFrame()
		}
	}

}
