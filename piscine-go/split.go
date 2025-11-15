package piscine

func Split(s, sep string) []string {
	var result []string
	sepLen := len(sep)
	word := ""

	for i := 0; i < len(s); {
		if sepLen > 0 && i+sepLen <= len(s) && s[i:i+sepLen] == sep {
			result = append(result, word)
			word = ""
			i += sepLen
		} else {
			word += string(s[i])
			i++
		}
	}

	result = append(result, word)
	return result
}
