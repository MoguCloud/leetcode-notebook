// Your runtime beats 97.30 % of golang submissions (8 ms)
// Your memory usage beats 35.14 % of golang submissions (5.4 MB)
//
// DP
// dp 为 n*n 的数组
// dp[i][j] 为 matrix[i][j] 为终点的最小下降路径
// 状态转移方程 dp[i][j] = min(matrix[i][j]+dp[i-1][j],
//          	              matrix[i][j]+dp[i-1][j-1],
//                            matrix[i][j]+dp[i-1][j+1])
// 初始条件 dp[0][i] = matrix[0][i]
//
// 时间复杂度 O(n^2)
// 空间复杂度 O(n^2)

func minFallingPathSum(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix))
		dp[0][i] = matrix[0][i]
	}
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			m := matrix[i][j] + dp[i-1][j]
			if j > 0 {
				m = min(m, matrix[i][j]+dp[i-1][j-1])
			}
			if j < len(matrix)-1 {
				m = min(m, matrix[i][j]+dp[i-1][j+1])
			}
			dp[i][j] = m
		}
	}
	ret := 100 * len(matrix)
	for i := 0; i < len(matrix); i++ {
		ret = min(ret, dp[len(dp)-1][i])
	}
	return ret
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
