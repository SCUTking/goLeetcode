package random

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	//二分查找默认得是有序的
	//
	//左闭右闭  left==right是有意义的  然后right与left在边界处理的时候都进行加一或减一
	for left <= right {
		middle := left + (right-left)/2 //防止溢出
		if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			return nums[middle]
		}
	}

	return -1

}
