// Binary Search
// Time complexity  : O(log n)
// Space complexity : O(1)

package solution

var pick int

// Forward declaration of guess API.
// param  num   your guess
// return
//
//	-1 if num is higher than the picked number
//	1 if num is lower than the picked number
//	otherwise return 0
func guess(num int) int {
	if num > pick {
		return -1
	} else if num < pick {
		return 1
	}
	return 0
}

func guessNumber(n int) int {
	left, right := 1, n
	var mid int

	for left <= right {
		mid = (left + right) / 2
		switch guess(mid) {
		case -1:
			right = mid - 1
		case 1:
			left = mid + 1
		default:
			return mid
		}
	}

	return 0
}
