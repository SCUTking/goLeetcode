package hash

func isHappy(n int) bool {
	m := make(map[int]bool) //用于判断是否出现循环

	for true {
		cur := square(n)
		if cur == 1 {
			return true
		}

		_, ok := m[cur]
		if ok {
			return false
		} else {
			m[cur] = true
		}
	}
	return false
}
func square(n int) int {
	ans := 0
	for n > 0 {
		cur := n % 10
		ans += cur * cur
		n /= 10
	}
	return ans
}
