package piscine

func IterativePower(nb int, power int) int {
	if power >= 0 {
		count := 1
		for i := 0; i < power; i++ {
			count *= nb
		}
		return count
	}
	return 0
}
