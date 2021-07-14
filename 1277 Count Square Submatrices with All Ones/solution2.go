// Your runtime beats 98.28 % of golang submissions (60 ms)
// Your memory usage beats 65.52 % of golang submissions (7.2 MB)
//
// DP
//
// dp[i][j] 为 右下角为matrix[i][j]最成的正方形的个数
// 状态转移方程
// matrix[i][j] = 0 时 dp[i][j] = 0
// matrix[i][j] = 1 时 dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
// 初始状态
// dp[0][j] = matrix[0][j]
// dp[i][0] = matrix[i][0]
//
// 时间复杂度 O(M * N)
// 空间复杂度 O(M * N)

func countSquares(matrix [][]int) int {
	ret := 0
	dp := make([][]int, len(matrix))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = matrix[i][j]
			} else {
				if matrix[i][j] == 0 {
					dp[i][j] = 0
				} else {
					dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				}
			}
			ret += dp[i][j]
		}
	}
	return ret
}

func min(nums ...int) int {
	n := nums[0]
	for i := 1; i < len(nums); i++ {
		if n > nums[i] {
			n = nums[i]
		}
	}
	return n
}
