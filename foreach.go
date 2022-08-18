package piscine

import (
	"github.com/01-edu/z01"
)

func ForEach(f func(int), a []int) {
	for _, value := range a {
		f(value)
	}
	z01.PrintRune('\n')
}
