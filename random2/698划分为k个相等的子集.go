package random2

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	eachSum := sum / k
	n := len(nums)
	dp := 1 << n

	sort.Ints(nums)
	var dfs func(state int, cur int) bool
	dfs = func(state int, cur int) bool {
		if state == dp-1 {
			return true
		}
		if dp>>state&1 == 1 {
			return false
		}
		dp |= 1 << state
		//试用每一个，不行就换另一个
		for i := 0; i < len(nums); i++ {
			if nums[i]+cur > eachSum {
				return false
			}
			if state>>i&1 == 1 {
				continue
			}
			b := dfs(state|(1<<i), (cur+nums[i])%eachSum)
			if b {
				return true
			}
		}
		return false
	}

	dfs(0, 0)
	return true
}
