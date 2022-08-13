package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(args []string) {
	str_str := ""
	for elem := 0; elem <= len(args)-1; elem++ {
		for _, value := range args[elem] {
			str_str += string(value)
		}
		if elem != len(args)-1 {
			str_str += string("\n")
		}
	}

	for _, elem := range str_str {
		z01.PrintRune(elem)
	}
	z01.PrintRune('\n')
}
