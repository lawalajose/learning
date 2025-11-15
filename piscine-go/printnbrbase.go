package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(n int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	baseLen := len(base)
	if n == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	// handle negative numbers
	if n < 0 {
		z01.PrintRune('-')
		// Instead of doing n = -n (which overflows for MinInt)
		printBase(-(n / baseLen), base, baseLen)
		z01.PrintRune(rune(base[-(n % baseLen)]))
		return
	}

	printBase(n, base, baseLen)
}

// helper function to check base validity
func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	for i, c := range base {
		if c == '+' || c == '-' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if c == rune(base[j]) {
				return false
			}
		}
	}
	return true
}

// recursive printer
func printBase(n int, base string, baseLen int) {
	if n >= baseLen {
		printBase(n/baseLen, base, baseLen)
	}
	z01.PrintRune(rune(base[n%baseLen]))
}
