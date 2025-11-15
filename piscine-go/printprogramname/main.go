package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args[0]
	lastbackslash := -1
	for i, pp := range argument {
		if pp == '/' {
			lastbackslash = i
		}
	}
	for j, ch := range argument {
		if j > lastbackslash {
			z01.PrintRune(ch)
		}
	}
	z01.PrintRune('\n')
}
