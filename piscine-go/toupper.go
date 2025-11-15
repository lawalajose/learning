package piscine

func ToUpper(s string) string {
	result := []rune(s)
	for index, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			result[index] = ch - 32
		}
	}
	return string(result)
}
