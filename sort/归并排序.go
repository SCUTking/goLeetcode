package sort

import "fmt"

func MergeSort(nums []int, left int, right int) []int {
	if left == right {
		//划分到长度只有一个元素的时候
		return []int{nums[left]}
	}

	mid := left + (right-left)/2
	//递归调用  获取的是已经排好序的左右子数组
	leftArr := MergeSort(nums, left, mid) //执行减去操作 可能会出现越界
	rightArr := MergeSort(nums, mid+1, right)

	i, j := 0, 0
	ans := make([]int, 0)
	//比较大小  将数字小的塞进去（升序）
	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i] <= rightArr[j] {
			ans = append(ans, leftArr[i])
			i++
		} else {
			ans = append(ans, rightArr[j])
			j++
		}
	}
	//这两个循环作兜底的工作  如果有没添加完成的数  就加到结果中
	//因为上面的循环是先放小的
	for i < len(leftArr) {
		ans = append(ans, leftArr[i])
		i++
	}
	for j < len(rightArr) {
		ans = append(ans, rightArr[j])
		j++
	}
	return ans
}

func MS(nums []int) {
	sort := MyMergeSort(nums, 0, len(nums)-1)
	fmt.Println(sort)
}

func MyMergeSort(nums []int, left, right int) []int {
	if left == right {
		return []int{nums[left]}
	}
	mid := left + (right-left)/2
	sLeft := MyMergeSort(nums, left, mid)
	sRight := MyMergeSort(nums, mid+1, right)

	sMerge := make([]int, 0)
	i, j := 0, 0

	for i < len(sLeft) && j < len(sRight) {
		if sLeft[i] < sRight[j] {
			sMerge = append(sMerge, sLeft[i])
			i++
		} else {
			sMerge = append(sMerge, sRight[j])
			j++
		}
	}

	//跳出循环后，说明有个数组放完了
	//如果有没添加完的数  就加到结果中
	for i < len(sLeft) {
		sMerge = append(sMerge, sLeft[i])
		i++
	}
	for j < len(sRight) {
		sMerge = append(sMerge, sRight[j])
		j++
	}

	return sMerge

}
