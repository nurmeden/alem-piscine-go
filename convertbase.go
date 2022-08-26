package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	res1 := AtoiBase(nbr, baseFrom)
	lens := len(baseTo)
	runes := []rune(baseTo)
	str := ""
	for res1 != 0 {
		temp := res1 % lens
		str += string(runes[temp])
		res1 /= lens
	}
	res := ""
	for i := len(str) - 1; i >= 0; i-- {
		res += string(str[i])
	}
	return res
}
