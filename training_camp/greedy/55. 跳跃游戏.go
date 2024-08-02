package greedy

import "math"

func canJump(nums []int) bool {
	maxGet := make([]int, 0)
	maxG := math.MinInt32
	for i := 0; i < len(nums); i++ {
		if i > 0 && maxG < i {
			return false
		}
		maxGet[i] = nums[i] + i
		maxG = max(maxG, maxGet[i])
	}

	return true

}
