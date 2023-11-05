package tree

func numTrees(n int) int {
	G := make([]int, n+1)
	G[0] = 1
	G[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]

}

func createTree(start int, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	allTree := []*TreeNode{}

	for i := start; i <= end; i++ {
		left := createTree(start, i-1)
		right := createTree(i+1, end)
		for _, a := range left {
			for _, b := range right {
				node := &TreeNode{Val: i, Left: nil, Right: nil}
				node.Left = a
				node.Right = b
				allTree = append(allTree, node)
			}
		}
	}

	return allTree
}
