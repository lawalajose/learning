package piscine

func Capitalize(s string) string {
	result := []rune(s)
	newWord := true
	for i, ch := range result {
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
			if newWord {
				if ch >= 'a' && ch <= 'z' {
					result[i] = ch - 32
				}
				newWord = false
			} else {
				if ch >= 'A' && ch <= 'Z' {
					result[i] = ch + 32
				}
			}
		} else {
			newWord = true
		}
	}
	return string(result)
}
