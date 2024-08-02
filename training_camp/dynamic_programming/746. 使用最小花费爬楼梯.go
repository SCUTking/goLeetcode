package dynamic_programming

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}
func minCostClimbingStairs1(cost []int) int {
	n := len(cost)
	dp := make([]int, n)
	pre, cur := 0, 0
	for i := 2; i <= n; i++ {
		pre, cur = cur, min(cur+cost[i-1], pre+cost[i-2])
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
