package stack_and_queue

import (
	"container/heap"
	"math"
	"sort"
)

func maxSlidingWindow(nums []int, k int) []int {

	maxNum := make([]int, len(nums))
	temp := math.MinInt32
	for i, num := range nums {
		if num > temp {
			temp = num
		}
		maxNum[i] = temp
	}
	res := make([]int, len(nums)-k+1)

	for i := 0; i <= len(nums)-k; i++ {
		res[0] = max(maxNum[i], maxNum[i+k-1])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxSlidingWindow1(nums []int, k int) []int {
	maxNum, maxIndex := math.MinInt32, 0
	for i := 0; i < k; i++ {
		if maxNum <= nums[i] {
			maxNum = nums[i]
			maxIndex = i
		}
	}
	res := make([]int, 0)
	res = append(res, maxNum)
	for i := k; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
			maxIndex = i
		}

		if i-k == maxIndex {
			for j := i - k + 1; j <= i; j++ {
				if maxNum <= nums[j] {
					maxNum = nums[j]
					maxIndex = j
				}
			}
		}

		res = append(res, maxNum)

	}

	return res
}

var a []int

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {
	return a[h.IntSlice[i]] < a[h.IntSlice[j]]
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	h2 := *h
	res := h.IntSlice[len(h2.IntSlice)-1]
	h2.IntSlice = h.IntSlice[:len(h2.IntSlice)-1]
	return res
}

func maxSlidingWindow2(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}

	heap.Init(q)
	n := len(nums)
	ans := make([]int, 1, n-k-1)
	ans[0] = nums[q.IntSlice[0]]

	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}

		ans = append(ans, nums[q.IntSlice[0]])
	}

	return ans
}
