// Brute Force
// Time complexity  : O(len(nums)^2)
// Space complexity : O(1)

package solution

func twoSum_BruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
