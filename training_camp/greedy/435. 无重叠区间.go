package greedy

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[i][1]
	})
	res := 0
	maxRight := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < maxRight {
			res++
		} else {
			maxRight = intervals[i][1]
		}
	}
	return res
}
