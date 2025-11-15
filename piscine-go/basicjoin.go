package piscine

func BasicJoin(elems []string) string {
	result := ""
	for _, ch := range elems {
		result = result + ch
	}
	return result
}
