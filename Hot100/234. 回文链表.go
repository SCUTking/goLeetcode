package Hot100

func isPalindrome(head *ListNode) bool {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	for i := 0; i < len(arr)/2; i++ {
		if arr[i] == arr[len(arr)-i-1] {

		} else {
			return false
		}
	}
	return true
}
