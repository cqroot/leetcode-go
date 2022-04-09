package P0368_largest_divisible_subset

import (
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	var dp [][]int = make([][]int, len(nums))
	var maxLen int = 0
	var resIdx int = 0

	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})

	dp[0] = []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		dp[i] = []int{nums[i]}
		for j := 0; j < i; j++ {
			if nums[j]%nums[i] == 0 || nums[i]%nums[j] == 0 {
				if len(dp[j])+1 > len(dp[i]) {
					dp[i] = make([]int, len(dp[j])+1)
					copy(dp[i], append(dp[j], nums[i]))
					if len(dp[i]) > maxLen {
						maxLen = len(dp[i])
						resIdx = i
					}
				}
			}
		}
	}
	return dp[resIdx]
}
