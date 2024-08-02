package greedy

import (
	"math"
	"strconv"
)

func monotoneIncreasingDigits(n int) int {
	count := 0
	for i := 0; i < 9; i++ {
		if n/int(math.Pow10(i)) == 0 {
			count = i + 1
			break
		}
	}
	pre := math.MinInt32
	res := 0
	for i := count; i > 0; i-- {
		cur := n / int(math.Pow10(count-1))
		n -= int(math.Pow10(count - 1))
		if i != count {
			if cur < pre {
				res += int(math.Pow10(count - 1))
			} else {
				res += pre * int(math.Pow10(count))
			}
		}
		pre = cur
	}

	return res
}
func monotoneIncreasingDigits1(n int) int {

	s := []byte(strconv.Itoa(n))
	i := 1
	l := len(s)
	for i < l && s[i-1] <= s[i] {
		i++
	}
	if i == l {
		return n
	}

	for i > 0 && s[i] > s[i-1] {
		s[i-1]--
		i--
	}

	for j := i + 1; j < l; j++ {
		s[j] = '9'
	}
	atoi, _ := strconv.Atoi(string(s))
	return atoi

}
