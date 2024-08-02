package Tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	path1 := findPath(root, p.Val)
	path2 := findPath(root, q.Val)
	for i := len(path1) - 1; i >= 0; i-- {
		for j := len(path2) - 1; j >= 0; j-- {
			if path2[j].Val == path1[i].Val {
				return path2[j]
			}
		}
	}
	return nil
}
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	parents := make(map[int]*TreeNode)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			parents[node.Left.Val] = node
			dfs(node.Left)
		}
		if node.Right != nil {
			parents[node.Right.Val] = node
			dfs(node.Right)
		}
	}
	dfs(root)
	visits := make(map[*TreeNode]bool)
	for p != nil {
		visits[p] = true
		p = parents[p.Val]
	}

	for q != nil {
		_, ok := visits[q]
		if ok {
			return q
		}
		q = parents[q.Val]
	}

	return nil

}
func findPath(root *TreeNode, target int) []*TreeNode {

	res := make([]*TreeNode, 0)
	var dfs func(node *TreeNode, path []*TreeNode)
	dfs = func(node *TreeNode, path []*TreeNode) {
		if node == nil {
			return
		}

		path = append(path, node)
		if node.Val == target {
			res = path
			return
		}

		dfs(node.Left, path)
		dfs(node.Right, path)
	}
	dfs(root, []*TreeNode{})
	return res
}
