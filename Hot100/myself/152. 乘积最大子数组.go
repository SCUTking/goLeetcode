package myself

import (
	"goLeetcode/Hot100"
	"math"
)

func maxProduct(nums []int) int {
	temp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		temp[i] = nums[len(nums)-i-1]
	}

	return Hot100.Max(getMax(nums), getMax(temp))

}
func getMax(nums []int) int {
	n := len(nums)
	pre := make([]int, n+1)
	pre[0] = 1
	flag := true
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 && flag {
			pre[i+1] = 1
			continue

		} else {
			flag = false
			pre[i+1] = pre[i] * nums[i]
		}

	}

	res := math.MinInt32
	for i := 0; i < n+1; i++ {
		for j := i + 1; j < n+1; j++ {
			if pre[i] == 0 {
				continue
			}
			cur := pre[j] / pre[i]
			if cur > res {
				res = cur
			}
		}
	}

	return res
}
