package greedy

import "sort"

func findMinArrowShots(points [][]int) int {
	ans := 0
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	maxRight := points[0][1]
	for _, point := range points {
		if point[0] > maxRight {
			maxRight = point[1]
			ans++
		}

	}

	return ans
}
