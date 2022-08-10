package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	list_int := []int{}
	if n == 0 {
		list_int = append(list_int, 0)
	}
	for n > 0 {
		dig := n % 10
		list_int = append(list_int, dig)
		n = n / 10
	}
	for i := 0; i < len(list_int); i++ {
		for j := 0; j < len(list_int); j++ {
			if list_int[i] < list_int[j] {
				list_int[i], list_int[j] = list_int[j], list_int[i]
			}
		}
	}
	for i := 0; i < len(list_int); i++ {
		z01.PrintRune(rune(list_int[i] + 48))
	}
}
