package piscine

func ListPushFront(l *List, data interface{}) {
	newnode := &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = newnode
		l.Tail = newnode
	} else {
		newnode.Next = l.Head
		l.Head = newnode
	}
}
