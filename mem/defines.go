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
)

type cell struct {
	romMode byte
	expMode byte
	exp     byte
	rom     byte
	ram     byte
}

type bank struct {
	rom   bool
	start uint16
	data  []byte
}

type memoryMap [4]bank

// Memory :
type Memory struct {
	bank    memoryMap
	Kernal  []byte
	Basic   []byte
	CharGen []byte
	Stack   []byte
	Screen  []byte
	Color   []byte
	Vic     [4][]byte
}
