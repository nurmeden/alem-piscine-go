package piscine

func ConcatParams(args []string) string {
	str_str := ""
	for elem := 0; elem <= len(args)-1; elem++ {
		for _, value := range args[elem] {
			str_str += string(value)
		}
		if elem != len(args)-1 {
			str_str += string("\n")
		}
	}
	return str_str
}
