package main

// package piscine

import (
	"fmt"
)

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}

func Atoi(s string) int {
	isFalse := false
	haveminus := 0
	total := 0
	sChangeable := []rune(s)
	for i := 0; i < len(sChangeable); i++ {
		total *= 10
		number := int(rune(sChangeable[i]) - 48)
		if number >= 0 && number < 10 {
			total += number
		} else if number == -3 {
			isFalse = true
			haveminus++
		} else if haveminus == 1 {
			isFalse = true
		} else if (number == -3 || number == -5) || haveminus > 1 {
			continue
		} else {
			return 0
		}
	}
	if isFalse == true {
		total *= -1
		return total
	} else {
		return total
	}
}
