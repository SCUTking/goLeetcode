package random

func PermuteUnique(nums []int) [][]int {
	n := len(nums)
	i := [][]int{}
	i = append(i, []int{nums[0]})
	m := make(map[int][][]int)
	m[0] = i
	for i := 1; i < n; i++ {
		cur := nums[i]
		for _, ints := range m[i-1] {
			for i2, _ := range ints {
				newInts := make([]int, i2)
				copy(newInts, ints[:i2])
				newInts = append(newInts, cur)
				newInts = append(newInts, ints[i2:]...)
				m[i] = append(m[i], newInts)
			}
			m[i] = append(m[i], append(ints, cur))
		}

	}
	return m[n-1]

}
