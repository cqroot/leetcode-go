package P0027_RemoveElement

func removeElement(nums []int, val int) int {
	p1 := 0
	p2 := len(nums) - 1

	for p1 <= p2 {
		if nums[p1] == val {
			nums[p1] = nums[p2]
			nums[p2] = val
			p2--
		} else {
			p1++
		}
	}
	return p1
}
