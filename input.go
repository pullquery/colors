package main

import (
	"golang.org/x/term"
	"os"
)

func makeRaw() *term.State {
	old, err := term.MakeRaw(int(os.Stdin.Fd()))
	checkErr(err)
	return old
}

func restoreRaw(old *term.State) {
	err := term.Restore(int(os.Stdin.Fd()), old)
	checkErr(err)
}
