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
	cmdLine  = `^(?:([a-zA-Z]\w+:)\s*)?(?:(?:(\w+)\s*)?([^\s;]*)?\s*(?:;.*)?)?$`
	addrMode = `^\*([\+\-]\d+)|(\(?)(?:(#?)\$?([0-9a-fA-F]{2}|[0-9a-fA-F]{4})|([a-zA-Z]\w+))(,[XYxy]|,[Xx]\)|\),[Yy]|\))?$`
)

type Assembler struct {
	mem      *mem.Memory
	prgStart globals.Word
	prgCount globals.Word
	labels   map[string]globals.Word
	result   []string
}

func (A *Assembler) Init() {
	A.labels = make(map[string]globals.Word)
}

func (A *Assembler) setLabel(lab string) {
	fmt.Printf("SetLabel: %s\n", lab)
	A.labels[lab] = A.prgCount
}

func (A *Assembler) getLabel(lab string, relative bool) []string {
	var addr []string
	labelAddr := A.labels[lab]
	if relative {
		labelAddr = globals.Word((int(A.prgCount) - int(labelAddr)) * -1)
		addr = append(addr, fmt.Sprintf("%X", globals.Byte(labelAddr)))
		return addr
	}

	addr = append(addr, fmt.Sprintf("%X", globals.Byte(labelAddr)))
	addr = append(addr, fmt.Sprintf("%X", globals.Byte(labelAddr>>8)))
	return addr
}

func (A *Assembler) checkAddr(opCode, val string) (string, []string) {
	var suffix string
	var addr []string
	var isRelatif bool = false

	var relatives = map[string]bool{
		"BCC": true,
		"BCS": true,
		"BEQ": true,
		"BMI": true,
		"BNE": true,
		"BPL": true,
		"BVC": true,
		"BVS": true,
	}

	addrRe := regexp.MustCompile(addrMode)
	style := addrRe.FindStringSubmatch(val)
	if len(style) == 0 {
		return "", nil
	}
	if len(style[1]) > 0 {
		// Relatif
		if _, ok := relatives[opCode]; ok {
			suffix = "_REL"
			isRelatif = true
		} else {
			panic("Syntax Error")
		}
	} else if style[2] == "(" {
		// Indirect
		switch style[6] {
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
		switch style[6] {
		case ",X":
			suffix = "_ABX"
		case ",Y":
			suffix = "_ABY"
		default:
			if _, ok := relatives[opCode]; ok {
				suffix = "_REL"
				isRelatif = true
			} else if style[3] == "#" {
				suffix = "_IM"
			} else {
				suffix = "_ABS"
			}
		}
	}
	if len(style[5]) > 0 && len(style[4]) == 0 {
		addr = A.getLabel(style[5], isRelatif)
	} else {
		if isRelatif {
			addr = append(addr, fmt.Sprintf("%02X", style[1]))
		} else {
			addr = append(addr, fmt.Sprintf("%02X", style[4]))
		}
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

	return suffix, addr
}

// func (A *Assembler) addInstr(hexa globals.Byte, addr string) {
// 	A.mem.Data[int(A.prgCount)] = hexa
// 	A.prgCount++
// 	if len(addr) > 0 {
// 		A.addAddr(addr)
// 	}
// }

func (A *Assembler) addOpCode(opCode string, addrFormat string) string {
	var codeLine string

	suffix, addr := A.checkAddr(opCode, addrFormat)
	fmt.Printf("OpCode: %s%s - Addr: %s\n", opCode, suffix, addr)

	if val, ok := cpu.CodeAddr[opCode+suffix]; ok {
		codeLine = fmt.Sprintf("%02X", val)
		A.prgCount++
	} else {
		panic("Syntax Error")
	}

	codeLine = fmt.Sprintf("%s %s", codeLine, strings.Join(addr[:], " "))
	A.prgCount++
	return codeLine
}

func (A *Assembler) computeMacro(cmd []string) {
	tmp, _ := strconv.ParseUint(cmd[1], 16, 16)
	A.prgStart = globals.Word(tmp)
	A.prgCount = globals.Word(tmp)
}

func (A *Assembler) computeOpCode(cmd []string) {
	var res string
	fmt.Printf("---------------------\nLine: [%s]\n", cmd[0])
	if len(cmd[1]) > 0 {
		A.setLabel(cmd[1])
	}
	if len(cmd[2]) > 0 {
		res = A.addOpCode(cmd[2], cmd[3])
		A.result = append(A.result, res)
	}
	fmt.Printf("Hexa: [%s]\n", res)
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
	fmt.Println()
	log.Println("Assembling complete without error")
	fmt.Println()
	fmt.Printf("%s", A.result)
	return A.prgStart, nil
}
