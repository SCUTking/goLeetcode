package backtrack

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	var dfs func(pre []int, sum int, index int)
	dfs = func(pre []int, sum int, index int) {
		if sum > target || index == len(candidates)-1 {
			return
		}
		if sum == target {
			temp := make([]int, len(pre))
			copy(temp, pre)
			res = append(res, temp)
		}
		index++
		for i := index; i < len(candidates); i++ {
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			pre = append(pre, candidates[i])
			dfs(pre, sum+candidates[i], i)
			pre = pre[:len(pre)-1]
		}
	}
	sort.Ints(candidates)
	dfs([]int{}, 0, -1)
	return res
}
