package main

import (
	"fmt"
	"os"
)

func main() {
	search_01 := "01"
	search_galaxy := "galaxy"
	search_galaxy_01 := "galaxy 01"
	arguments := os.Args
	for _, value := range arguments {
		if search_01 == string(value) || search_galaxy == string(value) || search_galaxy_01 == string(value) {
			fmt.Println("Alert!!!")
		}
	}
}
