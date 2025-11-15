package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var x int
	for i := range a {
		if i > 0 {
			if f(a[i], a[i-1]) == 0 {
				continue
			}
			if x == 0 {
				x = f(a[i], a[i-1])
			} else if (x > 0 && f(a[i], a[i-1]) < 0) || (x < 0 && f(a[i], a[i-1]) > 0) {
				return false
			}
			x = f(a[i], a[i-1])
		}
	}
	return true
}
