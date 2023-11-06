package random1

func partitionLabels(s string) []int {
	m := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		u := s[i]
		m[u]++
	}

	curChars := make(map[uint8]struct{}, 0)
	num := 0
	ans := make([]int, 0)
	for i := 0; i < len(s); i++ {
		u := s[i]
		curChars[u] = struct{}{}
		m[u]--
		num++
		if m[u] == 0 {
			flag := true
			for k, _ := range curChars {
				if m[k] != 0 {
					flag = false
				}
			}
			if flag {
				ans = append(ans, num)
				num = 0
				//清空map
				for u2, _ := range curChars {
					delete(curChars, u2)
				}
			}
		}
	}
	return ans
}
