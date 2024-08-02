package Array

func removeElement(nums []int, val int) int {

	sum := 0
	for _, num := range nums {
		if num == val {
			sum++
		}
	}
	count := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == val {
			for j := n - count - 1; j >= 0; j-- {
				if nums[j] != val {
					nums[j], nums[i] = nums[i], nums[j]
					count++
					break
				}
			}
			if i == n-sum {
				break
			}
		}
	}

	return n - sum
}
func removeElement1(nums []int, val int) int {

	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
			right++
		} else {
			right++
		}
	}
	return left + 1
}
