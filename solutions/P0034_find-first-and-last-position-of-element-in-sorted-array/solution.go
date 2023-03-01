// Binary Search
// Time complexity  : O(log n)
// Space complexity : O(1)

package solution

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	var mid int
	result := make([]int, 2)

	search := func(l int, r int, left bool) int {
		var m int

		for l <= r {
			m = (l + r) / 2

			if nums[m] == target {
				if left {
					r = m - 1
				} else {
					l = m + 1
				}
			} else if nums[m] > target {
				r = m - 1
			} else {
				l = m + 1
			}
		}

		if left {
			return l
		} else {
			return r
		}
	}

	for left <= right {
		mid = (left + right) / 2

		if nums[mid] == target {
			result[0] = search(left, mid, true)
			result[1] = search(mid, right, false)
			return result
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return []int{-1, -1}
}
