package arr

var direction [][]int

func init() {
	direction = make([][]int, 4)
	direction[0] = []int{0, 1}
	direction[1] = []int{1, 0}
	direction[2] = []int{0, -1}
	direction[3] = []int{-1, 0}
}

// 算法要点：找到第一个非负数的索引下标，利用双指针分别遍历
// 时间复杂度：O（N）
// 空间复杂度：O（1）
// 未能考虑到的点：
// 新赋值两个变量，标记原方向下下一移动要到哪，用来提前判断是否需要转换方向，再实际的进行移动能够保证不越界
func matrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	i, j := 0, 0
	index := 0
	for cur := 1; cur <= n*n; cur++ {
		res[i][j] = cur
		nextI, nextJ := i+direction[index][0], j+direction[index][1]
		if nextI > n-1 || nextJ < 0 || nextJ > n-1 || nextI < 0 || res[nextI][nextJ] != 0 {
			index = (index + 1) % 4
		}
		i += direction[index][0]
		j += direction[index][1]
	}
	return res
}
