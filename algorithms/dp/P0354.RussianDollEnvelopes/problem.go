package P0354_RussianDollEnvelopes

import (
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	var dp []int = make([]int, len(envelopes))
	var res int = 0

	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0]
	})
	for i := 0; i < len(envelopes); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[i][0] > envelopes[j][0] && envelopes[i][1] > envelopes[j][1] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
