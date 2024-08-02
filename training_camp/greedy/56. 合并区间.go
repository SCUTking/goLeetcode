package greedy

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > right {
			res = append(res, []int{left, right})
			left = intervals[i][0]
			right = intervals[i][1]
			continue
		}
		right = max(intervals[i][1], right)
	}
	return res

}
