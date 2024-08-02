package Graph

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	grid := make([][]bool, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		ints := rooms[i]
		for _, next := range ints {
			grid[i][next] = true
		}
	}

	visited := make([]bool, n)
	visited[0] = true
	res := false

	var dfs func(i int, count int)
	dfs = func(i int, count int) {
		if !visited[i] {
			if count == n {
				res = true
			} else {
				return
			}
		}

		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] {
				visited[j] = true
				dfs(j, count+1)
				visited[j] = false
			}
		}

	}

	dfs(0, 1)
	return res
}
