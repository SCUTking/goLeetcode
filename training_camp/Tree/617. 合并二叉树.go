package Tree

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root1 == nil && root2 == nil {
		return nil
	}
	node := &TreeNode{}
	if root1 == nil {
		node.Val = root2.Val
		node.Left = mergeTrees(nil, root2.Left)
		node.Right = mergeTrees(nil, root2.Right)
	}
	if root2 == nil {
		node.Val = root1.Val
		node.Left = mergeTrees(nil, root1.Left)
		node.Right = mergeTrees(nil, root1.Right)
	}
	node.Val = root1.Val + root2.Val
	node.Left = mergeTrees(root1.Left, root2.Left)
	node.Right = mergeTrees(root1.Right, root2.Right)
	return node
}

func mergeTrees1(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

func mergeTrees2(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}
