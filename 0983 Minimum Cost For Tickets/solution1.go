// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 86.36 % of golang submissions (2.4 MB)
//
// DP
// dp[i]为到第i天的最低消费
// 当第i天不需要旅行 dp[i] = dp[i-1]
// 当第i天需要旅行   dp[i] = min{dp[i-1]+costs[0], dp[i-7]+costs[1], dp[i-30]+costs[2]
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func mincostTickets(days []int, costs []int) int {
	dp := make([]int, days[len(days)-1]+1)
	idx := 0
	for i := 1; i < len(dp); i++ {
		// 判断第i天是否需要出行
		if i == days[idx] {
			// 通过使用dp[max(0, i)]来避免i<0的情况
			dp[i] = min(dp[max(0, i-1)]+costs[0], dp[max(0, i-7)]+costs[1], dp[max(0, i-30)]+costs[2])
			idx++
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(dp)-1]
}

func min(i, j, k int) int {
	ret := i
	if i < ret {
		ret = i
	}
	if j < ret {
		ret = j
	}
	if k < ret {
		ret = k
	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
