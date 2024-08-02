package sliding_window

func longestOnes(nums []int, k int) int {
	n := len(nums)
	left, right := 0, 0
	zeroCount := 0
	res := 0
	for right < n {
		if nums[right] == 0 {
			zeroCount++
		} else {

		}
		if zeroCount < k {
			res = max(res, right-left+1)
		}
		for zeroCount > k {
			if nums[left] == 1 {

			} else {
				zeroCount--
			}
			left++
			if zeroCount < k {
				res = max(res, right-left+1)
			}
		}
		right++
	}

	return res
}
