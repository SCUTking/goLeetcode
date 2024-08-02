package Graph

func solve(board [][]byte) {
	var dfs func(i, j int)
	m := len(board)
	n := len(board[0])

	dfs = func(i, j int) {
		if i < 0 || i > m-1 || j < 0 || j > n-1 || board[i][j] != 'O' {
			return
		}

		if board[i][j] == 'O' {
			board[i][j] = 'Y'
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

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'Y' {
				board[i][j] = 'O'
			}
		}
	}

}
