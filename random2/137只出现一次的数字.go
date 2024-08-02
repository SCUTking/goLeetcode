package random2

func singleNumber(nums []int) int {

	res := 0
	for i := 0; i < 32; i++ {
		temp := 0
		for _, num := range nums {
			temp += num >> i & 1
		}
		if temp%3 > 0 {
			res += 1 << i
		}

	}
	return res
}
