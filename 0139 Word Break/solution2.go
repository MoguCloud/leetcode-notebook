// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 44.06 % of golang submissions (2.25 MB)
//
// DP
// dp[i] 为字符串s[0..i]是否可以被拆分
// 初始条件 所有值都为 false，哨兵节点 dp[0] = true
// 状态转移方程 当 dp[j]=true 并且 s[j:i] 在 wordDict 中，此时 dp[i]=true
//
// 时间复杂度 O(n^2)
// 空间复杂度 O(n)

func wordBreak(s string, wordDict []string) bool {
	// wordSet 用于O(1)时间内判断子串是否在 wordDict中
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	dp := make([]bool, len(s)+1)
	// 哨兵节点置为 true
	dp[0] = true
	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				// 当满足条件时，不用判断其他情况
				dp[i] = true
				break
			}
		}
	}
	return dp[len(dp)-1]
}
