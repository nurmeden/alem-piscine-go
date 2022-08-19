package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	result := false
	for _, value := range arguments[1] {
		int_value := int(value - '0')
		for i := 1; i <= int_value; i++ {
			count := 1
			for in := 1; in <= i; in++ {
				count *= 2
			}
			if count == int_value {
				result = true
			}
		}
	}
	fmt.Println(result)
}
