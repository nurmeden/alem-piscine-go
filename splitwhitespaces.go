package piscine

func SplitWhiteSpaces(str string) []string {
	var result []string
	word := ""
	for index, value := range str {
		if value == ' ' {
			result = append(result, word)
			word = ""
		}
		word = word + string(value)
		if index == len(str)-1 {
			result = append(result, word)
		}
	}
	return result
}
