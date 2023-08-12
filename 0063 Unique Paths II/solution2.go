// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 40.15 % of golang submissions (2.56 MB)
//
// DP
// dp[i][j]为走到第(i, j)格的路径数
// base case
// 当(0, 0)无障碍物时 dp[0][0] = 1
// 当(0, 0)有障碍物时 dp[0][0] = 0
// 状态转移方程
// 当(i, j)无障碍物时 dp[i][j] = dp[i-1][j] + dp[i][j-1] // 走到(i, j)的路径数为走到其上一个和左一个的路径数之和
// 当(i, j)有障碍物时 dp[i][j] = 0
//
// 时间复杂度 O(m*n)
// 空间复杂度 O(m*n)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			if i == 0 && j == 0 && obstacleGrid[i][j] == 0 {
				dp[i][j] = 1
			} else {
				if obstacleGrid[i][j] == 1 {
					continue
				}
				if i > 0 {
					dp[i][j] += dp[i-1][j]
				}
				if j > 0 {
					dp[i][j] += dp[i][j-1]
				}
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}
