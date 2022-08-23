package piscine

func ListSize(l *List) int {
	count := 0
	if l.Head != nil {
		l.Head = l.Head.Next
		count++
	}
	return count
}
