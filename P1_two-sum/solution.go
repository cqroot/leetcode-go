package solution

// Brute Force
// Time complexity: O(n^2)
// Space complexity: O(1)
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

// Hash Table
// Time complexity: O(n)
// Space complexity: O(n)
func twoSum_HashTable(nums []int, target int) []int {
	table := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if j, ok := table[target-nums[i]]; ok {
			return []int{j, i}
		}

		table[nums[i]] = i
	}
	return nil
}
