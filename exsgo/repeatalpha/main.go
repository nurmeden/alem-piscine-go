package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for _, value := range args[1] {
		for x := 0; x <= len_word(value); x++ {
			z01.PrintRune(value)
		}
	}
	z01.PrintRune('\n')
}

func len_word(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r-'a') + 1
	} else if r >= 'A' && r <= 'Z' {
		return int(r-'A') + 1
	} else {
		return 1
	}
}
