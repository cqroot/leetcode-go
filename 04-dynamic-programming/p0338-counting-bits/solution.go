// Dynamic Programming
// Time complexity  : O(n)
// Space complexity : O(1)

package solution

func countBits(n int) []int {
	if n <= 0 {
		return []int{0}
	}

	ans := make([]int, n+1)
	isOdd := true
	ans[0] = 0

	for i := 1; i <= n; i++ {
		if isOdd {
			ans[i] = ans[i-1] + 1
		} else {
			ans[i] = ans[i/2]
		}

		isOdd = !isOdd
	}

	return ans
}
