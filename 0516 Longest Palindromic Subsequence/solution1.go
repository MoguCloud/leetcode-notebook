// DP
//
// Your runtime beats 34.38 % of golang submissions (56 ms)
// Your memory usage beats 17.19 % of golang submissions (18.1 MB)
// dp[i][j] 为 字符串s的下标在[i, j]闭区间内的最大回文长度
// base case
// dp[i][i] = 1   0 <= i < len(s)
// 状态转移方程
// 1. 当 s[i] == s[j] 时
// dp[i][j] = dp[i+1][j-1] + 2
// 2. 当 s[1] != s[j] 时
// dp[i][j] = max(dp[i+1][j], dp[i][j-1])
// 填表顺序 从 i = 0, j = 1 开始 斜着对角线向右下层次遍历

func longestPalindromeSubseq(s string) int {
	ret := 1
	dp := make([][]int, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	for delta := 1; delta < len(dp); delta++ {
		for i, j := 0, delta; j < len(dp); i, j = i+1, j+1 {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				if dp[i+1][j] > dp[i][j-1] {
					dp[i][j] = dp[i+1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
			if dp[i][j] > ret {
				ret = dp[i][j]
			}
		}
	}
	return ret

}
