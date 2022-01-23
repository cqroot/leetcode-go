package p0003longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	maxLen := 0
	leftIdx := 0

	for idx, char := range s {
		if c, ok := m[char]; ok == true {
			if c >= leftIdx {
				if idx-leftIdx > maxLen {
					maxLen = idx - leftIdx
				}
				leftIdx = c + 1
			}
		}
		m[char] = idx
	}

	if len(s)-leftIdx > maxLen {
		maxLen = len(s) - leftIdx
	}
	return maxLen
}
