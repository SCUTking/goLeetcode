package coding

import "math"

// 从后往前遍历，用两个指针表示交换的索引，
// 记录最大的索引，一直遍历到0（因为最大的肯定要放到最前面），然后交换
func maximumSwap(num int) int {
	nums := make([]int, 0)
	for num > 0 {
		b := num % 10
		num = num / 10
		nums = append(nums, b)
	}
	for true {
		m := make(map[int]bool)
		var index int
		if len(m) == len(nums) {
			break
		}
		for i := 0; i < len(nums); i++ {
			if m[index] {
				continue
			}
			if nums[index] < nums[i] {
				index = i
			}
		}
		l := len(m)
		if l != index {
			nums[len(nums)-l-1], nums[index] = nums[index], nums[len(nums)-l-1]
			break
		} else {
			m[index] = true
		}
	}

	var res int
	for i := 0; i < len(nums); i++ {
		res += int(math.Pow10(i)) * nums[i]
	}

	return res
}
