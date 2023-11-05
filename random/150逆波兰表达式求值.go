package random

import "strconv"

func EvalRPN(tokens []string) int {

	stack := []int{}
	for _, i := range tokens {
		atoi, ok := strconv.Atoi(i)
		if ok != nil {
			//说明是运算符号
			i2 := stack[len(stack)-1]
			i3 := stack[len(stack)-2]

			stack = stack[:len(stack)-2]
			value := 0
			if i == "+" {
				value = i3 + i2
			} else if i == "/" {
				value = i3 / i2
			} else if i == "-" {
				value = i3 - i2
			} else if i == "*" {
				value = i3 * i2
			}

			stack = append(stack, value)
		} else {
			stack = append(stack, atoi)
		}

	}

	return stack[0]

}
