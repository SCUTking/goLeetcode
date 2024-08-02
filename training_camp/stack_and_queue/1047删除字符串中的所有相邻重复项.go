package stack_and_queue

func removeDuplicates(s string) string {
	runes := []rune(s)
	ifOver := false
	for !ifOver {
		pre := 0
		for i := 1; i < len(runes); i++ {
			if runes[pre] == runes[i] {
				runes = append(runes[:pre], runes[pre+1:]...)
				runes = append(runes[:i], runes[i+1:]...)
				break
			} else {
				pre++
			}
		}
		if pre == len(runes)-1 {
			ifOver = true
		}
	}

	return string(runes)

}
func removeDuplicates1(s string) string {

	stack := make([]int32, 0)
	for _, i := range s {
		if len(stack) > 0 && stack[len(stack)-1] == i {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, i)
		}
	}

	return string(stack)
}
