package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	upper := false
	start := 0

	// Check for --upper flag
	if args[0] == "--upper" {
		upper = true
		start = 1
	}

	for i := start; i < len(args); i++ {
		n := Atoi(args[i])
		if n >= 1 && n <= 26 {
			letter := rune('a' + n - 1)
			if upper {
				letter = letter - 32
			}
			z01.PrintRune(letter)
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

// Atoi converts string to int manually (without strconv)
func Atoi(s string) int {
	result := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 0
		}
		result = result*10 + int(ch-'0')
	}
	return result
}
