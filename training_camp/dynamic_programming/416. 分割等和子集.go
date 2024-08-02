package dynamic_programming

// 暴力超时
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num

	}
	if sum&1 != 0 {
		return false
	}

	var dfs func(index int, tempSum int) bool
	dfs = func(index int, tempSum int) bool {
		if tempSum == sum/2 {
			return true
		}
		if index == len(nums) {
			return false
		}
		return dfs(index+1, tempSum) || dfs(index+1, tempSum+nums[index])
	}

	return dfs(0, 0)
}

func canPartition1(nums []int) bool {
	//变种的背包问题
	//就是问背包的二维dp数组有无等于sum/2的
	//价值是nums数组
	//重量是1
	//背包的总数是nums的长度
	sum := 0
	for _, num := range nums {
		sum += num

	}
	if sum&1 != 0 {
		return false
	}

	n := len(nums)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, sum/2)
	}

	//dp[i][j]=dp[i-1][j-nums[i]]||dp[i-1][j]
	for j := 0; j < n; j++ {
		dp[j][0] = true
	}

	for i := 0; i < n; i++ {
		for j := 1; j < sum/2; j++ {
			if j < nums[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j-nums[i]] || dp[i-1][j]
			}
		}
	}

	return dp[n-1][sum/2]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
