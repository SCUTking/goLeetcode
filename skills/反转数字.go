package skills

// 每次获取原数的个位，然后将res结果乘以10，再加上个位，然后将原数的个位去掉
// 比如
// 123
// res=0   获取个位 3    res=0*10+3  12
// res=3   获取个位 2    res=3*10+2  1
// res=32   获取个位 1    res=32*10+1
func Rev(num int) int {
	rev := 0
	for ; num > 0; num /= 10 {
		rev = rev*10 + num%10
	}
	return rev
}
