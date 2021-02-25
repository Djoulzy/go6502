package assembler

import (
	"bufio"
	"fmt"
	"go6502/cpu"
	"go6502/globals"
	"go6502/mem"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	startMem = `\s*\*\s*=\s*#?\$(\w{4})`
	cmdLine  = `\s*(?:(\w*)\s+)?(\w+)\s*(\(?#?\$?\w{2,4}\)?(?:,[XY])?\)?)?`
	addrMode = `(\(?)(#?\$?)(\w{2,4})(,X\)?|\)?,Y|\))?`
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
	addrRe := regexp.MustCompile(addrMode)
	style := addrRe.FindStringSubmatch(val)

	addrSuffix := ""
	if style[1] == "(" { // Mode indirect
		if style[2] != "$" {
			panic("Parsing error")
		}
		if style[4] == "),Y" {
			if len(style[3]) != 2 {
				panic("Parsing error")
			}
			addrSuffix = "_INY"
		} else if style[4] == ",X)" {
			if len(style[3]) != 2 {
				panic("Parsing error")
			}
			addrSuffix = "_INX"
		} else if style[4] == ")" {
			addrSuffix = "_IND"
		} else {
			panic("Parsing error")
		}
	} else { // mode direct
		if style[2][:1] == "#" {
			if len(style[4]) > 0 {
				panic("Parsing error")
			}
			addrSuffix = "_IM"
		} else if style[2][:1] == "$" {
			if len(style[3]) == 2 {
				if style[4] == ",X" {
					addrSuffix = "_ZPX"
				} else if style[4] == ",Y" {
					addrSuffix = "_ZPY"
				} else if style[4] == "" {
					addrSuffix = "_ZP"
				} else {
					panic("Parsing error")
				}
			} else if len(style[3]) == 4 {
				addrSuffix = "_ABS"
			} else {
				panic("Parsing error")
			}
		}
	}

	fmt.Printf("Val: %s Addr: %s\n", val, addrSuffix)
	// if _, ok := A.labels[val]; ok {
	// 	A.getLabel(val)
	// 	return
	// }
	// if len(val) == 2 {
	// 	addr, _ := strconv.ParseUint(val, 16, 8)
	// 	A.mem.Data[int(A.prgCount)] = globals.Byte(addr)
	// 	A.prgCount++
	// } else if len(val) == 4 {
	// 	addr, _ := strconv.ParseUint(val, 16, 16)
	// 	A.mem.Data[int(A.prgCount)] = globals.Byte(addr)
	// 	A.prgCount++
	// 	A.mem.Data[int(A.prgCount)] = globals.Byte(addr >> 8)
	// 	A.prgCount++
	// } else {
	// 	panic("Parsing error")
	// }
}

func (A *Assembler) addInstr(hexa globals.Byte, addr string) {
	A.mem.Data[int(A.prgCount)] = hexa
	A.prgCount++
	if len(addr) > 0 {
		A.addAddr(addr)
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
		if len(txt) == 0 {
			continue
		}
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
			fmt.Printf("---------------------\nLine: %v\n", test)
			if test[1] != "" {
				if val, ok := cpu.CodeAddr[test[1]]; ok {
					A.addInstr(val, test[2])
				} else {
					A.setLabel(test[1])
				}
			}
			if test[2] != "" {
				if val, ok := cpu.CodeAddr[test[2]]; ok {
					A.addInstr(val, test[3])
				} else {
					A.getLabel(test[2])
				}
			}
			// if test[3] != "" {
			// 	A.addAddr(test[3])
			// }
		}
	}
	// A.mem.Dump(A.prgStart)
	log.Println("Assembling complete without error")
	return A.prgStart, nil
}
