package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	fmt.Printf(arguments[0][2:])
}
