package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func PrintNumber(n int) {
	number := '0'
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		PrintNumber(n / 10)
	}
	modulus := n % 10
	for i := 0; i < modulus; i++ {
		number++
	}
	z01.PrintRune(number)
}
