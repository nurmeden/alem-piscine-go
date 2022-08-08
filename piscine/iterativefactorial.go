package piscine

func IterativeFactorial(nb int) int {
	count := 1
	for index := 1; index <= nb; index++ {
		count *= index
	}
	return count
}
