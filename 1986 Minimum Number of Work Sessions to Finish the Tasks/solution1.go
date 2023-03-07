// 参考 https://leetcode.cn/problems/minimum-number-of-work-sessions-to-finish-the-tasks/solution/zhuang-ya-dp-c-by-uthinkufunny-80n1/

func minSessions(tasks []int, sessionTime int) int {
	n := len(tasks)
	dp := make([]int, 1 << n)
	cost := make([]int, 1 << n)
	for i := 0; i < len(dp); i++ {
		dp[i] = 15
		cost[i] = 15
	}
    dp[0] = 0
	for i := 1; i < len(dp); i++ {
		for j := 0; j < n; j++ {
			if (1 << j) & i > 0 {
				if cost[i^(1<<j)] + tasks[j] <= sessionTime {
					cost[i] = min(cost[i], cost[i^(1<<j)] + tasks[j])
					dp[i] = min(dp[i], dp[i^(1<<j)])
				} else {
					cost[i] = min(cost[i], tasks[j])
					dp[i] = min(dp[i], dp[i^(1<<j)]+1)
				}
			}
		}
	}
	return dp[len(dp)-1]
}

func min(i, j int) int {
	if i < j {
		return i 
	} else {
		return j
	}
}
