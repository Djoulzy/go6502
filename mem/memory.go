package mem

import (
	"fmt"
	"go6502/globals"
	"io/ioutil"
)

// Init :
func (m *Memory) Init() {
	m.Stack = m.Data[stackStart : stackEnd+1]
	m.Screen = m.Data[screenStart : screenEnd+1]
	m.Color = m.Data[colorStart : colorEnd+1]
	m.Kernal = m.Data[KernalStart : KernalEnd+1]
	m.Basic = m.Data[BasicStart : BasicEnd+1]
	m.CharGen = make([]globals.Byte, 4096)

	m.Vic[0] = m.Data[0x0000:0x3FFF]
	m.Vic[1] = m.Data[0x4000:0x7FFF]
	m.Vic[2] = m.Data[0x8000:0xBFFF]
	m.Vic[3] = m.Data[0xC000:0xFFFF]

	for i := range m.Data {
		m.Data[i] = 0x00
	}
	cpt := 0
	for i := range m.Color {
		m.Color[i] = globals.Byte(cpt)
		cpt++
		if cpt > 15 {
			cpt = 0
		}
	}
	for i := range m.Screen {
		m.Screen[i] = globals.Byte(i)
	}

	m.loadRom("roms/char.bin", 4096, m.CharGen)
	m.loadRom("roms/kernal.bin", 8192, m.Kernal)
	m.loadRom("roms/basic.bin", 8192, m.Basic)
}

func (m *Memory) loadRom(filename string, fileSize int, dest []globals.Byte) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	if len(data) != fileSize {
		panic("Bad ROM Size")
	}
	for i := 0; i < fileSize; i++ {
		dest[i] = globals.Byte(data[i])
	}
}

func (m *Memory) DumpChar(screenCode globals.Byte) {
	cpt := globals.Word(screenCode) << 3
	for j := 0; j < 4; j++ {
		for i := 0; i < 8; i++ {
			fmt.Printf("%04X : %08b\n", cpt, m.CharGen[cpt])
			cpt++
		}
		fmt.Println()
	}
}

func (m *Memory) Dump(startAddr globals.Word) {
	cpt := startAddr
	for j := 0; j < 10; j++ {
		fmt.Printf("%04X : ", cpt)
		for i := 0; i < 8; i++ {
			fmt.Printf("%02X", m.Data[cpt])
			cpt++
			fmt.Printf("%02X ", m.Data[cpt])
			cpt++
		}
		fmt.Println()
	}
}

func (m *Memory) String2screenCode(startMem globals.Word, message string) {
	runes := []rune(message)
	for i := 0; i < len(runes); i++ {
		m.Data[startMem+globals.Word(i)] = globals.Byte(runes[i])
	}
}
