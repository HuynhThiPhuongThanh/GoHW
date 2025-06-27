// Last updated: 6/27/2025, 4:41:39 PM
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for _, s := range strs[1:] {
		for len(prefix) > 0 && len(s) < len(prefix) || s[:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]
		}
		if prefix == "" {
			return ""
		}
	}

	return prefix
}

