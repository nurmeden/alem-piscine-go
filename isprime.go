package piscine

func IsPrime(nb int) bool {
	if nb > 2 {
		for elem := 2; elem < nb; elem++ {
			if nb%elem == 0 {
				return false
			}
		}
		return true
	}
	if nb == 2 {
		return true
	}
	return false
}
