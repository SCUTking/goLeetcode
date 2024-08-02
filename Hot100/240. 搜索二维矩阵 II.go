package Hot100

// 时间复杂度 O(m+n)
func searchMatrix(matrix [][]int, target int) bool {
	//z字型查找
	m, n := len(matrix), len(matrix[0])
	i, j := 0, n-1
	for i < m && j >= 0 {
		cur := matrix[i][j]
		if cur == target {
			return true
		}

		if cur > target {
			j--
		} else {
			i++
		}
	}

	return false
}
