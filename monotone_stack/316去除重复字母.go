package monotone_stack

func RemoveDuplicateLetters(s string) string {

	m := make(map[int32]int, 26)
	for _, i := range s {
		m[i]++
	}

	existStack := make(map[int32]int, 26)

	monostoneStack := []int32{}

	for _, each := range s {

		if existStack[each] == 1 {
			m[each]--
			continue
		}
		for len(monostoneStack) > 0 && monostoneStack[len(monostoneStack)-1] > each {
			if m[monostoneStack[len(monostoneStack)-1]] > 0 {
				existStack[monostoneStack[len(monostoneStack)-1]] = 0
				monostoneStack = monostoneStack[:len(monostoneStack)-1]
			} else {
				break
			}
		}
		//加入栈中
		monostoneStack = append(monostoneStack, each)
		m[each]--
		existStack[each] = 1
	}

	s2 := string(monostoneStack)
	return s2
}
