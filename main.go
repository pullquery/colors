package main

import (
	"log"
)

func main() {
	old := makeRaw()
	hideCursor()
	clearDisplay()

	defer func() {
		restoreRaw(old)
		restoreCursor()
		clearDisplay()
	}()

	color := WHITE
	fillColor(color)

	for {
		controlKey(getKey(), &color)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
