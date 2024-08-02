package Tree

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	dad := findDad(root, val)
	if dad.Val > val {
		if dad.Left != nil {
			left := dad.Left
			dad.Left = &TreeNode{Val: val, Left: left}
		} else {
			dad.Left = &TreeNode{Val: val}
		}
	}
	if dad.Val < val {
		if dad.Right != nil {
			Right := dad.Right
			dad.Right = &TreeNode{Val: val, Right: Right}
		} else {
			dad.Right = &TreeNode{Val: val}
		}
	}
	return root
}
func findDad(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > val {
		if root.Left == nil {
			return root
		}
		if root.Left != nil && root.Left.Val < val {
			return root
		} else {
			return findDad(root.Left, val)
		}

	}
	if root.Val < val {
		if root.Right == nil {
			return root
		}
		if root.Right != nil && root.Right.Val > val {
			return root
		} else {
			return findDad(root.Right, val)
		}

	}
	return nil
}
func insertIntoBST1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	var parent *TreeNode
	cur := root
	for cur != nil {
		if cur.Val > val {
			parent = cur
			cur = cur.Left
		} else if cur.Val < val {
			parent = cur
			cur = cur.Right
		} else {
			//说明有重复的元素
		}
	}

	cur = &TreeNode{Val: val}
	if val > parent.Val {
		parent.Right = cur
	} else {
		parent.Left = cur
	}

	return root
}
