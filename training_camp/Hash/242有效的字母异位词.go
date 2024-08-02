package Hash

func isAnagram(s string, t string) bool {
	m := make(map[int32]int)
	for _, i := range s {
		i2 := i - 'a'
		m[i2]++
	}

	for _, i := range t {
		i2 := i - 'a'
		m[i2]--
		if m[i2] == 0 {
			delete(m, i2)
		}

	}
	if len(m) > 0 {
		return false
	} else {
		return true
	}

}
