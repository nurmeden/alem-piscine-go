package piscine

func IsUpper(s string) bool {
	s_rune := []rune(s)
	count := 0
	for _, elem := range s_rune {
		if elem >= 'A' && elem <= 'Z' {
			count++
		}
	}
	return count == len(s)
}
