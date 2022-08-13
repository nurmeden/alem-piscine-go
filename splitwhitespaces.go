package piscine

func SplitWhiteSpaces(str string) []string {
	var result []string
	word := ""
	for index, value := range str {
		if (value == ' ' || value == '\n' || value == '\t') && len(word) > 0 {
			result = append(result, word)
			word = ""
		}
		if value != ' ' {
			word = word + string(value)
		}
		if index == len(str)-1 {
			result = append(result, word)
		}
	}
	return result
}
