package sliding_window

func NumSubarrayProductLessThanK(nums []int, k int) int {

	//res := 0

	//for t := 1; t <= len(nums); t++ {
	//	left, right := 0, 0
	//	curMulti := 1
	//	for right < len(nums) {
	//		i := nums[right]
	//		curMulti = i * curMulti
	//		right++
	//
	//		if curMulti < k && right-left == t {
	//			res++
	//		}
	//
	//		for right-left >= t {
	//			l := nums[left]
	//			curMulti = curMulti / l
	//			left++
	//		}
	//
	//	}
	//}
	//
	//return res

	left, right := 0, 0
	cur := 1
	res := 0
	for right < len(nums) {
		cur *= nums[right]
		right++

		for left < right && cur >= k {

			cur /= nums[left]
			left++
		}
		if cur < k {
			res += right - left
		}
	}

	return res
}
