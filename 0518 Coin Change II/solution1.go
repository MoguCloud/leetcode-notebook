// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 25.77 % of golang submissions (2.2 MB)
//
// DP
//
// 时间复杂度 O(amount * len(coins))
// 空间复杂度 O(n)

func change(amount int, coins []int) int {
	// dp[i] 金额为i的组合数量
	dp := make([]int, amount+1)
	dp[0] = 1
	// 外层循环是coins，内层循环是amount，可以避免金额为3时被分解成1+2和2+1
	// 这样循环是求组合数，内外层循环反过来是求排列数
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}
