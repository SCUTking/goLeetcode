package monotone_stack

import "math"

func FindUnsortedSubarray(nums []int) int {
	//这种方法不行  因为1 2 4 5 3
	//现在只记录了5 3 ，有序后整个不是有序的
	//max := math.MinInt32
	//for _, num := range nums {
	//	if num > max {
	//		max = num
	//	}
	//}
	//
	//pre := math.MinInt32
	//indexArr := []int{}
	//for i, num := range nums {
	//	if pre > num {
	//		indexArr = append(indexArr, i)
	//	} else if pre == num {
	//		if max != num {
	//			indexArr = append(indexArr, i)
	//		}
	//	}
	//	pre = num
	//}
	//if len(indexArr) > 1 {
	//	return indexArr[len(indexArr)-1] - indexArr[0] + 2
	//} else if len(indexArr) == 1 {
	//	return 2
	//} else {
	//	return 0
	//}

	stack := []int{nums[0]}

	flag := math.MaxInt32
	last := -1
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		temp := []int{}
		for len(stack) > 0 && stack[len(stack)-1] > num {

			last = i
			temp = append(temp, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}

		if last != -1 {
			if flag > len(stack) {
				flag = len(stack)
			}
		}

		stack = append(stack, num)
		for k := len(temp) - 1; k >= 0; k-- {
			stack = append(stack, temp[k])
		}
	}
	if last == -1 {
		return 0
	}
	return last - flag + 1

}
