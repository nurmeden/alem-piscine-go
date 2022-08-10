package piscine

func Join(elems []string, sep string) string {
	str := ""
	for i := 0; i <= len(elems)-1; i++ {
		str += elems[i]
		if len(elems)-1 != i {
			str += sep
		}
	}
	return str
}
