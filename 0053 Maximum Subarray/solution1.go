// Your runtime beats 86.25 % of golang submissions (91 ms)
// Your memory usage beats 73.98 % of golang submissions (8.32 MB)
//
// DP
// dp[i + 1] 为以 nums[i] 结尾的连续子数组的最大和
// Base Case
// dp[0] = 0
// 状态转移方程
// dp[i + 1] = max(dp[i] + nums[i], nums[i])
// 要么和之前的数组继续构成连续子数组， 要么不和之前的数组构成连续子数组，作为新数组
// 这两种情况取最大值
// 因为dp[i + 1]只和dp[i]有关系
// 所以可以状态压缩成 dp = max(dp + nums[i], nums[i])
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func maxSubArray(nums []int) int {
	ret := -10000
	dp := 0
	for i := 0; i < len(nums); i++ {
		dp = max(dp+nums[i], nums[i])
		ret = max(ret, dp)
	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
