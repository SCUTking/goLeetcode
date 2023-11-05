package sliding_window

func FindAnagrams(s string, p string) []int {
	sMap := make(map[int32]int)
	pMap := make(map[int32]int)

	for _, i := range p {
		ch := i - 'a'
		pMap[ch]++
	}
	res := []int{}
	vaildNum := 0

	left, right := 0, 0
	for right < len(s) {
		ch := s[right] - 'a'
		sMap[(int32(ch))]++
		right++
		if sMap[(int32(ch))] == pMap[(int32(ch))] {
			vaildNum++
		}

		for right-left >= len(p) {
			if vaildNum == len(pMap) {
				res = append(res, left)
			}

			u := s[left] - 'a'
			if sMap[(int32(u))] == pMap[(int32(u))] {
				vaildNum--
			}
			sMap[(int32(u))]--
			left++

		}

	}

	return res
}
