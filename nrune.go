package piscine

func NRune(s string, n int) rune {
	if n < 0 || n > len(s) {
		return 0
	} else {
		a := []rune(s)
		return a[n-1]
	}
}
