package piscine

// a = 0
// b = 1
// result a = 1 b = 0

func Swap(a *int, b *int) {
	var number int = *a
	*a = *b
	*b = number
}
