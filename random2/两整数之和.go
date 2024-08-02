package random2

func getSum(a int, b int) int {

	res := 0
	jinwei := 0
	for i := 0; i < 32; i++ {
		numOne := 0
		index1 := a & (1 << i)
		index2 := b & (1 << i)
		if index1 == 1 {
			numOne++
		}
		if index2 == 1 {
			numOne++
		}
		if jinwei == 1 {
			numOne++
		}
		jinwei = 0
		if numOne >= 2 {
			jinwei = 1
		}
		res += (numOne % 2) << i
	}

	return res

}

func getSum1(a int, b int) int {
	//我们赋给无符号类型一个超出它表示范围的值时
	//结果是初始值对无符号类型表示数值总数取模的余数
	for b != 0 {
		carry := uint(a&b) << 1
		a ^= b
		b = int(carry)
	}
	return a
}
