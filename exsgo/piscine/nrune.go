package piscine

func NRune(s string, n int) rune {
	var res_str rune
	for index, value := range s {
		if index+1 == n {
			res_str += value
		}
	}
	return res_str
}
