package dynamic_programming

func findTargetSumWays(nums []int, target int) int {
	////01背包问题
	////背包容量——未定
	////物品的价值可正可负
	//n := len(nums)
	//sum := 0
	//for _, num := range nums {
	//	sum += num
	//}
	//dp := make([][]bool, n) //背包容量为j，前i数，
	//for i := 0; i < n; i++ {
	//	dp[i] = make([]bool, sum+1)
	//}
	//res := 0
	//dp[0][nums[0]] = true
	//for i := 1; i < n; i++ {
	//	for j := 0; j < sum+1; j++ {
	//		if j < nums[i] {
	//			if j+nums[i] <= sum {
	//				dp[i][j] = dp[i-1][j+nums[i]]
	//			} else {
	//				dp[i][j] = false
	//			}
	//		} else if j+nums[i] > sum {
	//			if j >= nums[i] {
	//				dp[i][j] = dp[i-1][j-nums[i]]
	//			} else {
	//				dp[i][j] = false
	//			}
	//		} else {
	//			dp[i][j] = dp[i-1][j-nums[i]] || dp[i-1][j+nums[i]]
	//		}
	//
	//		if j == target {
	//			res++
	//		}
	//	}
	//}
	//
	return 1
}
func findTargetSumWays1(nums []int, target int) int {
	res := 0
	var dfs func(index int, sum int)
	dfs = func(index int, sum int) {
		if index == len(nums) {
			if sum == target {
				res++
			}
		}

		dfs(index+1, sum+nums[index])
		dfs(index+1, sum-nums[index])
	}
	dfs(0, 0)
	return res
}
