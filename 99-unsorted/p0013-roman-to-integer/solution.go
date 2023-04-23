// Simulation
// Time complexity  : O(n)
// Space complexity : O(1)

package solution

var symVals = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	ans := 0
	prev := 0

	for i := len(s) - 1; i >= 0; i-- {
		curr := symVals[s[i]]

		if curr < prev {
			ans -= curr
		} else {
			ans += curr
		}

		prev = curr
	}

	return ans
}
