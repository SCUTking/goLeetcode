package sort

func MergeSort(nums []int, left int, right int) []int {
	if left == right {
		return append([]int{}, nums[left])
	}

	mid := left + (right-left)/2
	//递归调用  获取的是已经排好序的左右子数组
	leftArr := MergeSort(nums, left, mid) //执行减去操作 可能会出现越界
	rightArr := MergeSort(nums, mid+1, right)

	i, j := 0, 0
	ans := make([]int, 0)
	//比较大小  将数字小的塞进去
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
