package sort

import "fmt"

// 建立大根堆，用来升序排序
func HeapSort(nums []int) {
	buildHeap(nums)

	for i := 0; i < len(nums); i++ {
		lastIndex := len(nums) - i - 1
		nums[0], nums[lastIndex] = nums[lastIndex], nums[0]
		siftDown(nums, lastIndex-1, 0)
	}
	fmt.Println(nums)

}

func buildHeap(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		siftDown(nums, len(nums)-1, i)
	}
}

// 对切片nums的索引index进行从0到n向下渗透
func siftDown(nums []int, n int, index int) {
	for {
		l := 2*index + 1
		r := 2*index + 2
		if l > n {
			//退出条件
			break
		}
		bigIndex := l //默认是左边的
		if r <= n && nums[bigIndex] < nums[r] {
			bigIndex = r
		}
		if nums[bigIndex] < nums[index] {
			//找到指定位置了
			break
		} else {
			//交换
			nums[bigIndex], nums[index] = nums[index], nums[bigIndex]
			index = bigIndex
		}
	}
}
