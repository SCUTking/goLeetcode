package backtrack

func Permute(nums []int) [][]int {
	n := len(nums)
	res := [][]int{}
	vis := make(map[int]bool)
	var backtracking func(index int, pre []int)
	backtracking = func(index int, pre []int) {
		if index == n {
			temp := make([]int, n)
			copy(temp, pre)
			res = append(res, temp)
			return
		}

		for i := 0; i < n; i++ {
			_, exist := vis[nums[i]]
			if !exist {
				pre = append(pre, nums[i])
				vis[nums[i]] = true
				backtracking(index+1, pre)
				pre = pre[:len(pre)-1]
				delete(vis, nums[i])
			}

		}
	}

	backtracking(0, []int{})

	return res
}
