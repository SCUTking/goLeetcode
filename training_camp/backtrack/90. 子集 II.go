package backtrack

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{})
	tempRes := make([]int, 0)
	var dfs func(index int)
	dfs = func(index int) {
		if index == len(nums) {
			temp := make([]int, len(tempRes))
			copy(temp, tempRes)
			res = append(res, temp)
			return
		}
		index++
		for i := index; i < len(nums); i++ {
			if i > 0 && nums[i-1] == nums[i] {
				continue
			}
			tempRes = append(tempRes, nums[i])
			dfs(i)
			tempRes = tempRes[:len(tempRes)-1]
		}
	}

	dfs(0)
	return res
}
