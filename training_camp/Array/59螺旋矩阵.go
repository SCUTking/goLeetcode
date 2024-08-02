package Array

func generateMatrix(n int) [][]int {
	direction := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}

	i, j, dir := 0, 0, 0
	for num := 1; num <= n*n; num++ {
		ans[i][j] = num
		if r, c := i+direction[dir][0], j+direction[dir][1]; r < 0 || r > n-1 || c < 0 || c > n-1 || ans[i][j] != 0 {
			dir = (dir + 1) % 4
		}
		i += direction[dir][0]
		j += direction[dir][1]
	}
	return ans
}
