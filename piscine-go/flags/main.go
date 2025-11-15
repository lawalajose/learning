package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		PrintHelp()
		return
	}

	var insertStr string
	var order bool
	var baseStr string

	for _, arg := range args {
		if len(arg) >= 9 && arg[:9] == "--insert=" {
			insertStr = arg[9:]
		} else if len(arg) >= 3 && arg[:3] == "-i=" {
			insertStr = arg[3:]
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else {
			baseStr = arg
		}
	}

	final := baseStr + insertStr

	if order {
		final = SortString(final)
	}

	for _, r := range final {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func PrintHelp() {
	text := `--insert
  -i
	 This flag inserts the string into the string passed as argument.
--order
  -o
	 This flag will behave like a boolean, if it is called it will order the argument.`
	for _, r := range text {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func SortString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}
