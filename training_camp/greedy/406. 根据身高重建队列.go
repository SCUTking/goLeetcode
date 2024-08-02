package greedy

import "sort"

func reconstructQueue(people [][]int) (ans [][]int) {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	for _, person := range people {
		idx := person[1]
		//从person中间穿过
		ans = append(ans[:idx], append([][]int{person}, ans[idx:]...)...)
	}
	return
}
