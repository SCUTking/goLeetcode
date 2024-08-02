package Tree

func sortedArrayToBST(nums []int) *TreeNode {
	//n := len(nums)
	//rootVal := nums[n/2]
	//root := &TreeNode{Val: rootVal}
	//
	//for i := 0; i < n; i++ {
	//	if nums[i] != rootVal {
	//		add(root, nums[i])
	//	}
	//}
	//return root

	return helper1(nums, 0, len(nums))
}

func helper1(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right-left)/2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper1(nums, left, mid-1)
	root.Right = helper1(nums, mid+1, right)
	return root
}

func add(root *TreeNode, val int) {
	var parent *TreeNode
	cur := root
	for cur != nil {
		if cur.Val > val {
			parent = cur
			cur = cur.Left
		} else if cur.Val < val {
			parent = cur
			cur = cur.Right
		}
	}
	t := &TreeNode{Val: val}
	if parent.Val < val {
		parent.Right = t
	} else {
		parent.Left = t
	}
}
