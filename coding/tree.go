package coding

//2.重建二叉树
//输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
//例如，给出
//mid->Left->Right
//前序遍历 preorder = [3,9,20,15,7]
//Left->mid->Right
//中序遍历 inorder = [9,3,15,20,7]
//返回如下的二叉树：
//3
/// \
//9  20
///  \
//15   7

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rebuild(nums1 []int, nums2 []int) *TreeNode {
	var dfs func(nums1 []int, nums2 []int) *TreeNode
	dfs = func(nums1 []int, nums2 []int) *TreeNode {
		if len(nums1) == 0 {
			return nil
		}
		n := nums1[0]

		root := &TreeNode{Val: n}

		flag := -1
		for i := 0; i < len(nums2); i++ {
			if nums2[i] == n {
				flag = i
			}
		}
		root.Left = dfs(nums1[1:1+flag], nums2[:flag])
		root.Right = dfs(nums1[flag+1:], nums2[flag+1:])

		return root
	}

	return dfs(nums1, nums2)
}
