package greedy

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	res := 1
	var dfs func(start int)
	dfs = func(start int) {
		if start >= len(nums)-1 || start+nums[start] >= len(nums)-1 {
			return
		}
		maxIndex := start + 1
		for i := start + 1; i <= start+nums[start]; i++ {
			if nums[i]+i >= nums[maxIndex]+maxIndex {
				maxIndex = i
			}
		}
		if maxIndex+nums[maxIndex] > start+nums[start] {
			res++

			dfs(maxIndex)
		} else {
			res++
			dfs(start + nums[start])
		}
	}

	dfs(0)
	return res
}
func jump1(nums []int) int {
	n := len(nums)
	end := 0
	maxP := 0
	steps := 0
	for i := 0; i < n-1; i++ {
		maxP = max(maxP, nums[i]+i)
		if i == end {
			end = maxP
			steps++
		}
	}
	return steps
}
