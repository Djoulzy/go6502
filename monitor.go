package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func (C *CPU) output() {
	err := termbox.Init()
	termbox.HideCursor()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go func() {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Ch == 'q' {
				os.Exit(0)
			}
		}
	}()

	tbprintf(1, 0, termbox.ColorDefault, termbox.ColorDefault, "hh:mm:ss |  PC  | SP | A  | X  | Y  | NV-BDIZC")
	tbprintf(1, 1, termbox.ColorDefault, termbox.ColorDefault, "         | #### | ## | ## | ## | ## | 00000000")

	for {
		C.refreshScreen()
		time.Sleep(time.Second / 2)
	}
}

// DisplayHub : Affiche l'etat du Hub
func (C *CPU) refreshScreen() {
	t := time.Now()
	status := fmt.Sprintf("%s", t.Format("15:04:05"))
	tbprintf(1, 0, termbox.ColorDefault, termbox.ColorDefault, "%s", status)
	status = fmt.Sprintf("| %04X | %02X | %02X | %02X | %02X | %08b", C.PC, C.SP, C.A, C.X, C.Y, C.S)
	tbprintf(10, 1, termbox.ColorDefault, termbox.ColorDefault, "%s", status)
	status = fmt.Sprintf("%s", C.opName)
	tbprintf(1, 1, termbox.ColorDefault, termbox.ColorDefault, "%s", status)
	termbox.SetCursor(0, 2)
	err := termbox.Flush()
	if err != nil {
		panic("display")
	}
}

// This function is often useful:
func tbprintf(x, y int, fg, bg termbox.Attribute, format string, vars ...interface{}) {
	for _, c := range fmt.Sprintf(format, vars...) {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}
