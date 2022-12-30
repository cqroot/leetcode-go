// Binary Search
// Time complexity  : O(logn)
// Space complexity : O(1)

package solution

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	var ans int

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			ans = mid + 1
			left = mid + 1
		} else {
			ans = mid
			right = mid - 1
		}
	}

	return ans
}
