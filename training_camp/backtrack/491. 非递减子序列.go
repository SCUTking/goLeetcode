package backtrack

import "sort"

func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	tempRes := make([]int, 0)
	var dfs func(index int)
	dfs = func(index int) {
		if index == len(nums) {
			ints := make([]int, len(tempRes))
			copy(ints, tempRes)
			res = append(res, ints)
			return
		}
		if len(tempRes) > 0 && tempRes[len(tempRes)-1] > nums[index] {
			return
		}
		tempRes = append(tempRes, nums[index])
		dfs(index + 1)
		tempRes = tempRes[:len(tempRes)-1]
		dfs(index + 1)
	}

	sort.Ints(nums)
	dfs(0)
	return res
}
