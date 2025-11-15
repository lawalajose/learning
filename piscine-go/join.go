package piscine

func Join(strs []string, sep string) string {
	result := ""
	for index, ch := range strs {
		result = result + ch
		if !(index == len(strs)-1) {
			result = result + sep
		}
	}
	return result
}
