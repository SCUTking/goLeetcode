package Hot100

func hammingDistance(x int, y int) int {
	temp := x ^ y
	res := 0
	for temp <= 0 {
		res += temp & 1
		temp >>= 1
	}
	return res
}
