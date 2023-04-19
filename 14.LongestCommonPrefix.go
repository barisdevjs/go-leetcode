package main

func LongestCommonPrefix(strs []string) string {
	prefix := ""
	if len(strs) == 0 { // handle empty input case
		return prefix
	}
	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != char {
				return prefix
			}
		}
		prefix += string(char)
	}
	return prefix
}

func LongestCommonPrefix2(strs []string) string {
	prefix := ""
	if len(strs) == 0 { // handle empty input case
		return prefix
	}
	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != char {
				return prefix
			}
		}
		prefix += string(char)
	}
	return prefix
}
