package Hot100

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	res := make([]int, 0)
	m := make(map[int]bool, n)
	for i := 0; i < n; i++ {
		m[nums[i]] = true
	}
	for i := 1; i <= n; i++ {
		if m[i] {

		} else {
			res = append(res, i)
		}
	}
	return res
}
