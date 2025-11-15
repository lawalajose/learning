package piscine

func AtoiBase(s string, base string) int {
	if !isValidBaseAtoi(base) {
		return 0
	}

	baseLen := len(base)
	result := 0

	for _, char := range s {
		index := indexInBase(char, base)
		if index == -1 {
			return 0
		}
		result = result*baseLen + index
	}

	return result
}

func isValidBaseAtoi(base string) bool {
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

func indexInBase(c rune, base string) int {
	for i, b := range base {
		if c == b {
			return i
		}
	}
	return -1
}
