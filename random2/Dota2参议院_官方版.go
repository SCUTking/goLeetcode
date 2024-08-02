package random2

func predictPartyVictory1(senate string) string {

	RIndex := make([]int, 0)
	DIndex := make([]int, 0)
	for index, i := range senate {
		if i == 'R' {
			RIndex = append(RIndex, index)
		} else {
			DIndex = append(DIndex, index)
		}
	}

	for len(RIndex) > 0 && len(DIndex) > 0 {
		if RIndex[0] > DIndex[0] {
			DIndex = append(DIndex, DIndex[0]+len(senate))
		} else {
			RIndex = append(RIndex, RIndex[0]+len(senate))
		}
		DIndex = DIndex[1:]
		RIndex = RIndex[1:]
	}

	if len(RIndex) > 0 {
		return "Radiant"
	}
	return "Dire"

}
