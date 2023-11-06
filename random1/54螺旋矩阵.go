package random1

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	iStart, iEnd, jStart, jEnd := 1, m-1, 0, n-1
	i, j := 0, 0
	flag := 0
	ans := make([]int, 0)
	ans = append(ans, matrix[0][0])
	for i >= iStart && i <= iEnd && j >= jStart && j <= jEnd {
		if flag == 0 {
			for j < jEnd {
				j++
				ans = append(ans, matrix[i][j])
			}
			flag = 1
			jEnd--
		} else if flag == 1 {
			for i < iEnd {
				i++
				ans = append(ans, matrix[i][j])
			}
			flag = 2
			iEnd--
		} else if flag == 2 {
			for j > jStart {
				j--
				ans = append(ans, matrix[i][j])
			}
			flag = 3
			jStart++
		} else {
			for i > iStart {
				i--
				ans = append(ans, matrix[i][j])
			}
			flag = 0
			iStart++
		}
	}

	return ans
}
