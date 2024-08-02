package dynamic_programming

func rob(nums []int) int {
	n := len(nums)
	dp1 := make([]int, n)
	dp2 := make([]int, n)

	//不要第一个
	dp1[0] = 0
	for i := 1; i < n; i++ {
		if i == 1 {
			dp1[i] = nums[i]
		} else {
			dp1[i] = max(dp1[i-2]+nums[i], dp1[i-1])
		}
	}

	//要第一个
	dp2[0] = nums[0]
	for i := 1; i < n; i++ {
		if i == 1 {
			dp2[i] = nums[0]
		} else if i == n-1 {
			dp2[i] = dp2[i-1]
		} else {
			dp2[i] = max(dp2[i-2]+nums[i], dp2[i-1])
		}
	}

	return max(dp2[n-1], dp1[n-1])
}

func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res0 := rob1(nums[1:])
	res1 := rob1(nums[:len(nums)-1])
	return max(res0, res1)
}

// 原始的打家劫舍
func rob1(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			dp[0] = nums[0]
		} else if i == 1 {
			dp[1] = max(dp[0], nums[1])
		} else {
			dp[i] = max(dp[i-2]+nums[i], dp[i-1])
		}
	}

	return dp[n-1]
}
