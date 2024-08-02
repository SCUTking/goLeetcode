package Hash

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {

	m := make(map[int]int)
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			m[num1+num2]++
		}
	}
	res := 0
	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			v, exist := m[-(num3 + num4)]
			if exist {
				res += v
			}
		}
	}

	return res

}
