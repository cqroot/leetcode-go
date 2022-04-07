package P0740_delete_and_earn

func deleteAndEarn(nums []int) int {
	var maxVal int = 0
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}
	var sum []int = make([]int, maxVal)
	for _, num := range nums {
		sum[num-1] += num
	}
	return rob(sum)
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
