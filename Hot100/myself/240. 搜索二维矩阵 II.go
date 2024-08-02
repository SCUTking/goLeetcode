package myself

import "sort"

func searchMatrix(matrix [][]int, target int) bool {

	for i := 0; i < len(matrix); i++ {
		first := matrix[i][0]
		if first > target {
			continue
		}
		index := sort.SearchInts(matrix[i], target)
		if target != matrix[i][index] || index == len(matrix) {
			//不存在
		} else {
			return true
		}
	}
	return false
}
