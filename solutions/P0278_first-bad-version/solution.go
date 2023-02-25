// Binary Search
// Time complexity  : O(log n)
// Space complexity : O(1)

package solution

var badVersion int

// Forward declaration of isBadVersion API.
// param:
//
//	version:your guess about first bad version
//
// return:
//
//	true if current version is bad
//	false if current version is good
func isBadVersion(version int) bool {
	if version >= badVersion {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	left, right := 0, n
	var mid int

	for left <= right {
		mid = (left + right) / 2

		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
