package Hot100

import "math"

// 递归复杂度太高了
func lengthOfLIS1(nums []int) int {
	var dfs func(i int, pre int, l int)
	res := math.MinInt32
	dfs = func(i int, pre int, l int) {
		if i == len(nums) {
			return
		}

		//选
		if nums[i] > pre {
			l++
			if l > res {
				res = l
			}
			dfs(i+1, nums[i], l)
		}
		//不选
		dfs(i+1, pre, l)
	}

	dfs(0, math.MinInt32, 0)
	return res
}
func lengthOfLIS(nums []int) int {
	dp := make([]int, 0) //前i哥元素，以i结尾的最长的子序列的长度的
	res := -1
	for i := 0; i < len(nums); i++ {
		maxPre := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if maxPre < dp[j] {
					maxPre = dp[j]
				}
			}
		}

		dp[i] = maxPre + 1
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}
