// Two Pointers
// Time complexity  : O(n)
// Space complexity : O(1)

package solution

func removeElement_TwoPointers2(nums []int, val int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		if nums[left] == val {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		} else {
			left++
		}
	}

	return left
}
