package Graph

func ladderLength(beginWord string, endWord string, wordList []string) int {
	n := len(wordList)
	grid := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		grid[i] = make([]bool, n+1)
	}

	var check func(a, b string) bool
	check = func(a, b string) bool {
		l := len(a)
		if l != len(b) {
			return false
		}
		count := 0
		for i := 0; i < l; i++ {
			//一个字节表示一个小写字母
			if a[i] == b[i] {

			} else {
				count++
			}
		}
		if count != 1 {
			return false
		}
		return true
	}

	newList := append([]string{beginWord}, wordList...)
	for i := 0; i < n+1; i++ {
		for j := 0; j < n+1; j++ {
			if newList[i] != newList[j] {
				b := check(newList[i], newList[j])
				grid[i][j] = b
				grid[j][i] = b
			}
		}
	}
	visited := make([]bool, n+1)
	var res int = 0
	queue := make([]int, 0)
	queue = append(queue, 0)
	for len(queue) > 0 {
		res++
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			x := queue[0]
			queue = queue[1:]
			visited[i] = true
			for j := 0; j < len(grid[x]); j++ {
				if grid[x][j] && !visited[i] {
					if newList[j] == endWord {
						return res
					} else {
						queue = append(queue, j)
					}
				}
			}
		}
	}

	return 0
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
