package piscine

func FindNextPrime(nb int) int {
	if -9223372036854775807 < nb && nb < 9223372036854775807 {
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
	return nb
}
