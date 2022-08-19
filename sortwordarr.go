package piscine

func SortWordArr(a []string) {
	for range a {
		for i := 1; i < len(a); i++ {
			if a[i] < a[i-1] {
				a[i], a[i-1] = a[i-1], a[i]
			}
		}
	}
}
