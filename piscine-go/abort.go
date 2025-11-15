package piscine

func Abort(a, b, c, d, e int) int {
	kk := []int{a, b, c, d, e}
	larg := len(kk)
	ll := (larg + 1) / 2
	for i := 0; i < larg-1; i++ {
		for j := 0; j < larg-i-1; j++ {
			if kk[j] > kk[j+1] {
				kk[j], kk[j+1] = kk[j+1], kk[j]
			}
		}
	}
	return kk[ll-1]
}
