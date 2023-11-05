package sliding_window

func LongestSubstring(s string, k int) int {
	//m := make(map[uint8]int)
	//
	//var checkVaild func() bool
	//checkVaild = func() bool {
	//	for _, v := range m {
	//		if v < k {
	//			return false
	//		}
	//	}
	//	return true
	//}
	//
	//left, right := 0, 0
	//max := -1
	//for right < len(s) {
	//
	//	u := s[right]
	//	right++
	//	_, exist := m[u]
	//	if exist {
	//		m[u]++
	//	} else {
	//		m[u] = 1
	//	}
	//
	//	for !checkVaild() {
	//		len1 := right - left
	//		if len1 > max {
	//			max = len1
	//		}
	//		m[s[left]]--
	//		left++
	//
	//	}
	//
	//}
	//
	//return max

	max := 0

	for i := 1; i <= 26; i++ {
		cnt := [26]int{}
		total := 0
		lessk := 0

		l := 0
		for r, v := range s {
			ch := v - 'a'
			if cnt[ch] == 0 {
				total++
				lessk++
			}
			cnt[ch]++
			if cnt[ch] == k {
				lessk--
			}

			for total > i {
				ch1 := s[l] - 'a'

				if cnt[ch1] == k {
					lessk++
				}
				cnt[ch1]--
				if cnt[ch1] == 0 {
					total--
					lessk--
				}
				l++

			}

			if lessk == 0 {
				max = max1(max, r-l+1)
			}
		}

	}

	return max
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
