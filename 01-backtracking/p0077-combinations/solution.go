// Backtracking
// Time complexity  : O()
// Space complexity : O(n)

package solution

func combine(n int, k int) [][]int {
	ans := make([][]int, 0, n*2)

	var dfs func(int, []int)
	dfs = func(i int, arr []int) {
		if len(arr) == k {
			newArr := make([]int, len(arr))
			copy(newArr, arr)
			ans = append(ans, newArr)
			return
		}
		// Pruning
		if len(arr)+n-i+1 < k {
			return
		}
		if i > n {
			return
		}
		dfs(i+1, arr)
		dfs(i+1, append(arr, i))
	}

	dfs(1, make([]int, 0))
	return ans
}
