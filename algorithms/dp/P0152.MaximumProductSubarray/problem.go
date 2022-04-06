package P0152_MaximumProductSubarray

func maxProduct(nums []int) int {
	var dpMax []int = make([]int, len(nums))
	var dpMin []int = make([]int, len(nums))
	dpMax[0] = nums[0]
	dpMin[0] = nums[0]
	var res int = nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			if dpMax[i-1] > 0 {
				dpMax[i] = dpMax[i-1] * nums[i]
			} else {
				dpMax[i] = nums[i]
			}
			if dpMin[i-1] <= 0 {
				dpMin[i] = dpMin[i-1] * nums[i]
			} else {
				dpMin[i] = nums[i]
			}
		} else if nums[i] < 0 {
			if dpMin[i-1] <= 0 {
				dpMax[i] = dpMin[i-1] * nums[i]
			} else {
				dpMax[i] = nums[i]
			}
			if dpMax[i-1] > 0 {
				dpMin[i] = dpMax[i-1] * nums[i]
			} else {
				dpMin[i] = nums[i]
			}
		} else {
			dpMax[i] = 0
			dpMin[i] = 0
		}
		if dpMax[i] > res {
			res = dpMax[i]
		}
	}
	return res
}
