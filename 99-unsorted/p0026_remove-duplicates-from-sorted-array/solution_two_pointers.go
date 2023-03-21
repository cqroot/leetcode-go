// Two Pointers
// Time complexity  : O(n)
// Space complexity : O(1)

package solution

func removeDuplicates_TwoPointers(nums []int) int {
	left, right := 0, 1

	for right < len(nums) {
		if nums[right] != nums[left] {
			left++
			nums[left] = nums[right]
		}
		right++
	}

	return left + 1
}
