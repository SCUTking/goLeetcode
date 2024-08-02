package Array

// 利用双指针
func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	i, j := 0, n-1
	for pos := n - 1; pos >= 0; pos-- {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			res[pos] = nums[i] * nums[i]
			i++
		} else {
			res[pos] = nums[j] * nums[j]
			j--
		}
	}
	return res
}
