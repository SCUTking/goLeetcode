package Hash

import "sort"

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	if n < 4 {
		return [][]int{}
	}

	res := make([][]int, 0)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			left, right := j+1, n-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for right-1 >= 0 && nums[right-1] == nums[right] {
						right = right - 1
					}
					right = right - 1
					for left < right && nums[left+1] == nums[left] {
						left = left + 1
					}
					left = left + 1
				} else if sum > target {
					right = right - 1
				} else {
					left = left + 1
				}
			}
		}

	}

	return res
}
