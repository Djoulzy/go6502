package mem

const (
	memorySize  = 65536
	stackStart  = 0x0100
	stackEnd    = 0x01FF
	screenStart = 0x0400
	screenEnd   = 0x07FF
	colorStart  = 0xD800
	colorEnd    = 0xDBFF
	intAddr     = 0xFFFA
	resetAddr   = 0xFFFC
	brkAddr     = 0xFFFE
	KernalStart = 0xE000
	KernalEnd   = 0xFFFF
	BasicStart  = 0xA000
	BasicEnd    = 0xC000
	vic2        = 0x4000
	vic3        = 0x8000
	vic4        = 0xC000
)

type latch struct {
	kernal   bool
	basic    bool
	char     bool
	io       bool
	disabled bool
}

type Cell struct {
	RomMode *bool
	ExpMode *bool
	Exp     byte
	Rom     byte
	Ram     byte
}

// Memory :
type Memory struct {
	latch   latch
	Mem     [memorySize]Cell
	Kernal  []Cell
	Basic   []Cell
	CharGen []Cell
	Stack   []Cell
	Screen  []Cell
	Color   []Cell
	Vic     [4][]Cell
}
