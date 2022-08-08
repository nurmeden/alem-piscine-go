package piscine

func Sqrt(nb int) int {
	if nb >= 0 && nb < 627558 {
		for nums := 2; nums <= nb; nums++ {
			if nb%nums == 0 {
				return nums
			}
			return 0
		}
	}
	return 0
}
