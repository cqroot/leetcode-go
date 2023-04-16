// Two Pointers
// Time complexity  : O(n)
// Space complexity : O(1)

package solution

func removeDuplicates(nums []int) int {
	left, right := 0, 1

	first := true
	for right < len(nums) {
		if nums[right] != nums[left] {
			left++
			nums[left] = nums[right]
			first = true
		} else if first {
			left++
			nums[left] = nums[right]
			first = false
		}

		right++
	}

	return left + 1
}
