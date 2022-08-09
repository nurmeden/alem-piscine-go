package piscine

func IsPrintable(s string) bool {
	strs := []rune(s)
	for i := range strs {
		if strs[i] >= 0 && strs[i] <= 31 {
			return false
		}
	}
	return true
}
