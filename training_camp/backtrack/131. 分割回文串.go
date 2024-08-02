package backtrack

func partition(s string) [][]string {
	res := make([][]string, 0)
	var pre []string
	var dfs func(PreIndex int, index int)
	dfs = func(PreIndex int, index int) {
		if index == len(s)+1 && PreIndex == index-1 {
			temp := make([]string, len(pre))
			copy(temp, pre)
			res = append(res, temp)
			return
		}
		tempStr := s[PreIndex:index]

		index++
		dfs(PreIndex, index)
		if isOk(tempStr) {
			pre = append(pre, tempStr)
			dfs(index-1, index)
		}
	}
	dfs(0, 1)
	return res
}
func isOk(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
