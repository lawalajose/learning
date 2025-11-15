package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args[1:]
	for _, ch := range argument {
		for _, bb := range ch {
			z01.PrintRune(bb)
		}
		z01.PrintRune('\n')
	}
}
