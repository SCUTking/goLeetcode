package sort

// 时间复杂度平均：O(N^1.3)
// 空间复杂度：O(1)
func ShellSort(arr []int) []int {
	//希尔排序——插入排序的优化版，先设置gap为数组长度，一直减半的对距离为gap的子数组进行插入排序
	gap := len(arr)
	for gap > 1 {
		gap = gap / 2
		for i := 0; i < len(arr)-1; i++ {
			cur := arr[i]
			j := i - gap
			for ; j >= 0; j -= gap {
				if cur < arr[j] {
					arr[j+gap] = arr[j]
				} else {
					break
				}
			}
			arr[j+gap] = cur
		}
	}

	return arr

}
