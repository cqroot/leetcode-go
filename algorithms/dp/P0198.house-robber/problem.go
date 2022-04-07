package P0198_house_robber

func rob(nums []int) int {
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
	dp[1] = nums[1]
	dp[2] = nums[2] + nums[0]
	if dp[0] > dp[1] && dp[0] > dp[2] {
		res = dp[0]
	} else if dp[1] > dp[0] && dp[1] > dp[2] {
		res = dp[1]
	} else {
		res = dp[2]
	}

	for i := 3; i < len(nums); i++ {
		if dp[i-2] > dp[i-3] {
			dp[i] = dp[i-2] + nums[i]
		} else {
			dp[i] = dp[i-3] + nums[i]
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
