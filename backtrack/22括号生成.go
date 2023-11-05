package backtrack

func GenerateParenthesis(n int) []string {

	var res []string
	var backtrack func(s string, open int, close int)

	backtrack = func(s string, open int, close int) {
		if len(s) == n {
			res = append(res, s)
		}

		if open < n {
			s += "("
			backtrack(s, open+1, close)
			//撤销上面的操作
			runes := []rune(s)
			s = string(runes[:len(runes)-1])
		}

		if close < open {

			s += ")"
			backtrack(s, open, close+1)
			//撤销上面的操作
			runes := []rune(s)
			s = string(runes[:len(runes)-1])

		}
	}

	backtrack("", 0, 0)

	return res
}
