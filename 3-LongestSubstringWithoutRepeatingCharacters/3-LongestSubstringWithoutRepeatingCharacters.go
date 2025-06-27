// Last updated: 6/27/2025, 4:38:45 PM
func lengthOfLongestSubstring(s string) int {
	charIndexMap := make(map[rune]int)
	maxLength := 0
	start := 0

	for end, char := range s {
		if lastIndex, found := charIndexMap[char]; found && lastIndex >= start {
			start = lastIndex + 1 
		}
		charIndexMap[char] = end
		maxLength = max(maxLength, end-start+1)
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}