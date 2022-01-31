package p0014longestcommonprefix

func longestCommonPrefix(strs []string) string {
	commonPrefix := strs[0]
	for i := 1; i < len(strs); i++ {
		currentStr := strs[i]
		if len(currentStr) < len(commonPrefix) {
			commonPrefix = commonPrefix[0:len(currentStr)]
		}
		for j := 0; j < len(commonPrefix); j++ {
			if commonPrefix[j] != currentStr[j] {
				commonPrefix = commonPrefix[0:j]
				break
			}
		}
	}
	return commonPrefix
}
