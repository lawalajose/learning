package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Step 1: Convert from baseFrom to decimal
	decimal := 0
	baseLenFrom := len(baseFrom)

	for _, r := range nbr {
		decimal = decimal*baseLenFrom + findIndex(r, baseFrom)
	}

	// Step 2: Convert from decimal to baseTo
	if decimal == 0 {
		return string(baseTo[0])
	}

	baseLenTo := len(baseTo)
	result := ""
	for decimal > 0 {
		remainder := decimal % baseLenTo
		result = string(baseTo[remainder]) + result
		decimal /= baseLenTo
	}

	return result
}

func findIndex(r rune, base string) int {
	for i, v := range base {
		if v == r {
			return i
		}
	}
	return -1
}
