package random1

import (
	"fmt"
)

func Subsets(nums []int) [][]int {

	//ans := make([][]int, 0)
	//
	//var getSubSet func(length int)
	//getSubSet = func(length int) {
	//	left, rigth := 0, 0
	//	temp := []int{}
	//	for rigth < len(nums) {
	//		temp = append(temp, nums[rigth])
	//		rigth++
	//		for rigth-left > length {
	//			temp = temp[1:]
	//			left++
	//		}
	//		if rigth-left == length {
	//			ans = append(ans, temp)
	//		}
	//	}
	//}
	//ans = append(ans, []int{})
	//for i := 1; i < len(nums); i++ {
	//	getSubSet(i)
	//}
	//return ans

	ans := make([][]int, 0)
	preNums := make([]int, 0)
	var backTrack func(length int)
	backTrack = func(length int) {
		if len(nums) == length {
			ints := make([]int, len(preNums))
			copy(ints, preNums)
			ans = append(ans, ints)
			return
		}

		preNums = append(preNums, nums[length])
		//加上当前的元素
		backTrack(length + 1)
		preNums = preNums[:len(preNums)-1]
		//不加上当前的元素
		backTrack(length + 1)
	}
	backTrack(0)
	// 创建一个空的 map，用于存储切片的字符串表示以判断是否重复
	uniqueMap := make(map[string]bool)

	// 创建一个空的切片，用于存储去重后的切片
	uniqueSliceOfSlice := [][]int{}

	// 遍历切片的切片
	for _, slice := range ans {
		// 将切片转换为字符串表示
		sliceString := fmt.Sprint(slice)

		// 如果切片的字符串表示未存在于 map 中，则将其添加到 map 和结果切片中
		if _, ok := uniqueMap[sliceString]; !ok {
			uniqueMap[sliceString] = true
			uniqueSliceOfSlice = append(uniqueSliceOfSlice, slice)
		}
	}
	return uniqueSliceOfSlice
}
