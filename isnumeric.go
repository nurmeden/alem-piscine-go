package piscine

func IsNumeric(s string) bool {
	s_rune := []rune(s)
	count := 0
	for _, elem := range s_rune {
		if elem >= '0' && elem <= '9' {
			count++
		}
	}
	return count == len(s)
}
