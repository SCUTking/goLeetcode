package greedy

import "sort"

func findContentChildren(g []int, s []int) int {
	sum := 0
	sort.Ints(s)
	for _, num := range g {
		index := sort.SearchInts(s, num)
		if index != len(s) {
			sum++
			s = append(s[:index], s[index+1:]...)
		}
	}
	return sum

}
