package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args[1:]
	for i := 0; i < len(argument)-1; i++ {
		for j := 0; j < len(argument)-i-1; j++ {
			if argument[j] > argument[j+1] { // ASCII comparison
				argument[j], argument[j+1] = argument[j+1], argument[j]
			}
		}
	}
	for _, ch := range argument {
		for _, bb := range ch {
			z01.PrintRune(bb)
		}
		z01.PrintRune('\n')
	}
}
