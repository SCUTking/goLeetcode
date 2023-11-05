package backtrack

func Solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	var dfs func(i int, j int)
	dfs = func(i int, j int) {
		if j < 0 || j > n-1 || i < 0 || i > m-1 || board[i][j] != 'O' {
			return
		}
		board[i][j] = 'Y'

		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	for i, bytes := range board {
		for i2, b := range bytes {
			if (i == 0 || i == m-1 || i2 == 0 || i2 == n-1) && b-'O' == 0 {
				dfs(i, i2)
			}
		}
	}

	for _, bytes := range board {
		for _, b := range bytes {
			if b == 'O' {
				b = 'X'
			} else if b == 'Y' {
				b = 'O'
			}
		}
	}

}
