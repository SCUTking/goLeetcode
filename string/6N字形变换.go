package string

func convert(s string, numRows int) (result string) {
	runes := []rune(s)
	var flag int = 1
	res := []string{}
	var i int = 0
	for _, v := range runes {
		res[i] += string(v)
		i += flag
		if i == numRows-1 {
			flag = -1
		}
	}

	for _, v := range res {
		result += v
	}
	return
}
