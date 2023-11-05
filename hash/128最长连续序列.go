package hash

func LongestConsecutive(nums []int) int {

	m := make(map[int]string, len(nums))

	for _, num := range nums {
		m[num] = "1"
	}

	var length int
	length = 1
	var res int
	res = 1
	for _, num := range nums {
		if m[num-1] == "1" {
			//存在hash表中
			continue
		}
		for length = 1; ; length++ {
			if m[num+length] != "1" {
				break
			}
		}
		if res < length {
			res = length
		}

	}
	return res

}
