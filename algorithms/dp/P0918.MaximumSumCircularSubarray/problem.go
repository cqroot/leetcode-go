package P0918_MaximumSumCircularSubarray

func maxSubarraySumCircular(nums []int) int {
	var dpMax []int = make([]int, len(nums))
	var dpMin []int = make([]int, len(nums))
	dpMax[0] = nums[0]
	dpMin[0] = nums[0]
	var max = nums[0]
	var min = nums[0]
	var sum = nums[0]

	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if dpMax[i-1] > 0 {
			dpMax[i] = dpMax[i-1] + nums[i]
		} else {
			dpMax[i] = nums[i]
		}
		if dpMin[i-1] < 0 {
			dpMin[i] = dpMin[i-1] + nums[i]
		} else {
			dpMin[i] = nums[i]
		}
		if dpMax[i] > max {
			max = dpMax[i]
		}
		if dpMin[i] < min {
			min = dpMin[i]
		}
	}
	if max < 0 || max > sum-min {
		return max
	} else {
		return sum - min
	}
}
