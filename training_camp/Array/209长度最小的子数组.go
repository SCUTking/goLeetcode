package Array

import "math"

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	n := len(nums)
	sum := 0
	res := math.MaxInt32
	for right < n {
		sum += nums[right]
		right++

		for sum >= target && left < right {
			res = min(res, right-left-1)
			left++
			sum -= nums[left]
		}
	}

	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
