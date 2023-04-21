// Backtracking
// Time complexity  : O(2^n)
// Space complexity : O(n)

package solution

func subsetXORSum(nums []int) int {
	ans := 0

	var dfs func(int, int)
	dfs = func(i, val int) {
		if i >= len(nums) {
			ans += val
			return
		}
		dfs(i+1, val)
		dfs(i+1, val^nums[i])
	}
	dfs(0, 0)

	return ans
}
