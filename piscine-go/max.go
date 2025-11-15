package piscine

func Max(a []int) int {
	var kk int
	if len(a) == 0 {
		return 0
	}
	for i := 0; i < len(a)-1; i++ {
		if a[i] > kk {
			kk = a[i]
		}
	}
	return kk
}
