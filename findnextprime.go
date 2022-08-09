package piscine

func FindNextPrime(nb int) int {
	if nb > 0 {
		for elem := 2; elem < nb; elem++ {
			if nb%elem == 0 {
				nb = nb + 1
				FindNextPrime(nb)
			}
		}
		return nb
	}
	nb = 2
	return FindNextPrime(nb)
}
