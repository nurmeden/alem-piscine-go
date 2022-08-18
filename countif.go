package piscine

func CountIf(f func(string) bool, tab []string) int {
	result_bool := false
	count := 0
	for _, value := range tab {
		result_bool = f(string(value))
		if result_bool == true {
			count++
		}
	}
	return count
}
