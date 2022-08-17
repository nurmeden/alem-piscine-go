package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	length := 0
	for l := range arguments {
		length = l + 1
	}

	if length == 0 {
		input, err := ioutil.ReadAll(os.Stdin)
		for _, j := range string(input) {
			z01.PrintRune(j)
		}
		if err != nil {
			b := "ERROR: " + err.Error()
			for _, e := range b {
				z01.PrintRune(e)
			}
			z01.PrintRune('\n')
			os.Exit(1)
		}
		return
	}

	for _, arg := range arguments {
		file, err := os.Open(arg)
		if err != nil {
			b := "ERROR: " + err.Error()
			for _, e := range b {
				z01.PrintRune(e)
			}
			z01.PrintRune('\n')
			os.Exit(1)
			return
		}

		f, err := ioutil.ReadAll(file)

		for _, text := range string(f) {
			z01.PrintRune(text)
		}
		if err != nil {
			b := "ERROR: " + err.Error()
			for _, e := range b {
				z01.PrintRune(e)
			}
			z01.PrintRune('\n')
			os.Exit(1)
		}

		file.Close()
	}
}
