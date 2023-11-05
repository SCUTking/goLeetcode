package monotone_stack

import "math"

func find132pattern(nums []int) bool {

	n := len(nums)
	candidate := []int{nums[n-1]}

	mask := math.MinInt32

	for i := n - 2; i >= 0; i-- {
		//是否可以作为1（最小的哪个）
		if nums[i] < mask {
			return true
		}

		for len(candidate) > 0 && nums[i] > candidate[len(candidate)-1] {
			mask = candidate[len(candidate)-1]
			candidate = candidate[:len(candidate)-1]
		}

		if nums[i] > mask {
			candidate = append(candidate, nums[i])
		}
	}

	return false

}
