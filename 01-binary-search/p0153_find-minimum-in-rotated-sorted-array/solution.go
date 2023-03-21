// Binary Search
// Time complexity  : O(log n)
// Space complexity : O(1)

package solution

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	var mid int

	if len(nums) == 1 {
		return nums[0]
	}

	for left < right {
		mid = (left + right) / 2

		if nums[left] < nums[right] {
			return nums[left]
		} else {
			if nums[mid] > nums[right] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}
	return nums[right]
}
