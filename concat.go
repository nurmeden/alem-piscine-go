package piscine

func Concat(str1 string, str2 string) string {
	runes_1 := []rune(str1)
	runes_2 := []rune(str2)
	return string(runes_1) + string(runes_2)
}
