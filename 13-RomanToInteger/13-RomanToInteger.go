// Last updated: 6/27/2025, 4:40:30 PM
func romanToInt(s string) int {
	romanValues := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	n := len(s)

	for i := 0; i < n; i++ {
		value := romanValues[rune(s[i])]
		if i < n-1 && romanValues[rune(s[i+1])] > value {
			total -= value 
		} else {
			total += value
		}
	}
	return total
}