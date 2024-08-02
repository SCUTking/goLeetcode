package backtrack

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	var dfs func(index int)
	tempRes := make([]int, 0)
	m := make([]bool, 0)
	for _, num := range nums {
		m[num] = true
	}
	dfs = func(index int) {
		if index == len(nums) {
			ints := make([]int, len(tempRes))
			copy(ints, tempRes)
			res = append(res, ints)
			return
		}
		for k, _ := range m {
			if m[k] {
				tempRes = append(tempRes, k)
				m[k] = false
				dfs(index + 1)
				tempRes = tempRes[:len(tempRes)-1]
				m[k] = true
			}
		}
	}
	dfs(0)
	return res
}
