package random2

import "math/bits"

func maxProduct(words []string) int {
	bitArr := make([]int, 0)
	for i := 0; i < len(words); i++ {
		s := words[i]
		temp := 0
		m := make(map[int32]struct{}, 0)
		for _, i2 := range s {
			i3 := i2 - 'a'
			_, exist := m[i3]
			if !exist {
				m[i3] = struct{}{}
				temp += 1 << i3
			}

		}
		bitArr = append(bitArr, temp)
	}

	maxMutli := 0
	for i, i2 := range bitArr {
		for i3, i4 := range bitArr {
			if i != i3 && i2&i4 == 0 {
				temp := len(words[i]) * bits.OnesCount(uint(i4))
				if temp > maxMutli {
					maxMutli = temp
				}
			}
		}
	}

	return maxMutli
}
