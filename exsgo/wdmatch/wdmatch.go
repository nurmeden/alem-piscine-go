package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	var res_str string
	var result string
	if len(arguments) == 3 {
		for _, value := range arguments[2] {
			for _, value1 := range arguments[1] {
				if value1 == value {
					res_str += string(value1)
				}
			}
		}
		for i := 0; i < len(res_str)-1; i++ {
			for j := 0; j < len(res_str)-1; j++ {
				if res_str[j] == res_str[j+1] {
					result += string(res_str[j])
				}
				result += string(res_str[j])
			}
		}
	}
	fmt.Println(result)
}
