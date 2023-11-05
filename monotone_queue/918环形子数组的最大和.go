package monotone_queue

func MaxSubarraySumCircular(nums []int) int {
	type pair struct {
		Index int
		Val   int
	}
	pre, res := nums[0], nums[0]

	queue := []pair{{Index: 0, Val: nums[0]}}

	n := len(nums)
	for i := 0; i < 2*n; i++ {
		for len(queue) > 0 && queue[0].Index < i-n {
			queue = queue[1:]
		}
		pre += nums[i%n]
		res = max(res, pre-queue[0].Val)

		for len(queue) > 0 && queue[len(queue)-1].Val > pre {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, pair{Index: i, Val: pre})

	}

	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
