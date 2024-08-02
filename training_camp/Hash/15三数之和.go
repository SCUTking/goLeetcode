package Hash

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	if n < 3 {
		return [][]int{}
	}

	res := make([][]int, 0)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for right-1 >= 0 && nums[right-1] == nums[right] {
					right = right - 1
				}
				right = right - 1
				for left+1 <= n-1 && nums[left+1] == nums[left] {
					left = left + 1
				}
				left = left + 1
			} else if sum > 0 {
				for right-1 >= 0 && nums[right-1] == nums[right] {
					right = right - 1
				}
				right = right - 1
			} else {
				for left+1 <= n-1 && nums[left+1] == nums[left] {
					left = left + 1
				}
				left = left + 1
			}
		}
	}

	return res
}
