package piscine

func UltimateDivMod(a *int, b *int) {
	*a = *a / *b
	number := *a / *b
	*b = number % *b
}
