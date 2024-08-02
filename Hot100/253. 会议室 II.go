package Hot100

import (
	"sort"
)

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	minRArr := make([]int, 1)
	minRArr[0] = intervals[0][1]
	for i := 0; i < len(intervals); i++ {
		curL := intervals[i][0]
		curR := intervals[i][1]
		isvalidIndex := -1
		for j := 0; j < len(minRArr); j++ {
			if minRArr[j] < curL {
				isvalidIndex = i
			}
		}

		if isvalidIndex == -1 {
			minRArr = append(minRArr, curR)
		} else {
			minRArr[isvalidIndex] = curR
		}
	}

	return len(minRArr)
}
