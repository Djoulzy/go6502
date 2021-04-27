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
	CDStart     = 0xC000
	CDEnd      = 0xDFFF
	EFStart     = 0xE000
	EFEnd      = 0xFFFF
	BasicStart  = 0xA000
	BasicEnd    = 0xC000
	vic2        = 0x4000
	vic3        = 0x8000
	vic4        = 0xC000
)

const (
	RAM    = 0
	KERNAL = 1
	BASIC  = 1
	IO     = 1
	CHAR   = 2
	CART   = 2
)

type latch struct {
	kernal    int
	basic     int
	char_io_r int
	char_io_w int
	ram       int
}

type Cell struct {
	read    *int
	write   *int
	Zone    [3]byte
	IsRead  bool
	IsWrite bool
}

// Memory :
type Memory struct {
	PLA     latch
	Mem     [memorySize]Cell
	CD_Rom  []Cell
	EF_Rom  []Cell
	Basic   []Cell
	CharGen []Cell
	Stack   []Cell
	Screen  []Cell
	Vic     [4][]Cell
}
