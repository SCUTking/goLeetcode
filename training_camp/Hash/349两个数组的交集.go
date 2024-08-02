package Hash

func intersection(nums1 []int, nums2 []int) []int {

	m := make(map[int]struct{})
	for _, i := range nums1 {
		m[i] = struct{}{}
	}
	res := make([]int, 0)

	for _, i := range nums2 {
		_, exist := m[i]
		if exist {
			res = append(res, i)
		}
	}
	return res
}
