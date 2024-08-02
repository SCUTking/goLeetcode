package random2

func predictPartyVictory(senate string) string {
	chars := make([]int32, 0)
	for _, i := range senate {
		chars = append(chars, i)
	}

	res := ""
	isOver := false
	index := 0
	for !isOver {
		isOver = true
		index = index % len(chars)
		cur := chars[index]
		for i := index; i < len(chars)+index; i++ {
			temp := i % len(chars)
			if chars[temp] == cur {

			} else {

				index++
				index = index % len(chars)
				chars = append(chars[:temp], chars[temp+1:]...)
				isOver = false
				break
			}
		}
	}

	if chars[0] == 'R' {
		res = "Radiant"
	} else {
		res = "Dire"
	}

	return res
}
