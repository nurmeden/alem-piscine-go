package piscine

func Max(a []int) int {
	max_value := 0
	for _, value := range a {
		if max_value < value {
			max_value = value
		}
	}
	return max_value
}
