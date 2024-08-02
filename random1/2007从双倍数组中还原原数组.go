package random1

import "sort"

func findOriginalArray(changed []int) []int {
	//n := len(changed)
	//sort.Ints(changed)
	//
	//ans := make([]int, 0)
	//for i := 0; i < n; i++ {
	//	if len(changed) == 0 {
	//		break
	//	}
	//	last := changed[len(changed)-1]
	//	changed = changed[:len(changed)-1]
	//	if last&1 != 0 {
	//		return []int{}
	//	}
	//
	//	index := binarySearch(changed, last/2)
	//	if index != -1 {
	//		ans = append(ans, changed[index])
	//		changed = append(changed[:index], changed[index+1:]...)
	//	} else {
	//		return []int{}
	//	}
	//}
	//
	//return ans

	sort.Ints(changed)

	res := make([]int, 0)
	queue := make([]int, 0)
	n := len(changed)
	if n&1 == 0 {
		return res
	}
	for i := 0; i < n; i++ {
		if len(queue) == 0 {
			queue = append(queue, changed[i])
		} else {
			header := queue[0]
			queue = queue[1:]
			if header*2 == changed[i] {
				res = append(res, header)
			} else {
				queue = append(queue, changed[i])
			}
		}
	}

	if len(queue) != 0 {
		return []int{}
	}
	return res
}

// 找到对应数字的索引位置
func binarySearch(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (right-left)/2 + left
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
