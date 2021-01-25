package main

// Init :
func (m *Memory) Init() {
	m.Stack = m.Data[stackStart:stackEnd]
	for i := range m.Data {
		m.Data[i] = NOP
	}
}

func (m *Memory) load() {
	m.Data[0x0010] = 0x00
	m.Data[0x0011] = 0xEE
	m.Data[0x0012] = 0xFF

	m.Data[0xFF00] = LDA_IM
	m.Data[0xFF01] = 0xF0
	m.Data[0xFF02] = LDA_ZP
	m.Data[0xFF03] = 0x10
	m.Data[0xFF04] = LDX_ZP
	m.Data[0xFF05] = 0x11
	m.Data[0xFF06] = LDY_ZP
	m.Data[0xFF07] = 0x12
	m.Data[0xFF09] = LDA_IM
	m.Data[0xFF0A] = 0xAA
	m.Data[0xFF08] = PHA
	m.Data[0xFF09] = LDA_IM
	m.Data[0xFF0A] = 0xAA
	m.Data[0xFF0B] = PLA

	m.Data[0xFF11] = JMP_ABS
	m.Data[0xFF12] = 0xFF
	m.Data[0xFF13] = 0x00
}
