package tree

func generateTrees(n int) []*TreeNode {

	return help(1, n)

}

func help(start int, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	allTree := []*TreeNode{}
	for i := start; i <= end; i++ {
		left := help(start, i-1)
		right := help(i+1, end)
		for _, l := range left {
			for _, r := range right {
				tmp := &TreeNode{i, nil, nil}
				tmp.Left = l
				tmp.Right = r
				allTree = append(allTree, tmp)
			}
		}
	}
	return allTree

}
