package backtrack

import "sort"

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	var dfs func(index int)
	tempRes := make([]int, 0)
	m := make([]bool, len(nums))
	for i, _ := range nums {
		m[i] = true
	}
	dfs = func(index int) {
		if index == len(nums) {
			ints := make([]int, len(tempRes))
			copy(ints, tempRes)
			res = append(res, ints)
			return
		}
		for i := 0; i < len(m); {
			if m[i] {
				tempRes = append(tempRes, nums[i])
				m[i] = false
				dfs(index + 1)
				tempRes = tempRes[:len(tempRes)-1]
				m[i] = true
				for i+1 < len(m) && nums[i+1] == nums[i] {
					i++
				}
			}

			i++
		}
	}
	sort.Ints(nums)
	dfs(0)
	return res
}
