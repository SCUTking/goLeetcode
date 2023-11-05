package tree

import "math"

func RecoverTree(root *TreeNode) {
	var pre *TreeNode
	var twoNumber [2]*TreeNode
	pre = &TreeNode{Val: math.MinInt64}
	findTwoNumber(root, pre, twoNumber)
	a := twoNumber[0].Val
	b := twoNumber[1].Val
	swap(a, b, root)

}

func findTwoNumber(root *TreeNode, pre *TreeNode, twoNumber [2]*TreeNode) {
	if root == nil {

		return
	}
	if root.Left != nil {
		findTwoNumber(root.Left, pre, twoNumber)
	}

	if root.Val < pre.Val {
		if twoNumber[1] == nil {
			twoNumber[0] = pre
		}

		if twoNumber[1] == nil {
			twoNumber[1] = root
		}
		if twoNumber[0] != nil {
			twoNumber[1] = root
		}
	}
	pre = root

	if root.Right != nil {
		findTwoNumber(root.Right, pre, twoNumber)
	}

}

func swap(a, b int, root *TreeNode) {

	if root == nil {
		return
	}
	swap(a, b, root.Right)

	if root.Val == a || root.Val == b {
		if root.Val == a {
			root.Val = b
		} else {
			root.Val = a
		}
	}
	swap(a, b, root.Left)

}
