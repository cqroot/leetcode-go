package P0673_NumberOfLongestIncreasingSubsequence

func findNumberOfLIS(nums []int) int {
	var dp []int = make([]int, len(nums))
	var cnt []int = make([]int, len(nums))
	var maxLen int = 1
	var res int
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		cnt[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					cnt[i] = cnt[j]
					dp[i] = dp[j] + 1
				} else if dp[j]+1 == dp[i] {
					cnt[i] += cnt[j]
				}
			}
		}
		if dp[i] > maxLen {
			res = cnt[i]
			maxLen = dp[i]
		} else if dp[i] == maxLen {
			res += cnt[i]
		}
	}
	return res
}
