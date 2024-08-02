package stack_and_queue

func isValid(s string) bool {
	stack := make([]int32, 0)
	m := make(map[int32]int32)
	m['}'] = '{'
	m[')'] = '('
	m[']'] = '['

	for _, i := range s {
		if i == '{' || i == '(' || i == '[' {
			stack = append(stack, i)
		} else {
			flag := m[i]
			//ifHas := false
			if stack[len(stack)-1] == flag {
				stack = append(stack[:len(stack)-1], stack[len(stack)-1+1:]...)
			} else {
				return false
			}
			//for j := len(stack) - 1; j >= 0; j-- {
			//	if stack[j] == flag {
			//		stack = append(stack[:j], stack[j+1:]...)
			//		ifHas = true
			//	}
			//}
			//if !ifHas {
			//	return false
			//}
		}
	}

	return len(stack) == 0

}
