package piscine

import (
	"fmt"
)

func ForEach(f func(int), a []int) {
	for _, value := range a {
		f(value)
	}
	fmt.Printf("\n")
}
