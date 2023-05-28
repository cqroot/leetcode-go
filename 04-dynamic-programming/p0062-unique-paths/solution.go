// Dynamic Programming
// Time complexity  : O(mn)
// Space complexity : O(mn)

package solution

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)

	handleCell := func(i, j int) {
		if i == 0 && j == 0 {
			dp[i][j] = 1
			return
		}

		if i == 0 {
			dp[i][j] = dp[i][j-1]
			return
		}

		if j == 0 {
			dp[i][j] = dp[i-1][j]
			return
		}

		dp[i][j] = dp[i][j-1] + dp[i-1][j]
	}

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			handleCell(i, j)
		}
	}

	return dp[m-1][n-1]
}
