package greedy

func partitionLabels(s string) []int {
	m := make(map[int32]int, len(s))
	for index, i := range s {
		m[i] = index
	}
	maxRight := -1
	res := make([]int, 0)
	for index, i := range s {
		r := m[i]
		if maxRight == index {
			temp := maxRight + 1
			if len(res) > 0 {
				sum := 0
				for _, re := range res {
					sum += re
				}
				temp = maxRight - sum + 1
			}
			res = append(res, temp)
		}
		if r > maxRight {
			maxRight = r
		}

	}
	return res
}
