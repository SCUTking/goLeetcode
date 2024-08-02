package arr

// 算法要点：利用前缀和减少不必要的加法运算
// 时间复杂度：O（n）
// 空间复杂度：O（n）
// 未能考虑到的点：
func sumInArea(nums []int, s, e int) int {
	preSum := make([]int, len(nums)+1)
	preSum[0] = 0
	for i := 1; i <= len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	return preSum[e+1] - preSum[s]

}
