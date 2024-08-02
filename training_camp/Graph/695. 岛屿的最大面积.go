package Graph

func maxAreaOfIsland(grid [][]int) int {
	var dfs func(i, j int, sum int) int
	m := len(grid)
	n := len(grid[0])

	dfs = func(i, j int, sum int) int {
		if i < 0 || i > m-1 || j < 0 || j > n-1 || grid[i][j] == '0' {
			return 0
		}

		if grid[i][j] == '1' {
			sum += 1
			grid[i][j] = '0'
		}
		sum += dfs(i-1, j, sum)
		sum += dfs(i, j-1, sum)
		sum += dfs(i+1, j, sum)
		sum += dfs(i, j+1, sum)

		return sum
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				t := dfs(i, j, 0)
				res = max(res, t)
			}
		}
	}

	return res
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
