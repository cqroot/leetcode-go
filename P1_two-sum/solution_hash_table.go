package solution

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
