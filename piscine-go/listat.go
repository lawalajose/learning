package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil || pos < 0 {
		return nil
	}
	if pos == 0 {
		return l
	}
	count := 0
	for l != nil {
		if count == pos {
			return l
		}
		l = l.Next
		count++
	}
	return l
}
