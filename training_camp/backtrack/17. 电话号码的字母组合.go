package backtrack

func letterCombinations(digits string) []string {

	result := make([]string, 0, 10)
	mapletter := make(map[uint8][]string, 8)
	mapletter[2] = []string{"a", "b", "c"}
	mapletter[3] = []string{"d", "e", "f"}
	mapletter[4] = []string{"g", "h", "i"}
	mapletter[5] = []string{"j", "k", "l"}
	mapletter[6] = []string{"m", "n", "o"}
	mapletter[7] = []string{"p", "q", "r", "s"}
	mapletter[8] = []string{"v", "t", "u"}
	mapletter[9] = []string{"w", "x", "y", "z"}
	if len(digits) == 0 {
		return nil
	}
	var dfs func(pre string, index int)
	dfs = func(pre string, index int) {
		if len(pre) == len(digits) {
			result = append(result, pre)
		}

		index++
		u := digits[index] - '0'
		for _, s := range mapletter[u] {
			dfs(pre+s, index)
		}

	}
	dfs("", -1)
	return result
}
