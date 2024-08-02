package arr

import "sort"

// 算法要点：返回target的索引下标,数组要是升序的
// 时间复杂度：O（nlogN）
// 空间复杂度：O（1）
// 未能考虑到的点：
// left <= right中等号与定义right的时候息息相关，如果right是len（nums）-1，就说明left==right的情况是有意义的
// 如果有重复元素时，会有多个答案
func removeTarget(nums []int, target int) int {
	sort.Ints(nums)
	search := sortSearch(nums, target)
	for search != -1 {
		nums = append(nums[:search], nums[search+1:]...)
		search = sortSearch(nums, target)
	}
	return len(nums)
}

// 快慢指针法 本体只是需要知道移除元素后的数组长度 所以能够使用此方法能将时间复杂度减低到O（N）
func removeElement(nums []int, target int) int {
	slow, fast := 0, 0
	for ; fast < len(nums); fast++ {
		if nums[fast] == target {
			continue
		}
		nums[slow] = nums[fast]
		slow++
	}

	return slow
}
