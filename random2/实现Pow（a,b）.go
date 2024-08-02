package random2

func myPow(x float64, n int) float64 {

	if x >= 0 {
		return QuickPow(x, n)
	}

	return 1 / QuickPow(x, n)

}

func QuickPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := QuickPow(x, n/2)
	if n&1 == 0 {
		//偶数直接平方即可
		return y * y
	}
	//奇数的时候  还要多乘一个
	return y * y * x
}
