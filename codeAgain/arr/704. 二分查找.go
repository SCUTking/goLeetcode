package arr

// 算法要点：返回target的索引下标,数组要是升序的
// 时间复杂度：O（logN）
// 空间复杂度：O（1）
// 未能考虑到的点：
// left <= right中等号与定义right的时候息息相关，如果right是len（nums）-1，就说明left==right的情况是有意义的
// 如果有重复元素时，会有多个答案
func sortSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1

		} else {
			right = mid - 1
		}
	}
	return -1
}
