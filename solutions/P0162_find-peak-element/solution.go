// Binary Search
// Time complexity  : O(log n)
// Space complexity : O(1)

package solution

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	var mid int

	for left < right {
		mid = (left + right) / 2

		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
