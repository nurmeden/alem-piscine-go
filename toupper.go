package piscine

func ToUpper(s string) string {
	s_rune := []rune(s)
	for index, elem := range s_rune {
		if elem >= 'a' && elem <= 'z' {
			s_rune[index] = elem - 32
		}
	}
	return string(s_rune)
}
