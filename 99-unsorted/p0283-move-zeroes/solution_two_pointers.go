// Two Pointers
// Time complexity  : O(n)
// Space complexity : O(1)

package solution

func moveZeroes_TwoPointers(nums []int) {
	left, right := 0, 0

	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}

		right++
	}
}
