package string

func FindRepeatedDnaSequences(s string) (res []string) {

	m := make(map[string]bool)
	m1 := make(map[string]bool)

	runes := []rune(s)
	length := len(runes)

	if length < 10 {
		return []string{}
	}
	for i := 0; i <= length-10; i++ {
		s := runes[i : i+10]
		s2 := string(s)

		_, exist := m[s2]
		if exist {
			m1[s2] = true

		}
		m[s2] = true
	}

	for k, _ := range m1 {
		res = append(res, k)
	}
	return
}
