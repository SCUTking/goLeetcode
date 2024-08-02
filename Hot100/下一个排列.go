package Hot100

func nextPermutation(nums []int) {
	n := len(nums)
	if n == 1 {
		return
	}

	left := -1
	for i := n - 1; i >= 1; i-- {
		if nums[i] > nums[i-1] {
			left = i - 1
			break
		}
	}

	if left == -1 {
		//直接倒序即可
		reverse(nums)
		return
	}
	right := 0
	for i := n - 1; i >= 1; i-- {
		if nums[i] > nums[left] {
			right = i
			break
		}
	}

	nums[left], nums[right] = nums[right], nums[left]

	reverse(nums[left+1:])

}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
