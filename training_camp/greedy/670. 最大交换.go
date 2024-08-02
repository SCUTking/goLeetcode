package greedy

import (
	"math"
	"strconv"
)

func maximumSwap1(num int) int {
	nums := make([]int, 0)
	for num > 0 {
		b := num % 10
		num = num / 10
		nums = append(nums, b)
	}
	m := make(map[int]bool)
	for true {

		var index int = -1
		if len(m) == len(nums) {
			break
		}
		for i := 0; i < len(nums); i++ {
			if m[i] {
				continue
			}
			if index == -1 {
				index = i
			}
			if nums[index] < nums[i] {
				index = i
			}
		}
		l := len(m)
		if len(nums)-l-1 != index {
			if nums[index] == nums[len(nums)-l-1] {
				m[len(nums)-l-1] = true
			} else {
				nums[len(nums)-l-1], nums[index] = nums[index], nums[len(nums)-l-1]
				break
			}
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
func maximumSwap2(num int) int {
	nums := []byte(strconv.Itoa(num))
	n := len(nums)
	maxIndex, index1, index2 := n-1, -1, -1
	for i := n - 1; i >= 0; i-- {
		//相等时不处理
		//因为相同的maxindex，选择索引更靠后边的

		if nums[i] > nums[maxIndex] {
			maxIndex = i
		} else if nums[i] < nums[maxIndex] {
			index1, index2 = i, maxIndex
		}
	}

	if index1 < 0 {
		return num
	}
	nums[index1], nums[index2] = nums[index2], nums[index1]
	v, _ := strconv.Atoi(string(nums))

	return v
}
