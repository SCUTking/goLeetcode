package string

func CheckValidString(s string) bool {
	if s == "" {
		return true
	}
	stack := make([]string, 0)
	for i := 0; i < len(s); i++ {
		temp := s[i]
		if temp != ')' {
			stack = append(stack, string(temp))
		} else {
			xinNum := 0
			for len(stack) > 0 && stack[len(stack)-1] != string('(') {
				if stack[len(stack)-1] == string("*") {
					xinNum++
				}
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				if xinNum == 0 {
					return false
				} else {
					for i := 0; i < xinNum-1; i++ {
						stack = append(stack, string('*'))
					}
				}
			} else {
				stack = stack[:len(stack)-1]
				for i := 0; i < xinNum; i++ {
					stack = append(stack, string('*'))
				}
			}

		}
	}

	if len(stack) != 0 {
		for i, s2 := range stack {
			if s2 == "(" {
				flag := false
				for j := i; j < len(stack); j++ {
					if stack[j] == "*" {
						stack[j] = "a"
						flag = true
						break
					}
				}
				if !flag {
					return false
				}
			}
		}
	}
	return true

}
