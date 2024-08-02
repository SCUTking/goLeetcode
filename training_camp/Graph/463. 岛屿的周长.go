package Graph

func islandPerimeter(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	res := 0
	direction := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				for k := 0; k < 4; k++ {
					i += direction[k][0]
					j += direction[k][1]
					if i >= m || j >= n || i < 0 || j < 0 {
						res += 1
					}
					if grid[i][j] == 0 {
						res += 1
					}
				}

			}
		}
	}

	return res

}
