package Hot100

// 暴力没做出来
func subarraySum(nums []int, k int) int {
	n := len(nums)
	count := 0
	for i := 1; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			sum := 0
			for k := j; k < j+i; k++ {
				sum += nums[k]
			}
			if sum == k {
				count++
			}
		}

	}
	return count

}

func subarraySum1(nums []int, k int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if preSum[j] == preSum[i]-k {
				count++
			}
		}

	}

	return count

}
