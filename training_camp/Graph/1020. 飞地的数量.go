package Graph

func numEnclaves(grid [][]int) int {
	var dfs func(i, j int)
	m := len(grid)
	n := len(grid[0])

	dfs = func(i, j int) {
		if i < 0 || i > m-1 || j < 0 || j > n-1 || grid[i][j] == 0 {
			return
		}

		if grid[i][j] == 1 {
			grid[i][j] = 0
		}
		dfs(i-1, j)
		dfs(i, j-1)
		dfs(i+1, j)
		dfs(i, j+1)
	}

	//先处理边界的
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				dfs(i, j)
			}
		}
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res++
			}
		}
	}

	return res

}
