// DP
//
// Your runtime beats 18.52 % of golang submissions (264 ms)
// Your memory usage beats 33.33 % of golang submissions (6.1 MB)
//
// 将信封按照长或者宽进行排序，然后套用 `LIS` 算法
//
// 时间复杂度 O(n^2)
// 空间复杂度 O(n)

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		} else {
			return false
		}
	})
	max := 0
	dp := make([]int, len(envelopes))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if envelopes[i][0] > envelopes[j][0] && envelopes[i][1] > envelopes[j][1] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
