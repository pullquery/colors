package main

import (
	"fmt"
	"strconv"
)

const (
	WHITE   = 231
	RED     = 196
	GREEN   = 28
	BLUE    = 19
	YELLOW  = 226
	MAGENTA = 198
	CYAN    = 45
)

func hideCursor() {
	fmt.Print("\033[?25l")
}

func restoreCursor() {
	fmt.Print("\033[?25h")
}

func clearDisplay() {
	fmt.Print("\033[2J" + "\033[H")
}

func fillColor(color int) {
	fmt.Printf("\033[48;5;%sm \033[0m", strconv.Itoa(color))
	moveBackward()
}

func eraseColor() {}
