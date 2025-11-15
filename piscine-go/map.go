package piscine

func Map(f func(int) bool, a []int) []bool {
	var kk []bool
	for _, r := range a {
		f(r)
		kk = append(kk, f(r))
	}
	return kk
}
