package vic

import (
	"fmt"
	"go6502/databus"
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

	firstHBlankCol  = 453
	lastHBlankCol   = 50
	visibleFirstCol = 92
	visibleLastCol  = 412
)

func (V *VIC) Init(dbus *databus.Databus, mem *mem.Memory) {
	V.graph = &graphic.SDLDriver{}
	V.graph.Init(winWidth, winHeight)

	V.ram = mem
	V.dbus = dbus
	V.ram.Mem[REG_EC].Rom = Lightblue
	V.ram.Mem[REG_B0C].Rom = Blue
	V.BA = true
}

func (V *VIC) saveRasterPos(val int) {
	V.ram.Mem[REG_RASTER].Rom = byte(val)
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
	} else {
		V.dbus.WaitBusHigh()
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
	V.dbus.WaitBusHigh()
	for {

		<-ticker.C
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

			// fmt.Printf("VIC\n")
			switch cycle {
			case 1:
				V.dbus.WaitBusHigh()
			case 2:
				V.dbus.WaitBusHigh()
			case 3:
				V.dbus.WaitBusHigh()
			case 4:
				V.dbus.WaitBusHigh()
			case 5:
				V.dbus.WaitBusHigh()
			case 6:
				V.dbus.WaitBusHigh()
			case 7:
				V.dbus.WaitBusHigh()
			case 8:
				V.dbus.WaitBusHigh()
			case 9:
				V.dbus.WaitBusHigh()
			case 10:
				V.dbus.WaitBusHigh()
			case 11: // Debut de la zone visible
				V.drawChar(beamX, beamY)
				V.dbus.WaitBusHigh()
			case 12:
				V.drawChar(beamX, beamY)
				V.dbus.WaitBusHigh()
			case 13:
				V.drawChar(beamX, beamY)
				V.dbus.WaitBusHigh()
			case 14:
				V.VC = V.VCBASE
				V.VMLI = 0
				if !V.BA {
					V.RC = 0
				}
				V.drawChar(beamX, beamY)
				V.dbus.WaitBusHigh()
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
				V.dbus.WaitBusHigh()
			case 56: // Debut de la zone visible
				V.drawChar(beamX, beamY)
				V.dbus.WaitBusHigh()
			case 57:
				V.drawChar(beamX, beamY)
				V.dbus.WaitBusHigh()
			case 58:
				if V.RC == 7 {
					V.VCBASE = V.VC
				}
				if V.displayArea {
					V.RC++
				}
				V.drawChar(beamX, beamY)
				V.dbus.WaitBusHigh()
			case 59:
				V.drawChar(beamX, beamY)
				V.dbus.WaitBusHigh()
			case 60:
				V.dbus.WaitBusHigh()
			case 61:
				V.dbus.WaitBusHigh()
			case 62:
				V.dbus.WaitBusHigh()
			case 63:
				V.dbus.WaitBusHigh()
			}
			beamX += 8
		}

		beamY++
		if beamY >= screenHeightPAL {
			beamY = 0
			V.VCBASE = 0
			V.graph.DisplayFrame()
		}
		// if beamY == firstDisplayLine+9 {
		// os.Exit(1)
		// }
	}
}