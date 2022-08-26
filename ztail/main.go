package main

import (
	"fmt"
	"os"
)

func main() {
	var need_exit bool = false
	c := BasicAtoi(os.Args[2])
	if len(os.Args) == 4 {
		file, err := os.ReadFile(os.Args[3])
		if err != nil {
			fmt.Println(err.Error())
			need_exit = true
		} else {
			if c > len(file) {
				fmt.Println(string(file))
			} else {
				fmt.Println(string(file[len(file)-c : len(file)-1]))
			}
		}
	} else {
		for k, i := range os.Args[3:] {

			file, err := os.ReadFile(i)
			if err != nil {
				fmt.Println(err.Error())
				need_exit = true
			} else {
				if k >= 1 {
					fmt.Println()
				}
				fmt.Printf("==> %v <==\n", string(i))
				if c > len(file) {
					fmt.Print(string(file))
				} else {
					fmt.Print(string(file[len(file)-c:]))
				}

			}
		}
	}

	if need_exit {
		os.Exit(1)
	}
}

func BasicAtoi(s string) int {
	str := []byte(s)
	var num int = 0
	for i := 0; i < len(str); i++ {
		num = num*10 + int(str[i]-48)
	}
	return num
}
