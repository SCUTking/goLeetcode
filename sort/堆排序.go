package sort

func HeapSort(arr []int) []int {
	var adjust func(arr []int, cur int, length int)
	adjust = func(arr []int, cur int, length int) {
		temp := arr[cur]
		for k := cur*2 + 1; k < length; k = k*2 + 1 {
			if k+1 < length && arr[k] < arr[k+1] {
				k += 1
			}
			if temp < arr[k] {
				arr[cur] = arr[k]
				cur = k
			} else {
				break
			}
		}
		arr[cur] = temp //标记交换的元素，用于后续放最开始的父节点
	}

	//建堆
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		adjust(arr, i, n)
	}

	//进行堆排序
	for i := 0; i < n; i++ {
		swap(arr, 0, n-1-i)
		adjust(arr, 0, n-2-i)
	}
	return arr
}

func swap(arr []int, source int, target int) {
	temp := arr[source]
	arr[source] = arr[target]
	arr[target] = temp

}
