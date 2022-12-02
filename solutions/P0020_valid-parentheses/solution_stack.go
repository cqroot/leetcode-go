package solution

type stack struct {
	data   []rune
	length int
}

func (s *stack) push(c rune) {
	s.data[s.length] = c
	s.length = s.length + 1
}

func (s *stack) pop() {
	s.length = s.length - 1
	if s.length < 0 {
		s.length = 0
	}
}

func (s stack) top() rune {
	if s.length != 0 {
		return s.data[s.length-1]
	} else {
		return rune(0)
	}
}

// Stack
// Time complexity  : O(len(s))
// Space complexity : O(len(s))
func isValid(s string) bool {
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := stack{
		data:   make([]rune, len(s)),
		length: 0,
	}

	for _, c := range s {
		if stack.length == 0 {
			stack.push(c)
			continue
		}

		if val, ok := bracketMap[c]; ok && stack.top() == val {
			stack.pop()
		} else {
			stack.push(c)
		}
	}

	if stack.length == 0 {
		return true
	} else {
		return false
	}
}
