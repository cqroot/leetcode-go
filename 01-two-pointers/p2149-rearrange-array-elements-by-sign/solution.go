// Two Pointers
// Time complexity  : O(n)
// Space complexity : O(1)

package solution

func rearrangeArray(nums []int) []int {
	ans := make([]int, len(nums))
	pos, neg := 0, 0

	for _, num := range nums {
		if num >= 0 {
			ans[pos*2] = num
			pos++
		} else {
			ans[neg*2+1] = num
			neg++
		}
	}

	return ans
}
