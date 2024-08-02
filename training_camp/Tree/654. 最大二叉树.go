package Tree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index := getMaxIndex(nums)
	root := &TreeNode{Val: nums[index]}

	if index >= 0 {
		root.Left = constructMaximumBinaryTree(nums[:index])
	}
	root.Right = constructMaximumBinaryTree(nums[index+1:])

	return root

}

func getMaxIndex(nums []int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[max] < nums[i] {
			max = i
		}
	}
	return max
}
