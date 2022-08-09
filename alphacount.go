package piscine

func AlphaCount(s string) int {
	count := 0
	runes := []rune(s)
	for _, value := range runes {
		if value >= 'a' && value <= 'z' || value >= 'A' && value <= 'Z' {
			count++
		}
	}
	return count
}
