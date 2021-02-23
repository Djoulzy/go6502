package assembler

import (
	"bufio"
	"go6502/cpu"
	"go6502/globals"
	"go6502/mem"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	startMem = `\s*\*\s*=\s*#?\$(\w{4})`
	cmdLine  = `\s*(?:(\w*)\s+)?(\w+)\s*(?:#?\$?(\w{2,4}))?`
)

type Assembler struct {
	mem      *mem.Memory
	prgStart globals.Word
	prgCount globals.Word
	labels   map[string]globals.Word
}

func (A *Assembler) Init(mem *mem.Memory) {
	A.mem = mem
	A.labels = make(map[string]globals.Word)
}

func (A *Assembler) addInstr(hexa globals.Byte) {
	A.mem.Data[int(A.prgCount)] = hexa
	A.prgCount++
}

func (A *Assembler) setLabel(lab string) {
	A.labels[lab] = A.prgCount
}

func (A *Assembler) getLabel(lab string) {
	A.mem.Data[int(A.prgCount)] = globals.Byte(A.labels[lab])
	A.prgCount++
	A.mem.Data[int(A.prgCount)] = globals.Byte(A.labels[lab] >> 8)
	A.prgCount++
}

func (A *Assembler) addAddr(val string) {
	if len(val) == 2 {
		addr, _ := strconv.ParseUint(val, 16, 8)
		A.mem.Data[int(A.prgCount)] = globals.Byte(addr)
		A.prgCount++
	} else if len(val) == 4 {
		addr, _ := strconv.ParseUint(val, 16, 16)
		A.mem.Data[int(A.prgCount)] = globals.Byte(addr)
		A.prgCount++
		A.mem.Data[int(A.prgCount)] = globals.Byte(addr >> 8)
		A.prgCount++
	} else {
		panic("Parsing error")
	}
}

func (A *Assembler) ReadCode(file string) (globals.Word, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	startRe := regexp.MustCompile(startMem)
	cmdRe := regexp.MustCompile(cmdLine)

	for scanner.Scan() {
		txt := strings.Trim(scanner.Text(), " ")
		if txt == ".END" {
			break
		}
		test := startRe.FindStringSubmatch(txt)
		if len(test) > 0 {
			tmp, _ := strconv.ParseUint(test[1], 16, 16)
			A.prgStart = globals.Word(tmp)
			A.prgCount = globals.Word(tmp)
		} else {
			test = cmdRe.FindStringSubmatch(txt)
			if test[1] != "" {
				if val, ok := cpu.CodeAddr[test[1]]; ok {
					A.addInstr(val)
				} else {
					A.setLabel(test[1])
				}
			}
			if test[2] != "" {
				if val, ok := cpu.CodeAddr[test[2]]; ok {
					A.addInstr(val)
				} else {
					A.getLabel(test[2])
				}
			}
			if test[3] != "" {
				A.addAddr(test[3])
			}
		}
	}
	A.mem.Dump(A.prgStart)
	return A.prgStart, nil
}
