package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	ar := os.Args
	if len(ar) == 2 {
		for _, x := range addPrime(ar[1]) {
			z01.PrintRune(x)
		}
	} else {
		z01.PrintRune('0')
	}
	z01.PrintRune('\n')
}

func addPrime(num string) string {
	n, _ := strconv.Atoi(num)
	sum := 0
	for x := 2; x <= n; x++ {
		if isPrime(x) {
			sum += x
		}
	}
	return strconv.Itoa(sum)
}

func isPrime(n int) bool {
	z := 0
	for x := 2; x <= n; x++ {
		if n%x == 0 {
			z++
		}
	}
	if z == 1 {
		return true
	}
	return false
}
