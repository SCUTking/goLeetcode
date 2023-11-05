package main

import "goLeetcode/sliding_window"

func main() {
	//sliding_window.LongestSubstring("ababacb", 3)
	//sliding_window.FindAnagrams("cbaebabacd", "abc")
	//sliding_window.CharacterReplacement("AABABBA", 1)
	nums := []int{10, 5, 2, 6}
	nums1 := []int{10, 5, 2, 6, 3}
	//sliding_window.FindClosestElements(nums, 4, 3)
	//

	//sliding_window.Check(nums, nums1)
	sliding_window.FindLength(nums, nums1)
	//sliding_window.NumSubarrayProductLessThanK(nums, 100)
}
