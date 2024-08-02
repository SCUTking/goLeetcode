package Tree

import "math"

func minDepth(root *TreeNode) int {
	return getMinHigh(root)
}

func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MinInt32
	if root.Left != nil {
		minD = Min(minDepth1(root.Left), minD)
	}
	if root.Right != nil {
		minD = Min(minDepth1(root.Right), minD)
	}

	return minD
}

func getMinHigh(node *TreeNode) int {
	if node == nil {
		return 0
	}
	h1 := getMinHigh(node.Right)
	h2 := getMinHigh(node.Left)
	if h1 == 0 {
		return h2 + 1
	}
	if h2 == 0 {
		return h1 + 1
	}
	return Min(h2, h1) + 1
}

func Min(a, b int) int {
	if a > b {
		return a
	}
	return b
}
