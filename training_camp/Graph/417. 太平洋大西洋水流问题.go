package Graph

import "math"

func pacificAtlantic(heights [][]int) [][]int {
	m := len(heights)
	n := len(heights[0])
	p := make([][]bool, m)
	for i := 0; i < m; i++ {
		p[i] = make([]bool, n)
	}

	x := make([][]bool, m)
	for i := 0; i < m; i++ {
		x[i] = make([]bool, n)
	}

	var pDfs func(i, j int, pre int) bool
	pDfs = func(i, j int, pre int) bool {
		if i < 0 || j < 0 {
			return true
		}
		if i >= m || j >= n || p[i][j] {
			return false
		}
		cur := heights[i][j]

		if cur > pre {
			return false
		}
		return pDfs(i+1, j, cur) || pDfs(i-1, j, cur) || pDfs(i, j-1, cur) || pDfs(i, j+1, cur)
	}

	var xDfs func(i, j int, pre int) bool
	xDfs = func(i, j int, pre int) bool {

		if i >= m || j >= n {
			return true
		}
		if i < 0 || j < 0 || x[i][j] {
			return false
		}
		cur := heights[i][j]

		if cur > pre {
			return false
		}
		return xDfs(i+1, j, cur) || xDfs(i-1, j, cur) || xDfs(i, j-1, cur) || xDfs(i, j+1, cur)
	}
	res := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			x[i][j] = true
			p[i][j] = true
			if xDfs(i, j, math.MaxInt32) && pDfs(i, j, math.MaxInt32) {
				res = append(res, []int{i, j})
			}
			x[i][j] = false
			p[i][j] = false

		}
	}

	return res
}

// 逆流而上
func pacificAtlantic1(heights [][]int) [][]int {
	m := len(heights)
	n := len(heights[0])
	p := make([][]bool, m)
	for i := 0; i < m; i++ {
		p[i] = make([]bool, n)
	}

	x := make([][]bool, m)
	for i := 0; i < m; i++ {
		x[i] = make([]bool, n)
	}

	var dfs func(i, j, pre int, nums [][]bool)
	dfs = func(i, j, pre int, nums [][]bool) {

		if i >= m || j >= n || i < 0 || j < 0 || nums[i][j] {
			return
		}
		cur := heights[i][j]

		if cur >= pre {
			nums[i][j] = true
			dfs(i+1, j, cur, nums)
			dfs(i-1, j, cur, nums)
			dfs(i, j+1, cur, nums)
			dfs(i, j-1, cur, nums)
		}
	}

	for i := 0; i < m; i++ {
		dfs(i, n-1, math.MinInt32, x)
		dfs(i, 0, math.MinInt32, p)
	}

	for i := 0; i < n; i++ {
		dfs(0, i, math.MinInt32, p)
		dfs(m-1, i, math.MinInt32, x)
	}
	res := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[i][j] && x[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}
