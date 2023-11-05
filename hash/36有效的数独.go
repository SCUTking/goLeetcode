package hash

func IsValidSudoku(board [][]byte) bool {

	var rows [9][9]int
	var columns [9][9]int
	var niu [3][3][9]int

	for i1, row := range board {
		for i2, row1 := range row {
			if row1 != '.' {
				index := row1 - '1'

				rows[i1][index]++
				columns[i2][index]++
				niu[i1/3][i2/3][index]++
				if rows[i1][index] > 1 || columns[i2][index] > 1 || niu[i1/3][i2/3][index] > 1 {
					return false

				}
			}

		}
	}
	return true
}

func checkIfValid(niu []byte) bool {
	set := make(map[byte]string, 9)

	for i := 0; i < len(niu); i++ {
		b := niu[i]
		if b != '.' {
			if set[b] != "" {
				return false
			}
			set[b] = "1"
		}

	}
	return true
}
