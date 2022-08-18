package piscine

func CountIf(f func(string) bool, tab []string) int {
	result_bool := false
	count := 0
	for _, value := range tab {
		if result_bool == true {
			count++
		}
		result_bool = f(string(value))
	}
	return count
}
