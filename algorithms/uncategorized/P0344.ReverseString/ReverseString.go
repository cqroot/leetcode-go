package P0344_ReverseString

func reverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		t := s[left]
		s[left] = s[right]
		s[right] = t
		left++
		right--
	}
}
