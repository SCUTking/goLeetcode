package Tree

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	index := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			index = i
			break
		}
	}
	rightNum := len(inorder) - index - 1
	root := &TreeNode{Val: rootVal}
	rightNums := postorder[len(postorder)-1-rightNum : len(postorder)-1]
	root.Left = buildTree(inorder[:index], rightNums)
	root.Right = buildTree(inorder[index+1:], postorder[:len(postorder)-1-rightNum])
	return root
}
