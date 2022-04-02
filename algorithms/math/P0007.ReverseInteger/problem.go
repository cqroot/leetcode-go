package p0007reverseinteger

import "math"

func reverse(x int) int {
	flag := 1
	if x < 0 {
		x = 0 - x
		flag = -1
	}

	ans := 0
	for x != 0 {
		ans = ans*10 + x%10
		x = x / 10
	}

	ans = ans * flag
	if ans > math.MaxInt32 || ans < math.MinInt32 {
		ans = 0
	}

	return ans
}
