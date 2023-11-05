package random1

import "sort"

func WiggleSort(nums []int) {
	n := len(nums)
	temp := append([]int{}, nums...)

	sort.Ints(temp)
	var getLeast func() int
	getLeast = func() int {
		least := temp[0]
		temp = temp[1:]
		return least
	}

	var getMost func() int
	getMost = func() int {
		least := temp[len(temp)-1]
		temp = temp[:len(temp)-1]
		return least
	}

	ans := []int{}
	for len(ans) != n {
		if len(ans)&1 == 0 {
			nums[len(ans)] = getLeast()
		} else {
			nums[len(ans)] = getMost()
		}
		ans = append(ans, 1)
	}

}
