package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args[1:]
	if len(os.Args[1:]) == 0 {
		fmt.Println("File name missing")
	} else if len(os.Args[1:]) > 1 {
		fmt.Println("Too many arguments")
	} else {
		data_in_file, err := ioutil.ReadFile(arguments[0])
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Print(string(data_in_file))
	}
}
