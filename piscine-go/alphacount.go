package piscine

func AlphaCount(s string) int {
	count := 0
	for _, ch := range s {
		if ch >= 65 && ch <= 90 || ch >= 97 && ch <= 122 {
			count++
		}
	}
	return count
}
