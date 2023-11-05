package hash

func SetZeroes(matrix [][]int) {

	rowMap := make(map[int]bool)
	colMap := make(map[int]bool)
	for i, m := range matrix {
		for j, n := range m {
			if n == 0 {
				rowMap[i] = true
				colMap[j] = true
			}
		}
	}

	for k, _ := range rowMap {
		ints := matrix[k]
		for i := 0; i < len(ints); i++ {
			ints[i] = 0
		}
		matrix[k] = ints
	}

	for k, _ := range colMap {
		for i := 0; i < len(matrix); i++ {
			ints := matrix[i]
			ints[k] = 0
			matrix[i] = ints
		}

	}

}
