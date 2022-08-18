package piscine

// func f(a, b int) int {
// 	if a > b {
// 		return 1
// 	} else if a < b {
// 		return -1
// 	} else {
// 		return 0
// 	}
// }

func IsSorted(f func(a, b int) int, a []int) bool {
	f1 := true
	f2 := true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			f1 = false
		}
		if f(a[i], a[i+1]) < 0 {
			f2 = false
		}
	}
	return f1 || f2
}
