package piscine

func ListPushFront(l *List, data interface{}) {
	result := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = result
		l.Tail = result
	} else {
		result := &NodeL{Data: data}
		result.Next = l.Head
		l.Head = result
	}
}
