// Binary Search
// Time complexity  : O(log n)
// Space complexity : O(1)

package solution

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	var mid int

	for left <= right {
		mid = (left + right) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[left] {
			if nums[mid] > target && nums[left] <= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
