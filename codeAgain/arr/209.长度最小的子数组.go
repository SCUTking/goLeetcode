package arr

import "math"

// 算法要点：全部是正数，所以可以利用滑动窗口方法
// 时间复杂度：O(N)
// 空间复杂度：O（1）
// 未能考虑到的点：
func SmallestSubArr(nums []int, s int) int {
	left, right := 0, 0
	n := len(nums)
	res := math.MaxInt32
	curSum := 0
	for right < n {
		curSum += nums[right]
		if curSum >= s {
			l := right - left + 1
			if l < res {
				res = l
			}
		}
		for curSum >= s {
			curSum -= nums[left]
			left++
			if curSum >= s {
				l := right - left + 1
				if l < res {
					res = l
				}
			}
		}

		right++
	}

	return res
}
