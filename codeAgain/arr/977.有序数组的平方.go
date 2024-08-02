package arr

// 算法要点：找到第一个非负数的索引下标，利用双指针分别遍历
// 时间复杂度：O（N）
// 空间复杂度：O（1）
// 未能考虑到的点：
// 不用找非负数的索引的下标，直接从两头开始就可以
func sortedArrSquare(nums []int) []int {
	//找到第一个非负数的索引下标，利用下标
	NonNegativeIndex := 0
	for nums[NonNegativeIndex] < 0 {
		NonNegativeIndex++
	}
	res := make([]int, 0)
	left, right := NonNegativeIndex-1, NonNegativeIndex
	for left >= 0 && right < len(nums) {
		cur := left
		if -nums[cur] > nums[right] {
			cur = right
			right++
		} else {
			left--
		}
		res = append(res, nums[cur]*nums[cur])
	}

	for left >= 0 {
		res = append(res, nums[left]*nums[left])
		left--
	}
	for right < len(nums) {
		res = append(res, nums[right]*nums[right])
		right++
	}
	return res
}

func best_sortedArrSquare(nums []int) []int {
	n := len(nums)
	left, right := 0, n-1
	res := make([]int, n)
	count := 1
	for left <= right {
		cur := left
		if nums[left]*nums[left] < nums[right]*nums[right] {
			cur = right
			right--
		} else {
			left++
		}
		res[n-count] = nums[cur] * nums[cur]
		count++
	}
	return res
}
