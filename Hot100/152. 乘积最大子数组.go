package Hot100

import "math"

func maxProduct(nums []int) int {
	minDp, maxDp, res := nums[0], nums[0], nums[0] //以i结尾的最大的数组的值
	for i := 1; i < len(nums); i++ {
		maxT, minT := maxDp, minDp
		cur := nums[i]
		minDp = minOf(maxT*cur, minT*cur, cur)
		maxDp = maxOf(maxT*cur, minT*cur, cur)
		res = maxOf(res, maxDp)
	}
	return res
}

func maxOf(args ...int) int {
	res := math.MinInt32
	for i := 0; i < len(args); i++ {
		if res < args[i] {
			res = args[i]
		}
	}
	return res
}
