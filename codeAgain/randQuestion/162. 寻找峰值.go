package randQuestion

import "math"

func findPeakElement(nums []int) int {
	n := len(nums)
	nums = append([]int{math.MinInt64}, nums...)
	nums = append(nums, math.MinInt64)

	for i := 1; i <= n; i++ {
		left, mid, right := nums[i-1], nums[i], nums[i+1]
		if mid > left && mid > right {
			return i
		}
	}
	return 0
}

// 最终版本：
// 人往高处爬，从一个位置不断地往高处走，一定能到达峰值（不一定是顶峰）
// 不可能走下坡洛，所以每次往一边高处走后，另一边的全部都被废除了（类似于二分法）
func findPeakElement1(nums []int) int {
	n := len(nums)
	var get func(index int) int64
	get = func(index int) int64 {
		if index == -1 || index == n {
			return math.MinInt64
		} else {
			return int64(nums[index])
		}

	}

	var helper func(left, right int) int
	helper = func(left, right int) int {
		mid := left + (right-left)/2
		if get(mid) > get(mid-1) && get(mid) > get(mid+1) {
			return mid
		} else if get(mid) < get(mid+1) {
			return helper(mid+1, right)
		} else {
			return helper(left, mid-1)
		}
	}

	return helper(-1, n)
}
