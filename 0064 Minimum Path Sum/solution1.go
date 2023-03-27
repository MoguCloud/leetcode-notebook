// Your runtime beats 34.97 % of golang submissions (10 ms)
// Your memory usage beats 43.17 % of golang submissions (4.4 MB)
//
// DP
// dp[i][j] 表示从坐标(0, 0)的到坐标(i, j)的最短路径和
// 状态转移方程 dp[i][j] = min(dp[i-1][j], dp[i][j-1])) + grid[i][j]
//
// 时间复杂度 O(m*n)
// 空间复杂度 O(m*n)

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	dp[0] = make([]int, len(grid[0]))
	dp[0][0] = grid[0][0]
	for i := 1; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[i]))
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}
	for i := 1; i < len(grid[0]); i++ {
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
