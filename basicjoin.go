package piscine

func BasicJoin(elems []string) string {
	str := ""
	for _, value := range elems {
		str += value
	}
	return str
}
