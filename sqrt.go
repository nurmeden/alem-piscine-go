package piscine

func Sqrt(nb int) int {
	if nb > 0 {
		for elem := 0; elem <= nb; elem++ {
			if elem*elem == nb {
				return elem
			}
		}
	}
	return 0
}
