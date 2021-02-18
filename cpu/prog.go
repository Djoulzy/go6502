package cpu

import (
	"go6502/globals"
	"go6502/mem"
)

func (C *CPU) load1(m *mem.Memory) {
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

func (C *CPU) load0(m *mem.Memory) {
	line := globals.Word(0xFF00)

	m.Data[0xFF00] = LDY_IM
	m.Data[0xFF01] = 0x90
	m.Data[0xFF02] = LDA_IM
	m.Data[0xFF03] = 0x00
	m.Data[0xFF04] = BRK
	m.Data[0xFF05] = BRK

	m.String2screenCode(0xEE00, "12345")

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
	m.Data[line] = globals.Byte(label)
	line++
	m.Data[line] = globals.Byte(label >> 8)

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
