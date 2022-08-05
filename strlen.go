package piscine

func StrLen(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		result++
	}
	return result
}
