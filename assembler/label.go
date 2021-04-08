package assembler

import (
	"errors"
	"fmt"
	"strconv"
)

type Labels struct {
	name    string
	size    int
	value16 uint16
	value8  byte
	export  string
}

func (L *Labels) setValueWord(val uint16) {
	L.value16 = uint16(val)
	L.size = 2
	L.export = fmt.Sprintf("%02X %02X", byte(L.value16), byte(L.value16>>8))
}

func (L *Labels) setValueByte(val byte) {
	L.value8 = byte(val)
	L.size = 1
	L.export = fmt.Sprintf("%02X", L.value8)
}

func (L *Labels) setValueString(val string) {
	var tmp uint64
	var err error

	if len(val) == 4 {
		if tmp, err = strconv.ParseUint(val, 16, 16); err != nil {
			panic("Parse error")
		}
		L.setValueWord(uint16(tmp))
	} else if len(val) == 2 {
		if tmp, err = strconv.ParseUint(val, 16, 8); err != nil {
			panic("Parse error")
		}
		L.setValueByte(byte(tmp))
	} else {
		panic("Bad Value")
	}
}

func (L *Labels) getRelative(pc uint16) (string, error) {
	if L.size != 2 {
		return "", errors.New(fmt.Sprintf("Bad value size : %s", L.export))
	}
	dist := ((int(pc) - int(L.value16)) * -1)
	// fmt.Printf("Calc : %d(%04X) - %d(%04X) * -1 = %d(%02X)\n", int(pc), pc, int(L.value16), L.value16, dist, byte(dist))
	if dist > 0x00FF {
		return "", errors.New(fmt.Sprintf("Branch out of bound : %04X / %d", dist, int(dist)))
	}
	return fmt.Sprintf("%02X", byte(dist)), nil
}

func (L *Labels) getString() string {
	return L.export
}
