package piscine

func StrRev(s string) string {
	b := []rune(s)
	n := len(b)
	for i := 0; i < n/2; i++ {
		b[i], b[n-1-i] = b[n-1-i], b[i]
	}
	return string(b)
}
