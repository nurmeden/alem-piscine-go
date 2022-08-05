package piscine

import "fmt"

func PrintStr(s string) {
	c := 0
	for i := 0; i < len(s); i++ {
		c++
	}
	fmt.Println(c)
}
