package Tree

import "math"

func isValidBST(root *TreeNode) bool {
	res := make([]int, 0)
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		b1 := dfs(node.Left)
		if b1 {
			return false
		}
		if len(res) > 0 {
			last := res[len(res)-1]
			if last > root.Val {
				return false
			}
		}
		res = append(res, root.Val)
		b2 := dfs(node.Right)
		if b2 {
			return false
		}
		return true
	}

	return dfs(root)

}
func isValidBST1(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}
func helper(root *TreeNode, lower int, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}
