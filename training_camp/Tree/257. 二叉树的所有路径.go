package Tree

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	var AddString func(old string, node *TreeNode)
	AddString = func(old string, node *TreeNode) {
		if node == nil {
			return
		}
		if len(old) > 0 {
			old += "->"
		}

		old += strconv.Itoa(node.Val)

		if node.Left == nil && node.Right == nil {
			res = append(res, old)
		}
		AddString(old, node.Left)
		AddString(old, node.Right)
	}
	AddString("", root)
	return res
}
