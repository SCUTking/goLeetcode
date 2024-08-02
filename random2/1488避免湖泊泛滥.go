package random2

import "sort"

func avoidFlood(rains []int) []int {
	lakeMap := make(map[int]int)
	numNoRainIndex := make([]int, 0)
	res := make([]int, len(rains))

	for i, rain := range rains {
		if rain == 0 {
			res[i] = 1
			numNoRainIndex = append(numNoRainIndex, i)
			//
		} else {
			res[i] = -1
			preIndex, exist := lakeMap[rain]
			if !exist {
				//不存在
				lakeMap[rain] = i
			} else {
				index := sort.SearchInts(numNoRainIndex, preIndex)
				if index == len(numNoRainIndex) {
					return []int{}
				}
				res[index+1] = rain
				lakeMap[rain] = i

				numNoRainIndex = append(numNoRainIndex[:index], numNoRainIndex[index+1:]...)
			}
		}
	}

	return res
}
