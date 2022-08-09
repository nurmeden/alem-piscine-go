package piscine

func FindNextPrime(nb int) int {
	if nb > 1 {
		if nb/10 >= 1 && nb%10 == 0 {
			return FindNextPrime(nb + 1)
		}
		if nb <= 9 {
			for i := 2; i <= 9; i++ {
				if nb%i == 0 && nb/2 >= i {
					return FindNextPrime(nb + 1)
				}
				if nb < i {
					return nb
				}
			}
		}
		for i := 2; true; i++ {
			if nb%i == 0 && nb/4 >= i {
				return FindNextPrime(nb + 1)
			} else if nb/10 >= 1 && nb%10 == 0 {
				return FindNextPrime(nb + 1)
			}
			if nb/9 < i {
				return nb
			}
		}
	}
	if nb == 3 {
		return nb
	}
	return 2
}
