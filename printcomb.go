package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	start := '0'
	end := '9'
	for digit1 := start; digit1 <= end; digit1++ {
		for digit2 := start; digit2 <= end; digit2++ {
			if digit1 < digit2 {
				for digit3 := start; digit3 <= end; digit3++ {
					if digit2 < digit3 {
						z01.PrintRune(digit1)
						z01.PrintRune(digit2)
						z01.PrintRune(digit3)
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}
