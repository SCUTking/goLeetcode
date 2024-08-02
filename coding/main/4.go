package main

import (
	"fmt"
)

func main() {
	n := 0
	k := 0

	nums := make([]int, 0)
	fmt.Scan(&n, &k)
	for n > 0 {
		n--
		var a int
		fmt.Scan(&a)
		nums = append(nums, a)
	}

	pre := make([]int, n+1)
	pre[0] = 1
	for i := 0; i < n+1; i++ {
		pre[i+1] = pre[i] * nums[i]
	}
	res := 0
	for i := 1; i < n; i++ {
		for j := i; j < n+1; j++ {
			i2 := pre[j] / pre[j-i]
			i3 := count0(pre[n] / i2)
			if i3 >= k {
				res++
			}
		}
	}
	fmt.Println(res)
}

func count0(num int) int {
	res := 0
	for num > 0 {
		i := num % 10
		if i == 0 {
			res++
		} else {
			break
		}
		num = num / 10
	}
	return res
}
