package piscine

func LastRune(s string) rune {
	y := []rune(s)
	return y[len(y)-1]
}
