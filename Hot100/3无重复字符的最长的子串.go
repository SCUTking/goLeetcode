package Hot100

func lengthOfLongestSubstring(s string) int {
	m := make(map[uint8]int, 0)
	left, right := 0, 0
	res := 0
	for right < len(s) {
		u := s[right] - 'a'
		index, exist := m[u]
		if exist {
			left = Max(index+1, left)
		}
		//不管咋样都要加
		m[s[right]-'a'] = right
		length := right - left + 1
		res = Max(res, length)

		right++
	}

	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
