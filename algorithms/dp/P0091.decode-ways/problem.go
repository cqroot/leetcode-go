package P0091_decode_ways

import (
	"strconv"
)

func numDecodings(s string) int {
	if s[0] == '0' || len(s) < 1 {
		return 0
	}
	var dp []int = make([]int, len(s))
	var res int = 1
	dp[0] = 1
	for i := 1; i < len(s); i++ {
		if s[i] != '0' {
			dp[i] = dp[i-1]
		}
		if s[i-1] != '0' {
			num, _ := strconv.Atoi(s[i-1 : i+1])
			if num <= 26 {
				if i == 1 {
					dp[i] += 1
				} else {
					dp[i] += dp[i-2]
				}
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return dp[len(s)-1]
}
