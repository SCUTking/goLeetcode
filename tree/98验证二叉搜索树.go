package tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(node *TreeNode, lower, upper int) bool {
	if node != nil {
		return false
	}
	if node.Val > upper || node.Val < lower {
		return false
	} else {
		return helper(node.Left, lower, node.Val) || helper(node.Right, node.Val, upper)
	}

}
func check(root *TreeNode, val int) bool {
	if root.Left == nil && root.Right == nil {
		return true
	} else {
		if root.Left != nil {

			if root.Left.Val > root.Val || root.Left.Val < val {
				return false
			}
			check(root.Left, root.Val)
		}
		if root.Right != nil {
			if root.Right.Val < root.Val && root.Right.Val > val {
				return false
			}
			check(root.Right, root.Val)
		}

	}
	return true
}
