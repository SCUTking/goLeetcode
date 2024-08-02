package dynamic_programming

import "fmt"

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1) //dp[j]表示到达背包为j的最小的硬币数量
	dp[0] = 1
	for j := 0; j < target+1; j++ {
		for i := 0; i < len(nums); i++ {
			if j >= nums[i] {
				dp[j] = dp[j] + dp[j-nums[i]]
			}
		}
		fmt.Println(dp)
	}

	return dp[target]

}
