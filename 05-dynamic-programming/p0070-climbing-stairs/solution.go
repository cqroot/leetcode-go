// Dynamic Programming
// Time complexity  : O(n)
// Space complexity : O(1)

package solution

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	prevPrev := 1 // dp[0]
	prev := 2     // dp[1]
	curr := 2

	for i := 2; i < n; i++ {
		curr = prevPrev + prev

		prevPrev = prev
		prev = curr
	}

	return curr
}
