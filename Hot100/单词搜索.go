package Hot100

func exist(board [][]byte, word string) bool {

	m, n := len(board), len(board[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	res := false
	var dfs func(index int, x, y int)
	dfs = func(index int, x, y int) {
		if x < 0 || y < 0 || x >= m || y >= n || word[index] != board[x][y] || visited[x][y] {
			return
		}
		if index == len(word)-1 {
			res = true
			return
		}

		visited[x][y] = true

		dfs(index+1, x-1, y)
		dfs(index+1, x+1, y)
		dfs(index+1, x, y+1)
		dfs(index+1, x, y-1)

		visited[x][y] = false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(0, i, j)
		}
	}

	return res

}
