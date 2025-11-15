package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		z01.PrintRune('\n')
		return
	}

	var text []rune
	for i, arg := range args {
		for _, r := range arg {
			text = append(text, r)
		}
		if i < len(args)-1 {
			text = append(text, ' ')
		}
	}

	var vowelPositions []int
	var vowels []rune
	for i, r := range text {
		if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
			r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' {
			vowelPositions = append(vowelPositions, i)
			vowels = append(vowels, r)
		}
	}

	if len(vowels) > 0 {
		for i, j := 0, len(vowels)-1; i < j; i, j = i+1, j-1 {
			vowels[i], vowels[j] = vowels[j], vowels[i]
		}
		for i, pos := range vowelPositions {
			text[pos] = vowels[i]
		}
	}

	for _, r := range text {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
