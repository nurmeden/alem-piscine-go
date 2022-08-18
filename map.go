package piscine

func Map(f func(int) bool, a []int) []bool {
	var bool_arr []bool
	for _, value := range a {
		bool_arr = append(bool_arr, f(value))
	}
	return bool_arr
}
