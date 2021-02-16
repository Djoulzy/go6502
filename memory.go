package main

import (
	"fmt"
	"io/ioutil"
)

// Init :
func (m *Memory) Init() {
	m.Stack = m.Data[stackStart : stackEnd+1]
	m.Screen = m.Data[screenStart : screenEnd+1]
	m.Color = m.Data[colorStart : colorEnd+1]

	m.Vic[0] = m.Data[0x0000:0x3FFF]
	m.Vic[1] = m.Data[0x4000:0x7FFF]
	m.Vic[2] = m.Data[0x8000:0xBFFF]
	m.Vic[3] = m.Data[0xC000:0xFFFF]

	for i := range m.Data {
		m.Data[i] = NOP
	}
	for i := range m.Color {
		m.Color[i] = 0x01
	}
	// for i := range m.Screen {
	// 	m.Screen[i] = Byte(i)
	// }
	m.Data[0x0400] = 0x10
	m.loadCharGenRom("char.bin")
}

func (m *Memory) loadCharGenRom(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	if len(data) != 4096 {
		panic("Bad ROM Size")
	}
	for i := 0; i < 4096; i++ {
		m.CharGen[i] = Byte(data[i])
	}
}

func (m *Memory) dump(startAddr Word) {
	cpt := startAddr
	for j := 0; j < 10; j++ {
		fmt.Printf("%04X : ", cpt)
		for i := 0; i < 8; i++ {
			fmt.Printf("%04X ", m.Data[cpt])
			cpt++
		}
		fmt.Println()
	}
}

func (m *Memory) load1() {
	m.Data[0x0010] = 0x00
	m.Data[0x0011] = 0xEE
	m.Data[0x0012] = 0xFF

	m.Data[0x0013] = 0x01
	m.Data[0x0014] = 0x02
	m.Data[0x0015] = 0x03
	m.Data[0x0016] = 0x04

	m.Data[0xEE00] = LDA_IM // LDA #$AA
	m.Data[0xEE01] = 0xAA
	m.Data[0xEE02] = RTS // RTS

	m.Data[0xFF00] = LDA_IM // LDA #$0C
	m.Data[0xFF01] = 0x0C
	m.Data[0xFF02] = PHA // PHA
	m.Data[0xFF03] = JSR // JSR $EE00
	m.Data[0xFF04] = 0x00
	m.Data[0xFF05] = 0xEE
	m.Data[0xFF06] = NOP
	m.Data[0xFF07] = PLA // PLA

	m.Data[0xFF10] = LDA_IM // LDA #$F0
	m.Data[0xFF11] = 0xF0
	m.Data[0xFF12] = LDA_ZP // LDA $10
	m.Data[0xFF13] = 0x10
	m.Data[0xFF14] = LDX_ZP // LDX $11
	m.Data[0xFF15] = 0x11
	m.Data[0xFF16] = LDY_ZP // LDY $12
	m.Data[0xFF17] = 0x12
	m.Data[0xFF18] = NOP
	m.Data[0xFF19] = LDA_IM // LDA #$AA
	m.Data[0xFF1A] = 0xAA
	m.Data[0xFF1B] = PHA    // PHA
	m.Data[0xFF1C] = LDA_IM // LDA #$AA
	m.Data[0xFF1D] = 0xAA
	m.Data[0xFF1E] = PLA // PLA

	m.Data[0xFF1F] = LDY_IM
	m.Data[0xFF20] = 0x03
	m.Data[0xFF21] = LDA_INY
	m.Data[0xFF22] = 0x13

	m.Data[0xFF23] = JMP_ABS // JMP $FF00
	m.Data[0xFF24] = 0x1F
	m.Data[0xFF25] = 0xFF
}

func (m *Memory) string2screenCode(startMem Word, message string) {
	runes := []rune(message)
	for i := 0; i < len(runes); i++ {
		m.Data[startMem+Word(i)] = Byte(runes[i])
	}
}

func (m *Memory) load0() {
	line := Word(0xFF00)

	m.string2screenCode(0xEE00, "12345")

	m.Data[line] = LDX_IM
	line++
	m.Data[line] = 0x00
	line++
	label := line
	m.Data[line] = LDA_ABX
	line++
	m.Data[line] = 0x00
	line++
	m.Data[line] = 0xEE
	line++
	m.Data[line] = STA_ABX
	line++
	m.Data[line] = 0x90
	line++
	m.Data[line] = 0x05
	line++
	m.Data[line] = INX
	line++
	m.Data[line] = CPX_IM
	line++
	m.Data[line] = 0x05
	line++
	m.Data[line] = BNE
	line++
	m.Data[line] = Byte(label)
	line++
	m.Data[line] = Byte(label >> 8)

	// line++
	// m.Data[line] = DMP
	// line++
	// m.Data[line] = 0x90
	// line++
	// m.Data[line] = 0x05

	line++
	m.Data[line] = JMP_ABS // JMP $FF00
	line++
	m.Data[line] = 0x00
	line++
	m.Data[line] = 0xFF
}

// init_text  ldx #$00         ; init X register with $00
// loop_text  lda line1,x      ; read characters from line1 table of text...
//            sta $0590,x      ; ...and store in screen ram near the center
//            lda line2,x      ; read characters from line2 table of text...
//            sta $05e0,x      ; ...and put 2 rows below line1

//            inx
//            cpx #$28         ; finished when all 40 cols of a line are processed
//            bne loop_text    ; loop if we are not done yet
