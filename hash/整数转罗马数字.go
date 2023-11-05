package hash

type simpleMap struct {
	value  int
	symbol string
}

func IntToRoman(num int) string {
	//有序的数组
	var valueSymbols = []simpleMap{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result string

	for num > 0 {
		for _, smap := range valueSymbols {
			if num >= smap.value {
				result += smap.symbol
				num -= smap.value
				break
			} else {
				continue
			}
		}
	}

	return result
}
