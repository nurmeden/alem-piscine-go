// package main

// import "github.com/01-edu/z01"

// func main() {
// 	start := '0'
// 	end := '9'
// 	next := false
// 	for digit1 := start; digit1 <= end; digit1++ {
// 		for digit2 := start; digit2 <= end; digit2++ {
// 			for digit3 := start; digit3 <= end; digit3++ {
// 				for digit4 := start; digit4 <= end; digit3++ {
// 					if next {
// 						z01.PrintRune(',')
// 						z01.PrintRune(' ')
// 					}

// 					next = true
// 					z01.PrintRune(digit1)
// 					z01.PrintRune(digit2)
// 					z01.PrintRune(' ')
// 					z01.PrintRune(digit3)
// 					z01.PrintRune(digit4)
// 				}
// 			}
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			f := j + 1
			for k := i; k <= '9'; k++ {
				for ; f <= '9'; f++ {
					z01.PrintRune((i))
					z01.PrintRune((j))
					z01.PrintRune(' ')
					z01.PrintRune((k))
					z01.PrintRune((f))
					if i < '9' || j < '8' || k < '9' || f < '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

				}
				f = '0'
			}
		}
	}
	z01.PrintRune('\n')
}
