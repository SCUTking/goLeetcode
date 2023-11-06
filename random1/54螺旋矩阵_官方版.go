package random1

func spiralOrder1(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])

	var (
		row, column    = 0, 0
		directions     = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
		directionIndex = 0
	)
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	ans := make([]int, 0)
	sum := m * n
	for len(ans) < sum {
		ans = append(ans, matrix[row][column])
		visited[row][column] = true
		nextRow, nextColumn := row+directions[directionIndex][0], column+directions[directionIndex][1]
		if nextRow < 0 || nextRow >= m || nextColumn < 0 || nextColumn >= n || visited[nextRow][nextColumn] {
			directionIndex = (directionIndex + 1) % 4
		}
		row += directions[directionIndex][0]
		column += directions[directionIndex][1]

	}

	return ans
}
