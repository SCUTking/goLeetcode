package sort

import "fmt"

type student struct {
	age   int
	grade int
}

// 建立大根堆，用来升序排序
func HeapSortMulti(nums []student) {
	buildHeapMulti(nums)

	for i := 0; i < len(nums); i++ {
		lastIndex := len(nums) - i - 1
		nums[0], nums[lastIndex] = nums[lastIndex], nums[0]
		siftDownMulti(nums, lastIndex-1, 0)
	}
	fmt.Println(nums)

}

func buildHeapMulti(nums []student) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		siftDownMulti(nums, len(nums)-1, i)
	}
}

// 对切片nums的索引index进行从0到n向下渗透
func siftDownMulti(nums []student, n int, index int) {
	for {
		l := 2*index + 1
		r := 2*index + 2
		if l > n {
			//退出条件
			break
		}
		bigIndex := l //默认是左边的
		if r <= n && nums[bigIndex].age < nums[r].age {
			bigIndex = r
		}
		if nums[bigIndex].age < nums[index].age {
			//找到指定位置了
			break
		} else {
			//交换
			nums[bigIndex], nums[index] = nums[index], nums[bigIndex]
			index = bigIndex
		}
	}
}
