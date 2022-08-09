package piscine

func IsLower(s string) bool {
	s_rune := []rune(s)
	count := 0
	for _, elem := range s_rune {
		if elem >= 'a' && elem <= 'z' {
			count++
		}
	}
	return count == len(s)
}
