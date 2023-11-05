package interval

import "strconv"

func summaryRanges(nums []int) []string {

	//if len(nums) == 0 {
	//	return []string{}
	//}
	//ans := []string{}
	//pre := 0
	//for i := 1; i < len(nums); i++ {
	//	if nums[i]-nums[i-1] > 1 {
	//		//说明要
	//		if i-pre == 1 {
	//			ans = append(ans, strconv.Itoa(nums[pre]))
	//		} else {
	//			var tmp string = strconv.Itoa(nums[pre]) + "->" + strconv.Itoa(nums[i-1])
	//			ans = append(ans, tmp)
	//		}
	//		pre = i
	//	}
	//
	//	if i == len(nums)-1 {
	//		var tmp string = strconv.Itoa(nums[pre]) + "->" + strconv.Itoa(nums[i])
	//		ans = append(ans, tmp)
	//	}
	//
	//}
	//
	//return ans
	ans := []string{}
	i, n := 0, len(nums)
	for i < n {
		start := i
		for i < n && nums[i]+1 == nums[i+1] {
			i += 1
		}
		var s string = strconv.Itoa(nums[i])
		if i-start > 0 {
			s += "->" + strconv.Itoa(nums[start])
		}
		ans = append(ans, s)
	}

	return ans
}
