package Tree

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	switch {
	case root == nil:
		return nil
	case root.Val > high:
		return trimBST(root.Left, low, high)
	case root.Val < low:
		return trimBST(root.Right, low, high)
	default:
		root.Left = trimBST(root.Left, low, high)
		root.Right = trimBST(root.Right, low, high)
		return root
	}
}
