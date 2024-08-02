package greedy

import "math"

func largestSumAfterKNegations(nums []int, k int) int {
	for k > 0 {
		minOne := math.MaxInt32
		for i, num := range nums {
			if num < minOne {
				minOne = i
			}
		}
		nums[minOne] = -nums[minOne]
		k--
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
