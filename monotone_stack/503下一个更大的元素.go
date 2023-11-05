package monotone_stack

func NextGreaterElements(nums []int) []int {
	//stack := []int{nums[0]}
	//instack := make(map[int]int)
	//instack[nums[0]] = 1
	//
	//for i := 1; i < len(nums); i++ {
	//	num := nums[i]
	//	if instack[num] == 0 {
	//		temp := []int{}
	//		if stack[len(stack)-1] < num {
	//			stack = append(stack, num)
	//			instack[num]++
	//			continue
	//		}
	//		for len(stack) > 0 && stack[len(stack)-1] > num {
	//			temp = append(temp, stack[len(stack)-1])
	//			stack = stack[:len(stack)-1]
	//		}
	//
	//		stack = append(stack, num)
	//		instack[num]++
	//		for j := len(temp) - 1; j >= 0; j-- {
	//			stack = append(stack, temp[j])
	//		}
	//
	//	} else {
	//		continue
	//	}
	//
	//}
	//ans := []int{}
	//for _, num := range nums {
	//	for i3, i2 := range stack {
	//		if i2 == num {
	//			if i3 < len(stack)-1 {
	//				ans = append(ans, stack[i3+1])
	//			} else {
	//				ans = append(ans, -1)
	//			}
	//
	//		}
	//	}
	//}
	//
	//return ans

	doubleArr := nums
	for _, i := range nums {
		doubleArr = append(doubleArr, i)
	}

	ans := []int{}
	for i := 0; i < len(nums); i++ {
		l := len(ans)
		for j := i + 1; j < len(doubleArr); j++ {
			if nums[i] < doubleArr[j] {
				ans = append(ans, doubleArr[j])
				break
			}
		}
		if l == len(ans) {
			ans = append(ans, -1)
		}
	}
	return ans
}
