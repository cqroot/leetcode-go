// Backtracking
// Time complexity  : O(1)
// Space complexity : O(1)

package solution

import "fmt"

func readBinaryWatch(turnedOn int) []string {
	result := make([]string, 0)

	hours := []int{1, 2, 4, 8, 0, 0, 0, 0, 0, 0}
	minutes := []int{0, 0, 0, 0, 1, 2, 4, 8, 16, 32}

	var backtrack func(int, int, int, int)
	backtrack = func(turnedOn int, index int, hour int, minute int) {
		if hour > 11 {
			return
		}
		if minute > 59 {
			return
		}
		if turnedOn == 0 {
			result = append(result, fmt.Sprintf("%d:%02d", hour, minute))
		}
		for ; index < 10; index++ {
			backtrack(turnedOn-1, index+1, hour+hours[index], minute+minutes[index])
		}
	}

	backtrack(turnedOn, 0, 0, 0)
	return result
}
