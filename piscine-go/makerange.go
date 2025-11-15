package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	gg := max - min
	lam := make([]int, gg)
	for i := 0; i < gg; i++ {
		lam[i] = min + i
	}
	return lam
}
