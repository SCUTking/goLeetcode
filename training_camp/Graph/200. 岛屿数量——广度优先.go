package Graph

func numIslands1(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	queue := make([][]int, 0)
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				queue = append(queue, []int{i, j})
				for len(queue) > 0 {
					ints := queue[0]
					queue = queue[1:]
					x := ints[0]
					y := ints[1]

					if x < 0 || x > m-1 || y < 0 || y > n-1 || grid[x][y] == '0' {
						continue
					} else {
						grid[x][y] = '0'
						queue = append(queue, []int{x - 1, j})
						queue = append(queue, []int{x + 1, j})
						queue = append(queue, []int{x, j - 1})
						queue = append(queue, []int{x, j + 1})
					}
				}
				res++
			}
		}
	}

	return res
}
