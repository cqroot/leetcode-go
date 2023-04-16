// Two Pointers
// Time complexity  : O(n)
// Space complexity : O(1)

package solution

func sortColors(nums []int) {
	i, left, right := 0, 0, len(nums)-1

	for i <= right {
		if nums[i] == 0 && i > left {
			nums[i], nums[left] = nums[left], nums[i]
			left++

		} else if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--

		} else {
			i++
		}
	}
}
