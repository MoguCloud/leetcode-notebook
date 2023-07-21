// Your runtime beats 7.25 % of golang submissions (30 ms)
// Your memory usage beats 53.62 % of golang submissions (3.76 MB)
//
// DP
// 比LIS需要多记录个数
// dp[i] 为 nums[i] 结尾的最长子序列长度
// 初始条件 dp[i] = 1，自身可以构成长度为1的子序列
// 状态转移方程 dp[i] = max(dp[i], dp[j] + 1)， 当nums[i] > nums[j]
// 即当nums[i] > nums[j]时，nums[i]放在nums[j]后即可构成一个新的子序列
// count[i] 为 nums[i] 结尾的最长子序列的个数
// 初始条件 count[i] = 1，自身构成长度为1的子序列的个数为1
// 状态转移方程 count[i] = sum(count[j])，当nums[i] > nums[j]，且dp[i] = dp[j] + 1
// 最终结果为 sum(count[i])， 其中 i 满足 dp[i] == lis
// 即所有长度为lis的个数和
//
// 时间复杂度 O(n^2)
// 空间复杂度 O(n)

func findNumberOfLIS(nums []int) int {
	// dp[i] 为 nums[i] 结尾的最长子序列长度
	// count[i] 为 nums[i] 结尾的最长子序列的个数
	lis := 0
	ret := 0
	dp := make([]int, len(nums))
	count := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		c := 0
		dp[i] = 1
		count[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] == dp[j]+1 {
				c += count[j]
			}
		}
		count[i] = max(count[i], c)
		lis = max(lis, dp[i])
	}
	for i := 0; i < len(nums); i++ {
		if dp[i] == lis {
			ret += count[i]
		}
	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
