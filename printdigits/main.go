package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Print(i)
	}
	z01.PrintRune('\n')
}
