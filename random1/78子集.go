package random1

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
		}

		preNums = append(preNums, nums[length])
		//加上当前的元素
		backTrack(length + 1)
		preNums = preNums[:len(preNums)-1]
		//不加上当前的元素
		backTrack(length + 1)
	}

	return ans
}
