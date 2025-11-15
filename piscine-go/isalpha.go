package piscine

func IsAlpha(s string) bool {
	for _, ch := range s {
		if !((ch >= 'A' && ch <= 'Z') ||
			(ch >= 'a' && ch <= 'z') ||
			(ch >= '0' && ch <= '9')) {
			return false
		}
	}
	return true
}
