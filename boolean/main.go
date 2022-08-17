package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	result := true
	if nbr%2 != 0 {
		return result
	} else {
		result = false
		return result
	}
}

func main() {
	arguments := os.Args[1:]
	OddMsg := "I have an odd number of arguments"
	EvenMsg := "I have an even number of arguments"
	lengthOfArg := len(arguments)
	if isEven(lengthOfArg) == true {
		printStr(OddMsg)
	} else {
		printStr(EvenMsg)
	}
}
