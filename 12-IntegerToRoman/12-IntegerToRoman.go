// Last updated: 6/27/2025, 4:40:31 PM
func intToRoman(num int) string {
	var romanSymbols = []struct {
		value  int
		symbol string
	}{
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

	result := ""
	for _, r := range romanSymbols {
		for num >= r.value {
			result += r.symbol
			num -= r.value
		}
	}

	return result
}
