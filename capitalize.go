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

func Capitalize(s string) string {
	result := true
	s_rune := []rune(ToLower(s))
	for i := 0; i < len(s_rune); i++ {
		if s_rune[i] >= 'a' && s_rune[i] <= 'z' && result == true {
			s_rune[i] = s_rune[i] - 32
			result = false
		}
		if s_rune[i] < 'A' || (s_rune[i] > 'Z' && s_rune[i] < 'a') || s_rune[i] > 'z' || (s_rune[i] >= '0' && s_rune[i] <= '9') {
			if s_rune[i] >= '0' && s_rune[i] <= '9' {
				result = false
			} else {
				result = true
			}
		}
	}
	return string(s_rune)
}
