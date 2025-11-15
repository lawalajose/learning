package piscine

func Rot14(s string) string {
	var bee []rune
	for _, r := range s {
		bee = append(bee, ConV(r))
	}
	return string(bee)
}

func ConV(r rune) rune {
	counter := 14
	if r >= 65 && r <= 90 {
		for i := 0; i < counter; i++ {
			if r < 90 {
				r = r + 1
			} else {
				r = 65
			}
		}
	}
	if r >= 97 && r <= 122 {
		for i := 0; i < counter; i++ {
			if r < 122 {
				r = r + 1
			} else {
				r = 97
			}
		}
	}
	return r
}
