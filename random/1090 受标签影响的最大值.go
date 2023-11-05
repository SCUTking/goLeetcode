package random

import (
	"sort"
)

func LargestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	labelMap := make(map[int]int)
	curNum := 0

	indexArr := []int{}
	for i := 0; i < len(values); i++ {
		indexArr = append(indexArr, i)
		labelMap[values[i]] = 0
	}

	sort.Slice(indexArr, func(i, j int) bool {
		return values[indexArr[i]] > values[indexArr[j]]
	})
	ans := 0
	for _, i := range indexArr {
		value := values[i]
		label := labels[i]
		if labelMap[label]+1 > useLimit {
			continue
		}
		if curNum < numWanted {
			labelMap[label]++
			ans += value
			curNum++
		}
	}

	return ans

}
