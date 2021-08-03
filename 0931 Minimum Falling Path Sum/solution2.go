// Your runtime beats 97.30 % of golang submissions (8 ms)
// Your memory usage beats 35.14 % of golang submissions (5.4 MB)
//
// DP + 空间压缩
// dp 为 n*n 的数组
// dp[i][j] 为 matrix[i][j] 为终点的最小下降路径
// 状态转移方程 dp[i][j] = min(matrix[i][j]+dp[i-1][j],
//          	              matrix[i][j]+dp[i-1][j-1],
//                            matrix[i][j]+dp[i-1][j+1])
// 初始条件 dp[0][i] = matrix[0][i]
// 空间压缩
// 因为每一行的结果只与上一行的结果有关，所以可以用一维数组进行空间压缩
//
// 时间复杂度 O(n^2)
// 空间复杂度 O(n)

func minFallingPathSum(matrix [][]int) int {
	dp := make([]int, len(matrix))
	copy(dp, matrix[0])
	for i := 1; i < len(matrix); i++ {
		newDp := make([]int, len(matrix))
		for j := 0; j < len(matrix); j++ {
			m := matrix[i][j] + dp[j]
			if j > 0 {
				m = min(m, matrix[i][j]+dp[j-1])
			}
			if j < len(matrix)-1 {
				m = min(m, matrix[i][j]+dp[j+1])
			}
			newDp[j] = m
		}
		dp = newDp
	}
	ret := 100 * len(matrix)
	for i := 0; i < len(matrix); i++ {
		ret = min(ret, dp[i])
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
