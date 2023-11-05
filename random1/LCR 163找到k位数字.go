package random1

import (
	"math"
	"strconv"
)

func FindKthNumber(k int) int {
	nums := []int{0, 9}
	for i := 2; i <= 10; i++ {
		temp := 9 * (int(math.Pow10(i - 1))) * (i)
		nums = append(nums, nums[i-1]+temp)
	}

	index := 0
	for i, each := range nums {
		if k <= each {
			index = i
			break
		}
	}

	// 开始的那个数
	num := nums[index-1] + 1
	gap := k - num

	indexNum := (int(math.Pow10(index - 1))) + gap/(index)
	modNum := gap % (index)

	s := strconv.Itoa(indexNum)

	ans, _ := strconv.Atoi(string(s[modNum]))

	return ans

}
