package monotone_queue

func LongestSubarray(nums []int, limit int) (ans int) {
	//queue := []int{nums[0]}
	//
	//res := 1
	//n := len(nums)
	//for i := 1; i < n; i++ {
	//
	//	abs := math.Abs(float64(nums[i] - queue[len(queue)-1]))
	//	if abs > (float64(limit)) {
	//		queue = []int{}
	//		queue = append(queue, nums[i])
	//		continue
	//	}
	//
	//	index := -1
	//	for j := i - len(queue); j < i; j++ {
	//		if nums[j] < (nums[i] - limit) {
	//			index = j
	//		}
	//	}
	//	if index != -1 {
	//		queue = []int{}
	//		ints := nums[index+1 : i]
	//		sort.Ints(ints)
	//		for _, each := range ints {
	//			queue = append(queue, each)
	//		}
	//
	//	}
	//
	//	//for len(queue) > 0 && math.Abs(float64(nums[i]-queue[0])) > float64(limit) {
	//	//	queue = queue[1:]
	//	//}
	//
	//	temp := []int{}
	//	for len(queue) > 0 && queue[len(queue)-1] > nums[i] {
	//		temp = append(temp, queue[len(queue)-1])
	//		queue = queue[:len(queue)-1]
	//	}
	//	queue = append(queue, nums[i])
	//	for k := len(temp) - 1; k >= 0; k-- {
	//		queue = append(queue, temp[k])
	//	}
	//
	//	res = max(res, len(queue))
	//
	//}
	//return res

	var minQ, maxQ []int
	left := 0
	res := 0
	for right, num := range nums {
		for len(minQ) > 0 && minQ[len(minQ)-1] < num {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, num)
		for len(maxQ) > 0 && maxQ[len(maxQ)-1] > num {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, num)

		for len(minQ) > 0 && len(maxQ) > 0 && maxQ[0]-minQ[0] > limit {
			if nums[left] == minQ[0] {
				minQ = minQ[1:]
			}
			if nums[left] == maxQ[0] {
				maxQ = maxQ[1:]
			}
			left++
		}

		res = max(res, right-left+1)
	}
	return res

}
