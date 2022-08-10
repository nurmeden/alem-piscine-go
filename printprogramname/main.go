package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for _, value := range arguments[0][2:] {
		z01.PrintRune(rune(value))
	}
	z01.PrintRune('\n')
}
