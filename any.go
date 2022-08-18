package piscine

func Any(f func(string) bool, a []string) bool {
	result_bool := false
	for _, value := range a {
		result_bool = f(string(value))
		if result_bool == true {
			break
		}
	}
	return result_bool
}
