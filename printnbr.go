package piscine

import (
	"github.com/01-edu/z01"
)

func check(r int) {
	count := '0'
	if r == 0 {
		z01.PrintRune(count)
		return
	}
	for i := 1; i <= r%10; i++ {
		count++
	}
	for i := -1; i >= r%10; i-- {
		count++
	}
	if r/10 != 0 {
		check(r / 10)
	}
	z01.PrintRune(count)
	return
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	check(n)
}
