package mem

import (
	"fmt"
	"go6502/globals"
	"io/ioutil"
)

// Init :
func (m *Memory) Init() {
	m.Stack = m.bank[0].data[stackStart : stackEnd+1]
	m.Screen = m.bank[0].data[screenStart : screenEnd+1]
	m.Color = m.bank[2].data[colorStart : colorEnd+1]
	m.Kernal = m.bank[3].data
	m.Basic = m.bank[1].data
	m.CharGen = make([]globals.Byte, 4096)

	m.Vic[0] = m.bank[0].data
	m.Vic[1] = m.bank[1].data
	m.Vic[2] = m.bank[2].data
	m.Vic[3] = m.bank[3].data

	// for i := range m.Data {
	// 	m.Data[i] = 0x00
	// }
	// cpt := 0
	// for i := range m.Color {
	// 	m.Color[i] = globals.Byte(cpt)
	// 	cpt++
	// 	if cpt > 15 {
	// 		cpt = 0
	// 	}
	// }
	// for i := range m.Screen {
	// 	m.Screen[i] = globals.Byte(i)
	// }

	m.loadRom("roms/char.bin", 4096, m.CharGen)
	m.loadRom("roms/kernal.bin", 8192, m.Kernal)
	m.loadRom("roms/basic.bin", 8192, m.Basic)
}

func (m *Memory) getBank(addr globals.Word) int {
	return int(addr / 0xA000)
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
	bank := m.getBank(startAddr)
	cpt := startAddr
	for j := 0; j < 10; j++ {
		fmt.Printf("%04X : ", cpt)
		for i := 0; i < 8; i++ {
			fmt.Printf("%02X", m.bank[bank].data[cpt])
			cpt++
			fmt.Printf("%02X ", m.bank[bank].data[cpt])
			cpt++
		}
		fmt.Println()
	}
}

func (m *Memory) Read(addr globals.Word) globals.Byte {
	bank := m.getBank(addr)
	return m.bank[bank].data[addr]
}

func (m *Memory) Write(addr globals.Word, value globals.Byte)  {
	bank := m.getBank(addr)
	m.bank[bank].data[addr] = value
}

// func (m *Memory) String2screenCode(startMem globals.Word, message string) {
// 	runes := []rune(message)
// 	for i := 0; i < len(runes); i++ {
// 		m.Data[startMem+globals.Word(i)] = globals.Byte(runes[i])
// 	}
// }
