package stack_and_queue

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]string, 0)
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			front, _ := strconv.Atoi(stack[len(stack)-2])
			back, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-2]
			var temp string
			switch tokens[i] {
			case "+":
				temp = strconv.Itoa(front + back)
			case "-":
				temp = strconv.Itoa(front - back)
			case "*":
				temp = strconv.Itoa(front * back)
			case "/":
				temp = strconv.Itoa(front / back)
			}
			stack = append(stack, temp)

		} else {
			stack = append(stack, tokens[i])
		}
	}
	itoa, _ := strconv.Atoi(stack[0])
	return itoa
}
