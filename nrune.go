package piscine

func NRune(s string, n int) rune {
	if n > 0 {
		a := []rune(s)
		for index := range a {
			if index == n-1 {
				return rune(a[n-1])
			}
		}
	}
	return 0
}
