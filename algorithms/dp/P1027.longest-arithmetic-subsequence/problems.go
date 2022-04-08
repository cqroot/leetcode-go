package P1027_longest_arithmetic_subsequence

func longestArithSeqLength(nums []int) int {
	var dp [][]int = make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 1001)
	}
	var res int = 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			var diff int = nums[i] - nums[j] + 500
			if dp[j][diff] == 0 {
				dp[j][diff] = 1
			}
			if dp[j][diff]+1 > dp[i][diff] {
				dp[i][diff] = dp[j][diff] + 1
				if dp[i][diff] > res {
					res = dp[i][diff]
				}
			}
		}
	}
	return res
}
