package piscine

func IterativeFactorial(nb int) int {
	if nb > 0 {
		count := 1
		for index := 1; index <= nb; index++ {
			count *= index
		}
		return count
	}
	return 0
}
