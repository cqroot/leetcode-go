// Two Pointers
// Time complexity  : O(n)
// Space complexity : O(1)

package solution

func removeElement_TwoPointers1(nums []int, val int) int {
	left, right := 0, 0

	for right < len(nums) {
		if nums[right] != val {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}

		right++
	}

	return left
}
