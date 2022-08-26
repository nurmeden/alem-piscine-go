package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	var minus bool = false
	var max_num bool = false
	if !IsUniqueForNbrbase(base) || len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
	} else {
		if nbr < 0 {
			if nbr == -9223372036854775808 {
				nbr = -9223372036854775807
				max_num = true
			}
			nbr = nbr * -1
			minus = true
		}
		res := BaseForNbrbase(nbr, len(base))
		if minus {
			z01.PrintRune('-')
		}
		for i := len(res) - 1; i >= 0; i-- {
			if max_num && i == 0 {
				z01.PrintRune(rune(base[res[i]+1]))
			} else {
				z01.PrintRune(rune(base[res[i]]))
			}
		}
	}
}

func IsUniqueForNbrbase(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == '+' || s[i] == '-' {
			return false
		}
		for j := len(s) - 1; j > i; j-- {
			if s[j] == s[i] {
				return false
			}
		}
	}
	return true
}

func BaseForNbrbase(num int, base int) []int {
	var temp int
	res := []int{}
	for num > 0 {
		temp = num % base
		num = num / base
		res = append(res, temp)
	}
	return res
}
