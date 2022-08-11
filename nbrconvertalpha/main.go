package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for i := 1; i <= len(arguments)-1; i++ {
		for _, value := range arguments[i] {
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
	}
}
