package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	started := false
	for _, ch := range s {
		if ch == '-' && !started {
			sign = -1
		}
		if ch >= '0' && ch <= '9' {
			started = true
			result = result*10 + int(ch-'0')
		}
	}
	return result * sign
}
