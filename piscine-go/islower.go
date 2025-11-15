package piscine

func IsLower(s string) bool {
	for _, ch := range s {
		if ch < 97 || ch > 122 {
			return false
		}
	}
	return true
}
