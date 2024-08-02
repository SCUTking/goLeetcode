package Tree

import "math"

func getMinimumDifference(root *TreeNode) int {
	res := make([]int, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	pre := res[0]
	min := math.MaxInt32
	for i := 1; i < len(res); i++ {
		if res[i]-pre < min {
			min = res[i] - pre
		}
		pre = res[i]
	}
	return min

}
