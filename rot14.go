package piscine

func Rot14(s string) string {
	var result_str string
	for _, value := range s {
		if (value >= 65 && value <= 90) || (value >= 97 && value <= 122) {
			if (value <= 76 && value >= 65) || (value >= 97 && value <= 108) {
				result_str += string(value + 14)
			} else if (value >= 77 && value <= 90) || (value >= 109 && value <= 122) {
				result_str += string(value - 12)
			}
		} else {
			result_str += string(value)
		}
	}
	return result_str
}
