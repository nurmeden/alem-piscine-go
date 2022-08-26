// package main

// import (
// 	"fmt"
// )

// func main() {
// 	s := "HelloHAhowHAareHAyou?"
// 	fmt.Printf("%#v\n", Split(s, "HA"))
// }

// func Split(str, charset string) []string {
// 	st, size := div(str, charset, 1) // HelloHAhowHAareHAyou? HA 1
// 	word := ""                       // zhanadan string ashyp alamyz
// 	res := make([]string, size)      //
// 	n := 0
// 	for _, x := range st {
// 		if x != '\v' {
// 			word += string(x)
// 		} else {
// 			res[n] = word
// 			n++
// 			word = ""
// 		}
// 	}
// 	res[n] = word
// 	return res
// }

// func div(s, c string, n int) (string, int) {
// 	cl := len(c) // 2
// 	sl := len(s) // 21
// 	if sl < cl { // esli splittin uzyndigi stringten ulken bolsa
// 		return s, n
// 	}

// 	if s[0:cl] == c { // s[0:2] == HA
// 		return div("\v"+s[cl:], c, n+1)
// 	}
// 	st, k := div(s[1:], c, n)
// 	return s[0:1] + st, k
// }

// 2-variant
package split

import "fmt"

func Split(s, sep string) []string {
	res := []string{}
	b := []rune(s)
	fmt.Println(s[0:2])
	for len(b) > 0 {
		index := Index(string(b), sep)
		if index > -1 { // 5 > -1
			value := string(b[:index]) //
			fmt.Println(value)
			if value != "" {
				res = append(res, string(b[:index]))
			}

			b = b[index+len(sep):]
		} else {
			res = append(res, string(b[0:]))
			break
		}
	}

	return res
}

func Index(s string, toFind string) int {
	res := -1

	if len(toFind) <= 0 {
		return 0
	}

	for index := 0; index <= len(s)-len(toFind); index++ { // 19
		if s[index:len(toFind)+index] == toFind { // s[0:2],s[1:3] He el ll lo "HA"(s[5:7])
			res = index // res == 5
			break
		}
	}
	return res // 5
}
