package random

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func RightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	leftH := treeHeight(root.Left, 0)
	rightH := treeHeight(root.Right, 0)

	size := max(leftH, rightH) + 1

	ans := []int{}
	var help func(root *TreeNode, cur int)
	help = func(root *TreeNode, cur int) {
		if root == nil {
			return
		}
		if len(ans) < cur {
			ans = append(ans, root.Val)
		}
		cur++
		if root.Right != nil {
			help(root.Right, cur)
		}
		if root.Left != nil {
			help(root.Left, cur)
		}
	}

	var leftHelp func(root *TreeNode, h int, leftsize int)
	leftHelp = func(root *TreeNode, h int, leftsize int) {
		if root == nil {
			return
		}
		if size-h < leftsize {
			if len(ans) < h {
				ans = append(ans, root.Val)
			}
		}
		h--
		if root.Right != nil {
			leftHelp(root.Right, h, leftsize)
		}
		if root.Left != nil {
			leftHelp(root.Right, h, leftsize)
		}
	}

	help(root, 1)
	//if leftH <= rightH {
	//	help(root, 1)
	//} else {
	//	help(root, 1)
	//	//leftHelp(root.Left, 2, leftH-rightH)
	//
	//}

	return ans
}

func treeHeight(root *TreeNode, h int) int {
	if root == nil {
		return h
	}
	if root != nil {
		h++
	}
	return max(treeHeight(root.Left, h), treeHeight(root.Right, h))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
