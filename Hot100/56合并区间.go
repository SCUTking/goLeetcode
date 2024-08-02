package Hot100

import "sort"

func merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	mergeArr := make([][]int, 0)
	for _, interval := range intervals {
		if len(mergeArr) > 0 {
			ints := mergeArr[len(mergeArr)-1]

			back := ints[1]

			front := interval[0]
			if back >= front {
				mergeArr = mergeArr[:len(mergeArr)-1]
				mergeArr = append(mergeArr, []int{ints[0], max(interval[1], back)})
			} else {
				mergeArr = append(mergeArr, interval)
			}
		} else {
			mergeArr = append(mergeArr, interval)
		}
	}
	return mergeArr
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
