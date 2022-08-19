package piscine

func Atoi(s string) int {
	y := []rune(s)
	var n int
	count := len(y)
	for b := range y {
		if b == 0 && y[0] == '+' || b == 0 && y[0] == '-' || b == 0 && y[b] == '0' {
			continue
		}
		if y[b] == '+' || y[b] == '-' || y[b] == ' ' || y[b] > 57 || y[b] < 48 {
			n = 0
			return n
		}
		n = n*10 + int(y[b]-48)
		if b == count-1 && y[0] == '-' {
			n = n * -1
		}
	}
	return n
}
