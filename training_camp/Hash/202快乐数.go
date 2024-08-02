package Hash

func isHappy(n int) bool {
	occured := make(map[int]struct{})
	for n != 1 {
		n = square(n)
		_, exist := occured[n]
		if exist {
			return false
		} else {
			occured[n] = struct{}{}
		}
	}
	return true
}

func square(num int) int {
	res := 0
	for num != 0 {
		i := num % 10
		num = num / 10
		res += i * i
	}
	return res
}
