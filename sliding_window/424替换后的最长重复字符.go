package sliding_window

func CharacterReplacement(s string, k int) int {
	if len(s) < 2 {
		return len(s)
	}

	left, right := 0, 0
	ans := 0
	maxCount := 0
	fnt := [26]int{}
	for right < len(s) {
		u := s[right] - 'A'
		fnt[u]++
		right++
		if fnt[u] > maxCount {
			maxCount = fnt[u]
		}

		for right-left > maxCount+k {
			fnt[s[left]-'A']--
			left++
			max := 0
			for _, e := range fnt {
				if e > max {
					max = e
				}
			}
			maxCount = max
		}

		if right-left+1 > ans {
			ans = right - left
		}

	}
	return ans
}
