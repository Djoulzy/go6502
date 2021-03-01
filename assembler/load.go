package assembler

import (
	"fmt"
	"go6502/globals"
	"go6502/mem"
	"strconv"
	"strings"
)

func LoadFile(mem *mem.Memory, file string) {

}

func LoadHex(mem *mem.Memory, code string) (globals.Word, error) {
	data := strings.Fields(code)
	tmp, _ := strconv.ParseUint(data[0], 16, 16)
	start := globals.Word(tmp)
	fmt.Printf("Start: %04X", start)

	for i, val := range data[1:] {
		index := start + globals.Word(i)
		value, _ := strconv.ParseUint(val, 16, 8)
		mem.Data[index] = globals.Byte(value)
	}
	return start, nil
}
