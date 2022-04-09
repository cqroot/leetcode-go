package P1055_shortest_way_to_form_string

func shortestWay(source string, target string) int {
	var dp []int = make([]int, len(target))
	var idx int = 0
	var found bool = false

	dp[0] = 0

	for i := 0; i < len(target); i++ {
		if i == 0 {
			dp[i] = 1
		} else {
			dp[i] = dp[i-1]
		}
		for j := idx; j < len(source); j++ {
			if target[i] == source[j] {
				idx = j + 1
				found = true
				break
			}
		}
		if !found {
			idx = 0
			dp[i] += 1
			for j := idx; j < len(source); j++ {
				if target[i] == source[j] {
					idx = j + 1
					found = true
					break
				}
			}
			if !found {
				return -1
			}
		}
		found = false
	}
	return dp[len(target)-1]
}
