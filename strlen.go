package piscine

func StrLen(s string) int {
	name := []rune(s)
	result := 0
	for i := range name {
		i++
		result++
	}
	return result
}
