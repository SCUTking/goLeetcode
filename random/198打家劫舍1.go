package random

func rob1(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	if n == 1 {
		return nums[0]
	} else if n == 2 {
		return max(nums[0], nums[1])
	} else {
		dp[0] = nums[0]
		dp[1] = max(nums[0], nums[1])
		for i := 2; i < n; i++ {
			cur := dp[i-2] + nums[i]
			dp[i] = max(cur, dp[i-1])
		}
		return dp[n-1]
	}
}
