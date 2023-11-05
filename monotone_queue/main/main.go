package main

import "goLeetcode/monotone_queue"

func main() {
	nums := []int{1, 5, 6, 7, 8, 10, 6, 5, 6}
	//monotone_queue.MaxSubarraySumCircular(nums)
	monotone_queue.LongestSubarray(nums, 4)
}
