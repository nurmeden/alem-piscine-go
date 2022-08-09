package piscine

func IsPrintable(s string) bool {
	s_rune := []rune(s)
	count := 0
	for _, elem := range s_rune {
		if elem >= 'a' && elem <= 'z' || elem >= '0' && elem <= '9' || elem >= 'A' && elem <= 'Z' {
			count++
		}
	}
	return count == len(s)
}
