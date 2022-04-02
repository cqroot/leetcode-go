package P0053_MaximumSubarray

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var dpLast int = nums[0]
	var dpCurrent int
	var res int = nums[0]

	for i := 1; i < len(nums); i++ {
		dpCurrent = nums[i]
		if dpLast > 0 {
			dpCurrent = dpLast + nums[i]
		}
		if dpCurrent > res {
			res = dpCurrent
		}
		dpLast = dpCurrent
	}
	return res
}
