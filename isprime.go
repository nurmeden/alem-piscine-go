package piscine

func IsPrime(nb int) bool {
	if nb > 1 && nb < 100 {
		for elem := 2; elem < nb; elem++ {
			if nb%elem == 0 {
				return false
			} else {
				return true
			}
		}
	}
	return false
}
