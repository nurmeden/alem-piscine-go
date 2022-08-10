package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args

	for _, v := range arguments[1:] {
		for _, letter := range v {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}

	/*
		msstring := ""
		arguments := os.Args
		for i := 0; i < len(arguments); i++ {
			msstring = arguments[i]
			rune_m := []rune(msstring)
			z01.PrintRune(rune(rune_m))
		}
		z01.PrintRune('\n')*/
}
