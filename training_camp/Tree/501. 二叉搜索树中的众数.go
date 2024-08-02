package Tree

func findMode(root *TreeNode) []int {

	var base, count, maxCount int
	answer := make([]int, 0)
	update := func(val int) {
		if val == base {
			count++
		} else {
			base = val
			count = 1
		}

		if count == maxCount {
			answer = append(answer, base)
		}
		if count > maxCount {
			maxCount = count
			answer = []int{base}
		}
	}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		update(node.Val)
		dfs(node.Right)
	}
	dfs(root)

	return answer
}
