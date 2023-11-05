package backtrack

func numIslands(grid [][]byte) int {

	m := len(grid)
	n := len(grid[0])
	var backHelper func(i, j int)
	var inArea func(i, j int) bool
	inArea = func(i, j int) bool {
		return 0 <= i && i < m && 0 <= j && j < n
	}

	backHelper = func(i, j int) {
		if !inArea(i, j) {
			return
		}
		if grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'

		backHelper(i+1, j)
		backHelper(i, j+1)
		backHelper(i-1, j)
		backHelper(i, j-1)
	}
	ans := 0
	for i, bytes := range grid {
		for j, e := range bytes {
			if e == '1' {
				backHelper(i, j)
				ans++
			}
		}
	}

	return ans

}
