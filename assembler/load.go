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
	macroDef = `^([a-zA-Z]\w+|\*)\s*([=:])(?:\s*\$([0-9a-fA-F]{2}|[0-9a-fA-F]{4})$|$)`
	cmdLine  = `^(?:([a-zA-Z]\w+):\s*)?(?:(?:(\w+)\s*)?([^\s;]*)?\s*(?:;.*)?)?$`
	addrMode = `^\*([\+\-]\d+)|(\(?)(#?)(?:\$([0-9a-fA-F]{2}|[0-9a-fA-F]{4})|([a-zA-Z]\w+))(,[XYxy]|,[Xx]\)|\),[Yy]|\))?$`
)

type Assembler struct {
	mem      *mem.Memory
	prgStart globals.Word
	prgCount globals.Word
	labels   map[string]*Labels
	result   []string
}

func (A *Assembler) Init() {
	A.labels = make(map[string]*Labels)
}

func (A *Assembler) checkAddr(opCode, val string) (string, string) {
	var err error
	var suffix string
	var addr string
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
		return "", ""
	}
	if len(style[1]) > 0 {
		// Relatif (sans label)
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
			// Relatif (avec label)
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
		if isRelatif {
			if addr, err = A.labels[style[5]].getRelative(A.prgCount); err != nil {
				panic("Bad label")
			}
		} else {
			addr = A.labels[style[5]].getString()
		}
	} else {
		if isRelatif {
			addr = fmt.Sprintf("%02X", style[1])
		} else {
			addr = style[4]
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

	codeLine = fmt.Sprintf("%s %s", codeLine, addr)
	A.prgCount++
	return codeLine
}

func (A *Assembler) computeMacro(cmd []string) {
	var tmp uint64
	var err error

	if cmd[1] == "*" {
		if tmp, err = strconv.ParseUint(cmd[2], 16, 16); err != nil {
			panic("Parse error")
		}
		A.prgStart = globals.Word(tmp)
		A.prgCount = globals.Word(tmp)
		fmt.Printf("Start Code at $%04X\n", A.prgStart)
	} else {
		if cmd[2] == "=" {
			newLab := Labels{name: cmd[1]}
			newLab.setValueString(cmd[3])
			A.labels[cmd[1]] = &newLab
			fmt.Printf("New Label: %s (%s)\n", A.labels[cmd[1]].name, A.labels[cmd[1]].export)
		}
	}
}

func (A *Assembler) computeOpCode(cmd []string) {
	var res string
	fmt.Printf("---------------------\nLine: [%s]\n", cmd[0])
	if len(cmd[2]) > 0 {
		res = A.addOpCode(cmd[2], cmd[3])
		A.result = append(A.result, res)
	}
	fmt.Printf("Hexa: [%s]\n", res)
}

func (A *Assembler) secondPass(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	cmdRe := regexp.MustCompile(cmdLine)

	fmt.Println("==== Second Pass ====")
	for scanner.Scan() {
		txt := strings.TrimSpace(scanner.Text())
		if len(txt) == 0 {
			continue
		}
		if txt == ".END" {
			break
		}

		cmd := cmdRe.FindStringSubmatch(txt)
		if len(cmd) > 0 {
			A.computeOpCode(cmd)
		}
	}
	// A.mem.Dump(A.prgStart)
	fmt.Println()
	log.Println("Assembling complete without error")
	fmt.Println()
	fmt.Printf("%s", A.result)
	return nil
}

func (A *Assembler) firstPass(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	macroRe := regexp.MustCompile(macroDef)
	cmdRe := regexp.MustCompile(cmdLine)

	fmt.Println("==== First Pass ====")
	for scanner.Scan() {
		txt := strings.TrimSpace(scanner.Text())
		if len(txt) == 0 {
			continue
		}
		if txt == ".END" {
			break
		}
		cmd := macroRe.FindStringSubmatch(txt)
		if len(cmd) > 0 {
			A.computeMacro(cmd)
		} else {
			cmd = cmdRe.FindStringSubmatch(txt)
			if len(cmd[1]) > 0 {
				newLab := Labels{name: cmd[1]}
				newLab.setValueWord(A.prgCount)
				A.labels[cmd[1]] = &newLab
			}
			if len(cmd[2]) > 0 {
				A.prgCount++
			}
			if len(cmd[3]) > 0 {
				A.prgCount++
			}
		}
	}
	fmt.Printf("Program Count: %04X\n", A.prgCount)
	A.prgCount = A.prgStart
	return nil
}

func (A *Assembler) Assemble(file string) error {
	A.firstPass(file)
	fmt.Printf("%v\n", A.labels)
	A.secondPass(file)

	return nil
}
