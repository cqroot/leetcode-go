package p0001twosum

func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int)
	for index, num := range nums {
		i, ok := hashTable[target-num]
		if ok == true {
			return []int{i, index}
		}
		hashTable[num] = index
	}
	return nil
}
