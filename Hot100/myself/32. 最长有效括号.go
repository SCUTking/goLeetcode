package myself

func longestValidParentheses(s string) int {
	left, right := 0, 0
	leftCount := 0
	//var temp string
	res := 0
	for right < len(s) {
		u := s[right]
		if u == '(' {
			leftCount++
		} else {
			if leftCount > 0 {
				leftCount--
				//得从后往前遍历
				var rightCount int
				for i := right; i >= left; i-- {
					if s[i] == ')' {
						rightCount++
					} else {
						rightCount--
						if rightCount < 0 {
							l1 := right - i
							if res < l1 {
								res = l1
							}
							break
						}
					}
					if i == left && rightCount >= 0 {
						l1 := right - left + 1
						if res < l1 {
							res = l1
						}
					}
				}

			} else {
				//不符合情况
				right++
				left = right
				continue
			}
		}
		//temp = string(append([]byte(temp), u))
		right++
	}

	//还得再最后处理一下

	return res
}
