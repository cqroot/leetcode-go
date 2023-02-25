// Binary Search
// Time complexity  : O(log n)
// Space complexity : O(1)

package solution

func mySqrt(x int) int {
	left, right := 0, x
	var mid int

	for left <= right {
		mid = (left + right) / 2

		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

// func mySqrt(x int) int {
// 	left, right := 0, x
// 	var mid int
// 	ans := -1
//
// 	for left <= right {
// 		mid = (left + right) / 2
//
// 		if mid*mid == x {
// 			return mid
// 		} else if mid*mid < x {
// 			ans = mid
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}
// 	}
//
// 	return ans
// }
