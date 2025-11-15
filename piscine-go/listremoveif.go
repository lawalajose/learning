package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}

	// Remove matching nodes at the head
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	// If list becomes empty after removing head nodes
	if l.Head == nil {
		l.Tail = nil
		return
	}

	// Traverse the rest of the list
	prev := l.Head
	curr := l.Head.Next

	for curr != nil {
		if curr.Data == data_ref {
			prev.Next = curr.Next
			// If we removed the tail, update tail reference
			if curr.Next == nil {
				l.Tail = prev
			}
		} else {
			prev = curr
		}
		curr = curr.Next
	}
}
