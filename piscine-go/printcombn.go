package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	var digits [9]int
	for i := 0; i < n; i++ {
		digits[i] = i
	}

	for {
		for i := 0; i < n; i++ {
			z01.PrintRune(rune(digits[i] + '0'))
		}

		if digits[0] == 10-n {
			break
		}

		z01.PrintRune(',')
		z01.PrintRune(' ')

		for i := n - 1; i >= 0; i-- {
			if digits[i] < 9-(n-1-i) {
				digits[i]++
				for j := i + 1; j < n; j++ {
					digits[j] = digits[j-1] + 1
				}
				break
			}
		}
	}
	z01.PrintRune('\n')
}
