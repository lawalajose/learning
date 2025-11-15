package piscine

func Atoi(s string) int {
	sign := 1
	result := 0
	for i, ch := range s {
		if i == 0 && (ch == '-' || ch == '+') {
			if ch == '-' {
				sign = -1
			}
			continue
		}
		if ch < '0' || ch > '9' {
			return 0
		}
		result = result*10 + int(ch-'0')
	}
	return result * sign
}
