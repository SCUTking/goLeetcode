package Hot100

func removeInvalidParentheses(s string) []string {
	//得利用回溯的算法
	res := make(map[string]bool, 0)
	left, right := getLeastLeftAndRight(s)
	nlen := len(s)
	var helper func(pre string, index int)
	helper = func(pre string, index int) {
		if nlen-1-index < right+left || index == nlen-1 {
			return
		}
		//nlen-(left+right) == len(pre)
		if left == 0 && right == 0 {
			if ifVaild(pre) && !res[pre] {
				res[pre] = true
			}
			return
		}
		bytes := []byte(pre)
		bytes = append(bytes, s[index])
		helper(string(bytes), index+1)

		if s[index] == '(' && left > 0 {
			left--
		} else if s[index] == ')' && right > 0 {
			right--
		}
		helper(pre, index+1)
	}

	helper("", 0)

	ret := make([]string, 0)
	for k, _ := range res {
		ret = append(ret, k)
	}
	return ret

}

func getLeastLeftAndRight(s string) (left, right int) {
	left, right = 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else if s[i] == ')' {
			if left > 0 {
				left--
			} else {
				right++
			}
		}
	}
	return
}

func ifVaild(s string) bool {
	left, right := getLeastLeftAndRight(s)
	if left == 0 && right == 0 {
		return true
	}
	return false

}
