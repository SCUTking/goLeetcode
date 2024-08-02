package Hash

import "sort"

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	for _, v := range nums {
		m[v] = struct{}{}
	}
	sort.Ints(nums)
	if len(nums) == 0 {
		return 0
	}
	res := 1
	temp := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			temp++
			res = max(res, temp)
		} else if nums[i] == nums[i-1] {

		} else {
			temp = 1
		}
	}

	return res

}

// O(N^2)
func longestConsecutive1(nums []int) int {
	m := make(map[int]struct{})
	for _, v := range nums {
		m[v] = struct{}{}
	}

	if len(nums) == 0 {
		return 0
	}
	res := 1
	for i := 0; i < len(nums); i++ {
		_, ok := m[nums[i]]
		if ok {
			sum := 0
			temp := nums[i]
			for j := 0; j < len(nums); j++ {
				_, exist := m[temp+1]
				if exist {
					temp += 1
					sum++
				} else {
					break
				}
			}
			res = max(sum, res)
		}

	}

	return res

}
func longestConsecutive2(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}

	if len(nums) == 0 {
		return 0
	}
	res := 1
	for _, num := range nums {
		if !m[num-1] {
			sum := 0
			temp := num
			for m[temp] {
				sum++
				temp = temp + 1
			}
			res = max(res, sum)
		}
	}

	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
