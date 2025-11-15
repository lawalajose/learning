package piscine

func FindNextPrime(nb int) int {
	bolle := true
	if nb < 2 {
		bolle = false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			bolle = false
		}
	}
	if bolle == true {
		return nb
	} else {
		nb = FindNextPrime(nb + 1)
		return nb
	}
}
