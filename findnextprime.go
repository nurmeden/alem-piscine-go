package piscine

func FindNextPrime(nb int) int {
	if nb > 2 {
		for elem := 2; elem < nb; elem++ {
			if nb%elem == 0 {
				nb = nb + 1
				FindNextPrime(nb)
			}
		}
		return nb
	}
	if nb == 2 {
		return nb
	}
	nb = nb + 1
	return FindNextPrime(nb)
}
