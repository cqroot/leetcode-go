package P0873_length_of_longest_fibonacci_subsequence

func lenLongestFibSubseq(arr []int) int {
	var dp [][]int = make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = make([]int, len(arr))
	}
	var res int = 0

	var intMap map[int]int = make(map[int]int)
	for i := 0; i < len(arr); i++ {
		intMap[arr[i]] = i
	}

	dp[0][1] = 0
	if arr[0]+arr[1] == arr[2] {
		dp[1][2] = 3
	} else {
		dp[1][2] = 0
	}
	for i := 3; i < len(arr); i++ {
		for j := 1; j < i; j++ {
			if k, ok := intMap[arr[i]-arr[j]]; ok {
				if k < j {
					if dp[k][j]+1 > 3 {
						dp[j][i] = dp[k][j] + 1
					} else {
						dp[j][i] = 3
					}
					if dp[j][i] > res {
						res = dp[j][i]
					}
				}
			}
		}
	}
	return res
}
