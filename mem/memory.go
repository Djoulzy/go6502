package mem

import (
	"fmt"
	"io/ioutil"
)

// Init :
func (m *Memory) Init() {
	m.latch = latch{
		kernal: true,
		basic:  true,
		io:     true,
		char:   false,
	}

	m.Stack = m.Mem[stackStart : stackEnd+1]
	m.Screen = m.Mem[screenStart : screenEnd+1]
	m.Color = m.Mem[colorStart : colorEnd+1]
	m.Kernal = m.Mem[KernalStart : KernalEnd+1]
	m.Basic = m.Mem[BasicStart : BasicEnd+1]
	m.CharGen = make([]cell, 4096)

	m.Vic[0] = m.Mem[0x0000 : vic2-1]
	m.Vic[1] = m.Mem[vic2 : vic3-1]
	m.Vic[2] = m.Mem[vic3 : vic4-1]
	m.Vic[3] = m.Mem[vic4:0xFFFF]

	m.Mem[0].Ram = 0x2F // Processor port data direction register
	m.Mem[1].Ram = 0x37 // Processor port / memory map configuration

	// for i := range m.Data {
	// 	m.Data[i] = 0x00
	// }
	// cpt := 0
	// for i := range m.Color {
	// 	m.Color[i] = byte(cpt)
	// 	cpt++
	// 	if cpt > 15 {
	// 		cpt = 0
	// 	}
	// }
	// for i := range m.Screen {
	// 	m.Screen[i] = byte(i)
	// }

	m.loadRom("roms/char.bin", 4096, m.CharGen, &m.latch.char)
	m.loadRom("roms/kernal.bin", 8192, m.Kernal, &m.latch.kernal)
	m.loadRom("roms/basic.bin", 8192, m.Basic, &m.latch.basic)
}

func (m *Memory) loadRom(filename string, fileSize int, dest []cell, mode *bool) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	if len(data) != fileSize {
		panic("Bad ROM Size")
	}
	for i := 0; i < fileSize; i++ {
		dest[i].romMode = mode
		dest[i].Rom = byte(data[i])
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
	final := addr.(uint16)
	return m.Mem[final].Ram
}

func (m *Memory) Write(addr interface{}, value byte) {
	final := addr.(uint16)
	m.Mem[final].Ram = value
}

func (m *Memory) Dump(startAddr uint16) {
	cpt := startAddr
	for j := 0; j < 10; j++ {
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
