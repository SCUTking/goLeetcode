package tree

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	res := &TreeNode{Val: preorder[0], Left: nil, Right: nil}

	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	res.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	res.Right = buildTree(preorder[len(inorder[i+1:])+1:], inorder[i+1:])

	return res
}
