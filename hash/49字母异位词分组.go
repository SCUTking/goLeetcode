package hash

import "sort"

func GroupAnagrams(strs []string) [][]string {

	var result [][]string

	m := make(map[[26]int][]string)
	s := strs[0]
	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})

	for _, i2 := range strs {
		ints := [26]int{}
		for _, i3 := range i2 {
			ints[i3-'a']++
		}
		m[ints] = append(m[ints], i2)
	}

	for _, strings := range m {
		result = append(result, strings)
	}

	return result

}
