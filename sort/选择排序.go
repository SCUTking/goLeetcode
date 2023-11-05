package sort

// 每次遍历，记录区间的最大值下标,然后与数组最后一个数交换
// 优化：一次遍历可以获取到最大值或者最小值

// 时间复杂度平均：O(N^1.3)
// 空间复杂度：O(1)
func SelectSort(arr []int) []int {
	n := len(arr)
	begin, end := 0, n-1
	for begin < end {
		minI, maxI := begin, begin

		for i := begin; i <= end; i++ {
			if arr[i] > arr[maxI] {
				maxI = i
			}
			if arr[i] < arr[minI] {
				minI = i
			}
		}

		temp := arr[minI]
		arr[minI] = arr[begin]
		arr[begin] = temp

		//如果最大值再第一个，但是第一个刚刚给人换走了，所以要记录换走之后的东西
		if begin == maxI {
			maxI = minI
		}

		temp1 := arr[maxI]
		arr[maxI] = arr[end]
		arr[end] = temp1

		begin++
		end--
	}

	return arr

}
