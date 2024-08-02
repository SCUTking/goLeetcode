package main

func HeapSort(arr []int) []int {
	var adjust func(arr []int, cur int, length int)
	adjust = func(arr []int, cur int, length int) {
		temp := arr[cur]
		for k := cur*2 + 1; k < length; k = k*2 + 1 {
			//找到子节点最大的一个
			if k+1 < length && arr[k] < arr[k+1] {
				k += 1
			}
			//如果最大子节点大于cur当前节点
			if temp < arr[k] {
				arr[cur] = arr[k] //交换节点
				cur = k           //交换后，将cur变为交换后的索引下标
			} else {
				break
			}
		}
		arr[cur] = temp //找到了应该放的位置 标记交换的元素，用于后续放最开始的父节点
	}

	//建堆
	n := len(arr)
	//从最后一个父节点开始
	for i := n/2 - 1; i >= 0; i-- {
		adjust(arr, i, n)
	}

	//进行堆排序
	for i := 0; i < n; i++ {
		//2.将堆顶元素与末尾元素进行交换，使末尾元素最大
		swap(arr, 0, n-1-i)
		//注意要缩小范围，因为最后那些已经排好序的了
		adjust(arr, 0, n-2-i)
	}
	return arr
}

func swap(arr []int, source int, target int) {
	temp := arr[source]
	arr[source] = arr[target]
	arr[target] = temp

}
