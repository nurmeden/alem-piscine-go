package piscine

func ToLower(s string) string {
	s_rune := []rune(s)
	for index, elem := range s_rune {
		if elem >= 'A' && elem <= 'Z' {
			s_rune[index] = elem + 32
		}
	}
	return string(s_rune)
}
