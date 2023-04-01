// Backtracking
// Time complexity  : O()
// Space complexity : O(len(digits))

package solution

func letterCombinations(digits string) []string {
	diglen := len(digits)
	ans := make([]string, 0)
	letters := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var dfs func(int, []rune)
	dfs = func(i int, lets []rune) {
		if i >= diglen {
			ans = append(ans, string(lets))
			return
		}
		for _, let := range letters[digits[i]] {
			dfs(i+1, append(lets, let))
		}
	}

	if len(digits) != 0 {
		dfs(0, make([]rune, 0))
	}
	return ans
}
