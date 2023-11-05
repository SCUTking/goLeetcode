package sliding_window

import (
	"reflect"
)

func FindLength(nums1 []int, nums2 []int) int {

	len1 := len(nums1)
	len2 := len(nums2)
	ans := 0
	for i := 0; i < len1; i++ {
		i2 := min(len2, len1-i)
		i3 := maxLen(nums1, nums2, i, 0, i2)
		ans = max(i3, ans)
	}

	for i := 0; i < len2; i++ {
		i2 := min(len1, len2-i)
		i3 := maxLen(nums1, nums2, 0, i, i2)
		ans = max(i3, ans)
	}

	return ans
	//str1, str2 := "", ""
	//for i := 0; i < len(nums1); i++ {
	//	str1 += strconv.Itoa(nums1[i])
	//}
	//for i := 0; i < len(nums2); i++ {
	//	str2 += strconv.Itoa(nums2[i])
	//}
	//
	//left, right := 0, 0
	//max := 0
	//cur := []int{}
	//for right < len(nums1) {
	//	cur = append(cur, nums1[right])
	//	right++
	//
	//	for !Check(cur, nums2) {
	//		cur = cur[1:]
	//		left++
	//	}
	//	if max < right-left {
	//		max = right - left
	//	}
	//}
	//
	//return max
}

func maxLen(num1, num2 []int, a, b, len int) int {
	maxlen1 := 0
	for i := 0; i < len; i++ {
		if num1[a+i] == num2[b+i] {
			maxlen1++
		} else {
			break
		}
	}

	return maxlen1

}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}

}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}

}

func Check(num1 []int, num2 []int) bool {
	length := len(num1)
	for i := 0; i <= len(num2)-length; i++ {
		numTemp := num2[i : i+length]
		if reflect.DeepEqual(numTemp, num1) {
			return true
		}
	}
	return false

}
