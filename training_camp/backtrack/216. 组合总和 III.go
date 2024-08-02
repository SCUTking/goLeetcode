package backtrack

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	var helper func(s []int, i int)
	helper = func(s []int, i int) {
		if len(s) == k {
			sum := 0
			for j := 0; j < len(s); j++ {
				sum += s[j]
			}
			if sum == n {
				temp := make([]int, k)
				copy(temp, s)
				res = append(res, temp)
				return
			}
		}
		if i == 9 {
			return
		}
		i++
		helper(s, i)
		s = append(s, i)
		helper(s, i)

	}

	helper([]int{}, 0)
	return res
}
