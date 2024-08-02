package sort

import "fmt"

// InsertSort 时间复杂度：最坏情况下为O(N*N)，此时待排序列为逆序，或者说接近逆序
// 最好情况下为O(N)，此时待排序列为升序，或者说接近升序。
// 空间复杂度：O(1)
func InsertSort(arr []int) []int {
	for i, _ := range arr {

		start := i
		if i == 0 {
			continue

		}
		cur := arr[i]
		for cur < arr[start-1] {
			arr[start] = arr[start-1]
			start--
			if start == 0 {
				break
			}
		}
		arr[start] = cur
	}
	return arr

}

func InsertionSort(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		cur := nums[i]
		tempIndex := i - 1
		for tempIndex >= 0 && cur < nums[tempIndex] {
			nums[tempIndex+1] = nums[tempIndex] //不用进行元素交换，让元素往后移动一下就好
			tempIndex--
		}
		nums[tempIndex+1] = cur
	}
	fmt.Println(nums)
}
