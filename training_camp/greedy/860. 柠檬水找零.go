package greedy

func lemonadeChange(bills []int) bool {
	m := make(map[int]int, 0)
	for _, bill := range bills {
		if m[5] < 0 {
			return false
		}

		switch bill {
		case 5:
			m[5]++
		case 10:
			m[5]--
			m[10]++
		case 20:
			if m[10] > 0 {
				if m[5] > 0 {
					m[10]--
					m[5]--
				} else {
					return false
				}
			} else {
				m[5] = m[5] - 3
			}
		}
	}
	return true
}
