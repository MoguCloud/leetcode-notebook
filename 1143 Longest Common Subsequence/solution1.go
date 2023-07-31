// Your runtime beats 64.75 % of golang submissions (15 ms)
// Your memory usage beats 23.39 % of golang submissions (19.4 MB)
//
// DP
// dp[i+1][j+1] 为 text1在区间[0, i] text2在区间[0,j] 内的最长公共子序列长度
// 初始状态
// dp[0][..] = 0  因为长度为0的text1子序列和任意长度的text2子序列的公共子序列长度为0
// dp[..][0] = 0  因为任意长度的text1子序列和长度为0的text2子序列的公共子序列长度为0
// 状态转移方程
// 当 text1[i] == text2[j] 时
// dp[i+1][j+1] = dp[i][j] + 1
// 当 text1[i] != text2[j] 时
// dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1]) text1[0..i]和text2[0..j]的公共子序列长度为 text1[0..i-1]和text2[0..j]的公共子序列长度 和 text1[0..i]和text2[0..j-1]的公共子序列长度 的最大值
//
// 时间复杂度 O(m*n)
// 空间复杂度 O(m*n)

func longestCommonSubsequence(text1 string, text2 string) int {
	ret := 0
	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
				ret = max(ret, dp[i+1][j+1])
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return ret
}

func max(nums ...int) int {
	ret := nums[0]
	for _, num := range nums {
		if num > ret {
			ret = num
		}
	}
	return ret
}
