package P0032_longest_valid_parentheses

func longestValidParentheses(s string) int {
	var res int = 0
	if len(s) <= 1 {
		return 0
	}

	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			continue
		}
		var countRight int = 1
		var maxLen int = 0
		for j := i - 1; j >= 0; j-- {
			if countRight < 0 {
				break
			}
			if s[j] == '(' {
				countRight -= 1
			} else {
				countRight += 1
			}
			if countRight == 0 {
				maxLen = i - j + 1
			}
		}
		if maxLen > res {
			res = maxLen
		}
	}
	return res
}
