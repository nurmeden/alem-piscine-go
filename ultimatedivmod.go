package piscine

func UltimateDivMod(a *int, b *int) {
	number := *a
	*a = *a / *b
	*b = number % *b
}
