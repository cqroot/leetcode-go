// XOR
// Time complexity  : O(len(nums))
// Space complexity : O(1)

package solution

func singleNumber(nums []int) int {
	var result int = 0

	for _, num := range nums {
		result = result ^ num
	}

	return result
}
