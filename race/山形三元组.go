package race

import "math"

func MinimumSum(nums []int) int {
	left := []int{}
	right := []int{}
	for i := 0; i < len(nums); i++ {

		cur := nums[i]
		minL := math.MaxInt32
		mI := -1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < minL && nums[j] < cur {
				minL = nums[j]
				mI = j
			}
		}
		if minL == math.MaxInt32 {
			left = append(left, -1)
		} else {
			left = append(left, mI)
		}

		minR := math.MaxInt32
		mRI := -1
		for j := i; j < len(nums); j++ {
			if nums[j] < minR && nums[j] < cur {
				minR = nums[j]
				mRI = j
			}
		}

		if minR == math.MaxInt32 {
			right = append(right, -1)
		} else {
			right = append(right, mRI)
		}

	}

	ans := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		cur := 0
		l := left[i]
		r := right[i]
		if l != -1 && r != -1 {
			cur = nums[i] + nums[l] + nums[r]
		} else {
			cur = -1
		}

		if cur >= 0 {
			ans = min(ans, cur)
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b

}
