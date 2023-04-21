// Backtracking
// Time complexity  : O(n)
// Space complexity : O(n)

package solution

func generateParenthesis(n int) []string {
	ans := make([]string, 0)

	var dfs func(string, int, int)
	dfs = func(str string, left int, right int) {
		if right > left || left > n {
			return
		}
		if right == n && left == n {
			ans = append(ans, str)
			return
		}

		dfs(str+"(", left+1, right)
		dfs(str+")", left, right+1)
	}

	dfs("", 0, 0)
	return ans
}
