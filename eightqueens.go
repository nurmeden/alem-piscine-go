package piscine

import "github.com/01-edu/z01"

const N = 8

var position = [N]int{}

func isSafe(queen_number, row_position int) bool {
	for i := 0; i < queen_number; i++ {
		other_row_pos := position[i]

		if other_row_pos == row_position || other_row_pos == row_position-(queen_number-i) || other_row_pos == row_position+(queen_number-i) {
			return false
		}
	}
	return true
}

func solve(k int) {
	if k == N {
		for i := 0; i < N; i++ {
			gg := position[i] + 1
			for gg > 0 {
				t := gg % 10
				z01.PrintRune(rune(int(rune(t)) + 48))
				gg = gg / 10
			}
		}
	} else {
		for i := 0; i < N; i++ {
			if isSafe(k, i) {
				position[k] = i
				solve(k + 1)
			}
		}
	}
}

func EightQueens() {
	solve(0)
}
