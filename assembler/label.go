package assembler

import (
	"errors"
	"fmt"
	"go6502/globals"
	"strconv"
)

type Labels struct {
	name    string
	size    int
	value16 globals.Word
	value8  globals.Byte
	export  string
}

func (L *Labels) setValueWord(val globals.Word) {
	L.value16 = globals.Word(val)
	L.size = 2
	L.export = fmt.Sprintf("%02X %02X", globals.Byte(L.value16), globals.Byte(L.value16>>8))
}

func (L *Labels) setValueByte(val globals.Byte) {
	L.value8 = globals.Byte(val)
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
		L.setValueWord(globals.Word(tmp))
	} else if len(val) == 2 {
		if tmp, err = strconv.ParseUint(val, 16, 8); err != nil {
			panic("Parse error")
		}
		L.setValueByte(globals.Byte(tmp))
	} else {
		panic("Bad Value")
	}
}

func (L *Labels) getRelative(pc globals.Word) (string, error) {
	if L.size != 2 {
		return "", errors.New(fmt.Sprintf("Bad value size : %s", L.export))
	}
	dist := ((int(pc) - int(L.value16)) * -1)
	// fmt.Printf("Calc : %04X - %04X * -1 = %d\n", pc, L.value16, dist)
	if dist > 0x00FF {
		return "", errors.New(fmt.Sprintf("Branch out of bound : %04X / %d", dist, int(dist)))
	}
	return fmt.Sprintf("%02X", globals.Byte(dist)), nil
}

func (L *Labels) getString() string {
	return L.export
}
