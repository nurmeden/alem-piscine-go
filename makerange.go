package piscine

func MakeRange(min, max int) []int {
	length := max - min
	if max > min {
		answer := make([]int, length)
		for i := 0; i <= length-1; i++ {
			answer[i] = i + min
		}
		return answer
	} else {
		return nil
	}
}
