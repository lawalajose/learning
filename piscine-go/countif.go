package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	var kk bool
	for _, r := range tab {
		kk = f(r)
		if kk {
			count++
		}
	}
	return count
}
