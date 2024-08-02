package Monotone_stack

func dailyTemperatures(temperatures []int) []int {

	//单调递减栈
	stack := make([]int, 0)
	res := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {

		for len(stack) > 0 && stack[len(stack)-1] < temperatures[i] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[idx] = i - idx
		}
		stack = append(stack, i)

		if i == len(temperatures)-1 {
			for j := 0; j < len(stack); j++ {
				res[stack[j]] = 0
			}
		}
	}

	return res
}
