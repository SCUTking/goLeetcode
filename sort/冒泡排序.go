package sort

//时间复杂度：最坏情况：O(N^2)
//最好情况：O(N)
//空间复杂度：O(1)

// BubbleSort 冒泡排序：每次循环都将（从0开始将元素与相邻元素进行比对————冒泡的来由）
// ，将最大的数放到最后，进行N遍后数组就变成有序了
// 优化点： 如果一趟下来都没有交换，说明已经是有序的了，无需再进行循环进去
func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {

		flag := false
		for j := 1; j < n-i; j++ {
			if arr[j-1] > arr[j] {
				temp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
				flag = true
			}
		}
		if !flag {
			break
		}
	}

	return arr

}
