package sliding_window

import "sort"

func FindClosestElements(arr []int, k int, x int) []int {
	if x >= arr[len(arr)-1] {
		return arr[len(arr)-k:]
	}
	if x <= arr[0] {
		return arr[:k]
	}

	firstIndex := 0
	for i, e := range arr {
		if e >= x {
			firstIndex = i
			break
		}
	}

	res := []int{}
	left, right := firstIndex, firstIndex
	leftArr := arr[:left]
	rightArr := arr[right:]
	for k > 0 {
		if len(leftArr) > 0 && len(rightArr) > 0 {
			abs1 := x - leftArr[len(leftArr)-1]
			abs2 := rightArr[0] - x
			if abs2 >= abs1 {
				res = append(res, leftArr[len(leftArr)-1])
				leftArr = leftArr[:len(leftArr)-1]
			} else {
				res = append(res, rightArr[0])
				rightArr = rightArr[1:]
			}
		} else if len(leftArr) > 0 {
			res = append(res, leftArr[len(leftArr)-1])
			leftArr = leftArr[:len(leftArr)-1]
		} else if len(rightArr) > 0 {
			res = append(res, rightArr[0])
			rightArr = rightArr[1:]
		}
		k--
	}

	sort.Ints(res)
	return res

}
