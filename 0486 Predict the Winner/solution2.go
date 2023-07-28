// Your runtime beats 83.33 % of golang submissions (1 ms)
// Your memory usage beats 70.83 % of golang submissions (2.10 MB)
//
// DP
// 根据递归如何变成求解子问题可以转换成动态规划
// 记忆化的时候存储了不同i和j的情况下先手与后手相差多少分
// 所以 dp[i][j] 为 nums在区间[i:j]内 先手与后手相差的得分
// 递归的终止条件是 i == j 时，返回 nums[i]
// 所以 dp 的初始条件 dp[i][i] = nums[i]
// 递归的每一步是在求解选了开头或者结尾的得分会比后手相差多少
// 递归返回值为 max(nums[i] - helper(i+1, j), nums[j] - helper(i, j-1))
// 所以 dp 的状态转移方程为 dp[i][j] = max(nums[i] - dp[i+1][j], nums[j] - dp[i][j-1]
//
// 时间复杂度 O(n)
// 空间复杂度 O(n^2)

func PredictTheWinner(nums []int) bool {
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
		dp[i][i] = nums[i]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	return dp[0][len(nums)-1] >= 0
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
