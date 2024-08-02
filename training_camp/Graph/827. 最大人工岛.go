package Graph

func largestIsland(grid [][]int) int {
	//用于记录标记区域的面积
	m := make(map[int]int, 0)

	n := len(grid)
	mark := 2 //从二开始因为 01 被用掉了
	//参数是 坐标，标记的索引，面积
	var dfs func(x, y int, mark int, s int)
	dfs = func(x, y int, mark int, s int) {
		if x < 0 || y < 0 || x >= n || y >= n || grid[x][y] != 1 {
			return
		}
		s++
		grid[x][y] = mark
		m[mark] = s
		dfs(x+1, y, mark, s)
		dfs(x, y-1, mark, s)
		dfs(x, y+1, mark, s)
		dfs(x-1, y, mark, s)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			//等于1就继续深度优先  如果刚好只有一个  就直接为1即可
			if grid[i][j] == 1 {
				dfs(0, 0, mark, 0)
			}
		}
	}

	m1 := make(map[int]int, 0) //存储mark号 用于去重
	res := 0
	var add func(x, y int) int
	add = func(x, y int) int {
		if x < 0 || y < 0 || x >= n || y >= n || grid[x][y] != 1 {
			return 0
		}
		temp, ok := m[grid[x][y]]
		if ok {
			m1[grid[x][y]] = 1
			return temp
		}
		return 0

	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			//等于0  判断上下左右是否为标记的区域
			if grid[i][j] == 0 {
				temp := 0
				add(i+1, j)
				add(i-1, j)
				add(i, j+1)
				add(i, j-1)

				for k, _ := range m1 {
					temp += m[k]
				}

				res = max(temp, res)
			}
		}
	}

	return res

}
