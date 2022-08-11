package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	car := false
	for _, b := range arguments[1:] {
		for _, v := range b {
			if v != 0 {
				car = true
			}
		}
	}
	if car {
		for _, b := range arguments[1:] {
			numv := 0
			for _, v := range b {
				numv = numv*10 + int(v-'0')
			}
			if numv >= 1 && numv <= 26 {
				if arguments[1] == "--upper" {
					z01.PrintRune(rune(numv + 64))
				} else {
					z01.PrintRune(rune(numv + 96))
				}
			} else if arguments[1] == "--upper" {
				continue
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
