package P0213_house_robber_ii

func _rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}

	var dp []int = make([]int, len(nums))
	var res int

	dp[0] = nums[0]
	if nums[0] > nums[1] {
		dp[1] = nums[0]
		res = nums[0]
	} else {
		dp[1] = nums[1]
		res = nums[1]
	}

	for i := 2; i < len(nums); i++ {
		if dp[i-1] > dp[i-2]+nums[i] {
			dp[i] = dp[i-1]
		} else {
			dp[i] = dp[i-2] + nums[i]
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res1 := _rob(nums[0 : len(nums)-1])
	res2 := _rob(nums[1:len(nums)])
	if res1 > res2 {
		return res1
	} else {
		return res2
	}
}
