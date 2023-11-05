package race

import "sort"

func MinGroupsForValidAssignment(nums []int) int {

	m := make(map[int][]int)

	for i := 0; i < len(nums); i++ {
		_, ok := m[nums[i]]
		if ok {
			m[nums[i]] = append(m[nums[i]], i)
		} else {
			m[nums[i]] = []int{i}
		}
	}

	ints := []int{}
	for _, i2 := range m {
		ints = append(ints, len(i2))
	}

	sort.Ints(ints)
	min := ints[0]

	for ints[len(ints)-1]-min > 1 {
		value := ints[len(ints)-1]
		ints = ints[:len(ints)-1]

		minL := min - 1
		minR := min + 1
		if value%min == 0 && min != 1 {
			num := value / min
			for i := 0; i < num; i++ {
				ints = append(ints, min)
			}
		} else if value%minR == 0 && minR != 0 {
			num := value / minR
			for i := 0; i < num; i++ {
				ints = append(ints, minR)
			}
		} else if minL != 0 && minL != 1 && value%minL == 0 {
			num := value / minL
			for i := 0; i < num; i++ {
				ints = append(ints, minL)
			}
		} else if value&1 == 0 {
			//偶数
			ints = append(ints, value/2)
			ints = append(ints, value/2)
		} else {
			ints = append(ints, value/2)
			ints = append(ints, value/2+1)
		}
		sort.Ints(ints)
		min = ints[0]
	}
	return len(ints)
}
