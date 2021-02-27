package assembler

import (
	"bufio"
	"fmt"
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
	cmdLine  = `^(?:([a-zA-Z]\w+:)\s*)?(?:(?:(\w+)\s*)?([^\s;]*)?\s*(?:;.*)?)?$`
	addrMode = `^(\(?)(?:#?\$?([0-9a-fA-F]{2}|[0-9a-fA-F]{4})|([a-zA-Z]\w+))(,[XYxy]|,[Xx]\)|\),[Yy]|\))?$`
)

type Assembler struct {
	mem      *mem.Memory
	prgStart globals.Word
	prgCount globals.Word
	labels   map[string]globals.Word
}

func (A *Assembler) Init() {
	A.labels = make(map[string]globals.Word)
}

func (A *Assembler) setLabel(lab string) {
	fmt.Printf("SetLabel: %s\n", lab)
	A.labels[lab] = A.prgCount
}

func (A *Assembler) getLabel(lab string) string {
	// A.mem.Data[int(A.prgCount)] = globals.Byte(A.labels[lab])
	// A.prgCount++
	// A.mem.Data[int(A.prgCount)] = globals.Byte(A.labels[lab] >> 8)
	// A.prgCount++
	addr := fmt.Sprintf("%X %X", globals.Byte(A.labels[lab]), globals.Byte(A.labels[lab] >> 8)
}

func (A *Assembler) addOpCode(opCode string) {
	fmt.Printf("OpCode: %s\n", opCode)
}

func (A *Assembler) addAddr(val string) (string, string) {
	var suffix, addr string

	addrRe := regexp.MustCompile(addrMode)
	style := addrRe.FindStringSubmatch(val)
	fmt.Printf("Addr: %v\n", style)

	if style[1] == "(" {
		// Indirect
		switch style[4] {
		case ",X)":
			suffix = "_INX"
		case "),Y":
			suffix = "_INY"
		case ")":
			suffix = "_IND"
		default:
			suffix = ""
			panic("Parsing error")
		}
	} else {
		// Direct
		switch style[4] {
		case ",X":
			suffix = "_ABX"
		case ",Y":
			suffix = "_ABY"
		default:
			suffix = "_ABS"
		}
	}
	if len(style[3]) > 0 && len(style[2]) == 0 {
		addr = A.labels[style[3]]
	} else {
		addr = style[2]
	}
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

	fmt.Printf("Addr: %s - Prefix: %s\n", addr, suffix)
	return suffix, addr
}

func (A *Assembler) addInstr(hexa globals.Byte, addr string) {
	A.mem.Data[int(A.prgCount)] = hexa
	A.prgCount++
	if len(addr) > 0 {
		A.addAddr(addr)
	}
}

func (A *Assembler) computeMacro(cmd []string) {
	tmp, _ := strconv.ParseUint(cmd[1], 16, 16)
	A.prgStart = globals.Word(tmp)
	A.prgCount = globals.Word(tmp)
}

func (A *Assembler) computeOpCode(cmd []string) {
	fmt.Printf("---------------------\nLine: [%s]\n", cmd[0])
	if len(cmd[1]) > 0 {
		A.setLabel(cmd[1])
	}
	if len(cmd[2]) > 0 {
		A.addOpCode(cmd[2])
	}
	if len(cmd[3]) > 0 {
		A.addAddr(cmd[3])
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
		txt := strings.TrimSpace(scanner.Text())
		if len(txt) == 0 {
			continue
		}
		if txt == ".END" {
			break
		}
		cmd := startRe.FindStringSubmatch(txt)
		if len(cmd) > 0 {
			A.computeMacro(cmd)
		} else {
			cmd = cmdRe.FindStringSubmatch(txt)
			if len(cmd) > 0 {
				A.computeOpCode(cmd)
			} else {
				panic("Parsing error")
			}
		}
	}
	// A.mem.Dump(A.prgStart)
	log.Println("\n\nAssembling complete without error\n\n")
	return A.prgStart, nil
}
