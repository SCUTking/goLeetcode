package backtrack

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	var helper func(s []int, i int)
	helper = func(s []int, i int) {
		if len(s) == k {
			temp := make([]int, k)
			copy(temp, s)
			res = append(res, temp)
			return
		}
		if i == n {
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
