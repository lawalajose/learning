package piscine

func Any(f func(string) bool, a []string) bool {
	var kk bool
	for _, r := range a {
		kk = f(r)
		if kk {
			return true
		}
	}
	return false
}
