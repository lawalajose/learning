package piscine

func AppendRange(min, max int) []int {
	var lam []int
	for i := min; i < max; i++ {
		lam = append(lam, i)
	}
	return lam
}
