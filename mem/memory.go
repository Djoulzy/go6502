package mem

import (
	"fmt"
	"io/ioutil"
)

// Init :
func (m *Memory) Init() {
	m.latch = latch{
		kernal:   true,
		basic:    true,
		io:       false,
		char:     false,
		disabled: false,
	}

	m.Stack = m.Mem[stackStart : stackEnd+1]
	m.Screen = m.Mem[screenStart : screenEnd+1]
	m.Color = m.Mem[colorStart : colorEnd+1]
	m.Kernal = m.Mem[KernalStart : KernalEnd+1]
	m.Basic = m.Mem[BasicStart : BasicEnd+1]
	m.CharGen = make([]Cell, 4096)

	m.Vic[0] = m.Mem[0x0000 : vic2-1]
	m.Vic[1] = m.Mem[vic2 : vic3-1]
	m.Vic[2] = m.Mem[vic3 : vic4-1]
	m.Vic[3] = m.Mem[vic4:0xFFFF]

	m.Mem[0].Ram = 0x2F // Processor port data direction register
	m.Mem[1].Ram = 0x37 // Processor port / memory map configuration

	cpt := 0
	fill := byte(0x00)
	for i := range m.Mem {
		m.Mem[i].RomMode = &m.latch.disabled
		m.Mem[i].ExpMode = &m.latch.disabled
		m.Mem[i].Ram = fill
		cpt++
		if cpt == 0x40 {
			fill = ^fill
			cpt = 0
		}
	}

	for i := range m.Color {
		m.Color[i].Ram = 0
	}
	// for i := range m.Screen {
	// 	m.Screen[i] = byte(i)
	// }

	m.loadRom("roms/char.bin", 4096, m.CharGen, &m.latch.char)
	m.loadRom("roms/kernal.bin", 8192, m.Kernal, &m.latch.kernal)
	m.loadRom("roms/basic.bin", 8192, m.Basic, &m.latch.basic)

	m.setRomMode(0xDC00, 0xFF, &m.latch.io)
	m.setRomMode(0xDD00, 0xFF, &m.latch.io)
}

func (m *Memory) loadRom(filename string, fileSize int, dest []Cell, mode *bool) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	if len(data) != fileSize {
		panic("Bad ROM Size")
	}
	for i := 0; i < fileSize; i++ {
		dest[i].RomMode = mode
		dest[i].Rom = byte(data[i])
	}
}

func (m *Memory) setRomMode(start uint16, size uint16, mode *bool) {
	for i := start; i <= start+size; i++ {
		m.Mem[i].RomMode = mode
	}
}

func (m *Memory) DumpChar(screenCode byte) {
	cpt := uint16(screenCode) << 3
	for j := 0; j < 4; j++ {
		for i := 0; i < 8; i++ {
			fmt.Printf("%04X : %08b\n", cpt, m.CharGen[cpt])
			cpt++
		}
		fmt.Println()
	}
}

func (m *Memory) Read(addr interface{}) byte {
	var Cell Cell

	switch typed := addr.(type) {
	case uint16:
		Cell = m.Mem[typed]
	case uint8:
		Cell = m.Mem[typed]
	}
	if *Cell.RomMode {
		return Cell.Rom
	}
	return Cell.Ram
}

func (m *Memory) Write(addr interface{}, value byte) bool {
	var Cell *Cell

	switch typed := addr.(type) {
	case uint16:
		Cell = &m.Mem[typed]
	case uint8:
		Cell = &m.Mem[typed]
	}
	if *Cell.RomMode {
		return false
	}
	Cell.Ram = value
	return true
}

func (m *Memory) DumpStack(SP byte, nbline int) {
	if nbline == 0 {
		nbline = 16
	}
	cpt := 255
	for y := 0; y < nbline; y++ {
		for x := 0; x < 16; x++ {
			if byte(cpt) == SP {
				fmt.Printf(">%02X", m.Stack[cpt].Ram)
			} else {
				fmt.Printf(" %02X", m.Stack[cpt].Ram)
			}
			cpt--
		}
		if y < nbline-1 {
			fmt.Printf("\n")
		}
	}
}

func (m *Memory) Dump(startAddr uint16) {
	cpt := startAddr
	for j := 0; j < 16; j++ {
		fmt.Printf("%04X : ", cpt)
		for i := 0; i < 8; i++ {
			fmt.Printf("%02X", m.Read(cpt))
			cpt++
			fmt.Printf("%02X ", m.Read(cpt))
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
