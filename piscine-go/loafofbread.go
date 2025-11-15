package piscine

func LoafOfBread(str string) string {
	runes := []rune(str)
	nonSpaceCount := 0

	for _, r := range runes {
		if r != ' ' {
			nonSpaceCount++
		}
	}

	if str == "" || nonSpaceCount == 0 {
		return "\n"
	}
	if nonSpaceCount < 5 {
		return "Invalid Output\n"
	}

	var groups []string
	var buf []rune
	i := 0

	for i < len(runes) {
		r := runes[i]
		if r == ' ' {
			i++
			continue
		}
		buf = append(buf, r)
		i++
		if len(buf) == 5 {
			groups = append(groups, string(buf))
			buf = buf[:0]
			if i < len(runes) {
				i++
			}
		}
	}

	if len(buf) > 0 {
		groups = append(groups, string(buf))
	}

	out := ""
	for idx, g := range groups {
		if idx > 0 {
			out += " "
		}
		out += g
	}
	out += "\n"
	return out
}
