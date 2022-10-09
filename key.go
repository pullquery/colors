package main

import (
	"fmt"
	"os"
)

func getKey() byte {
	b := make([]byte, 1)

	_, err := os.Stdin.Read(b)
	checkErr(err)
	return b[0]
}

func controlKey(key byte, color *int) {
	if key == CtrlC {
		os.Exit(0)
	}

	switch key {
	case 'W', 'w':
		moveUp()
		fillColor(*color)
	case 'S', 's':
		moveDown()
		fillColor(*color)
	case 'A', 'a':
		moveBackward()
		fillColor(*color)
	case 'D', 'd':
		moveForward()
		fillColor(*color)
	case '`':
		*color = WHITE
	case '1':
		*color = RED
	case '2':
		*color = GREEN
	case '3':
		*color = BLUE
	case '4':
		*color = YELLOW
	case '5':
		*color = MAGENTA
	case '6':
		*color = CYAN
	}
}

func moveUp()       { fmt.Print("\033[A") }
func moveDown()     { fmt.Print("\033[B") }
func moveForward()  { fmt.Print("\033[C") }
func moveBackward() { fmt.Print("\033[D") }
