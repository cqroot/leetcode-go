package P0413_arithmetic_slices

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	var res int = 0
	for i := len(nums) - 1; i >= 2; i-- {
		var d int = nums[i] - nums[i-1]
		var j int = i - 2
		for j >= 0 {
			if nums[j+1]-nums[j] != d {
				break
			}
			j -= 1
		}
		if i-j > 2 {
			res += i - j - 2
		}
	}
	return res
}
