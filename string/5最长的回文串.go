package string

func LongestPalindrome(s string) string {
	length := len(s)
	var temp [][]int
	//初始化
	for i := 0; i < length; i++ {
		ints := make([]int, length)
		temp = append(temp, ints)
	}
	res := ""
	for i := 0; i < length; i++ {
		ints := temp[i]
		ints[i] = 1
		temp[i] = ints
		//res = string(r[i])
	}

	//length := len(s)
	r := []rune(s)
	for i := 1; i <= length; i++ {
		for j := 0; j <= length-i; j++ {
			if i == 1 {
				ints := temp[j]
				ints[j] = 1
				temp[j] = ints
				res = string(r[i])
				continue
			}
			r3 := r[j]
			r4 := r[j+i-1]
			if i == 2 {
				if r3 == r4 {
					ints := temp[j]
					ints[j+1] = 1
					temp[j] = ints
					res = string(r[i : i+2])
				}
				continue
			}
			ints := temp[j+1]
			i2 := ints[j+i-2]
			if i2 == 1 && r3 == r4 {
				ints1 := temp[j]
				ints1[j+i-1] = 1
				temp[j] = ints1
				res = string(r[j : j+i])
			}

		}
	}

	return res
}
func reverseString(s string) string {
	runes := []rune(s)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}

	return string(runes)
}
