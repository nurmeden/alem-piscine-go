package piscine

func ListReverse(l *List) {
	var massivInterface []interface{}

	for i := l.Head; i != nil; i = i.Next {
		massivInterface = append(massivInterface, i.Data)
	}

	links := l.Head

	for i := len(massivInterface) - 1; i >= 0; i-- {
		links.Data = massivInterface[i]
		links = links.Next
	}
}
