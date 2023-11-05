package random

func rob(nums []int) int {
	//n := len(nums)
	//jiSum, ouSum := 0, 0
	//for i := 0; i < n; i += 2 {
	//	jiSum += nums[i]
	//}
	//for i := 1; i < n; i += 2 {
	//	9
	//	ouSum += nums[i]
	//}
	//if n&1 == 0 {
	//	return max5(jiSum, ouSum)
	//
	//} else {
	//	outFjiSum := jiSum - nums[0]
	//	outLjiSum := jiSum - nums[len(nums)]
	//	jiSum = max5(outLjiSum, outFjiSum)
	//	return max5(jiSum, ouSum)
	//}
	n := len(nums)

	fSlice := nums[1:]
	sSlice := nums[:n-1]
	i := rob2(fSlice)
	i1 := rob2(sSlice)
	return max(i, i1)

}

func rob2(nums []int) int {
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
		return dp[n]
	}
}
