package P0300_LongestIncreasingSubsequence

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var maxAns int = 1
	var dp []int = make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
		if dp[i] > maxAns {
			maxAns = dp[i]
		}
	}
	return maxAns
}
