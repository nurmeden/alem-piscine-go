package piscine

func StrLen(s string) int {
	count := 0
	for index := range s {
		index++
		count++
	}
	return count
}
