package assembler

import (
	"bufio"
	"fmt"
	"go6502/cpu"
	"go6502/globals"
	"go6502/mem"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	macroDef = `^([a-zA-Z]\w+|\*)\s*([=:])(?:\s*\$([0-9a-fA-F]{2}|[0-9a-fA-F]{4}))?(?:\s.*)?$`
	cmdLine  = `^(?:(?:[a-zA-Z]\w+):\s*)?(?:(?:(\w+)\s*)?([^\s;=]*)?\s*(?:;.*)?)?$`
	addrMode = `^\*([\+\-]\d+)|(\(?)(#?)(?:\$([0-9a-fA-F]{2}|[0-9a-fA-F]{4})|([a-zA-Z]\w+))(,[XYxy]|,[Xx]\)|\),[Yy]|\))?$`
)

type Assembler struct {
	mem      *mem.Memory
	line     int
	prgStart globals.Word
	prgCount globals.Word
	labels   map[string]*Labels
	result   []string
}

func (A *Assembler) Init() {
	A.labels = make(map[string]*Labels)
}

func (A *Assembler) parseError(format string, vars ...interface{}) {
	fmt.Printf("ERR Line %d : %s\n", A.line, fmt.Sprintf(format, vars...))
	os.Exit(1)
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
	if _, ok := relatives[opCode]; ok {
		isRelatif = true
	}

	addrRe := regexp.MustCompile(addrMode)
	style := addrRe.FindStringSubmatch(val)

	if len(style[5]) > 0 && len(style[4]) == 0 {
		if isRelatif {
			var label *Labels
			var ok bool
			if label, ok = A.labels[style[5]]; !ok {
				A.parseError("Bad label -> %s", style[5])
			}
			if addr, err = label.getRelative(A.prgCount); err != nil {
				A.parseError("%s", err)
			}
		} else {
			addr = A.labels[style[5]].getString()
		}
	} else {
		if isRelatif {
			val, _ := strconv.ParseInt(style[1], 10, 8)
			addr = fmt.Sprintf("%02X", globals.Byte(val))
		} else {
			addr = style[4]
		}
	}

	if len(style) == 0 {
		return "", ""
	}
	if len(style[1]) > 0 {
		// Relatif (sans label)
		if isRelatif {
			suffix = "_REL"
		} else {
			A.parseError("Syntax Error -> %s", style[1])
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
			A.parseError("Bad Indirect addressing -> %s", style[6])
		}
	} else {
		// Direct
		switch style[6] {
		case ",X":
			if len(addr) > 2 {
				suffix = "_ABX"
			} else {
				suffix = "_ZPX"
			}
		case ",Y":
			if len(addr) > 2 {
				suffix = "_ABY"
			} else {
				suffix = "_ZPY"
			}
		default:
			// Relatif (avec label)
			if isRelatif {
				suffix = "_REL"
			} else if style[3] == "#" {
				suffix = "_IM"
			} else {
				if len(addr) > 2 {
					suffix = "_ABS"
				} else {
					suffix = "_ZP"
				}
			}
		}
	}
	return suffix, addr
}

func (A *Assembler) addOpCode(opCode string, addrFormat string) string {
	var codeLine string

	suffix, addr := A.checkAddr(opCode, addrFormat)
	// fmt.Printf("OpCode: %s%s - Addr: %s\n", opCode, suffix, addr)

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

	switch cmd[2] {
	case "=":
		if cmd[1] == "*" {
			if tmp, err = strconv.ParseUint(cmd[3], 16, 16); err != nil {
				panic("Parse error: Bad start address code")
			}
			A.prgStart = globals.Word(tmp)
			A.prgCount = globals.Word(tmp)
			fmt.Printf("Start Code at $%04X\n", A.prgStart)
		} else {
			newLab := Labels{name: cmd[1]}
			newLab.setValueString(cmd[3])
			A.labels[cmd[1]] = &newLab
			fmt.Printf("New Label: %s (%s)\n", A.labels[cmd[1]].name, A.labels[cmd[1]].export)
		}
	case ":":
		newLab := Labels{name: cmd[1]}
		newLab.setValueWord(A.prgCount)
		A.labels[cmd[1]] = &newLab
		fmt.Printf("New Label: %s (%s)\n", A.labels[cmd[1]].name, A.labels[cmd[1]].export)
	default:
		panic("Parse error")
	}
}

func (A *Assembler) computeOpCode(cmd []string) {
	var res string
	if len(cmd[1]) > 0 {
		fmt.Printf("Line %d: [%s] - ", A.line, cmd[0])
		res = A.addOpCode(cmd[1], cmd[2])
		A.result = append(A.result, res)
		fmt.Printf("Hexa: [%s]\n", res)
	}
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
	A.line = 0
	for scanner.Scan() {
		A.line++
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
	A.line = 0
	for scanner.Scan() {
		A.line++
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
			if len(cmd) == 0 {
				continue
			}
			// if len(cmd[1]) > 0 {
			// 	newLab := Labels{name: cmd[1]}
			// 	newLab.setValueWord(A.prgCount)
			// 	A.labels[cmd[1]] = &newLab
			// }
			if len(cmd[1]) > 0 {
				A.prgCount++
			}
			if len(cmd[2]) > 0 {
				A.prgCount++
			}
		}
	}
	A.prgCount = A.prgStart
	return nil
}

func (A *Assembler) Assemble(file string) string {
	if err := A.firstPass(file); err != nil {
		A.parseError("Error First pass Assembling: %s", err)
	}
	if err := A.secondPass(file); err != nil {
		A.parseError("Error Second pass Assembling: %s", err)
	}
	fmt.Println("==== End of process ====")
	log.Printf("Assembling complete without error\n")

	hexFile := strings.TrimSuffix(file, filepath.Ext(file)) + ".hex"
	tmp, err := os.Create(hexFile)
	if err != nil {
		A.parseError("%s", err)
	}
	content := fmt.Sprintf("%04X: %s", A.prgStart, strings.Join(A.result, " "))
	tmp.Write([]byte(content))
	tmp.Close()
	log.Printf("Dumping to %s\n", hexFile)

	return content
}
