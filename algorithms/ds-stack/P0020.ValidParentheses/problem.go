package p0020validparentheses

func isValid(s string) bool {
	parenthesesMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := make([]rune, 0)

	for _, c := range s {
		if len(stack) == 0 || parenthesesMap[stack[len(stack)-1]] != c {
			stack = append(stack, c)
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
