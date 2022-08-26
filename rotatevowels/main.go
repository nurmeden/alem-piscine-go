package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	result := []rune{}
	positions := []int{}
	vowels := []rune{}
	var what_pos int = 0

	for i := 1; i < len(os.Args); i++ {
		arg := []rune(os.Args[i])
		for j := 0; j < len(arg); j++ {

			if arg[j] == 'a' || arg[j] == 'e' || arg[j] == 'o' || arg[j] == 'u' || arg[j] == 'i' || arg[j] == 'A' || arg[j] == 'E' || arg[j] == 'O' || arg[j] == 'U' || arg[j] == 'I' {
				positions = append(positions, what_pos)
				vowels = append(vowels, arg[j])
			}
			result = append(result, arg[j])
			what_pos++
		}
		result = append(result, ' ')
		what_pos++
	}
	MirrorForVowels(result, positions, vowels)

	for i := 0; i < len(result); i++ {
		z01.PrintRune(result[i])
	}
	z01.PrintRune('\n')
}

func MirrorForVowels(res []rune, pos []int, letter []rune) {
	for i := 0; i < len(pos)/2; i++ {
		res[pos[i]] = letter[len(letter)-1-i]
		res[pos[len(pos)-1-i]] = letter[i]
	}
}
