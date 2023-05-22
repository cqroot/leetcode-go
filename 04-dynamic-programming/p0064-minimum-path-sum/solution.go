// Dynamic Programming
// Time complexity  : O(mn)
// Space complexity : O(1)

package solution

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))

	handleCell := func(i, j int) {
		if i == 0 && j == 0 {
			dp[i][j] = grid[i][j]
			return
		}

		if i == 0 {
			dp[i][j] = dp[i][j-1] + grid[i][j]
			return
		}

		if j == 0 {
			dp[i][j] = dp[i-1][j] + grid[i][j]
			return
		}

		if dp[i-1][j] < dp[i][j-1] {
			dp[i][j] = dp[i-1][j] + grid[i][j]
		} else {
			dp[i][j] = dp[i][j-1] + grid[i][j]
		}
	}

	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))

		for j := 0; j < len(grid[0]); j++ {
			handleCell(i, j)
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}
