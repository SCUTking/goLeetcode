package random2

func canConstruct(s string, k int) bool {
	even, odd := 0, 0
	m := make(map[int32]int)
	for _, temp := range s {
		m[temp]++
	}
	for _, v := range m {
		if v&1 == 0 {
			even++
		} else {
			odd++
		}
	}

	if odd > k || len(s) < k {
		return false
	}
	return true

}
