package mem

import (
	"fmt"
	"io/ioutil"
)

// Init :
func (m *Memory) Init() {
	m.PLA = latch{
		Kernal:    KERNAL,
		Basic:     BASIC,
		Char_io_r: CHAR,
		Char_io_w: RAM,
		Ram:       RAM,
	}

	m.Stack = m.Mem[stackStart : stackEnd+1]
	m.Screen = m.Mem[screenStart : screenEnd+1]
	m.Color = m.Mem[colorStart : colorEnd+1]
	m.Kernal = m.Mem[KernalStart : KernalEnd+1]
	m.Basic = m.Mem[BasicStart : BasicEnd+1]
	m.CharGen = m.Mem[charStart : charEnd+1]

	cpt := 0
	fill := byte(0x00)
	for i := range m.Mem {
		m.Mem[i].Read = &m.PLA.Ram
		m.Mem[i].Write = &m.PLA.Ram
		m.Mem[i].Zone[RAM] = fill
		m.Mem[i].IsRead = false
		m.Mem[i].IsWrite = false
		cpt++
		if cpt == 0x40 {
			fill = ^fill
			cpt = 0
		}
	}

	m.Vic[0] = m.Mem[0x0000 : vic2-1]
	m.Vic[1] = m.Mem[vic2 : vic3-1]
	m.Vic[2] = m.Mem[vic3 : vic4-1]
	m.Vic[3] = m.Mem[vic4:0xFFFF]

	m.loadRom("roms/kernal.bin", 8192, m.Kernal, &m.PLA.Kernal, &m.PLA.Ram)
	m.loadRom("roms/basic.bin", 8192, m.Basic, &m.PLA.Basic, &m.PLA.Ram)
	m.loadRom("roms/char.bin", 4096, m.CharGen, &m.PLA.Char_io_r, &m.PLA.Ram)
	m.PLA.Char_io_r = IO

	m.Mem[0].Zone[RAM] = 0x2F // Processor port data direction register
	m.Mem[1].Zone[RAM] = 0x37 // Processor port / memory map configuration

	for i := range m.Color {
		m.Color[i].Zone[RAM] = 0xF6
	}
}

func (m *Memory) loadRom(filename string, fileSize int, dest []Cell, rmode *int, wmode *int) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	if len(data) != fileSize {
		panic("Bad ROM Size")
	}
	for i := 0; i < fileSize; i++ {
		dest[i].Read = rmode
		dest[i].Write = wmode
		dest[i].Zone[*rmode] = byte(data[i])
	}
}

func (m *Memory) DumpChar(screenCode byte) {
	cpt := uint16(screenCode) << 3
	for j := 0; j < 4; j++ {
		for i := 0; i < 8; i++ {
			fmt.Printf("%04X : %08b\n", cpt, m.CharGen[cpt].Zone[CHAR])
			cpt++
		}
		fmt.Println()
	}
}

func (m *Memory) Read(addr uint16) byte {
	cell := &m.Mem[addr]
	cell.IsRead = true
	return cell.Zone[*cell.Read]
}

func (m *Memory) Write(addr uint16, value byte) {
	cell := &m.Mem[addr]
	cell.IsWrite = true
	cell.Zone[*cell.Write] = value
}

func (m *Memory) DumpStack(SP byte, nbline int) {
	if nbline == 0 {
		nbline = 16
	}
	cpt := 0
	for y := 0; y < nbline; y++ {
		for x := 0; x < 16; x++ {
			if byte(cpt) == SP {
				fmt.Printf(">%02X", m.Stack[cpt].Zone[RAM])
			} else {
				fmt.Printf(" %02X", m.Stack[cpt].Zone[RAM])
			}
			cpt++
		}
		if y < nbline-1 {
			fmt.Printf("\n")
		}
	}
}

func (m *Memory) DumpCIA() {
	cia1 := 0xDC00
	fmt.Printf("\n")

	fmt.Printf("CIA1 RAM: ")
	for i := 0; i < 16; i++ {
		fmt.Printf("%02X ", m.Mem[cia1+i].Zone[RAM])
	}
	fmt.Printf("\nCIA1 IO : ")
	for i := 0; i < 16; i++ {
		fmt.Printf("%02X ", m.Mem[cia1+i].Zone[IO])
	}
}

func (m *Memory) Dump(startAddr uint16, zone int) {
	cpt := startAddr
	fmt.Printf("\n")
	for j := 0; j < 16; j++ {
		fmt.Printf("%04X : ", cpt)
		for i := 0; i < 16; i++ {
			fmt.Printf("%02X ", m.Mem[cpt].Zone[zone])
			cpt++
		}
		fmt.Println()
	}
}

// func (m *Memory) String2screenCode(startMem uint16, message string) {
// 	runes := []rune(message)
// 	for i := 0; i < len(runes); i++ {
// 		m.Data[startMem+uint16(i)] = byte(runes[i])
// 	}
// }
