package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]

	if checkInput(arguments) == true {
		error_msg := "Error"
		table := [9][9]rune{}
		table = fill_input_Table(table, arguments)

		if Settle(&table) == true {
			for y := 0; y < 9; y++ {
				for x := 0; x < 9; x++ {
					if x != 8 {
						z01.PrintRune(rune(table[y][x]))
						z01.PrintRune(32)
					} else {
						z01.PrintRune(rune(table[y][x]))
					}
				}
				z01.PrintRune(10)
			}
		} else {
			for _, value := range error_msg {
				z01.PrintRune(value)
			}
			z01.PrintRune('\n') // checkInput is not True
		}
	}
}

func checkInput(args []string) bool {
	error_message := "Error"
	if len(args) != 9 {
		for _, v := range error_message {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
		return false
	}
	for i := 0; i < len(args); i++ {
		if len(args[i]) != 9 {
			for _, v := range error_message {
				z01.PrintRune(v)
			} //  Input length is out of range
			z01.PrintRune('\n')
			return false
		}
	}
	for i := 0; i < len(args); i++ {
		for _, value := range args[i] {
			if value == 47 || value == 48 {
				for _, v := range error_message {
					z01.PrintRune(v)
				} // Input is not correct
				z01.PrintRune('\n')
				return false
			} else if value < 46 || value > 57 {
				for _, v := range error_message {
					z01.PrintRune(v)
				} // Input is not correct
				z01.PrintRune('\n')
				return false
			}
		}
	}
	return true
}

// fills with input arguments
func fill_input_Table(table [9][9]rune, args []string) [9][9]rune {
	for i := range args {
		for j := range args[i] {
			table[i][j] = rune(args[i][j])
		}
	}
	return table
}

// counts remaining empty cells
func counts_empty_cells(table *[9][9]rune) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j] == '.' {
				return true
			}
		}
	}
	return false
}

// check if it fits

func Suit(table *[9][9]rune, x int, y int, z rune) bool {
	// check double int
	for i := 0; i < 9; i++ {
		if z == table[i][x] {
			return false
		}
	}

	for j := 0; j < 9; j++ {
		if z == table[y][j] {
			return false
		}
	}

	// square check
	a := x / 3
	b := y / 3

	for k := 3 * a; k < 3*(a+1); k++ {
		for l := 3 * b; l < 3*(b+1); l++ {
			if z == table[l][k] {
				return false
			}
		}
	}
	return true
}

// backtracking solver

func Settle(table *[9][9]rune) bool {
	if !counts_empty_cells(table) {
		return true
	}
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if table[y][x] == '.' {
				for z := '1'; z <= '9'; z++ {
					if Suit(table, x, y, z) {
						table[y][x] = z
						if Settle(table) {
							return true
						}
					}
					table[y][x] = '.'
				}
				return false
			}
		}
	}
	return false
}

// checkInput is not True
