package Hot100

import (
	"reflect"
	"sort"
)

// 超时了
func threeSum(nums []int) [][]int {

	sort.Ints(nums)
	ans := make([][]int, 0)

	var backFun func(count int, slice []int)
	backFun = func(count int, slice []int) {
		if len(slice) == 3 || count == len(nums) {
			if len(slice) == 3 {
				if slice[0]+slice[1]+slice[2] == 0 {
					//判断重复
					for _, t := range ans {
						equal := reflect.DeepEqual(t, slice)
						if equal {
							return
						}
					}
					temp := make([]int, 3)
					copy(temp, slice)
					ans = append(ans, temp)
					return
				} else {
					return
				}
			} else {
				return
			}

		}

		//不加这个
		backFun(count+1, slice)

		slice = append(slice, nums[count])

		backFun(count+1, slice)

	}
	backFun(0, make([]int, 0))

	return ans

}
func threeSum1(nums []int) [][]int {
	ans := make([][]int, 0)
	n := len(nums)
	sort.Ints(nums)
	for first := 0; first < len(nums); first++ {

		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		if nums[first] > 0 {
			break
		}

		third := n - 1
		target := -1 * nums[first]

		for second := first + 1; second < n; second++ {
			if second > 0 && nums[second] == nums[second-1] {
				continue
			}

			for second < third && nums[third]+nums[second] > target {
				third--
			}

			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}

			if nums[third]+nums[second] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}

	}

	return ans
}
