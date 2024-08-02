package backtrack

import "sort"

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	tempRes := make([]int, 0)
	var dfs func(index int)
	dfs = func(index int) {
		if index == len(nums) {
			temp := make([]int, len(tempRes))
			copy(temp, tempRes)
			res = append(res, temp)
			return
		}

		tempRes = append(tempRes, nums[index])

		dfs(index + 1)
		tempRes = tempRes[:len(tempRes)-1]

		for index+1 < len(nums) && nums[index+1] == nums[index] {
			index++
		}
		dfs(index + 1)
	}
	sort.Ints(nums)
	dfs(0)
	return res
}
