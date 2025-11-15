package piscine

func ToLower(s string) string {
	result := []rune(s)
	for index, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			result[index] = ch + 32
		}
	}
	return string(result)
}
