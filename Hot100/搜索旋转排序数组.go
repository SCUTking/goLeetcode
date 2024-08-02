package Hot100

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	if left == right && nums[left] == target {
		return 1
	}
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] < nums[mid] {
			if target > nums[mid] || target < nums[left] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		if nums[mid] < nums[right] {
			if target < nums[mid] || target > nums[right] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}
